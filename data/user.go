package data

import (
	"database/sql"

	"github.com/alowayed/go-error/errors"
)

type (
	User struct{}

	UserRepository interface {
		Find(userID int64) (*User, errors.Error)
	}

	SimpleUserRepository struct{}
)

func NewUserRepository() UserRepository {
	return &SimpleUserRepository{}
}

func (*SimpleUserRepository) Find(userID int64) (*User, errors.Error) {

	user, err := sqlSelectUser(userID)
	if err != nil {
		return nil, categorizeDBError(err).WithInfo("user ID = %v", userID)
	}

	return user, nil
}

// Stub function for sql.Select(userID, selectUserByIDSQL)
func sqlSelectUser(userID int64) (*User, error) {
	return nil, sql.ErrNoRows
}
