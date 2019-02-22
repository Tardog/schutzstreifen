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
	Latitude    float32   `db:"latitude"`
	Longitude   float32   `db:"longitude"`
	User        User      `belongs_to:"user"`
	UserID      uuid.UUID `db:"user_id"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

// Hazards is a collection of road hazards
type Hazards []Hazard
