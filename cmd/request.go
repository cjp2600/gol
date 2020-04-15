package cmd

import (
	"context"
	"github.com/go-resty/resty"
	"github.com/rs/zerolog"
	"gopkg.in/workanator/go-floc.v2"
	"strings"
)

type RequestTranslator struct {
	Job    Job
	Ctx    context.Context
	client *resty.Client
	logger zerolog.Logger
}

func NewRequestTranslator(job Job, logger zerolog.Logger) *RequestTranslator {
	return &RequestTranslator{Job: job, client: resty.New(), logger: logger}
}

func (r *RequestTranslator) FlocExecute() func(ctx floc.Context, ctrl floc.Control) error {
	return func(ctx floc.Context, ctrl floc.Control) error {
		resp, err := r.Execute(ctx)
		if err != nil {
			return err
		}
		ctx.AddValue(r.Job.Id, resp.String())
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
	if r.Job.Header != nil {
		req.SetHeaders(r.Job.GetHeaders(ctx))
	}
	if r.Job.Method != "" {
		switch strings.ToLower(r.Job.Method) {
		case "get":
			response, err = req.Get(r.Job.GetUrl(ctx))
		case "post":
			response, err = req.Post(r.Job.GetUrl(ctx))
		case "put":
			response, err = req.Put(r.Job.GetUrl(ctx))
		case "delete":
			response, err = req.Delete(r.Job.GetUrl(ctx))
		}
	}
	if err != nil {
		return nil, err
	}
	if e := r.logger.Debug(); e.Enabled() {
		e.Str("type", "request").
			Str("method", r.Job.Method).
			Str("url", r.Job.GetUrl(ctx)).
			Dur("connTime", req.TraceInfo().ConnTime).
			Dur("serverTime", req.TraceInfo().ServerTime).
			Dur("responseTime", req.TraceInfo().ResponseTime).
			Dur("totalTime", req.TraceInfo().TotalTime).
			Msgf(string(response.Body()))
	}
	return response, nil
}
