package actions

import (
	"fmt"

	"github.com/tardog/schutzstreifen/models"
)

// Test that a list of existing hazard types is rendered
func (as *ActionSuite) Test_HazardTypesResource_List() {
	as.AdminLogin()
	as.LoadFixture("hazard types")

	res := as.HTML("/hazard_types").Get()

	as.Equal(200, res.Code)
	body := res.Body.String()

	as.Contains(body, "No cycle lane")
	as.Contains(body, "Cycling infrastructure is absent here.")
}

// Test that an existing hazard type can be displayed
func (as *ActionSuite) Test_HazardTypesResource_Show() {
	as.AdminLogin()
	as.LoadFixture("hazard types")

	res := as.HTML("/hazard_types/0111a456-3e1d-482d-85c3-5eee7df8816e").Get()

	as.Equal(200, res.Code)
	body := res.Body.String()

	as.Contains(body, "No cycle lane")
	as.Contains(body, "Cycling infrastructure is absent here.")
}

// Test that the form for new hazard types is shown
func (as *ActionSuite) Test_HazardTypesResource_New() {
	as.AdminLogin()

	res := as.HTML("/hazard_types/new").Get()
	as.Equal(200, res.Code)
}

// Test that a new hazard type can be created
func (as *ActionSuite) Test_HazardTypesResource_Create() {
	as.AdminLogin()

	label := "foo"
	description := "foobar"

	hazardType := &models.HazardType{
		Label:       label,
		Description: description,
	}

	res := as.HTML("/hazard_types").Post(hazardType)
	as.Equal(302, res.Code)

	newType := &models.HazardType{}
	err := as.DB.Last(newType)
	as.NoError(err)

	as.NotZero(newType.ID)
	as.Equal(label, newType.Label)
	as.Equal(description, newType.Description)
}

// Test that empty values are rejected
func (as *ActionSuite) Test_HazardTypesResource_Create_Validation() {
	as.AdminLogin()

	hazardType := &models.HazardType{
		Label:       "",
		Description: "",
	}

	res := as.HTML("/hazard_types").Post(hazardType)
	as.Equal(422, res.Code)
}

// Test that the edit form is shown for an existing hazard type
func (as *ActionSuite) Test_HazardTypesResource_Edit() {
	as.AdminLogin()
	as.LoadFixture("hazard types")

	res := as.HTML("/hazard_types/0111a456-3e1d-482d-85c3-5eee7df8816e/edit").Get()

	as.Equal(200, res.Code)
	body := res.Body.String()

	as.Contains(body, "No cycle lane")
	as.Contains(body, "Cycling infrastructure is absent here.")
}

// Test that an existing hazard type can be updated
func (as *ActionSuite) Test_HazardTypesResource_Update() {
	as.AdminLogin()
	as.LoadFixture("hazard types")

	id := "0111a456-3e1d-482d-85c3-5eee7df8816e"
	newLabel := "foobar"
	newDescription := "baz"

	hazardType := &models.HazardType{}
	as.DB.Find(hazardType, id)

	hazardType.Label = newLabel
	hazardType.Description = newDescription

	res := as.HTML("/hazard_types/" + id).Put(hazardType)
	as.Equal(302, res.Code)

	updatedHazardType := &models.HazardType{}
	err := as.DB.Find(updatedHazardType, id)
	as.NoError(err)

	fmt.Println(updatedHazardType)

	as.Equal(id, updatedHazardType.ID.String())
	as.Equal(newLabel, updatedHazardType.Label)
	as.Equal(newDescription, updatedHazardType.Description)
}

// Test that a hazard type can be deleted
func (as *ActionSuite) Test_HazardTypesResource_Destroy() {
	as.AdminLogin()
	as.LoadFixture("hazard types")

	res := as.HTML("/hazard_types/0111a456-3e1d-482d-85c3-5eee7df8816e").Delete()
	as.Equal(302, res.Code)

	hazardType := &models.HazardType{}
	err := as.DB.Find(hazardType, "0111a456-3e1d-482d-85c3-5eee7df8816e")

	as.Error(err)
}
