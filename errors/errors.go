package errors

import "errors"

type errorType int

const (
	errNotFound errorType = iota + 1
	errConflict
	errConnection
	errAlreadyExists
	errGeneric
	errWatch
	errUnauthorized
	errForbidden
)

var _ error = Error{}

type Error struct {
	errType errorType
	cause   error
}

func (c Error) Error() string {
	return c.cause.Error()
}

func NewGenericError(err error) Error {
	return Error{
		errType: errGeneric,
		cause:   err,
	}
}

func NewNotFoundError(err error) Error {
	return Error{
		errType: errNotFound,
		cause:   err,
	}
}

func NewConflictError(err error) Error {
	return Error{
		errType: errConflict,
		cause:   err,
	}
}

func NewConnectionError(err error) Error {
	return Error{
		errType: errConnection,
		cause:   err,
	}
}

func NewAlreadyExistsError(err error) Error {
	return Error{
		errType: errAlreadyExists,
		cause:   err,
	}
}

func NewWatchError(err error) Error {
	return Error{
		errType: errWatch,
		cause:   err,
	}
}

func NewUnauthorizedError(err error) Error {
	return Error{
		errType: errUnauthorized,
		cause:   err,
	}
}

func NewForbiddenError(err error) Error {
	return Error{
		errType: errForbidden,
		cause:   err,
	}
}

func isError(err error, t errorType) bool {
	var e Error
	if ok := errors.As(err, &e); !ok {
		return false
	}

	if e.errType == t {
		return true
	}

	return false
}

func IsGenericError(err error) bool {
	return isError(err, errGeneric)
}

func IsNotFoundError(err error) bool {
	return isError(err, errNotFound)
}

func IsConflictError(err error) bool {
	return isError(err, errConflict)
}

func IsConnectionError(err error) bool {
	return isError(err, errConnection)
}

func IsAlreadyExistsError(err error) bool {
	return isError(err, errAlreadyExists)
}

func IsWatchError(err error) bool {
	return isError(err, errWatch)
}

func IsUnauthorizedError(err error) bool {
	return isError(err, errUnauthorized)
}

func IsForbiddenError(err error) bool {
	return isError(err, errForbidden)
}
