package firebasecli

import (
	"errors"
)

// Version of the firebasecli.
const Version = "0.0.1"

var (
	// ErrFailedToParseArgs is returned when args are not parseable.
	ErrFailedToParseArgs = errors.New("failed to parse args")

	// ErrUnknownCommand is returned when a given command is unknown/undefined.
	ErrUnknownCommand = errors.New("unknown command")

	// ErrFailedToParseHashConfig is returned when hash config file does not follow the format.
	ErrFailedToParseHashConfig = errors.New("failed to parse hash_config")
)
