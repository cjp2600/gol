package cmd

import (
	"context"
	"encoding/json"
	pb "github.com/cjp2600/gol/proto"
	"github.com/go-resty/resty"
	"github.com/rs/zerolog"
	"github.com/yalp/jsonpath"
	"gopkg.in/workanator/go-floc.v2"
	"strings"
)

type RequestTranslator struct {
	Job    *pb.Job
	Ctx    context.Context
	client *resty.Client
	logger zerolog.Logger
}

func NewRequestTranslator(job *pb.Job, logger zerolog.Logger) *RequestTranslator {
	return &RequestTranslator{Job: job, client: resty.New(), logger: logger}
}

func (r *RequestTranslator) setType(ctx floc.Context, n, t string, val interface{}) {
	switch strings.ToLower(t) {
	case "string":
		ctx.AddValue("$"+n, val.(string))
	case "int":
		ctx.AddValue("$"+n, val.(int))
	case "int32":
		ctx.AddValue("$"+n, val.(int32))
	case "int64":
		ctx.AddValue("$"+n, val.(int64))
	}
}

func (r *RequestTranslator) FindVarByJPath(js string, jPath string) (interface{}, error) {
	var store interface{}
	err := json.Unmarshal([]byte(js), &store)
	if err != nil {
		return nil, err
	}
	val, err := jsonpath.Read(store, jPath)
	if err != nil {
		return nil, err
	}
	return val, nil
}

func (r *RequestTranslator) FlocExecute() func(ctx floc.Context, ctrl floc.Control) error {
	return func(ctx floc.Context, ctrl floc.Control) error {
		resp, err := r.Execute(ctx)
		if err != nil {
			return err
		}
		ctx.AddValue(r.Job.Id, resp.String())
		if r.Job.Var != nil {
			for _, v := range r.Job.Var {
				if len(v.JPath) > 0 {
					val, err := r.FindVarByJPath(resp.String(), v.JPath)
					if err != nil {
						return err
					}
					r.setType(ctx, v.Name, v.Type, val)
				}
			}
		}
		return nil
	}
}

func (r *RequestTranslator) Execute(ctx floc.Context) (*resty.Response, error) {
	var err error
	var response *resty.Response
	req := r.client.R()
	req.EnableTrace()

	if r.Job.Body != nil {
		req.SetBody(r.Job.Body)
	}
	/*	if r.Job.Header != nil {
		req.SetHeaders(r.Job.GetHeaders(ctx))
	}*/
	switch r.Job.Method {
	case pb.Methods_get:
		response, err = req.Get(r.Job.GetUrl())
	case pb.Methods_post:
		response, err = req.Post(r.Job.GetUrl())
	case pb.Methods_put:
		response, err = req.Put(r.Job.GetUrl())
	case pb.Methods_patch:
		response, err = req.Delete(r.Job.GetUrl())
	}
	if err != nil {
		return nil, err
	}
	if e := r.logger.Debug(); e.Enabled() {
		e.Str("type", "request").
			Str("method", r.Job.Method.String()).
			Str("url", r.Job.GetUrl()).
			Dur("connTime", req.TraceInfo().ConnTime).
			Dur("serverTime", req.TraceInfo().ServerTime).
			Dur("responseTime", req.TraceInfo().ResponseTime).
			Dur("totalTime", req.TraceInfo().TotalTime).
			Msgf(string(response.Body()))
	}
	return response, nil
}
