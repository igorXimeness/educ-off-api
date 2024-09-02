package dao

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/igorXimeness/educ-off-api/internal/model"
	"log"
)

type LessonDAO struct {
	db *pgxpool.Pool
}

func NewLessonDAO(db *pgxpool.Pool) LessonDAO {
	return LessonDAO{
		db: db,
	}
}

// FetchLessons busca todas as lições para um módulo específico
func (dao LessonDAO) FetchLessons(ctx context.Context, moduleID int) ([]model.Lesson, error) {
	rows, err := dao.db.Query(ctx, "SELECT lesson_id, module_id, title, content FROM lessons WHERE module_id = $1", moduleID)
	if err != nil {
		log.Printf("Error querying lessons: %v", err)
		return nil, err
	}
	defer rows.Close()

	var lessons []model.Lesson
	for rows.Next() {
		var lesson model.Lesson
		if err := rows.Scan(&lesson.LessonID, &lesson.ModuleID, &lesson.Title, &lesson.Content); err != nil {
			log.Printf("Error scanning lesson: %v", err)
			return nil, err
		}
		lessons = append(lessons, lesson)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Error iterating over rows: %v", err)
		return nil, err
	}

	return lessons, nil
}
