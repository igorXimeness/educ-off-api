package model

import (
	"github.com/google/uuid"
)

type User struct {
    UserID    uuid.UUID `json:"id"`        
    FirstName string   `json:"first_name"`
    LastName  string   `json:"last_name"`
    Email     string   `json:"email"`
    Password  string   `json:"password"`
}




// UUID = 314fa231
// ID = 312
type LoginForm struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}