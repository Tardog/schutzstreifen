package actions

import (
	"fmt"

	"github.com/gofrs/uuid"
	"github.com/tardog/schutzstreifen/models"
)

// Test that a list of existing hazards is rendered
func (as *ActionSuite) Test_HazardsResource_List() {
	as.UserLogin()
	as.LoadFixture("hazard types")
	as.LoadFixture("hazards")

	res := as.HTML("/hazards").Get()

	as.Equal(200, res.Code)
	body := res.Body.String()

	as.Contains(body, "Collision")
	as.Contains(body, "Watch out!")
}

// Test that an existing hazard can be displayed
func (as *ActionSuite) Test_HazardsResource_Show() {
	as.UserLogin()
	as.LoadFixture("hazard types")
	as.LoadFixture("hazards")

	res := as.HTML("/hazards/1111a456-3e1d-482d-85c3-5eee7df8816e").Get()

	as.Equal(200, res.Code)
	body := res.Body.String()

	as.Contains(body, "Collision")
	as.Contains(body, "Watch out!")
}

// Test that the form for new hazards is shown
func (as *ActionSuite) Test_HazardsResource_New() {
	as.UserLogin()

	res := as.HTML("/hazards/new").Get()
	as.Equal(200, res.Code)
}

// Test that a new hazard can be created
func (as *ActionSuite) Test_HazardsResource_Create() {
	as.UserLogin()
	as.LoadFixture("hazard types")

	typeID, _ := uuid.FromString("0111a456-3e1d-482d-85c3-5eee7df8816e")

	label := "Danger danger"
	description := "High voltage"
	lat := 1.05
	lon := 2.05

	hazard := &models.Hazard{
		Label:        label,
		Description:  description,
		Lat:          lat,
		Lon:          lon,
		Visible:      true,
		HazardTypeID: typeID,
	}

	res := as.HTML("/hazards").Post(hazard)
	as.Equal(302, res.Code)

	newHazard := &models.Hazard{}
	err := as.DB.Last(newHazard)
	as.NoError(err)

	as.NotZero(newHazard.ID)
	as.Equal(label, newHazard.Label)
	as.Equal(description, newHazard.Description)
	as.Equal(lat, newHazard.Lat)
	as.Equal(lon, newHazard.Lon)
}

// Test that the edit form is shown for an existing hazard
func (as *ActionSuite) Test_HazardsResource_Edit() {
	as.UserLogin()
	as.LoadFixture("hazard types")
	as.LoadFixture("hazards")

	res := as.HTML("/hazards/1111a456-3e1d-482d-85c3-5eee7df8816e/edit").Get()

	as.Equal(200, res.Code)
	body := res.Body.String()

	as.Contains(body, "Collision")
	as.Contains(body, "Watch out!")
}

// Test that an existing hazard can be updated
func (as *ActionSuite) Test_HazardsResource_Update() {
	as.UserLogin()
	as.LoadFixture("hazard types")
	as.LoadFixture("hazards")

	id := "1111a456-3e1d-482d-85c3-5eee7df8816e"
	newLabel := "foobar"
	newDescription := "baz"

	hazard := &models.Hazard{}
	as.DB.Find(hazard, id)

	hazard.Label = newLabel
	hazard.Description = newDescription

	res := as.HTML("/hazards/" + id).Put(hazard)
	as.Equal(302, res.Code)

	updatedHazard := &models.Hazard{}
	err := as.DB.Find(updatedHazard, id)
	as.NoError(err)

	fmt.Println(updatedHazard)

	as.Equal(id, updatedHazard.ID.String())
	as.Equal(newLabel, updatedHazard.Label)
	as.Equal(newDescription, updatedHazard.Description)
}

// Test that a hazard can be deleted
func (as *ActionSuite) Test_HazardsResource_Destroy() {
	as.UserLogin()
	as.LoadFixture("hazard types")
	as.LoadFixture("hazards")

	res := as.HTML("/hazards/1111a456-3e1d-482d-85c3-5eee7df8816e").Delete()
	as.Equal(302, res.Code)

	hazardType := &models.HazardType{}
	err := as.DB.Find(hazardType, "1111a456-3e1d-482d-85c3-5eee7df8816e")

	as.Error(err)
}
