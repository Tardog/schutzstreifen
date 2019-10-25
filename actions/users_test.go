package actions

import (
	"fmt"

	"github.com/tardog/schutzstreifen/models"
)

// Test that existing users are displayed in the list
func (as *ActionSuite) Test_UsersResource_List() {
	as.LoadFixture("users")

	admin := &models.User{}
	err := as.DB.First(admin)
	as.NoError(err)

	as.Session.Set("current_user_id", admin.ID)

	res := as.HTML("/users/").Get()
	as.Equal(200, res.Code)

	body := res.Body.String()

	as.Contains(body, "Carol Danvers")
	as.Contains(body, "captain.marvel@schutzstreifen.dev")
	as.NotContains(body, "foobar")

	as.Contains(body, "King T&#39;Challa")
	as.Contains(body, "black.panther@schutzstreifen.dev")
	as.NotContains(body, "barfoo")

	as.Session.Clear()
}

// Test that an existing user can be displayed
func (as *ActionSuite) Test_UsersResource_Show() {
	as.LoadFixture("users")

	admin := &models.User{}
	err := as.DB.First(admin)
	as.NoError(err)

	as.Session.Set("current_user_id", admin.ID)

	res := as.HTML("/users/" + admin.ID.String()).Get()
	as.Equal(200, res.Code)

	body := res.Body.String()

	as.Contains(body, "admin@schutzstreifen.dev")
}

// Test that the create user form can be rendered
func (as *ActionSuite) Test_UsersResource_New() {
	as.LoadFixture("users")

	admin := &models.User{}
	err := as.DB.First(admin)
	as.NoError(err)

	as.Session.Set("current_user_id", admin.ID)

	res := as.HTML("/users/new").Get()
	as.Equal(200, res.Code)
}

// Test that a user is created with the values submitted by POST
func (as *ActionSuite) Test_UsersResource_Create() {
	as.LoadFixture("users")

	admin := &models.User{}
	err := as.DB.First(admin)
	as.NoError(err)

	as.Session.Set("current_user_id", admin.ID)

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
	err = as.DB.Last(newUser)
	as.NoError(err)

	as.NotZero(newUser.ID)
	as.Equal(name, newUser.Name)
	as.Equal(email, newUser.Email)
	as.NotZero(newUser.PasswordHash)
}

// Test that empty values are rejected
func (as *ActionSuite) Test_UsersResource_Create_Validation() {
	as.LoadFixture("users")

	admin := &models.User{}
	err := as.DB.First(admin)
	as.NoError(err)

	as.Session.Set("current_user_id", admin.ID)

	user := &models.User{
		Name:     "",
		Email:    "",
		Password: "",
	}

	res := as.HTML("/users").Post(user)
	as.Equal(422, res.Code)
}

// Test that the edit form renders with existing user data
func (as *ActionSuite) Test_UsersResource_Edit() {
	as.LoadFixture("users")

	admin := &models.User{}
	err := as.DB.First(admin)
	as.NoError(err)

	as.Session.Set("current_user_id", admin.ID)

	res := as.HTML("/users/afd08b38-7b5d-4ca4-919d-2ec2f358c187/edit").Get()
	as.Equal(200, res.Code)

	body := res.Body.String()

	as.Contains(body, "Carol Danvers")
	as.Contains(body, "captain.marvel@schutzstreifen.dev")
	as.NotContains(body, "foobar")
}

// Test that a user can be updated with new data
func (as *ActionSuite) Test_UsersResource_Update() {
	as.LoadFixture("users")

	admin := &models.User{}
	err := as.DB.First(admin)
	as.NoError(err)

	as.Session.Set("current_user_id", admin.ID)

	id := "afd08b38-7b5d-4ca4-919d-2ec2f358c187"
	newName := "foobar"
	newEmail := "foo@baz.bar"

	user := &models.User{
		Name:  newName,
		Email: newEmail,
	}

	res := as.HTML("/users/" + id).Put(user)
	as.Equal(200, res.Code)

	updatedUser := &models.User{}
	err = as.DB.Find(updatedUser, id)
	as.NoError(err)

	fmt.Println(updatedUser)

	as.Equal(id, updatedUser.ID.String())
	as.Equal(newName, updatedUser.Name)
	as.Equal(newEmail, updatedUser.Email)
	as.NotZero(updatedUser.PasswordHash)
}

// Test that a user can be deleted
func (as *ActionSuite) Test_UsersResource_Destroy() {
	as.LoadFixture("users")

	admin := &models.User{}
	err := as.DB.First(admin)
	as.NoError(err)

	as.Session.Set("current_user_id", admin.ID)

	res := as.HTML("/users/afd08b38-7b5d-4ca4-919d-2ec2f358c187").Delete()
	as.Equal(302, res.Code)

	user := &models.User{}
	err = as.DB.Find(user, "afd08b38-7b5d-4ca4-919d-2ec2f358c187")

	as.Error(err)
}
