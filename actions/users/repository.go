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
