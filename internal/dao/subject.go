package dao

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/igorXimeness/educ-off-api/internal/model"
	"github.com/jackc/pgx/v4/pgxpool"
)

type SubjectDao struct {
	subjectDao *pgxpool.Pool
}

func NewSubjectDAO(subjectDao *pgxpool.Pool) SubjectDao {
	return SubjectDao{
		subjectDao: subjectDao,
	}
}


func (dao SubjectDao) DeleteModule(ctx context.Context, moduleID string) error {
	// Excluir as lições associadas ao módulo
	_, err := dao.subjectDao.Exec(ctx, "DELETE FROM lesson WHERE module_id = $1", moduleID)
	if err != nil {
		return fmt.Errorf("failed to delete lessons associated with module: %w", err)
	}

	// Excluir o módulo
	_, err = dao.subjectDao.Exec(ctx, "DELETE FROM modules WHERE module_id = $1", moduleID)
	if err != nil {
		return fmt.Errorf("failed to delete module: %w", err)
	}

	return nil
}


func (dao SubjectDao) CreateModule(ctx context.Context, module model.Modules) error {
    query := `
        INSERT INTO modules (module_name, subject_id, done)
        VALUES ($1, $2, $3)
    `
    _, err := dao.subjectDao.Exec(ctx, query, module.ModuleName, module.SubjectID, module.Done)
    if err != nil {
        return fmt.Errorf("failed to create module: %w", err)
    }
    return nil
}


func (dao SubjectDao) DeleteSubject(ctx context.Context, subjectID string) error {
    _, err := dao.subjectDao.Exec(ctx, "DELETE FROM subjects WHERE subject_id = $1", subjectID)
    if err != nil {
        return fmt.Errorf("failed to delete subject with ID %s: %w", subjectID, err)
    }
    return nil
}


func (dao SubjectDao) CreateSubject(ctx context.Context, subject model.Subject) error {
	_, err := dao.subjectDao.Exec(ctx, "INSERT INTO subjects (name) VALUES ($1)", subject.Name)
	if err != nil {
		return fmt.Errorf("failed to insert subject: %w", err)
	}
	return nil
}


func (dao SubjectDao) FetchModules(ctx context.Context, subjectName string) ([]model.Modules, error) {
	var subjectID int
	err := dao.subjectDao.QueryRow(ctx, "SELECT subject_id FROM subjects WHERE name = $1", subjectName).Scan(&subjectID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch subject_id: %w", err)
	}

	rows, err := dao.subjectDao.Query(ctx, "SELECT module_id, module_name, done FROM modules WHERE subject_id = $1", subjectID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch modules: %w", err)
	}
	defer rows.Close()

	var modules []model.Modules
	for rows.Next() {
		var module model.Modules
		err := rows.Scan(&module.ModulesID, &module.ModuleName, &module.Done)
		if err != nil {
			return nil, fmt.Errorf("failed to scan module: %w", err)
		}
		// Remove spaces from module_name
		module.ModuleName = strings.TrimSpace(module.ModuleName)
		modules = append(modules, module)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("rows error: %w", rows.Err())
	}

	return modules, nil
}

func (dao SubjectDao) FetchSubjects(ctx context.Context) ([]model.Subject, error) {
	rows, err := dao.subjectDao.Query(ctx, "SELECT subject_id, name FROM subjects")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch subjects: %w", err)
	}
	defer rows.Close()

	var subjects []model.Subject
	for rows.Next() {
		var subject model.Subject
		err := rows.Scan(&subject.SubjectID, &subject.Name)
		if err != nil {
			return nil, fmt.Errorf("failed to scan subject: %w", err)
		}
		// Remove espaços extras do nome do assunto
		subject.Name = strings.TrimSpace(subject.Name)
		subjects = append(subjects, subject)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("rows error: %w", rows.Err())
	}

	return subjects, nil
}

func (dao SubjectDao) FinishModule(ctx context.Context, moduleID string) error {
	_, err := dao.subjectDao.Exec(ctx, "UPDATE modules SET done = TRUE WHERE module_id = $1", moduleID)
	if err != nil {
		return fmt.Errorf("failed to update module: %w", err)
	}
	return nil
}

func (dao SubjectDao) FetchSubjectsWithDoneModules(ctx context.Context) ([]model.Subject, error) {
	query := `
        SELECT DISTINCT s.subject_id, s.name 
        FROM subjects s
        JOIN modules m ON s.subject_id = m.subject_id
        WHERE m.done = TRUE
    `

	var subjects []model.Subject

	rows, err := dao.subjectDao.Query(ctx, query) // Use Query com o contexto
	if err != nil {
		return nil, fmt.Errorf("failed to query subjects: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var subject model.Subject
		if err := rows.Scan(&subject.SubjectID, &subject.Name); err != nil {
			return nil, fmt.Errorf("failed to scan subject: %w", err)
		}
		subjects = append(subjects, subject)
	}

	if len(subjects) == 0 {
		return nil, sql.ErrNoRows
	}

	return subjects, nil
}
