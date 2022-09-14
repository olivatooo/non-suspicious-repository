package users

import (
	"fmt"
	"net/http"

	"account/actions"

	"github.com/gobuffalo/buffalo"
)

func Router(app *buffalo.App) {
	group := app.Group("/users/")
	group.POST("", handleCreation)
	group.POST("auth", handleLogin)
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func handleCreation(c buffalo.Context) error {
	credentials := Credentials{}
	_, err := actions.GetJSONParametersFromBody(&c.Request().Body, &credentials)
	if err != nil {
		return c.Render(http.StatusBadRequest, actions.R.JSON("Invalid email or password"))
	}

	err = validateCredentials(&credentials)

	if err != nil {
		return c.Render(http.StatusBadRequest, actions.R.JSON("Invalid email or password"))
	}

	// Call direct from repository since we don't have any business logic
	err = hashPasswordAndCreateUser(credentials.Password, credentials.Email)
	if err != nil {
		return c.Render(http.StatusInternalServerError, actions.R.JSON("Could not create user"))
	}

	return c.Render(http.StatusCreated, actions.R.JSON(""))
}

func handleLogin(c buffalo.Context) error {
	credentials := Credentials{}
	_, err := actions.GetJSONParametersFromBody(&c.Request().Body, &credentials)
	if err != nil {
		return c.Render(http.StatusBadRequest, actions.R.JSON("Invalid email or password"))
	}

	err = validateCredentials(&credentials)
	if err != nil {
		return c.Render(http.StatusBadRequest, actions.R.JSON("Invalid email or password"))
	}

	user, err := getUserByEmail(credentials.Email)
	if err != nil {
		return c.Render(http.StatusNotFound, actions.R.JSON("User not found"))
	}

	err = checkPassword(user.Password, credentials.Password)
	if err != nil {
		return c.Render(http.StatusUnauthorized, actions.R.JSON("Unauthorized"))
	}

	tokenString, err := generateJWT(credentials.Email)
	fmt.Println(err)
	if err != nil {
		return c.Render(http.StatusInternalServerError, actions.R.JSON("Something went wrong"))
	}

	type token struct {
		Token string `json:"token"`
	}
	authToken := token{Token: tokenString}

	return c.Render(http.StatusOK, actions.R.JSON(authToken))
}
