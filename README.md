# Golang error handling proposal and proof of concept

This repo seeks to build upon the existing [proposal](https://docs.google.com/document/d/18wX3UAzpL1SDAdwx0wze6mtH54Zkq4zhhqRMaop8wrE/edit) and [PoC](https://github.com/ugizashinje/poc) by Tihomir and Shiggy. In it, we support the problems and general architecture of the original proposal, and add some feedback on areas we think can be improved.

To run this PoC, use: 
```go run *.go```

This will print out errors coming from the `data`, `service`, and `resource` layers respectively. Look at `main.go`'s `run()` method to see example of how each of these errors are handled using the new error handling style (`repoErrExample()`, `serviceErrExample()`, `resourceErrExample()`)

## Pain Points
As outlined in the original [proposal](https://docs.google.com/document/d/18wX3UAzpL1SDAdwx0wze6mtH54Zkq4zhhqRMaop8wrE/edit), there are 2 main problems with our current error handling:
1. **Lack of isolation**: `sql.ErrNoRows` is returned from the data layer to the servicer layer, creating a need to import the `sql` package in the service layer. The same applies for 3rd part specific errors (Stripe, Okta, etc) that are returned without wrapping
1. **Too many lines to handle a single error and too little visibility**: errors are currently wrapped with more information. Each wrap requires at least 3 lines of code all the way up the chain. And the only way to debug an errors is searching the sentences in the error.

## Existing Proposal and PoC Solutions
1. **Wrap original errs in Super Errors**: `sql` errors should never leave the data layer, and instead are wrapped in a `SuperError` with the appropriate type. This applies to all 3rd party errors, and is implemented using functions such `SUBSCRIBER_NOT_FOUND_ERROR()` functions in the original proposal, and `ErrNotFound` category in this proposal
1. **Add stacktrace and info on error creation**: the original PoC provides stack traces with all errors created, which means errors can simply be passed up the callstack without loss of information. There is still a `WithInfo(...)` method to add context, but it can be used on the same line with ease.

## Areas of Improvement

1. **Response status set on error creation**
The original PoC attaches an HTTP status code to the error upon its creation, which we may want to change in the resource layer. The method creating an error is not aware of the context of the request and should not choose the HTTP status. For example, an authentication request that fails due to the user not being found in the database should not return a 404 status. The resource layer should get the cause or type of the original error and return a more appropriate response. We propose using constant categories (`ErrNotFound`, `ErrDBConnDone`, `CategoryChargeCardExpired`, etc) to specify the type of the original error. The resource layer can check the type to return the proper status.

1. **Error types are too specific**
Using error creators such as `SUBSCRIBER_DOES_NOT_EXSIS` is too specific in our opinion. With over 50 services and repos and at 10 error types per repo, we would have 100 error types for future developers to choose from. We propose more general types such as `ErrNotFound` that accepts the original error as an argument. This way, the original error is retained, and there are fewer types for parent methods to check on return.

1. **The use of atomic transactions should not be a part of this proposal**
The PoC and proposal present possible implementations of atomic transactions on the resource level through the use of a `wrap(funcHandler...)` method. Although this is good to see, and a great future improvement to our backend, we propose focusing on the error handling aspect alone in this proposal. In doing so, we have removed any atomic transaction code from this PoC, and modeled it after the current state of our backend with only the addition of the proposed error handling changes.

## Next Steps
Kudos to Shiggy and Tihi for pushing out the proposal and PoC! We're looking forward to moving forward by collecting feedback from the greater team, then planning the rollout of this error handling in the backend. 

I'm personally glad we won't need to resort to `Ctrl` + `F` the repo to find the path taken for a bug to happen!

