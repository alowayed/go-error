package main

// import (
// 	"fmt"

// 	"github.com/pkg/errors"
// )

// type (
// 	Error struct {
// 		err    error
// 		notify bool
// 		status int
// 	}

// 	SuperError interface {
// 		WithInfo(format string, a ...interface{}) SuperError
// 		WithStatus(status int) SuperError
// 		Notify() SuperError
// 		String() string
// 		Error() string
// 	}
// )

// func NewError(err error, format string, a ...interface{}) SuperError {
// 	return &Error{
// 		err:    errors.Wrapf(err, format, a...),
// 		notify: false,
// 		status: 0,
// 	}
// }

// func (e *Error) WithInfo(format string, a ...interface{}) SuperError {
// 	e.err = errors.Wrapf(e.err, format, a...)
// 	return e
// }

// func (e *Error) WithStatus(status int) SuperError {
// 	e.status = status
// 	return e
// }

// func (e *Error) Notify() SuperError {
// 	e.notify = true
// 	return e
// }

// func (e *Error) String() string {
// 	return fmt.Sprintf("%+v", e.err)
// }

// func (e *Error) Error() string {
// 	return fmt.Sprintf("%+v", e.err)
// }
