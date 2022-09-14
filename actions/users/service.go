package users

import (
	"account/models"
	"encoding/json"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
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

func generateJWT(email string) (string, error) {
	var sampleSecretKey = []byte("SecretYouShouldHide")

	// TODO: Find a better jwt library
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": json.Number(strconv.FormatInt(time.Now().Add(time.Hour*time.Duration(1)).Unix(), 10)),
		"iat": json.Number(strconv.FormatInt(time.Now().Unix(), 10)),
	})

	return token.SignedString(sampleSecretKey)
}
