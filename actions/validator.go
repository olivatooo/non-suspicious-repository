package actions

import (
	"errors"
	"net/mail"
	"regexp"
	"strconv"
)

func IsEmail(possibleEmail string) (string, error) {
	address, err := mail.ParseAddress(possibleEmail)
	if err != nil {
		return "", err
	}
	return address.Address, nil
}

func IsString(input string, maxSize int) (string, error) {
	// Error supressed reason: Regex always valid
	re := regexp.MustCompile(`^[a-zA-Z0-9.-]*$`)

	if len(input) > maxSize || len(input) <= 1 {
		return "", errors.New("Invalid String")
	}

	if re.MatchString(input) {
		return input, nil
	}
	return "", errors.New("Invalid String")
}

func IsBigNumber(input string, maxSize int) (string, error) {
	// Error supressed reason: Regex always valid
	re := regexp.MustCompile(`^[0-9.]*$`)

	if len(input) > maxSize || len(input) < 1 {
		return "", errors.New("Invalid BigNumber")
	}

	if re.MatchString(input) {
		return input, nil
	}
	return "", errors.New("Invalid BigNumber")
}

// Get string
func GetInt(input string) (int, error) {
	value, err := strconv.Atoi(input)
	if err != nil {
		return 0, errors.New("Invalid Number")
	}
	return value, err
}
