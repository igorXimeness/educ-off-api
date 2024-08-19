package dao

import (
	"context"
	"fmt"

	"github.com/igorXimeness/educ-off-api/internal/model"
	"github.com/jackc/pgx/v4/pgxpool"
	"golang.org/x/crypto/bcrypt"
	 "github.com/sirupsen/logrus"
)

type UserDAO struct {
	db pgxpool.Pool
}

// Signup implements service.UserRepository.
func (UserDAO) Signup(context.Context, model.User) error {
	panic("unimplemented")
}

func NewUserDAO(db *pgxpool.Pool) UserDAO {
	return UserDAO{
		db: *db,
	}
}
var log = logrus.New()

func (dao UserDAO) CreateUser(ctx context.Context, user model.User) error {
    _, err := dao.db.Exec(ctx, "INSERT INTO users (email, password, first_name, last_name) VALUES ($1, $2, $3, $4)", user.Email, user.Password, user.FirstName, user.LastName)
    if err != nil {
        log.WithFields(logrus.Fields{
            "user":  user,
            "error": err,
        }).Error("Failed to create user")
        return fmt.Errorf("failed to create user: %w", err)
    }
    return nil
}
func (dao UserDAO) FindUserByEmail(ctx context.Context, email string) (model.User, error) {
	user := model.User{}
	err := dao.db.QueryRow(ctx, "SELECT email, password FROM users WHERE email = $1", email).Scan(&user.Email, &user.Password)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}
func (dao UserDAO) Login(ctx context.Context, user model.User) error {
    // Encontrar o usuário pelo e-mail
    storedUser, err := dao.FindUserByEmail(ctx, user.Email)
    if err != nil {
        return err
    }
    err = bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
    if err != nil {
        return err 
    }

    return nil
}