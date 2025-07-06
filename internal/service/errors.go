package service

import "github.com/go-errors/errors"

var (
	ErrValidation         = errors.New("validation error")
	ErrNotFound           = errors.New("not found")
	ErrAlreadyDeleted     = errors.New("already deleted")
	ErrNotDeleted         = errors.New("not deleted")
	ErrNotAuthorized      = errors.New("not authorized")
	ErrUniqueKeyViolation = errors.New("unique key violation")
	ErrConflict           = errors.New("conflict error")
)
