package users

import (
	"log"
	"net/http"
	"os"
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

func (as *ActionSuite) Test_handleUserLogin() {
	as.LoadFixture("credentials")

	credential := Credentials{}

	credential.Email = "testing-non-email"
	res := as.JSON("/users/auth").Post(credential)
	as.Equal(http.StatusBadRequest, res.Code)

	credential.Email = "test.found@test.com"
	credential.Password = "testing-password"
	res = as.JSON("/users/auth").Post(credential)
	as.Equal(http.StatusNotFound, res.Code)

	credential.Email = "test@test.com"
	credential.Password = "testing-password"
	res = as.JSON("/users/auth").Post(credential)
	as.Equal(http.StatusUnauthorized, res.Code)

	credential.Email = "test@test.com"
	credential.Password = "password"
	res = as.JSON("/users/auth").Post(credential)
	as.Contains(res.Body.String(), "token")
}

func Test_ActionSuite(t *testing.T) {
	app := actions.App()
	Router(app)

	action, err := suite.NewActionWithFixtures(app, os.DirFS("../../fixtures"))
	if err != nil {
		t.Fatal(err)
	}

	as := &ActionSuite{
		Action: action,
	}
	suite.Run(t, as)
}
