package cmd

import (
	"context"
	"fmt"
	"github.com/rs/cors"
	"github.com/rs/zerolog/log"
	"github.com/valyala/fasthttp"
	"net/http"

	"runtime/debug"
	"strings"
	"sync/atomic"
	"time"
)

type ctxKeyRequestID int

const RequestIDKey ctxKeyRequestID = 0

var reqID uint64
var prefix = "gol"

func ErrorConverter(err ErrorSlug) (string, int, ErrorSlug) {
	info := ErrList()
	if v, ok := info[err]; ok {
		return v.Message, v.HttpCode, err
	}
	return ErrText(ErrSlugInternalError), fasthttp.StatusInternalServerError, err
}

func Handle(h http.Handler) http.Handler {
	h = cors.AllowAll().Handler(h)
	if false {
		h = addRequestID(addLogger(h))
	}
	return recoverer(h)
}

func recoverer(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rvr := recover(); rvr != nil {
				debug.PrintStack()
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func addRequestID(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		myid := atomic.AddUint64(&reqID, 1)
		ctx := r.Context()
		ctx = context.WithValue(ctx, RequestIDKey, fmt.Sprintf("%s-%06d", prefix, myid))
		h.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getReqID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	if reqID, ok := ctx.Value(RequestIDKey).(string); ok {
		return reqID
	}
	return ""
}

func addLogger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		var scheme string

		if r.Header.Get("X-Liveness-Probe") == "Healthz" {
			h.ServeHTTP(w, r)
			return
		}
		id := getReqID(ctx)
		if r.TLS != nil {
			scheme = "https"
		} else {
			scheme = "http"
		}
		uri := strings.Join([]string{scheme, "://", r.Host, r.RequestURI}, "")

		log.Debug().Msgf("request started: (request-id: %s, http-scheme: %s, http-proto: %s, http-method: %s, remote-addr: %s, user-agent: %s, uri: %s)",
			id, scheme, r.Proto, r.Method, r.RemoteAddr, r.UserAgent(), uri)
		t1 := time.Now()
		h.ServeHTTP(w, r)
		log.Debug().Msgf("request completed: (request-id: %s, http-scheme: %s, http-proto: %s, http-method: %s, remote-addr: %s, user-agent: %s, uri: %s, elapsed-ms: %g)",
			id, scheme, r.Proto, r.Method, r.RemoteAddr, r.UserAgent(), uri, float64(time.Since(t1).Nanoseconds())/1000000.0)
	})
}
