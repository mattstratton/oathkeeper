package models

import (
	"encoding/json"
	"time"

	"github.com/markbates/pop"
	"github.com/markbates/validate"
	"github.com/markbates/validate/validators"
	"github.com/satori/go.uuid"
)

type Event struct {
	ID             uuid.UUID `json:"id" db:"id"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
	Name           string    `json:"name" db:"name"`
	CfpOpenDate    time.Time `json:"cfp_open_date" db:"cfp_open_date"`
	CfpCloseDate   time.Time `json:"cfp_close_date" db:"cfp_close_date"`
	EventStartDate time.Time `json:"event_start_date" db:"event_start_date"`
	EventEndDate   time.Time `json:"event_end_date" db:"event_end_date"`
	CfpUrl         string    `json:"cfp_url" db:"cfp_url"`
	EventUrl       string    `json:"event_url" db:"event_url"`
}

// String is not required by pop and may be deleted
func (e Event) String() string {
	je, _ := json.Marshal(e)
	return string(je)
}

// Events is not required by pop and may be deleted
type Events []Event

// String is not required by pop and may be deleted
func (e Events) String() string {
	je, _ := json.Marshal(e)
	return string(je)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (e *Event) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: e.Name, Name: "Name"},
		&validators.TimeIsPresent{Field: e.CfpOpenDate, Name: "CfpOpenDate"},
		&validators.TimeIsPresent{Field: e.CfpCloseDate, Name: "CfpCloseDate"},
		&validators.TimeIsPresent{Field: e.EventStartDate, Name: "EventStartDate"},
		&validators.TimeIsPresent{Field: e.EventEndDate, Name: "EventEndDate"},
		&validators.StringIsPresent{Field: e.CfpUrl, Name: "CfpUrl"},
		&validators.StringIsPresent{Field: e.EventUrl, Name: "EventUrl"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (e *Event) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (e *Event) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
