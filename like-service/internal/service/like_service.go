package service

import (
	"strconv"

	"liker/internal/dal"
)

type LikeServiceInterface interface {
	ToLike(post_id int, user_idStr string) error
	ToUnlike(post_id int, user_idStr string) error
	GetLikesCount(post_idStr string) (string, error)
	GetLikesList(post_idStr string) (string, error)
}

type likeService struct {
	likeDal dal.LikeDalInterface
}

func NewLikeService(likeDal dal.LikeDalInterface) *likeService {
	return &likeService{likeDal: likeDal}
}

func (s *likeService) ToLike(post_id int, user_idStr string) error {
	// post_id, err := strconv.Atoi(post_idStr)
	// if err != nil {
	// 	return err
	// }

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

func (s *likeService) ToUnlike(post_id int, user_idStr string) error {
	// post_id, err := strconv.Atoi(post_idStr)
	// if err != nil {
	// 	return err
	// }

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

func (s *likeService) GetLikesCount(post_idStr string) (string, error) {
	post_id, err := strconv.Atoi(post_idStr)
	if err != nil {
		return "", err
	}
	return s.likeDal.SelectLikesCount(post_id)
}

func (s *likeService) GetLikesList(post_idStr string) (string, error) {
	post_id, err := strconv.Atoi(post_idStr)
	if err != nil {
		return "", err
	}
	return s.likeDal.SelectLikesList(post_id)
}
