package errors

import (
	"fmt"
	"strings"
)

type (
	Err struct {
		err      error
		Notify   bool
		Text     string   `json:"error"`
		Stack    []string `json:"-"`
		category Category
	}

	Category string

	Error interface {
		Error() string
		String() string
		WithInfo(fmt string, args ...interface{}) Error
		Category() Category
		JsonResponse() interface{}
	}
)

func (e *Err) Error() string {
	return e.String()
}

func (e *Err) String() string {

	prefix := "[WARN]"
	if e.Notify {
		prefix = "[ERROR]"
	}

	return fmt.Sprintf(prefix+" %v \n%v", e.Text, strings.Join(e.Stack, ""))
}

func (e *Err) WithInfo(format string, args ...interface{}) Error {
	format = e.Text + ": " + format
	e.Text = fmt.Sprintf(format, args...)
	return e
}

func (e *Err) Category() Category {
	return e.category
}

func (e *Err) JsonResponse() interface{} {

	return map[string]interface{}{
		"error":        e.err.Error(),
		"errorMessage": e.Text,
	}

}
