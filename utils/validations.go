package utils

import (
	"errors"
	"regexp"

	"github.com/Similadayo/models"
)

func Validate(user *models.User) error {
	if user.FirstName == "" {
		return errors.New("firstname is required")
	}
	if user.LastName == "" {
		return errors.New("lastname is required")
	}
	if user.UserName == "" {
		return errors.New("username is required")
	}
	if user.Email == "" {
		return errors.New("email is required")
	}
	if user.Password == "" {
		return errors.New("password is required")
	}
	if user.Address == "" {
		return errors.New("address is required")
	}
	if user.Phone == "" {
		return errors.New("phone is required")
	}
	return nil
}

// Checking if email is valid using regexp
func IsEmailValid(email string) bool {
	emailRegexp := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return emailRegexp.MatchString(email)
}

func IsPasswordValid(password string) bool {
	return len(password) >= 8
}
