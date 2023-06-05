package model

import (
	"errors"
	"fmt"
	"ginbasic/helper"
	"html"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	//uid      uint   `gorm:"size:100;not null;primaryKey"`
	Email    string `gorm:"unique"`
	Uname    string
	PassHash string
}

func (user *User) Save() (*User, error) {
	err := helper.Database.Create(&user).Error
	if err != nil {
		fmt.Println("ennamo thappu nadakuthu shivaji")
		return &User{}, err
	}
	return user, nil
}

func (user *User) BeforeSave(*gorm.DB) error {

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.PassHash), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.PassHash = string(passwordHash)
	user.Email = html.EscapeString(strings.TrimSpace(user.Email))
	return nil
}

func (user *User) Validate(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(user.PassHash), []byte(password))
}

func FindUserByEmail(email string) (*User, error) {
	var user User
	err := helper.Database.Where("email = ?", email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
