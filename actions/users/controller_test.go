package users

import (
	"net/http"
	"testing"

	"account/actions"

	"github.com/gobuffalo/suite/v4"
)

type ActionSuite struct {
	*suite.Action
}

func (as *ActionSuite) Test_GetPing() {
	res := as.HTML("/users/ping").Get()

	as.Equal(http.StatusAccepted, res.Code)
	as.Contains(res.Body.String(), "deu bom")
}

// Test to validate account creation using email and password
func (as *ActionSuite) Test_handleUserCreation() {
	type createUser struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	user := createUser{}
	user.Email = "jorge@jorge.com"
	user.Password = "testingPassword"
	res := as.JSON("/users").Post(user)

	as.Equal(http.StatusCreated, res.Code)
}

func Test_ActionSuite(t *testing.T) {
	app := actions.App()
	Router(app)
	as := &ActionSuite{suite.NewAction(app)}
	suite.Run(t, as)
}
