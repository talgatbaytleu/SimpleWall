package apperrors

import (
	"errors"
)

var (
	ErrNotFound       = errors.New("NOT found")
	ErrLikeNotDeleted = errors.New("like-service: dal: like NOT deleted")
)
