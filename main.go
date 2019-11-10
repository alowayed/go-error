package main

import (
	"fmt"
	"log"

	"github.com/alowayed/go-error/service"

	"github.com/alowayed/go-error/data"

	"github.com/alowayed/go-error/errors"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("%+v\n", err)
	}
}

func run() errors.SuperError {

	var err errors.SuperError

	err = repoErrExample()
	fmt.Printf("\n# Repo\n\n%v\n", err)

	err = serviceErrExample()
	fmt.Printf("\n# Service\n\n%v\n", err)

	err = resourceErrExample()
	fmt.Printf("\n# Resource\n\n%v\n", err)

	return nil
}

// Because of the stacktrace, repository layer errors can simply be
// passed up the chain without wrapping
func repoErrExample() errors.SuperError {

	userRepository := data.NewUserRepository()

	userID := int64(7)
	_, err := userRepository.Find(userID)
	if err != nil {
		return err
	}

	return nil
}

func serviceErrExample() errors.SuperError {

	subscriptionService := service.NewSubscriptionService()

	subscriptionID := int64(7)
	_, err := subscriptionService.Renew(subscriptionID)
	if err != nil {
		return err
	}

	return nil
}

func resourceErrExample() errors.SuperError {

	// Error from 3rd party service
	// user := &data.User{}
	// _, err := auth.authenticate(user)
	// if err != nil {
	// 	return err
	// }

	// How resource layer handlers can check errors causes, and set status
	// switch err.Ca {
	// case CategoryUnauthorized: // Wrong password
	// 	fmt.Print(http.StatusNotFound) // This is where we set the http status
	// 	return err
	// case CategoryNotFound: // User not in system
	// 	fmt.Print(http.StatusNotFound) // This is where we set the http status
	// 	return err
	// }

	return nil
}
