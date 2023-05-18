package errno

import (
	"errors"
	"fmt"
)

type ErrNo struct {
	Code ErrNoCode
	Msg  string
}

func New(code ErrNoCode, msg string) *ErrNo {
	return &ErrNo{
		Code: code,
		Msg:  msg,
	}
}

func (e *ErrNo) Error() string {
	return fmt.Sprintf("error code: %d, error msg: %s\n", e.Code, e.Msg)
}

func Convert(err error) *ErrNo {
	errno := &ErrNo{}
	if errors.As(err, &errno) {
		return errno
	}

	s := ServiceError
	s.Msg = err.Error()
	return s
}
