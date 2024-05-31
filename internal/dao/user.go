package dao

import (
    "context"
    "github.com/igorXimeness/educ-off-api/internal/model"
    "github.com/jackc/pgx/v4/pgxpool"
)

type UserDAO struct {
    db *pgxpool.Pool
}

func NewUserDAO(db *pgxpool.Pool) *UserDAO {
    return &UserDAO{
        db: db,
    }
}

func (dao *UserDAO) CreateUser(ctx context.Context, user model.User) error {
    _, err := dao.db.Exec(ctx, "INSERT INTO users (email, password) VALUES ($1, $2)", user.Email, user.Password)
    return err
}

func (dao *UserDAO) FindUserByEmail(ctx context.Context, email string) (*model.User, error) {
    user := &model.User{}
    err := dao.db.QueryRow(ctx, "SELECT email, password FROM users WHERE email = $1", email).Scan(&user.Email, &user.Password)
    if err != nil {
        return nil, err
    }
    return user, nil
}
