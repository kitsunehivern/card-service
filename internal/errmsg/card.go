package errmsg

import "errors"

var (
	CardAlreadyExists          = errors.New("card already exists for user")
	CardNotFound               = errors.New("card not found")
	UnknownCardStatus          = errors.New("unknown card status")
	CardInvalidStateTransition = errors.New("invalid state transition")
)
