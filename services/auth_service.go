package services

import (
	"errors"
	"sspanel-metron-go/models"
	"sspanel-metron-go/utils"
)

func Authenticate(username, password string) (*models.User, error) {
	user, err := models.GetUserByUsername(username)
	if err != nil {
		return nil, errors.New("用户不存在")
	}

	if !utils.CheckPasswordHash(password, user.PasswordHash) {
		return nil, errors.New("密码错误")
	}

	return user, nil
}
