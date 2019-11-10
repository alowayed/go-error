package errors

import (
	"fmt"
	"strings"
)

type (
	Error struct {
		err      error
		Notify   bool
		Text     string   `json:"error"`
		Stack    []string `json:"-"`
		Category Category
	}

	Category string

	SuperError interface {
		Error() string
		String() string
		WithInfo(fmt string, args ...interface{}) SuperError
	}
)

func (e *Error) Error() string {
	return e.String()
}

func (e *Error) String() string {

	prefix := "[WARN]"
	if e.Notify {
		prefix = "[ERROR]"
	}

	return fmt.Sprintf(prefix+" %v \n%v", e.Text, strings.Join(e.Stack, ""))
}

func (e *Error) WithInfo(format string, args ...interface{}) SuperError {
	format = e.Text + ": " + format
	e.Text = fmt.Sprintf(format, args...)
	return e
}
