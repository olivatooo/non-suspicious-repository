package users

import (
	"account/actions"
)

func validateCredentials(credentials *Credentials) error {
	// Validate Password
	_, err := actions.IsString(credentials.Password, 60)
	if err != nil {
		return err
	}

	// Validate Email
	_, err = actions.IsEmail(credentials.Email)

	if err != nil {
		return err
	}

	return nil
}

func validateSecret(secret string) error {
	_, err := actions.IsString(secret, 36)
	if err != nil {
		return err
	}
	return nil
}
