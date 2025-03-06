package apperrors

import (
	"errors"
)

var (
	ErrNotFound          = errors.New("NOT found")
	ErrNotAllowed        = errors.New("comment-service: operation not allowed")
	ErrCommentNotDeleted = errors.New("comment-service: dal: comment NOT deleted")
	ErrCommentNotCreated = errors.New("comment-service: dal: comment NOT created")
	ErrCommentNotUpdated = errors.New("comment-service: dal: comment NOT updated")
	ErrNoPostId          = errors.New("comment-service: handler: post_id is required")
)
