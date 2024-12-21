package twitter

import "errors"

var (
	ErrBadCredentials = errors.New("email/password wrong credentials")
	ErrNotFound = errors.New("not founc")
	ErrValidation = errors.New("validation error")
)
