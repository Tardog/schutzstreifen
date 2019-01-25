package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/tardog/schutzstreifen/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
