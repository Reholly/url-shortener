package postgres

import "errors"

var (
	NotFoundErr      = errors.New("not found")
	AlreadyExistsErr = errors.New("item already exists")
)
