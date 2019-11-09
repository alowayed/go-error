package main

import (
	"fmt"
	"runtime"
)

type (
	Error struct {
		err      error
		notify   bool
		Text     string   `json:"error"`
		Stack    []string `json:"-"`
		category Category
	}

	Category string
)

const (
	CategoryNotFound      Category = "not found"
	CategoryDB            Category = "DB error"
	CategoryUnauthorized  Category = "unauthorized error"
	CategoryAlreadyExists Category = "already exists error"
)

func create() *Error {
	pcs := make([]uintptr, 64)
	height := runtime.Callers(2, pcs)
	frames := runtime.CallersFrames(pcs)
	stack := []string{}
	for frame, more := frames.Next(); more && height > 0; frame, more = frames.Next() {
		height--
		stack = append(stack, fmt.Sprintf("%s %s:%d\n", frame.Function, frame.File, frame.Line))
	}

	return &Error{Stack: stack,
		notify: false,
	}
}

func (e *Error) Error() string {
	return fmt.Sprintf("[WARN] %v\n%v", e.Text, e.Stack)
}

func (e *Error) WithInfo(format string, args ...interface{}) SuperError {
	format = e.Text + ": " + format
	e.Text = fmt.Sprintf(format, args...)
	return e
}

func (e *Error) Is(category Category) bool {
	return e.category == category
}

func (e *Error) GetCategory() Category {
	return e.category
}

func (e *Error) Notify() SuperError {
	e.notify = true
	return e
}

func (e *Error) String() string {
	return fmt.Sprintf("[WARN] %v %v", e.Text, e.Stack)
}

type SuperError interface {
	Error() string
	WithInfo(fmt string, args ...interface{}) SuperError
	Is(category Category) bool
	GetCategory() Category
	Notify() SuperError
	String() string
}

func newError(err error, text string, category Category) SuperError {
	instance := create()
	instance.err = err
	instance.Text = text
	instance.category = category
	return instance
}

// ----- Errors

func Err(err error) SuperError {
	return newError(err, "unkown error", "")
}

func ErrNotFound(err error) SuperError {
	return newError(err, err.Error(), CategoryNotFound)
}

func ErrDB(err error) SuperError {
	return newError(err, err.Error(), CategoryDB)
}

func ErrUnauthorized(err error) SuperError {
	return newError(err, err.Error(), CategoryUnauthorized)
}

func ErrAlreadyExists(err error) SuperError {
	return newError(err, err.Error(), CategoryAlreadyExists)
}
