//1) conter a lógica de negócios para o cadastro e login
//2) validar os usuários
//3) criptografar senhas
//4) verificar credenciais

package service

import (
	"context"
	"errors"

	"github.com/igorXimeness/educ-off-api/internal/model" // Ajuste para o caminho correto do seu pacote model
	"golang.org/x/crypto/bcrypt"
)

// UserService representa o serviço que lida com a lógica de negócios para usuários
type UserService struct {
   userRepository UserRepository 
}

type UserRepository interface{ 
	Signup(context.Context, model.User)
	Login(context.Context, model.User)
}


// NewUserService cria uma nova instância de UserService
func NewUserService() *UserService {
    return &UserService{
        // Inicialize aqui os campos necessários
    }
}

// Signup lida com o cadastro de novos usuários
func (s *UserService) Signup(user *model.User) error {
    // 1) Validar os usuários
    if user.Email == "" || user.Password == "" {
        return errors.New("email and password are required")
    }

    // 2) Criptografar senhas
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    user.Password = string(hashedPassword)

    // 3) Salvar o usuário no banco de dados
    // Aqui você chamaria o método do seu repositório DAO para salvar o usuário
    // Exemplo: err = s.userRepo.CreateUser(user)
    if err != nil {
        return err
    }

    return nil
}

// Login verifica as credenciais do usuário
func (s *UserService) Login(email, password string) error {
    // 1) Verificar se o usuário existe
    // Exemplo: user, err := s.userRepo.GetUserByEmail(email)
    if err != nil {
        return err
    }

    // 2) Verificar credenciais
    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
    if err != nil {
        // Senha não confere
        return errors.New("invalid credentials")
    }

    return nil
}