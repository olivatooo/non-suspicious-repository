package users

import (
	"net/http"

	"account/actions"

	"github.com/gobuffalo/buffalo"
)

func AuthMiddleware(next buffalo.Handler) buffalo.Handler {
	return func(c buffalo.Context) error {
		secret := c.Request().Header.Get("secret")
		err := validateSecret(secret)
		if err != nil {
			return c.Render(http.StatusUnauthorized, actions.R.JSON("Unauthorized"))
		}
		user, err := FindBySecret(secret)
		c.Set("user_id", user.ID)
		if err != nil {
			return c.Render(http.StatusUnauthorized, actions.R.JSON("Unauthorized"))
		}
		return nil
	}
}

func Router(app *buffalo.App) {
	app.Use(AuthMiddleware)
	group := app.Group("/users/")
	group.POST("", handleCreation)
	group.POST("auth", handleLogin)
	group.Middleware.Skip(AuthMiddleware, handleCreation, handleLogin)
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
	err = hashPasswordAndCreateUser(credentials.Email, credentials.Password)
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

	secretString, err := generateSecret(user.ID)
	if err != nil {
		return c.Render(http.StatusInternalServerError, actions.R.JSON("Something went wrong"))
	}

	type secret struct {
		Secret string `json:"secret"`
	}
	authSecret := secret{Secret: secretString}

	return c.Render(http.StatusOK, actions.R.JSON(authSecret))
}
