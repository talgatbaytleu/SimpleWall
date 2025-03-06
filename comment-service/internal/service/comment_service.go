package service

import (
	"encoding/json"
	"io"
	"strconv"

	"commenter/internal/dal"
	"commenter/pkg/apperrors"
	"commenter/pkg/models"
)

type CommentServiceInterface interface {
	CreateComment(body io.Reader, user_idStr string) error
	UpdateComment(body io.Reader, comment_idStr string, user_idStr string) error
	GetCommentById(comment_idStr string) (string, error)
	GetPostComments(post_idStr string) (string, error)
	DeleteComment(comment_idStr string, user_idStr string) error
}

type commentService struct {
	commentDal dal.CommentDalInterface
}

func NewCommentService(commentDal dal.CommentDalInterface) *commentService {
	return &commentService{commentDal: commentDal}
}

func (s *commentService) CreateComment(body io.Reader, user_idStr string) error {
	user_id, err := strconv.Atoi(user_idStr)
	if err != nil {
		return err
	}

	var comment models.CommentType

	err = json.NewDecoder(body).Decode(&comment)
	if err != nil {
		return err
	}

	return s.commentDal.InsertComment(comment.PostID, user_id, comment.Content)
}

func (s *commentService) UpdateComment(
	body io.Reader,
	comment_idStr string,
	user_idStr string,
) error {
	comment_id, err := strconv.Atoi(comment_idStr)
	if err != nil {
		return err
	}

	user_id, err := strconv.Atoi(user_idStr)
	if err != nil {
		return err
	}

	user_idOfComment, err := s.commentDal.SelectUserOfComment(comment_id)
	if err != nil {
		return err
	}

	if user_id != user_idOfComment {
		return apperrors.ErrNotAllowed
	}

	var comment models.CommentType

	err = json.NewDecoder(body).Decode(&comment)
	if err != nil {
		return err
	}

	return s.commentDal.UpdateComment(comment_id, comment.Content)
}

func (s *commentService) GetCommentById(comment_idStr string) (string, error) {
	comment_id, err := strconv.Atoi(comment_idStr)
	if err != nil {
		return "", err
	}

	return s.commentDal.SelectComment(comment_id)
}

func (s *commentService) GetPostComments(post_idStr string) (string, error) {
	post_id, err := strconv.Atoi(post_idStr)
	if err != nil {
		return "", err
	}

	return s.commentDal.SelectPostComments(post_id)
}

func (s *commentService) DeleteComment(comment_idStr string, user_idStr string) error {
	comment_id, err := strconv.Atoi(comment_idStr)
	if err != nil {
		return err
	}

	user_id, err := strconv.Atoi(user_idStr)
	if err != nil {
		return err
	}

	user_idOfComment, err := s.commentDal.SelectUserOfComment(comment_id)
	if err != nil {
		return err
	}

	if user_id != user_idOfComment {
		return apperrors.ErrNotAllowed
	}

	return s.commentDal.DeleteComment(comment_id, user_id)
}
