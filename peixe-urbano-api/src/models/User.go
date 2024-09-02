package models

import (
	"errors"
	"strings"
	"time"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedOn time.Time `json:"createdOn,omitempty"`
}

// Prepare call methods to validate and format fields
func (user *User) Prepare() error {
	if err := user.validate(); err != nil {
		return err
	}
	user.format()
	return nil
}

// validate check empty fields
func (user *User) validate() error {
	if user.Name == "" {
		return errors.New("the name is required and must not empty")
	}
	if user.Nick == "" {
		return errors.New("the nick is required and must not empty")
	}
	if user.Email == "" {
		return errors.New("the email is required and must not empty")
	}
	if user.Password == "" {
		return errors.New("the name is required and must not empty")
	}
	return nil
}

// format remove spaces
func (user *User) format() {
	user.Name = strings.TrimSpace((user.Name))
	user.Nick = strings.TrimSpace((user.Nick))
	user.Email = strings.TrimSpace((user.Email))
}
