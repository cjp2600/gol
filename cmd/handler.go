package cmd

import (
	"bytes"
	"encoding/json"
	routing "github.com/qiangxue/fasthttp-routing"
	"net/http"
)

type ConvertErrorFunc func(*routing.Context, error) error
const ResponseSlug = "RESPONSE"

type Handler struct {
	api *Rest
}

func NewHandler(api *Rest) *Handler {
	return &Handler{api: api}
}

func (h *Handler) JSON(status int, ctx *routing.Context, v interface{}) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(true)
	if err := enc.Encode(v); err != nil {
		ctx.Response.Header.Set("Content-Type", "text/plain; charset=utf-8")
		ctx.Response.Header.Set("X-Content-Type-Options", "nosniff")
		ctx.SetStatusCode(http.StatusInternalServerError)
		ctx.SetBody([]byte(err.Error()))
		return
	}
	ctx.Response.Header.Set("Content-Type", "application/json; charset=utf-8")
	if status > 0 {
		ctx.SetStatusCode(status)
	}
	ctx.SetBody(buf.Bytes())
}

func (h *Handler) Response(status int, c *routing.Context, data interface{}, code string, message string) error {
	if len(code) == 0 {
		code = ResponseSlug
	}
	if data == nil {
		data = []string{}
	}
	resp := struct {
		Status  int         `json:"status"`
		Code    string      `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data"`
	}{}
	resp.Status = status
	resp.Data = data
	resp.Code = code
	resp.Message = message
	h.JSON(status, c, resp)

	return nil
}

func (h *Handler) ErrorHandler(errorf ...ConvertErrorFunc) routing.Handler {
	return func(c *routing.Context) error {
		err := c.Next()
		if err == nil {
			return nil
		}
		if len(errorf) > 0 {
			err = errorf[0](c, err)
		}
		if err != nil {
			mess, status, err := ErrorConverter(ErrToSlug(err), c)
			_ = h.Response(status, c, nil, err.String(), mess)
			c.Abort()
		}
		return nil
	}
}

func (h *Handler) ExecuteHandler(ctx *routing.Context) error {
	return ErrSlugEmptyQuery.Error()
}
