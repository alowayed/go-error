# Golang error handling proposal and proof of concept

This repo seeks to build upon the existing [proposal](https://docs.google.com/document/d/18wX3UAzpL1SDAdwx0wze6mtH54Zkq4zhhqRMaop8wrE/edit) and [PoC](https://github.com/ugizashinje/poc) by Tihomir and Shiggy by presenting alternatives to some of the design ... TODO

IThe original proposal and PoC by Tihomir and Shiggy addresses all current painpoints simply and consicely. But we believe some areas can be improved. We'll first outline these areas then provide what we believe to be an enhancment over the existing design. 

## Areas of Improvement

**Response status set on error creation**
The method creating an error is not aware of the context of the request and should not choose status is. For example, an authentication request that fails due to the user not being found in the database should not return a 404 status. The resource layer should get the cause or type of the original error and return a more appropriate response. 

We propose using constant categories, in this PoC they are strings, to specify the type of the original error. The resource layer can use the types to return the proper status.

**Error types are too specific**
Using error creators such as `SUBSCRIBER_DOES_NOT_EXSIS` is too specific in our opinion. With over 50 services and repos and at 10 error types per repo, we would have 100 error types for future developers to choose from. We propose more general types such as `ErrNotFound` that accepts the original error as an argument. This way, the original error is retained, and there are fewer types for parents methods to check on return.

**The use of Context should not be a part of this proposal**
The PoC and proposal present possible implementations of Context use in our backend. Although it is great to see, it does not impact the current or proposed error handling architecture. We have removed any Context use from this PoC to focus on presenting a toy example of the backend error handling changes.

## Kudos to Shiggy and Tihi, and Next Steps
We're looking forward to moving collecting feedback and moving forward with rolling this out in the backend. The added information from the stacktrace, and the ability to reduce error handling code will decrease debugging time, reduce code clutter, and speed up development. I'm personally glad we won't need to resort to Ctrl + F the repo to find the path taken for a bug to happen!

