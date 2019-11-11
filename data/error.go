package data

import (
	"database/sql"

	"github.com/alowayed/go-error/errors"
)

// Replace sql.Err with Error equivalent
func categorizeDBError(err error) errors.Error {

	superErr := errors.New(err, errors.CategoryOther)

	switch err {
	case sql.ErrNoRows:
		superErr = errors.New(err, errors.CategoryNotFound)
	case sql.ErrConnDone:
		superErr = errors.New(err, errors.CategoryDBConnDone)
	case sql.ErrTxDone:
		superErr = errors.New(err, errors.CategoryDBTxDone)
	}

	return superErr
}
