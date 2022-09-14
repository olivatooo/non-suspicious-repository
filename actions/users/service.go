package users

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPasswordAndCreateUser(password string, email string) error {
	hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return err
	}
	hashedPassword := string(hashedPasswordBytes)
	err = Create(email, hashedPassword)
	return err
}
