package main

import (
	"database/sql"
)

type (
	User struct {
	}
)

// import (
// 	"database/sql"
// 	"errors"
// )

// Fake repo call to get a user by id
func userRepoGetUser(userID int) (*User, SuperError) {

	// sql.Select(userID) failed with err of type sql.ErrNoRows
	user, err := selectUser(userID)
	if err != nil {
		return nil, categorizeDBError(err).WithInfo("user ID = %v", userID)
	}

	return user, nil
}

func categorizeDBError(err error) SuperError {

	superErr := Err(err)

	switch err {
	case sql.ErrNoRows:
		superErr = ErrNotFound(err)
	case sql.ErrConnDone:
		superErr = ErrDB(err)
	}

	return superErr

}

// Fake DB call (sql.Select(...))
func selectUser(userID int) (*User, error) {
	return nil, sql.ErrNoRows
}
