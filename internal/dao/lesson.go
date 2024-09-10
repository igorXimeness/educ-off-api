package dao

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/igorXimeness/educ-off-api/internal/model"
	"github.com/jackc/pgx/v4/pgxpool"
)

type LessonDAO struct {
	db *pgxpool.Pool
}

func NewLessonDAO(db *pgxpool.Pool) LessonDAO {
	return LessonDAO {
		db: db,
	}
}

func (dao LessonDAO) FetchLesson(ctx context.Context, moduleName string) (model.Lesson, error) {
    var lesson model.Lesson

    // Executa a consulta
    err := dao.db.QueryRow(ctx, `
        SELECT l.lesson_id, l.module_id, l.title, l.content 
        FROM lesson l
        JOIN modules m ON l.module_id = m.module_id
        WHERE m.module_name = $1
    `, moduleName).Scan(&lesson.LessonID, &lesson.ModuleID, &lesson.Title, &lesson.Content)

    // Verifica se houve erro ao buscar a lição
    if err != nil {
        if err == sql.ErrNoRows {
            return model.Lesson{}, fmt.Errorf("no rows in result set for module_name: %s", moduleName)
        }
        return model.Lesson{}, fmt.Errorf("failed to fetch lesson for module_name: %s, error: %w", moduleName, err)
    }

    return lesson, nil
}
