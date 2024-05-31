package service

import (
    "context"
    "errors" // Importar o pacote de erros
    "database/sql" // Importar o pacote SQL
    "github.com/igorXimeness/educ-off-api/internal/model"
    "github.com/igorXimeness/educ-off-api/internal/dao"
    "golang.org/x/crypto/bcrypt"
)

type UserService struct {
    userDAO dao.UserDAO
}

func NewUserService(userDAO dao.UserDAO) *UserService {
    return &UserService{
        userDAO: userDAO,
    }
}

func (s *UserService) Signup(ctx context.Context, user model.User) error {
    // Verificar se o e-mail já está em uso
    _, err := s.userDAO.FindUserByEmail(ctx, user.Email)
    if err == nil {
        return errors.New("email already exists")
    } else if err != nil && err != sql.ErrNoRows {
        return err
    }
    
    // Criar o usuário
    err = s.userDAO.CreateUser(ctx, user)
    if err != nil {
        return err
    }

    return nil
}

func (s *UserService) Login(ctx context.Context, email, password string) (*model.User, error) {
    // Encontrar o usuário pelo e-mail
    user, err := s.userDAO.FindUserByEmail(ctx, email)
    if err != nil {
        return nil, err
    }
    
    // Verificar a senha
    err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
    if err != nil {
        return nil, err
    }

    return user, nil
}
