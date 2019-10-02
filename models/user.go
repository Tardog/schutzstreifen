package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
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
	Admin        bool      `json:"-" db:"admin"`
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
	errors := validate.Validate(
		&validators.StringIsPresent{Name: "Name", Field: u.Name, Message: "You can't be nameless."},
		&validators.EmailIsPresent{Name: "Email", Field: u.Email, Message: "Please provide a valid e-mail address."},
		&validators.StringIsPresent{Name: "Password", Field: u.Password, Message: "You need to choose a password."},
	)

	return errors, nil
}

func (u *User) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	errors := validate.Validate(
		&validators.StringIsPresent{Name: "Name", Field: u.Name, Message: "You can't be nameless."},
		&validators.EmailIsPresent{Name: "Email", Field: u.Email, Message: "Please provide a valid e-mail address."},
	)

	return errors, nil
}

func (u *User) BeforeCreate(tx *pop.Connection) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return errors.WithStack(err)
	}

	u.PasswordHash = string(hash)
	u.Admin = false

	return nil
}
