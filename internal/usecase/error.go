package usecase

import (
	"errors"
)

var (
	ErrInvalidPriority = errors.New("invalid priority value")
	ErrInvalidTaskID   = errors.New("invalid task ID")
)
