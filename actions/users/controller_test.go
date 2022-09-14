package users

import (
	"log"
	"net/http"
	"testing"
	"time"

	"account/actions"

	"github.com/gobuffalo/suite/v4"
)

type ActionSuite struct {
	*suite.Action
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
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

	user = createUser{}
	user.Email = "not_an_email"
	user.Password = "testingPassword"
	res = as.JSON("/users").Post(user)
	as.Equal(http.StatusBadRequest, res.Code)

	user = createUser{}
	user.Email = ""
	user.Password = ""
	res = as.JSON("/users").Post(user)
	as.Equal(http.StatusBadRequest, res.Code)

	user = createUser{}
	res = as.JSON("/users").Post(user)
	as.Equal(http.StatusBadRequest, res.Code)
}

func Test_ActionSuite(t *testing.T) {
	app := actions.App()
	Router(app)
	as := &ActionSuite{suite.NewAction(app)}
	suite.Run(t, as)
	as = &ActionSuite{suite.NewAction(app)}
	suite.Run(t, as)
}
