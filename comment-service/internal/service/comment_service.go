package service

import "commenter/internal/dal"

type CommentServiceInterface interface {
	CreateComment()
}

type commentService struct {
	commentDal dal.CommentDalInterface
}

func NewCommentService(commentDal dal.CommentDalInterface) *commentService {
	return &commentService{commentDal: commentDal}
}
