package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
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
	UserID       uuid.UUID  `db:"user_id"`
	HazardType   HazardType `json:"hazard_type" has_one:"hazard_type"`
	HazardTypeID uuid.UUID  `db:"hazard_type_id"`
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

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (h *Hazard) Validate(tx *pop.Connection) (*validate.Errors, error) {
	errors := validate.NewErrors()

	if "" == h.Label {
		errors.Add("Label", "Label cannot be empty")
	}

	if "" == h.Description {
		errors.Add("Description", "Description cannot be empty")
	}

	if 0 == h.Lat {
		errors.Add("Lat", "Lat cannot be empty")
	}

	if 0 == h.Lon {
		errors.Add("Lon", "Lon cannot be empty")
	}

	if "" == h.HazardTypeID.String() {
		errors.Add("HazardTypeID", "Please select a type for this hazard")
	}

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
