package models

import "time"

type (
	FormUser struct {
		Id          int64     `json:"id" db:"id"`
		FirstName   string    `db:"first_name"`
		LastName    string    `db:"last_name"`
		Email       string    `db:"email"`
		PhoneNumber string    `db:"phone_number"`
		BirthDate   time.Time `db:"birth_date"`
		Gender      string    `db:"gender"`
	}

	FormPhone struct {
		PhoneNumber string `json:"phone_number"`
	}

	FormEmail struct {
		Email string `json:"email"`
	}
)
