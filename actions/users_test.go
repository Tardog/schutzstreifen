package actions

import (
	"github.com/tardog/schutzstreifen/models"
)

// Test that existing users are displayed in the list
func (as *ActionSuite) Test_UsersResource_List() {
	as.LoadFixture("users")

	res := as.HTML("/users/").Get()
	as.Equal(200, res.Code)

	body := res.Body.String()

	as.Contains(body, "Carol Danvers")
	as.Contains(body, "captain.marvel@babymarkt.de")
	as.NotContains(body, "foobar")

	as.Contains(body, "King T&#39;Challa")
	as.Contains(body, "black.panther@babymarkt.de")
	as.NotContains(body, "barfoo")
}

// Test that an existing user can be displayed
func (as *ActionSuite) Test_UsersResource_Show() {
	as.LoadFixture("users")

	res := as.HTML("/users/afd08b38-7b5d-4ca4-919d-2ec2f358c187").Get()
	as.Equal(200, res.Code)

	body := res.Body.String()

	as.Contains(body, "Carol Danvers")
	as.Contains(body, "captain.marvel@babymarkt.de")
}

// Test that the create user form can be rendered
func (as *ActionSuite) Test_UsersResource_New() {
	res := as.HTML("/users/new").Get()
	as.Equal(200, res.Code)
}

// Test that a user is created with the values submitted by POST
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
	err := as.DB.First(newUser)
	as.NoError(err)

	as.NotZero(newUser.ID)
	as.Equal(name, newUser.Name)
	as.Equal(email, newUser.Email)
	as.NotZero(newUser.PasswordHash)
}

// Test that empty values are rejected
func (as *ActionSuite) Test_UsersResource_Create_Validation() {
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

// Test that the edit form renders with existing user data
func (as *ActionSuite) Test_UsersResource_Edit() {
	as.LoadFixture("users")

	res := as.HTML("/users/afd08b38-7b5d-4ca4-919d-2ec2f358c187/edit").Get()
	as.Equal(200, res.Code)

	body := res.Body.String()

	as.Contains(body, "Carol Danvers")
	as.Contains(body, "captain.marvel@babymarkt.de")
	as.NotContains(body, "foobar")
}

// Test that a user can be updated with new data
func (as *ActionSuite) Test_UsersResource_Update() {
	as.LoadFixture("users")

	id := "afd08b38-7b5d-4ca4-919d-2ec2f358c187"
	name := "foobar"
	email := "foo@baz.bar"

	user := &models.User{
		Name:  name,
		Email: email,
	}

	res := as.HTML("/users/" + id).Put(user)
	as.Equal(200, res.Code)

	updatedUser := &models.User{}
	err := as.DB.Find(updatedUser, id)
	as.NoError(err)

	as.Equal(id, updatedUser.ID.String())
	as.Equal(name, updatedUser.Name)
	as.Equal(email, updatedUser.Email)
	as.NotZero(updatedUser.PasswordHash)
}

// Test that a user can be deleted
func (as *ActionSuite) Test_UsersResource_Destroy() {
	as.LoadFixture("users")

	res := as.HTML("/users/afd08b38-7b5d-4ca4-919d-2ec2f358c187").Delete()
	as.Equal(302, res.Code)

	user := &models.User{}
	err := as.DB.Find("afd08b38-7b5d-4ca4-919d-2ec2f358c187", user)

	as.Error(err)
}
