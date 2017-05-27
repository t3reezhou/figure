package errors

import (
	"fmt"
	"net/http"

	"github.com/juju/errors"
)

type Error struct {
	errors.Err `json:"_"`
	Code       int    `json:"_"`
	OutPut     string `json:"_"`
	Result     string `json:"result"`
	Message    string `json:"msg"`
}

func NewError(result string, code int) *Error {
	outPut := fmt.Sprintf("{\"Result\":\"%s\", \"msg\":\"%s\"}", result, result)
	err := &Error{errors.NewErr(outPut), code, outPut, result, result}
	err.SetLocation(2)
	return err
}

func (e *Error) Error() string {
	return e.OutPut
}

func (e *Error) MSG(message string) {
	e.Message = message
}

// err = errors.Annotatef(err, "context")
func Annotater(other error, message string) *Error {
	if e, ok := other.(*Error); ok {
		new := errors.Annotate(e, message).(*errors.Err)
		err := &Error{*new, e.Code, e.OutPut, e.Result, e.Result}
		err.SetLocation(1)
		return err
	}
	return nil
}

var (
	ErrFigureNotExist  = NewError("figureNotExist", http.StatusNotFound)
	ErrInvalidArgument = NewError("InvalidArgument", http.StatusForbidden)
)
