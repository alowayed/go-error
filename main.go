package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("%+v\n", err)
	}
}

func run() SuperError {

	return repoErrExample()

	return serviceErrExample()

	return resourceErrExample()

	return nil
}

func repoErrExample() SuperError {

	// Error from repo layer. Send up the chain
	userID := 7
	_, err := userRepoGetUser(userID)
	if err != nil {
		return err
	}

	return nil
}

func serviceErrExample() SuperError {

	email := "john@example.com"
	createSubscriber(email)

	return nil
}

func resourceErrExample() SuperError {

	// Error from 3rd party service
	user := &User{}
	_, err := authenticate(user)
	if err != nil {
		return err
	}

	// How resource layer handlers can check errors causes, and set status
	switch err.GetCategory() {
	case CategoryUnauthorized: // Wrong password
		fmt.Print(http.StatusNotFound) // This is where we set the http status
		return err
	case CategoryNotFound: // User not in system
		fmt.Print(http.StatusNotFound) // This is where we set the http status
		return err
	}

	return err
}
