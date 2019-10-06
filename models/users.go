package models

import (
	"time"
)

type (
	FormUserResponse struct {
		Id          int64     `json:"id" db:"id"`
		FirstName   string    `json:"first_name" db:"first_name"`
		LastName    string    `json:"last_name" db:"last_name"`
		Email       string    `json:"email" db:"email"`
		PhoneNumber string    `json:"phone_number" db:"phone_number"`
		BirthDate   *time.Time `json:"birth_date" db:"birth_date"`
		Gender      *string    `json:"gender" db:"gender"`
	}

	FormUserRequest struct {
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name"`
		Email       string `json:"email"`
		PhoneNumber string `json:"phone_number"`
		BirthDate   string `json:"birth_date"`
		Gender      string `json:"gender"`
	}

	FormPhone struct {
		PhoneNumber string `json:"phone_number"`
	}

	FormEmail struct {
		Email string `json:"email"`
	}
)
