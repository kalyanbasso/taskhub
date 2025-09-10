package controller

import (
	"errors"
)

type ErrorResponse struct {
	Message string `json:"message,omitempty"`
}

var (
	ErrInvalidTaskID = errors.New("invalid task ID")
)
