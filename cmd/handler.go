package cmd

import (
	"bytes"
	"encoding/json"
	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/rs/zerolog"
	"github.com/valyala/fasthttp"
	"net/http"

	"gopkg.in/workanator/go-floc.v2"
	"gopkg.in/workanator/go-floc.v2/run"
)

type ConvertErrorFunc func(*routing.Context, error) error

const ResponseSlug = "RESPONSE"

type Handler struct {
	api    *Rest
	logger zerolog.Logger
}

func NewHandler(api *Rest, logger zerolog.Logger) *Handler {
	return &Handler{api: api, logger: logger}
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

func (h *Handler) ErrorHandler(logger zerolog.Logger, errorf ...ConvertErrorFunc) routing.Handler {
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
			if e := logger.Debug(); e.Enabled() {
				e.Str("type", "error").Int("code", status).Msgf(mess)
			}
			_ = h.Response(status, c, nil, err.String(), mess)
			c.Abort()
		}
		return nil
	}
}

func (h *Handler) ExecuteHandler(c *routing.Context) error {
	var jobs map[string]interface{}
	var req Request
	var j []floc.Job
	if err := req.UnmarshalJSON(c.PostBody()); err != nil {
		return err
	}
	ctx := floc.NewContext()
	jobs = make(map[string]interface{})
	for _, sequence := range req.Sequence {
		switch sequence.Type {
		case Parallel:
			var pj []floc.Job
			for _, job := range sequence.Jobs {
				pj = append(pj, NewRequestTranslator(job, h.logger).FlocExecute())
			}
			j = append(j, run.Parallel(pj...))
		case Sync:
			var sj []floc.Job
			for _, job := range sequence.Jobs {
				sj = append(sj, NewRequestTranslator(job, h.logger).FlocExecute())
			}
			j = append(j, sj...)
		}
	}
	flow := run.Sequence(j...)
	_, _, err := floc.RunWith(ctx, floc.NewControl(ctx), flow)
	if err != nil {
		return err
	}
	for _, sequence := range req.Sequence {
		for _, job := range sequence.Jobs {
			if v, ok := ctx.Value(job.Id).(string); ok {
				jobs[job.Id] = v
			}
		}
	}
	return h.Response(fasthttp.StatusOK, c, jobs, "", "")
}

func (h *Handler) PingHandler(ctx *routing.Context) error {
	resp := struct {
		Version string `json:"version"`
	}{Version: "1.0"}
	return h.Response(fasthttp.StatusOK, ctx, resp, "PONG", "")
}
