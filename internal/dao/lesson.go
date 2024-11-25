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


func (dao LessonDAO) CreateLesson(ctx context.Context, lesson model.Lesson) (int, error) {
	// Insere a lição no banco de dados
	var lessonID int
	err := dao.db.QueryRow(ctx, `
		INSERT INTO lesson (module_id, title, content)
		VALUES ($1, $2, $3)
		RETURNING lesson_id
	`, lesson.ModuleID, lesson.Title, lesson.Content).Scan(&lessonID)

	if err != nil {
		return 0, fmt.Errorf("failed to create lesson: %w", err)
	}

	return lessonID, nil
}


func (dao LessonDAO) DeleteLesson(ctx context.Context, lessonID string) error {
	// Excluir as questões associadas a essa lição
	_, err := dao.db.Exec(ctx, `DELETE FROM question WHERE lessonId = $1`, lessonID)
	if err != nil {
		return fmt.Errorf("failed to delete questions associated with the lesson: %w", err)
	}

	// Agora, excluir a lição
	_, err = dao.db.Exec(ctx, `DELETE FROM lesson WHERE lesson_id = $1`, lessonID)
	if err != nil {
		return fmt.Errorf("failed to delete lesson: %w", err)
	}

	return nil
}


func (dao LessonDAO) CreateQuestion(ctx context.Context, question model.Question) (int, error) {
    var questionID int

    // Executando a consulta de inserção da questão com as opções e a resposta correta
    err := dao.db.QueryRow(ctx, `
        INSERT INTO question (lessonId, questionText, optionA, optionB, optionC, optionD, rightOption)
        VALUES ($1, $2, $3, $4, $5, $6, $7)
        RETURNING question_id`,
        question.LessonID, 
        question.QuestionText,
        question.OptionA, 
        question.OptionB, 
        question.OptionC, 
        question.OptionD, 
        question.RightOption).
        Scan(&questionID)

    if err != nil {
        return 0, fmt.Errorf("failed to insert question: %w", err)
    }

    return questionID, nil
}