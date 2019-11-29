package actions

// Test that the / route loads successfully
func (as *ActionSuite) Test_HomeHandler() {
	res := as.HTML("/").Get()
	as.Equal(200, res.Code)
	as.Contains(res.Body.String(), "main-map")
}

// Test that the /points route returns a JSON representation of all hazards
func (as ActionSuite) Test_PointsHandler() {
	as.LoadFixture("hazard types")
	as.LoadFixture("users")
	as.LoadFixture("hazards")

	res := as.JSON("/points").Get()
	as.Equal(200, res.Code)

	as.Contains(res.Body.String(), string(`"id":"1111a456-3e1d-482d-85c3-5eee7df8816e"`))
}
