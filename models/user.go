package models

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
)

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name" binding:"required`
	Email string `json:"email" binding:"required`
	Phone int    `json:"phone" binding:"required`
}

func NewUser(logger *logrus.Logger, name string, email string, phone int) (*User, error) {
	fmt.Println("name ", name)
	if len(name) == 0 {
		err := errors.New("name is required")
		logger.Error(err)
		return nil, err
	}

	// convert string password to hash
	user := &User{
		Name:  name,
		Email: email,
		Phone: phone,
	}
	return user, nil
}
