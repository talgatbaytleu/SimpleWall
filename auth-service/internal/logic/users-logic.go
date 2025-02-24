package logic

import (
	"encoding/json"
	"io"

	"auth-service/internal/dal"
	"auth-service/models"
	"auth-service/utils"
)

type UserLogicInterface interface {
	CreateUser(body io.Reader) error
	LoginUser(body io.Reader) (string, error)
	CheckToken()
}

type userLogic struct {
	userDal dal.UsersDalInterface
}

func NewUserLogic(userDal dal.UsersDalInterface) *userLogic {
	return &userLogic{userDal: userDal}
}

func (l *userLogic) CreateUser(body io.Reader) error {
	user := models.User{}

	err := json.NewDecoder(body).Decode(user)
	if err != nil {
		return err
	}

	err = l.userDal.InsertUser(user.Username, user.PaswordHash)
	if err != nil {
		return err
	}

	return nil
}

func (l *userLogic) LoginUser(body io.Reader) (string, error) {
	loginingUser := models.User{}
	user := models.User{}

	err := json.NewDecoder(body).Decode(loginingUser)
	if err != nil {
		return "", err
	}

	userJson, err := l.userDal.SelectUser(loginingUser.Username)
	err = json.Unmarshal([]byte(userJson), &user)
	if err != nil {
		return "", err
	}

	if !utils.CheckPassword(loginingUser.Password, user.PaswordHash) {
		return "", nil // error SHOULD BE HANDLED!!!
	}

	jwtToken, err := utils.GenerateJWT(user.ID)

	return jwtToken, err
}
