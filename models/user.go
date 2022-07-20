package models

import (
	"errors"

	"gorm.io/gorm"
)

type User struct {
	Id       int    `gorm:"primaryKey;autoIncrement:true;unique"`
	Username string `gorm:"not nul, unique"`
	Email    string `gorm:"not nul, unique"`
	Password string
	Status   int `gorm:"nullable"`
}

var ErrNotFound = errors.New("models: resource not found")

func UserById(id int) (*User, error) {
	var user User
	err := DB.Where("id = ?", id).First(&user).Error
	switch err {
	case nil:
		return &user, nil
	case gorm.ErrRecordNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func UserByEmail(email string) (*User, error) {
	var user User
	err := DB.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func UserByUsername(username string) (*User, error) {
	var user User
	err := DB.Where("username = ?", username).Find(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func NewUser(username, email, password string) (*User, error) {
	var user User
	user.Username = username
	user.Email = email
	user.Password = password
	err := DB.Create(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, err
}

// update user's status to 1
func VerifyUser(id int) error {
	var user User
	err := DB.Where("id = ?", id).Find(&user).Error
	if err != nil {
		return err
	}
	user.Status = 1
	err = DB.Save(&user).Error
	return err
}