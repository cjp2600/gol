package cmd

import (
	"context"
	pb "github.com/cjp2600/gol/proto"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"time"

	grpcRecover "github.com/grpc-ecosystem/go-grpc-middleware/recovery"

	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

type Server struct {
	handler *Handler
	logger  zerolog.Logger
	port    string
	grpc    string
}

func NewServer(logger zerolog.Logger, port string, grpc string) *Server {
	r := Server{logger: logger}
	r.handler = NewHandler(&r, logger)
	r.port = port
	r.grpc = grpc
	return &r
}

func (r *Server) RunHttp(ctx context.Context) error {
	customMarshaller := &runtime.JSONPb{
		OrigName:     true,
		EmitDefaults: true,
	}
	muxOpt := runtime.WithMarshalerOption(runtime.MIMEWildcard, customMarshaller)
	mux := runtime.NewServeMux(muxOpt)
	runtime.HTTPError = r.handler.CustomHTTPError
	opts := []grpc.DialOption{grpc.WithInsecure()}

	if err := pb.RegisterGolHandlerFromEndpoint(ctx, mux, "localhost:"+r.port, opts); err != nil {
		return err
	}
	r.logger.Info().Msg("Start http server port: " + r.port)
	s := &http.Server{
		Addr:    ":" + r.port,
		Handler: Handle(mux),
	}
	return s.ListenAndServe()
}

func (r *Server) RunGRPC() {
	lis, err := net.Listen("tcp", ":"+r.grpc)
	if err != nil {
		r.logger.Fatal().Msgf("failed to listen: %v", err)
	}
	r.logger.Info().Msgf("Start grpc server port: %s", r.grpc)
	opts := []grpcRecover.Option{
		grpcRecover.WithRecoveryHandler(InterceptorOptionHandler),
	}
	s := grpc.NewServer(
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: 5 * time.Minute,
		}),
		grpcMiddleware.WithUnaryServerChain(
			// error recover
			grpcRecover.UnaryServerInterceptor(opts...),
		),
		grpcMiddleware.WithStreamServerChain(
			// error recover
			grpcRecover.StreamServerInterceptor(opts...),
		),
	)
	pb.RegisterGolServer(s, r.handler)
	if err := s.Serve(lis); err != nil {
		r.logger.Fatal().Msgf("failed to serve: %v", err)
	}
}
