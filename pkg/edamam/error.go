package edamam

import (
	"bytes"
	"fmt"
	"github.com/artback/edamamWrapper/internal/edamam"
	"net/http"
)

type Error struct {
	Code    string
	Message string
	Op      string
	Err     error
}

func (e *Error) Error() string {
	var buf bytes.Buffer
	if e.Code != "" {
		_, _ = fmt.Fprintf(&buf, "%s ", e.Code)
	}
	if e.Op != "" {
		_, _ = fmt.Fprintf(&buf, "%s: ", e.Op)
	}
	if e.Err != nil {
		buf.WriteString(e.Err.Error())
	} else {
		buf.WriteString(e.Message)
	}
	return buf.String()
}

type HttpError struct {
	http.Response
	Err error
}

func (h HttpError) Error() string {
	switch h.StatusCode {
	case http.StatusNotFound:
		return NotFoundError{Message: h.Request.URL.Path}.Error()
	case http.StatusBadRequest:
		return BadRequestError{Err: h.Err}.Error()
	case http.StatusNotImplemented:
		return NotImplementedError{Err: h.Err}.Error()
	case http.StatusInternalServerError:
		return InternalError{Message: "internal problems", Err: h.Err}.Error()
	case http.StatusUnauthorized:
		return UnauthorizedError{Err: h.Err}.Error()
	}
	return InternalError{Err: h.Err}.Error()
}

type NotImplementedError struct {
	Op  string
	Err error
}

func (err NotImplementedError) Error() string {
	e := Error{Code: edamam.ENOTFOUND, Message: "not implemented", Op: err.Op, Err: err.Err}
	return e.Error()
}

type NotFoundError struct {
	Op      string
	Err     error
	Message string
}

func (err NotFoundError) Error() string {
	e := Error{Code: edamam.ENOTFOUND, Message: err.Message, Op: err.Op, Err: err.Err}
	return e.Error()
}

type BadRequestError struct {
	Op  string
	Err error
}

func (err BadRequestError) Error() string {
	e := Error{Code: edamam.ECONFLICT, Message: "bad request", Op: err.Op, Err: err.Err}
	return e.Error()
}

type UnauthorizedError struct {
	Op  string
	Err error
}

func (err UnauthorizedError) Error() string {
	e := Error{Code: edamam.EINVALID, Message: "unauthorized", Op: err.Op, Err: err.Err}
	return e.Error()
}

type InternalError struct {
	Message string
	Op      string
	Err     error
}

func (err InternalError) Error() string {
	e := Error{Code: edamam.EINTERNAL, Message: err.Message, Op: err.Op, Err: err.Err}
	return e.Error()
}
