package users

import (
	"account/models"

	"golang.org/x/crypto/bcrypt"
)

func hashPasswordAndCreateUser(email string, password string) error {
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}
	hashedPassword := string(hashedPasswordBytes)
	err = Create(email, hashedPassword)
	return err
}

func getUserByEmail(email string) (*models.User, error) {
	user, err := FindByEmail(email)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func checkPassword(hashedPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if err != nil {
		return err
	}

	return nil
}
