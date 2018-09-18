package user

import (
	"errors"
	gp "github.com/abaft/goPassword"
	"strings"
)

func AuthAttempt(username, password string) (User, error) {
	user, err := getUser(username)
	fmt.Println(user)
	if err != nil {
		return User{}, err
	}

	if gp.PasswordCheck(password, user.Password) {
		return user, nil
	} else {
		return User{}, errors.New("Incorrect Password")
	}
}

func CreateUser(username, password, email string) error {
	rawUser := User{
		Username: username,
		Password: gp.HashPassword(password),
		Email:    email,
	}

	if len(rawUser.Username) > 64 {
		return errors.New("Username too long")
	} else if len(rawUser.Email) > 128 {
		return errors.New("Email too long")
	}

	err := insertUser(rawUser)
	if err != nil {
		if strings.HasPrefix(err.Error(), "Error 1062") {
			return errors.New("Username Taken")
		} else {
			return err
		}
	}
	return nil
}
