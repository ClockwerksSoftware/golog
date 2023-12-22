package common

import (
	"errors"
)

var (
	ErrLogLevelUnknown       = errors.New("Unknown Log Level")
	ErrLogLevelAlreadyExists = errors.New("Log Level Already Exists")
)
