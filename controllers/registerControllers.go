package controllers

import (
	"github.com/raihaninfo/activita/models"
)

type RegisterError struct {
	Username string
	Email    string
	Password string
}

func RegisterMessage(username, email, password, confirmPassword string) RegisterError {
	var msg RegisterError
	if password != confirmPassword {
		msg.Password = "Password don't match"
	}

	if password == "" {
		msg.Password = "Password is required"
	}

	userName, _ := models.UserByUsername(username)

	if username == userName.Username {
		msg.Username = "username already exist"
	}
	if username == "" {
		msg.Username = "username is required"
	}

	user, _ := models.UserByEmail(email)
	// if err1 != nil {
	// 	log.Println(err1)
	// }

	if email == user.Email {
		msg.Email = "Email already exist"
	}
	if email == "" {
		msg.Email = "Email is required"
	}

	return msg
}
