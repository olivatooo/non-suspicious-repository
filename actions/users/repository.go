package users

import (
	"account/models"
)

// Receives email and hashed password
func Create(email string, hashedPassword string) error {
	user := models.User{}
	user.Email = email
	user.Password = hashedPassword
	err := models.DB.Create(&user)
	return err
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
