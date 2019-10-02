package models

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

// Hazard is a permanent infrastructural or environmental danger on the road for cyclists (and others)
type Hazard struct {
	ID           uuid.UUID  `json:"id" db:"id"`
	Label        string     `json:"label" db:"label"`
	Description  string     `json:"description" db:"description"`
	Lat          float64    `json:"lat" db:"lat"`
	Lon          float64    `json:"lon" db:"lon"`
	Visible      bool       `json:"visible" db:"visible"`
	User         User       `json:"user" belongs_to:"user"`
	UserID       uuid.UUID  `json:"-" db:"user_id"`
	HazardType   HazardType `json:"hazard_type" belongs_to:"hazard_type"`
	HazardTypeID uuid.UUID  `json:"-" db:"hazard_type_id"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (h Hazard) String() string {
	jh, _ := json.Marshal(h)
	return string(jh)
}

// Hazards is not required by pop and may be deleted
type Hazards []Hazard

// String is not required by pop and may be deleted
func (h Hazards) String() string {
	jh, _ := json.Marshal(h)
	return string(jh)
}

// LatitudeValidator checks if the given value is a valid latitude
type LatitudeValidator struct {
	Name  string
	Field float64
}

func (v *LatitudeValidator) IsValid(errors *validate.Errors) {
	if v.Field < -90.0 || v.Field > 90.0 {
		errors.Add(strings.ToLower(v.Name), "Your latitude is out of bounds (-90 to 90).")
	}
}

// LongitudeValidator checks if the given value is a valid longitude
type LongitudeValidator struct {
	Name  string
	Field float64
}

func (v *LongitudeValidator) IsValid(errors *validate.Errors) {
	if v.Field < -180.0 || v.Field > 180.0 {
		errors.Add(strings.ToLower(v.Name), "Your longitude is out of bounds (-180 to 180).")
	}
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (h *Hazard) Validate(tx *pop.Connection) (*validate.Errors, error) {
	errors := validate.Validate(
		&validators.StringIsPresent{Name: "Label", Field: h.Label, Message: "A label is required."},
		&validators.StringIsPresent{Name: "Description", Field: h.Description, Message: "Please provide a brief description."},
		&LatitudeValidator{Name: "Lat", Field: h.Lat},
		&LongitudeValidator{Name: "Lon", Field: h.Lon},
		&validators.UUIDIsPresent{Name: "HazardTypeID", Field: h.HazardTypeID, Message: "Please select the type of this hazard."},
	)

	return errors, nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (h *Hazard) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (h *Hazard) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
