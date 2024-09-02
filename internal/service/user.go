package service

import (
	"context"
	//"database/sql" // Importar o pacote SQL
	"errors"       // Importar o pacote de erros
	"fmt"

	"github.com/google/uuid"
	"github.com/igorXimeness/educ-off-api/internal/model"
	//"golang.org/x/crypto/bcrypt"
)

type UserService struct {
    userRepository UserRepository
}

type UserRepository interface {
    Login(context.Context, model.User) error 
    CreateUser(context.Context, model.User) error 
    FindUserByEmail(context.Context, string) (model.User, error)   
}


func NewUserService(userRepository UserRepository ) UserService {
    return UserService{
        userRepository: userRepository,
    }
}

func (s UserService) Signup(ctx context.Context, user model.User) error {
    _, err := s.userRepository.FindUserByEmail(ctx, user.Email)
    if err == nil {
        return errors.New("email already exists")
    }
    
    user.UserID = uuid.New()

    err = s.userRepository.CreateUser(ctx, user)
    if err != nil {
        fmt.Printf("Error in Signup: %v\n", err)
        return err
    }

    return nil
}


func (s UserService) Login(ctx context.Context, email, password string) (model.User, error) {
    user, err := s.userRepository.FindUserByEmail(ctx, email)
    if err != nil {
        return model.User{}, err
    }

    if user.Password != password {
        return model.User{}, fmt.Errorf("invalid password")
    }

    return user, nil
}
