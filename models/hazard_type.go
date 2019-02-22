package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
)

// HazardType is a type of road hazard
type HazardType struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Label       string    `json:"label" db:"label"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (h HazardType) String() string {
	jh, _ := json.Marshal(h)
	return string(jh)
}

// HazardTypes is not required by pop and may be deleted
type HazardTypes []HazardType

// String is not required by pop and may be deleted
func (h HazardTypes) String() string {
	jh, _ := json.Marshal(h)
	return string(jh)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (h *HazardType) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (h *HazardType) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (h *HazardType) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
