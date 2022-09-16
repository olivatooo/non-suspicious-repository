package users

import (
	"account/models"

	"github.com/gofrs/uuid"
)

// Receives email and hashed password
func Create(email string, hashedPassword string) error {
	user := models.User{}
	user.Email = email
	user.Secret, _ = uuid.NewV4()
	user.Password = hashedPassword
	err := models.DB.Create(&user)
	return err
}

func UpdateSecret(userID int) (string, error) {
	user, err := FindByID(userID)
	if err != nil {
		return "", err
	}
	user.Secret, _ = uuid.NewV4()
	models.DB.ValidateAndSave(user)
	return user.Secret.String(), nil
}

func FindBySecret(secret string) (*models.User, error) {
	user := models.User{}
	query := models.DB.Where("secret = ?", secret)
	err := query.First(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func FindByID(userID int) (*models.User, error) {
	user := models.User{}
	query := models.DB.Where("id = ?", userID)
	err := query.First(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func FindByEmail(email string) (*models.User, error) {
	user := models.User{}
	query := models.DB.Where("email = ?", email)
	err := query.First(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
