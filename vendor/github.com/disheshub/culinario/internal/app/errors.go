package app

import (
	"github.com/pkg/errors"
)

// Errors
var (
	ErrNotUser      = errors.New("User not found.")
	ErrPassNotMatch = errors.New("User password did not match")
)
