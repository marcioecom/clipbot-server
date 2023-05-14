package models

import "time"

type Users struct {
	ID        string     `sql:"primary_key"`
	Email     string     `validate:"required,email" json:"email,omitempty"`
	Password  string     `validate:"required" json:"password,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
