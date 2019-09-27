package grifts

import (
	"errors"

	"github.com/gobuffalo/pop"
	"github.com/markbates/grift/grift"
	"github.com/tardog/schutzstreifen/models"
)

var _ = grift.Namespace("db", func() {

	grift.Desc("seed", "Seeds a database")
	grift.Add("seed", func(c *grift.Context) error {
		tx, err := pop.Connect("development")
		if err != nil {
			return err
		}

		user := &models.User{
			Name:     "admin",
			Email:    "admin@schutzstreifen.dev",
			Password: "dirtysecret",
			Admin:    true,
		}

		verrs, err := tx.ValidateAndCreate(user)
		if err != nil || verrs.HasAny() {
			return errors.New("failed to create the admin user")
		}

		return nil
	})

})
