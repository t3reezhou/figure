package errors

import (
	"fmt"
	"net/http"
)

type Error struct {
	Result  string        `json:"result"`
	Message string        `json:"msg"`
	Code    int           `json:"-"`
	Args    []interface{} `json:"-"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("{Result:%s, Message:%s}", e.Result, e.Message)
}

func NewError(result string, code int) *Error {
	err := new(Error)
	err.Result = result
	err.Message = result
	err.Code = code
	return err
}

var (
	ErrFigureNotExist  = NewError("figureNotExist", http.StatusNotFound)
	ErrInvalidArgument = NewError("InvalidArgument", http.StatusForbidden)
)
