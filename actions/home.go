package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/pkg/errors"
	"github.com/tardog/schutzstreifen/models"
)

// HomeHandler is a default handler to serve up a home page.
func HomeHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("index.html"))
}

// PointsHandler serves JSON data for marker points on the map
func PointsHandler(c buffalo.Context) error {
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}

	hazards := &models.Hazards{}

	if err := tx.Where("visible = true").All(hazards); err != nil {
		return errors.WithStack(err)
	}

	return c.Render(200, r.JSON(hazards))
}
