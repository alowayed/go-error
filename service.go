package main

type (
	Subscriber struct{}
)

func createSubscriber(email string) (*Subscriber, SuperError) {

	subscriber := &Subscriber{}

	// TODO

	return subscriber, nil
}
