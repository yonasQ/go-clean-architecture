package dto

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
)

type User struct {
	// ID is the unique identifier of the user.
	// It is automatically generated when the user is created.
	ID uuid.UUID `json:"id"`
	// FirstName is the first name of the user.
	FirstName string `json:"first_name,omitempty"`
	// MiddleName is the middle name of the user.
	MiddleName string `json:"middle_name,omitempty"`
	// LastName is the last name of the user.
	LastName string `json:"last_name,omitempty"`
	// Email is the email of the user.
	Email string `json:"email,omitempty"`
	// Status is the status of the user.
	// It is set to active by default after successful registration.
	Status string `json:"status,omitempty"`
	// CreatedAt is the time when the user is created.
	// It is automatically set when the user is created.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// DeletedAt is the time the user was deleted.
	DeletedAt time.Time `json:"deleted_at,omitempty"`
	// UpdatedAt is the time the user was last updated.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type RegisterUser struct {
	// FirstName is the first name of the user.
	FirstName string `json:"first_name,omitempty"`
	// MiddleName is the middle name of the user.
	MiddleName string `json:"middle_name,omitempty"`
	// LastName is the last name of the user.
	LastName string `json:"last_name,omitempty"`
	// Email is the email of the user.
	Email string `json:"email,omitempty"`
}

func (u RegisterUser) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.FirstName, validation.Required.Error("first name is required")),
		validation.Field(&u.MiddleName, validation.Required.Error("middle name is required")),
		validation.Field(&u.LastName, validation.Required.Error("last name is required")),
		validation.Field(&u.Email, validation.Required.Error("email is required"), is.EmailFormat.Error("email is not valid")),
	)
}

type UpdateUser struct {
	// FirstName is the first name of the user.
	FirstName string `json:"first_name,omitempty"`
	// MiddleName is the middle name of the user.
	MiddleName string `json:"middle_name,omitempty"`
	// LastName is the last name of the user.
	LastName string `json:"last_name,omitempty"`
}

func (u UpdateUser) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.FirstName, validation.Required.Error("first name is required")),
		validation.Field(&u.MiddleName, validation.Required.Error("middle name is required")),
		validation.Field(&u.LastName, validation.Required.Error("last name is required")),
	)
}
