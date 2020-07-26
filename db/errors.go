package db

import "errors"

var (
	ErrInvalidDbConnection = errors.New("invalid db connection")
	ErrNotSupported        = errors.New("not supported")
	ErrInvalidData         = errors.New("invalid input data")
	ErrAlreadyExist        = errors.New("already exist")
)
