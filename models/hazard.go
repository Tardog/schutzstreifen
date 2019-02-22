package models

import (
	"time"

	"github.com/gofrs/uuid"
)

// Hazard describes any kind of road safety issue for cyclists (and others)
type Hazard struct {
	ID          uuid.UUID `db:"id"`
	Label       string    `db:"label"`
	Description string    `db:"description"`
	Location    Geography `db:"location"`
	Visible     bool      `db:"visible"`
	User        User      `belongs_to:"user"`
	UserID      uuid.UUID `db:"user_id"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

// Hazards is an array of road hazards
type Hazards []Hazard

// Geography describes a geographical point
type Geography struct {
	Latitude  float32
	Longitude float32
}
