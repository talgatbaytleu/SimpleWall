package service

import (
	"strconv"

	"liker/internal/dal"
)

type LikeServiceInterface interface {
	ToLike(post_idStr string, user_idStr string) error
	ToUnlike(post_idStr string, user_idStr string) error
	GetLikeCount(post_idStr string) (*string, error)
}

type likeService struct {
	likeDal dal.LikeDalInterface
}

func NewLikeService(likeDal dal.LikeDalInterface) *likeService {
	return &likeService{likeDal: likeDal}
}

func (s *likeService) ToLike(post_idStr string, user_idStr string) error {
	post_id, err := strconv.Atoi(post_idStr)
	if err != nil {
		return err
	}

	user_id, err := strconv.Atoi(user_idStr)
	if err != nil {
		return err
	}

	err = s.likeDal.InsertLike(post_id, user_id)
	if err != nil {
		return err
	}

	return nil
}

func (s *likeService) ToUnlike(post_idStr string, user_idStr string) error {
	post_id, err := strconv.Atoi(post_idStr)
	if err != nil {
		return err
	}

	user_id, err := strconv.Atoi(user_idStr)
	if err != nil {
		return err
	}

	err = s.likeDal.DeleteLike(post_id, user_id)
	if err != nil {
		return err
	}

	return nil
}

func (s *likeService) GetLikeCount(post_idStr string) (*string, error) {
	post_id, err := strconv.Atoi(post_idStr)
	if err != nil {
		return nil, err
	}
	return s.likeDal.SelectLikeCount(post_id)
}
