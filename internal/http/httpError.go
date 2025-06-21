package http

import "net/http"

type HTTPError struct {
	Code    int
	Message string
	Err     error
}

func (e *HTTPError) Error() string {
	return e.Message
}

func (e *HTTPError) Unwrap() error {
	return e.Err
}

func NewHTTPError(code int, msg string, err error) *HTTPError {
	return &HTTPError{Code: code, Message: msg, Err: err}
}

var (
	ErrBadRequest = func(msg string, err error) *HTTPError {
		return NewHTTPError(http.StatusBadRequest, msg, err)
	}
	ErrInternal = func(msg string, err error) *HTTPError {
		return NewHTTPError(http.StatusInternalServerError, msg, err)
	}
)
