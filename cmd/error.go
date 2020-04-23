package cmd

import (
	"errors"
	"fmt"
	"github.com/valyala/fasthttp"
)

const (
	ErrSlugInternalError ErrorSlug = "INTERNAL_ERROR"
	ErrSlugEmptyQuery    ErrorSlug = "EMPTY_QUERY"
	ErrSlugNotFound      ErrorSlug = "Not Found"
)

type ErrorSlug string
type ErrorText = map[ErrorSlug]string

type ErrorInfo struct {
	HttpCode int
	Message  string
}

func ErrText(slug ErrorSlug) string {
	m := make(map[ErrorSlug]string)
	m[ErrSlugInternalError] = "unknown error"
	m[ErrSlugEmptyQuery] = "request has not been sent"
	m[ErrSlugNotFound] = "resource not found"
	if v, ok := m[slug]; ok {
		return v
	}
	return "nil"
}

func (e ErrorSlug) String() string {
	return string(e)
}

func (e ErrorSlug) Error() error {
	return errors.New(e.String())
}

func ErrToSlug(e error) ErrorSlug {
	return ErrorSlug(e.Error())
}

func ErrList() map[ErrorSlug]ErrorInfo {
	info := make(map[ErrorSlug]ErrorInfo)
	info[ErrSlugInternalError] = ErrorInfo{
		HttpCode: fasthttp.StatusInternalServerError,
		Message:  ErrText(ErrSlugInternalError),
	}
	info[ErrSlugEmptyQuery] = ErrorInfo{
		HttpCode: fasthttp.StatusInternalServerError,
		Message:  ErrText(ErrSlugEmptyQuery),
	}
	info[ErrSlugNotFound] = ErrorInfo{
		HttpCode: fasthttp.StatusNotFound,
		Message:  ErrText(ErrSlugNotFound),
	}
	return info
}

func InterceptorOptionHandler(p interface{}) error {
	return errors.New(fmt.Sprintf("%s", p))
}
