package errmsg

import "errors"

var (
	CardAlreadyExists          = errors.New("card already exists for user")
	CardNotFound               = errors.New("card not found")
	CardInvalidStateTransition = errors.New("invalid state transition")
	CardMissingFieldInBody     = errors.New("missing field(s) in body")
)
