package models

import (
	"time"

	"github.com/gofrs/uuid"
)

// User represents a single account with access to the application
type User struct {
	ID        uuid.UUID `db:"id"`
	Name      string    `db:"name"`
	Email     string    `db:"email"`
	Password  string    `db:"password"`
	Hazards   Hazards   `has_many:"hazards" order_by:"updated_at desc"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
