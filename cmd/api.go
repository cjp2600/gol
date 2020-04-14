package cmd

import (
	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/rs/zerolog"
	"github.com/valyala/fasthttp"
)

type Rest struct {
	handler *Handler
	logger  zerolog.Logger
	port string
}

func NewRest(logger zerolog.Logger, port string) *Rest {
	r := Rest{logger: logger}
	r.handler = NewHandler(&r)
	r.port = port
	return &r
}

func (r *Rest) CreateRoute() *routing.Router {
	router := routing.New()
	router.Use(r.handler.ErrorHandler())
	router.Use(PanicHandler())
	return router
}

func (r *Rest) Routes() *routing.Router {
	router := r.CreateRoute()
	execute := router.Group("/execute")
	r.executeRoute(execute)
	return router
}

func (r *Rest) executeRoute(user *routing.RouteGroup) {
	user.Get("", r.handler.ExecuteHandler)
}

func (r *Rest) Run() error {
	if len(r.port) == 0 {
		r.port = "8081"
	}
	r.logger.Info().Msgf("listening port %v ...", r.port)
	h := r.Routes().HandleRequest

	h = fasthttp.CompressHandlerLevel(h, fasthttp.CompressBestSpeed)
	serv := fasthttp.Server{
		Handler:            h,
		MaxRequestBodySize: 100 * 1024 * 1024 * 1024,
	}
	if e := r.logger.Debug(); e.Enabled() {
		e.Msgf("start debug mode")
	}
	return serv.ListenAndServe(":" + r.port)
}
