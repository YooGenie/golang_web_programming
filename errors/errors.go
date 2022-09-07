package errors

import (
	"fmt"
)

type ServiceError struct {
	err error
	msg string
	typ string
}

func NewServiceError(err error, msg, typ string) *ServiceError {
	return &ServiceError{
		err: err,
		msg:msg,
		typ:typ,
	}
}

func (s *ServiceError) Error() string {
	return fmt.Sprintf("[error] %s: %s", s.typ,s.msg)
}
