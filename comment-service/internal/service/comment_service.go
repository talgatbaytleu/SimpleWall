package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"

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
	SendNotification(body io.Reader, user_idStr string) error
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

func (s *commentService) SendNotification(body io.Reader, user_idStr string) error {
	var comment models.CommentType
	kafkaURL := os.Getenv("KAFKA_ADDR")

	post_idStr := strconv.Itoa(comment.CommentID)

	err := json.NewDecoder(body).Decode(&comment)
	if err != nil {
		return err
	}

	kafkaWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{kafkaURL},
		Topic:    "comments-notifications",
		Balancer: &kafka.LeastBytes{},
	})

	defer kafkaWriter.Close()

	message := fmt.Sprintf(
		"User %s commented on post %d: %s",
		user_idStr,
		comment.PostID,
		comment.Content,
	)

	err = kafkaWriter.WriteMessages(context.Background(), kafka.Message{
		Key:   []byte(post_idStr),
		Value: []byte(message),
		Time:  time.Now(),
	})
	if err != nil {
		return err
	}

	return nil
}
