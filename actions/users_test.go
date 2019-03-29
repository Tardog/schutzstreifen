package actions

import (
	"github.com/tardog/schutzstreifen/models"
)

func (as *ActionSuite) Test_UsersResource_List() {
}

func (as *ActionSuite) Test_UsersResource_Show() {
}

func (as *ActionSuite) Test_UsersResource_New() {
	res := as.HTML("/users/new").Get()

	as.Equal(200, res.Code)
}

func (as *ActionSuite) Test_UsersResource_Create() {
	name := "foobar"
	email := "foo@baz.bar"
	password := "password12345"

	user := &models.User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	res := as.HTML("/users").Post(user)
	as.Equal(302, res.Code)

	newUser := &models.User{}

	// Read the first user from the DB
	err := as.DB.First(newUser)
	as.NoError(err)

	// Assert that the created user has the expected values
	as.NotZero(newUser.ID)
	as.Equal(name, newUser.Name)
	as.Equal(email, newUser.Email)
	as.Equal(password, newUser.Password)
}

func (as *ActionSuite) Test_UsersResource_Create_ValidateEmptyInput() {
	user := &models.User{
		Name:     "",
		Email:    "",
		Password: "",
	}

	res := as.HTML("/users").Post(user)
	as.Equal(422, res.Code)

	err := as.DB.First(user)
	as.Error(err)

	body := res.Body.String()
	as.Contains(body, "Name cannot be empty")
	as.Contains(body, "Email cannot be empty")
	as.Contains(body, "Password cannot be empty")
}

func (as *ActionSuite) Test_UsersResource_Edit() {
}

func (as *ActionSuite) Test_UsersResource_Update() {
}

func (as *ActionSuite) Test_UsersResource_Destroy() {
}
