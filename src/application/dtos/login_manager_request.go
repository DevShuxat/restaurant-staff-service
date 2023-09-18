package dtos

import (
	"errors"
	"fmt"
)

type LoginManagerRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r LoginManagerRequest) Validate() error {
	if r.Email == "" {
		return fmt.
			Errorf("invalid or missing email: %s", r.Email)
	}
	if r.Password == "" {
		return errors.New("invalid or missing password")
	}
	return nil
}
