package error

import (
	"errors"
)

var (
	ErrInvalidParam   = errors.New("invalid parameters")
	ErrNotifierClosed = errors.New("notifier was closed")
	ErrNotFound       = errors.New("not found")

	ErrOutOfRange     = errors.New("index out of rang")
	ErrActorNotFound  = errors.New("cann't found actor")
	ErrMethodNotFound = errors.New("cann't found method in actor")
)
