package main

import "errors"

func authenticate(user *User) (*string, SuperError) {

	token, err := callOkta(user)
	if err != nil {
		return nil, ErrUnauthorized(err)
	}

	return token, nil
}

func callOkta(user *User) (*string, error) {
	return nil, errors.New("okta error")
}
