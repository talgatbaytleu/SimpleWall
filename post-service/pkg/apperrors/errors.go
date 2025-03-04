package apperrors

import (
	"errors"
)

var (
	ErrNotFound     = errors.New("NOT found")
	ErrNotAllowed   = errors.New("post-service: operation not allowed")
	ErrS3NotCreated = errors.New("post-service: bucket/image NOT created on S3 service")
	ErrS3NotDeleted = errors.New("post-service: bucket/image NOT deleted on S3 service")
)
