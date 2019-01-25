package models

import "github.com/gofrs/uuid"

// Hazard is a type of road safety issue
type Hazard struct {
	ID    uuid.UUID `db:"id"`
	Label string    `db:"label"`
}
