package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/uuid"
)

// User represents an authenticated user account
type User struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	Email     string    `json:"email" db:"email"`
	Password  string    `db:"password"`
	Hazards   Hazards   `json:"hazards" has_many:"hazards" order_by:"updated_at desc"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type Users []User

func (u Users) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}
