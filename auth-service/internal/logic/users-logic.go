package logic

import (
	"encoding/json"
	"fmt"
	"io"

	"auth-service/internal/dal"
	"auth-service/pkg/apperrors"
	"auth-service/pkg/models"
	"auth-service/pkg/utils"
)

type UserLogicInterface interface {
	CreateUser(body io.Reader) error
	LoginUser(body io.Reader) (string, error)
	CheckToken(token string) (string, error)
}

type userLogic struct {
	userDal dal.UsersDalInterface
}

func NewUserLogic(userDal dal.UsersDalInterface) *userLogic {
	return &userLogic{userDal: userDal}
}

func (l *userLogic) CreateUser(body io.Reader) error {
	user := models.User{}

	err := json.NewDecoder(body).Decode(&user)
	if err != nil {
		return err
	}

	// Validate username
	err = utils.ValidateUsername(user.Username)
	if err != nil {
		return err
	}

	// Validate password
	err = utils.ValidatePassword(user.Password)
	if err != nil {
		return err
	}

	user.PasswordHash, err = utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	err = l.userDal.InsertUser(user.Username, user.PasswordHash)
	if err != nil {
		return err
	}

	return nil
}

func (l *userLogic) LoginUser(body io.Reader) (string, error) {
	loginingUser := models.User{}
	user := models.User{}

	err := json.NewDecoder(body).Decode(&loginingUser)
	if err != nil {
		return "", err
	}

	userJson, err := l.userDal.SelectUser(loginingUser.Username)
	err = json.Unmarshal([]byte(userJson), &user)
	if err != nil {
		return "", err
	}

	if !utils.CheckPassword(loginingUser.Password, user.PasswordHash) {
		return "", apperrors.ErrIncorrectPswd // error SHOULD BE HANDLED!!!
	}

	jwtToken, err := utils.GenerateJWT(user.ID)

	return jwtToken, err
}

func (l *userLogic) CheckToken(token string) (string, error) {
	user_id, err := utils.ValidateJWT(token)
	if err != nil {
		return "", err
	}

	fmt.Println("user_id: ", user_id)

	err = l.userDal.CheckUser(user_id)
	if err != nil {
		return "", err
	}

	return user_id, nil
}
