package users

import (
	"net/http"

	"account/actions"

	"github.com/gobuffalo/buffalo"
)

func Router(app *buffalo.App) {
	group := app.Group("/users/")
	group.POST("", handleCreation)
}

func handleCreation(c buffalo.Context) error {
	type Req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	req := Req{}
	_, err := actions.GetJSONParametersFromBody(&c.Request().Body, &req)
	if err != nil {
		return c.Render(http.StatusBadRequest, actions.R.JSON("Invalid email or password"))
	}
	// Validate Password
	password, err := actions.IsString(req.Password, 60)
	if err != nil {
		return c.Render(http.StatusBadRequest, actions.R.JSON("Invalid email or password"))
	}

	// Validate Email
	email, err := actions.IsEmail(req.Email)
	if err != nil {
		return c.Render(http.StatusBadRequest, actions.R.JSON("Invalid email or password"))
	}

	// Call direct from repository since we don't have any business logic
	err = HashPasswordAndCreateUser(password, email)
	if err != nil {
		return c.Render(http.StatusInternalServerError, actions.R.JSON("Could not create user"))
	}

	return c.Render(http.StatusAccepted, actions.R.JSON(""))
}
