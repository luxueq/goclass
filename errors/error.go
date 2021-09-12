package errors

import (
	"errors"
	"fmt"
)

type Error struct {
	Code int
	Message string
	Reason string
	Metadata map[string]string
}

func (e *Error) Error() string {
	return fmt.Sprintf("error code:%d;reason:%s;message:%s;metadata:%v", e.Code, e.Reason, e.Message, e.Metadata)
}

func New(code int, reason string, message string) *Error {
	return  &Error{
		Code: code,
		Reason: reason,
		Message: message,
	}
}

func (e *Error) Is(err error) bool {
	if  ori:= new(Error); errors.As(err, &ori) {
		return ori.Reason == e.Reason
	}
	return false
}