package customerr

import (
	"github.com/pkg/errors"
)

type stackTracer interface {
	StackTrace() errors.StackTrace
}

type Error struct {
	err error
}

func (e Error) Error() string {
	return e.err.Error()
}

func (e Error) StackTrace() errors.StackTrace {
	if e.err == nil {
		return nil
	}
	st, ok := e.err.(stackTracer)
	if !ok {
		return nil
	}
	return st.StackTrace()
}

func StackTrace(err error) error {
	if err == nil {
		return err
	}

	_, ok := err.(stackTracer)
	if !ok {
		err = Error{
			errors.WithStack(err),
		}
	}

	return err
}
