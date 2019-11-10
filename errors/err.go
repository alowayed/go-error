package errors

import (
	"fmt"
	"runtime"
)

const (
	CategoryOther Category = "other"

	// Database errors
	CategoryDBConnDone Category = "DB connection done"
	CategoryDBTxDone   Category = "DB transaction done"

	CategoryChargeInsufficientFunds Category = "insufficient funds"
	CategoryChargeCardExpired       Category = "card expired"

	CategoryUnauthorized  Category = "unauthorized error"
	CategoryNotFound      Category = "not found"
	CategoryAlreadyExists Category = "already exists error"
)

// Capture a stacktrace and
func New(err error, category Category) SuperError {

	// Capture stack
	stack := []string{}
	pcs := make([]uintptr, 64)
	height := runtime.Callers(2, pcs)
	frames := runtime.CallersFrames(pcs)
	for frame, more := frames.Next(); more && height > 0; frame, more = frames.Next() {
		height--
		stack = append(stack, fmt.Sprintf("%s %s:%d\n", frame.Function, frame.File, frame.Line))
	}

	return &Error{
		err:      err,
		Text:     err.Error(),
		Stack:    stack,
		Category: category,
	}
}
