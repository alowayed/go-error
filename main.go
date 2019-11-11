package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/alowayed/go-error/data"
	"github.com/alowayed/go-error/errors"
	"github.com/alowayed/go-error/mock"
	"github.com/alowayed/go-error/service"
)

func main() {
	if err := run(); err != nil {
		log.Fatalf("%+v\n", err)
	}
}

func run() errors.Error {

	var err errors.Error

	fmt.Printf("\n# Repo\n\n")
	err = repoErrExample()
	fmt.Printf("%v\n", err)

	fmt.Printf("\n# Service\n\n")
	err = serviceErrExample()
	fmt.Printf("%v\n", err)

	fmt.Printf("\n# Resource\n\n")
	err = resourceErrExample()
	fmt.Printf("%v\n", err)

	return nil
}

// Because of the stacktrace, repository layer errors can simply be
// passed up the chain without wrapping
func repoErrExample() errors.Error {

	userRepository := data.NewUserRepository()

	userID := int64(7)
	_, err := userRepository.Find(userID)
	if err != nil {
		return err
	}

	return nil
}

func serviceErrExample() errors.Error {

	subscriptionService := service.NewSubscriptionService()

	subscriptionID := int64(7)
	_, err := subscriptionService.Renew(subscriptionID)
	if err != nil {
		return err
	}

	return nil
}

func resourceErrExample() errors.Error {

	c := &mock.Context{}

	userRepository := data.NewUserRepository()

	userID := int64(7)
	user, err := userRepository.Find(userID)
	if err != nil {

		// Resource layer can check error category and respond with appropriate HTTP status
		switch err.Category() {
		case errors.CategoryNotFound:
			c.JSON(http.StatusNotFound, err.JsonResponse())
		case errors.CategoryDBConnDone:
			c.JSON(http.StatusInternalServerError, err.JsonResponse())
		case errors.CategoryDBTxDone:
			c.JSON(http.StatusInternalServerError, err.JsonResponse())
		}

		return err
	}

	c.JSON(http.StatusOK, user)
	return nil
}
