package actions

import (
	"testing"

	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/suite"
	"github.com/gofrs/uuid"
)

type ActionSuite struct {
	*suite.Action
}

func Test_ActionSuite(t *testing.T) {
	action, err := suite.NewActionWithFixtures(App(), packr.New("../fixtures", "../fixtures"))
	if err != nil {
		t.Fatal(err)
	}

	as := &ActionSuite{
		Action: action,
	}
	suite.Run(t, as)
}

func (as *ActionSuite) AdminLogin() {
	as.LoadFixture("users")

	adminID, _ := uuid.FromString("a3e9c741-357b-495e-a376-54d4cdb6f7ac")
	as.Session.Set("current_user_id", adminID)
}

func (as *ActionSuite) UserLogin() {
	as.LoadFixture("users")

	userID, _ := uuid.FromString("afd08b38-7b5d-4ca4-919d-2ec2f358c187")
	as.Session.Set("current_user_id", userID)
}
