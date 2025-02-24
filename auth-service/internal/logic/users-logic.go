package logic

import (
	"auth-service/internal/dal"
	"io"
)

type UserLogicInterface interface{
  CreateUser (body io.Reader) error
  LoginUser (body io.Reader) (string, error)
  CheckToken ()
}

type userLogic struct {
	userDal dal.UsersDalInterface
}

func NewUserLogic(userDal dal.UsersDalInterface) *userLogic {
	return &userLogic{userDal: userDal}
}

func (*l useruserLogic) 
