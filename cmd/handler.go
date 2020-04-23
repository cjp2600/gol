package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	pb "github.com/cjp2600/gol/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"gopkg.in/workanator/go-floc.v2"
	"gopkg.in/workanator/go-floc.v2/run"
	"net/http"
)

type Handler struct {
	api    *Server
	logger zerolog.Logger
}

func NewHandler(api *Server, logger zerolog.Logger) *Handler {
	return &Handler{api: api, logger: logger}
}

func (h *Handler) CustomHTTPError(ctx context.Context, _ *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, req *http.Request, err error) {
	w.Header().Set("Content-type", marshaler.ContentType())
	resp := struct {
		Code    string `json:"code"`
		Message string `json:"message"`
		Status  int    `json:"status"`
	}{}
	message, status, code := ErrorConverter(ErrToSlug(fmt.Errorf(grpc.ErrorDesc(err))))
	resp.Code = code.String()
	resp.Message = message
	resp.Status = status

	h.logger.Error().Msg(fmt.Sprintf("[%s] - %s", resp.Code, resp.Message))

	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(true)
	if err := enc.Encode(resp); err != nil {
		http.Error(w, err.Error(), status)
		return
	}
	w.WriteHeader(status)
	w.Write(buf.Bytes())
}

func (h *Handler) Execute(c context.Context, request *pb.ExecuteRequest) (*pb.ExecuteResponse, error) {
	var response *pb.ExecuteResponse
	var jobs map[string]interface{}
	var j []floc.Job

	ctx := floc.NewContext()
	jobs = make(map[string]interface{})
	for _, sequence := range request.Sequence {
		switch sequence.Type {
		case pb.SequenceType_parallel:
			var pj []floc.Job
			for _, job := range sequence.Jobs {
				pj = append(pj, NewRequestTranslator(job, h.logger).FlocExecute())
			}
			j = append(j, run.Parallel(pj...))
		case pb.SequenceType_sync:
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
		return nil, err
	}
	for _, sequence := range request.Sequence {
		for _, job := range sequence.Jobs {
			if v, ok := ctx.Value(job.Id).(string); ok {
				jobs[job.Id] = v
			}
		}
	}
	return response, nil
}
