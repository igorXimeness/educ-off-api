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


type LoginForm struct {
	Email    string `json:"email" form:"email" query:"email"`
	Password string `json:"password" form:"password" query:"password"`
}
