package service

import (
	goErrors "errors"

	"github.com/alowayed/go-error/errors"
)

type (
	Subscription struct{}

	SubscriptionService interface {
		Renew(subscriptionID int64) (*Subscription, errors.Error)
	}

	SimpleSubscriptionService struct{}
)

func NewSubscriptionService() SubscriptionService {
	return &SimpleSubscriptionService{}
}

func (*SimpleSubscriptionService) Renew(subscriptionID int64) (*Subscription, errors.Error) {

	// Get subscription from repo layer
	// More business logic

	stripeCustomerID := int64(7)
	amount := 100
	err := stripeCharge(stripeCustomerID, amount)
	if err != nil {
		return nil, categorizeStripeError(err).WithInfo("stripe customer [%v] and amount [%v]", stripeCustomerID, amount)
	}

	return nil, nil
}

func stripeCharge(stripeCustomerID int64, amount int) error {
	return goErrors.New("stripe insufficient funds")
}

func categorizeStripeError(err error) errors.Error {

	superErr := errors.New(err, errors.CategoryOther)

	switch err.Error() {
	case "stripe insufficient funds":
		superErr = errors.New(err, errors.CategoryChargeInsufficientFunds)
	case "stripe card expired":
		superErr = errors.New(err, errors.CategoryChargeCardExpired)
	}

	return superErr

}
