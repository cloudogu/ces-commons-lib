package repository

import "errors"

//goland:noinspection GoUnusedGlobalVariable
var ErrUnauthorized = errors.New("401 unauthorized, please login to proceed")

//goland:noinspection GoUnusedGlobalVariable
var ErrForbidden = errors.New("403 forbidden, not enough privileges")

//goland:noinspection GoUnusedGlobalVariable
var ErrNotFound = errors.New("404 not found")
