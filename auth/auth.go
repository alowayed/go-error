package auth

import (
	"github.com/alowayed/go-error/data"
	"github.com/alowayed/go-error/errors"
)

func authenticate(user *data.User) (*string, errors.Error) {

	token, err := oktaAuthenticate(user)
	if err != nil {
		// return nil, errors.Err(err)
	}

	return nil, nil
}

func oktaAuthenticate(user *data.User) (*string, error) {
	return nil, errors.Err("okta error")
}
