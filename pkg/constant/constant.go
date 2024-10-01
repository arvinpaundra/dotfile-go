package constant

import "errors"

var (
	ErrKeyUserNotFound = errors.New("key user not found from context")
	ErrInvalidApiKey   = errors.New("invalid api key")
)
