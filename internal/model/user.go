package model

import (
	"github.com/google/uuid"
)

type User struct {
	UserID uuid.UUID `json:"id"`
	FistName string  `json: "first_name"`
	LastName string  `json: "last_name"`
	Email    string  `json: "email"`
	Password string  `json: "password"`
}

