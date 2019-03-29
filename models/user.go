package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

// User represents an authenticated user account
type User struct {
	ID           uuid.UUID `json:"id" db:"id"`
	Name         string    `json:"name" db:"name"`
	Email        string    `json:"email" db:"email"`
	PasswordHash string    `json:"password_hash" db:"password"`
	Password     string    `json:"-" db:"-" `
	Hazards      Hazards   `json:"hazards" has_many:"hazards" order_by:"updated_at desc"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}

type Users []User

func (u Users) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

func (u *User) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	errors := validate.NewErrors()

	if "" == u.Name {
		errors.Add("name", "Name cannot be empty")
	}

	if "" == u.Email {
		errors.Add("email", "Email cannot be empty")
	}

	if "" == u.Password {
		errors.Add("password", "Password cannot be empty")
	}

	return errors, nil
}

func (u *User) BeforeCreate(tx *pop.Connection) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.WithStack(err)
	}

	u.PasswordHash = string(hash)

	return nil
}
