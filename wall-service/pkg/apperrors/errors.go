package apperrors

import (
	"errors"
)

var (
	ErrNotFound = errors.New("NOT found")
	ErrNoPostId = errors.New("wal-service: handler: post_id is required")
)
