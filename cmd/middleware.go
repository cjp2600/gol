package cmd

import (
	"bytes"
	"fmt"
	routing "github.com/qiangxue/fasthttp-routing"
	"github.com/rs/zerolog"
	"github.com/valyala/fasthttp"
	"runtime"
)

func ErrorConverter(err ErrorSlug, c *routing.Context) (string, int, ErrorSlug) {
	info := ErrList(c)
	if v, ok := info[err]; ok {
		return v.Message, v.HttpCode, err
	}
	return ErrText(ErrSlugInternalError), fasthttp.StatusInternalServerError, err
}

func PanicHandler(logger zerolog.Logger) routing.Handler {
	return func(c *routing.Context) (err error) {
		defer func() {
			if e := recover(); e != nil {
				var ok bool
				if err, ok = e.(error); !ok {
					err = fmt.Errorf("%v", e)
					if err != nil {
						if e := logger.Debug(); e.Enabled() {
							e.Str("type", "panic").Msgf("%s, %v", err.Error(), getCallStack(4))
						}
					}
				}
			}
		}()
		return c.Next()
	}
}

func getCallStack(skip int) string {
	buf := new(bytes.Buffer)
	for i := skip; ; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		fmt.Fprintf(buf, "\n%s:%d", file, line)
	}
	return buf.String()
}
