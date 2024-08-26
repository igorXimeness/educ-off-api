package dao

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

type SubjectDao struct {
	subjectDao pgxpool.Pool
}



func NewSubjectDAO(subjectDao *pgxpool.Pool) SubjectDao {
	return SubjectDao{
		subjectDao: *subjectDao,
	}
}