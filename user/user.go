package user

import (
	"errors"
	gp "github.com/abaft/goPassword"
	"strings"
)

func AuthAttempt(username, password string) (User, error) {
	user, err := getUser(username)
	if err != nil {
		return User{}, err
	}

	if gp.PasswordCheck(password, user.Password) {
		return user, nil
	} else {
		return User{}, errors.New("Incorrect Password")
	}
}

func CreateUser(username, password, email string) (User, error) {
	rawUser := User{
		Username: username,
		Password: gp.HashPassword(password),
		Email:    email,
	}

	if len(rawUser.Username) > 64 {
		return User{}, errors.New("Username too long")
	} else if len(rawUser.Email) > 128 {
		return User{}, errors.New("Email too long")
	}

	err := insertUser(rawUser)
	if err != nil {
		if strings.HasPrefix(err.Error(), "Error 1062") {
			msg := strings.Split(err.Error(), "'")
			return User{}, errors.New(msg[3] + " already taken")
		} else {
			return User{}, err
		}
	}
	return rawUser, nil
}

func DeleteUser(user User) error {

	err := deleteUser(user)
	if err != nil {
		return err
	}
	return nil
}
