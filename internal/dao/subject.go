package dao

import (
    "context"
    "fmt"

    "github.com/igorXimeness/educ-off-api/internal/model"
    "github.com/jackc/pgx/v4/pgxpool"
    "strings"
)

type SubjectDao struct {
    subjectDao pgxpool.Pool
}

func NewSubjectDAO(subjectDao *pgxpool.Pool) SubjectDao {
    return SubjectDao{
        subjectDao: *subjectDao,
    }
}
func (dao SubjectDao) FetchModules(ctx context.Context, subjectName string) ([]model.Modules, error) {
    var subjectID int
    err := dao.subjectDao.QueryRow(ctx, "SELECT subject_id FROM subjects WHERE name = $1", subjectName).Scan(&subjectID)
    if err != nil {
        return nil, fmt.Errorf("failed to fetch subject_id: %w", err)
    }

    rows, err := dao.subjectDao.Query(ctx, "SELECT subject_id, module_id, module_name, done FROM modules WHERE subject_id = $1", subjectID)
    if err != nil {
        return nil, fmt.Errorf("failed to fetch modules: %w", err)
    }
    defer rows.Close()

    var modules []model.Modules
    for rows.Next() {
        var module model.Modules
        err := rows.Scan(&module.SubjectID, &module.ModulesID, &module.ModuleName, &module.Done)
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