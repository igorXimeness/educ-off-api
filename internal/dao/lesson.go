package dao

import (
	"context"

	"github.com/igorXimeness/educ-off-api/internal/model"
)

type LessonDAO struct {
	db pgxpool.Pool
}



func NewLessonDAO(db *pgxpool.Pool) LessonDAO {
	return LessonDAO{
		db: *db,
	}
}


func FetchLesson(context ctx, string module) {
	
}