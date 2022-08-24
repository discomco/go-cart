package errors

import "github.com/pkg/errors"

var (
	ErrStreamNotFound = errors.New(StreamNotFound)
)

const (
	StreamNotFound = "stream not found"
)
