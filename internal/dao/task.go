package dao

import (
	"context"
	"fmt"
	"github.com/igorXimeness/educ-off-api/internal/model"
	"github.com/jackc/pgx/v4/pgxpool"
	"time"
)

type TaskDao struct {
	db *pgxpool.Pool
}

func NewTaskDAO(db *pgxpool.Pool) TaskDao {
	return TaskDao{db: db}
}


func (dao TaskDao) FetchTasks(ctx context.Context) ([]model.Task, error) {
    rows, err := dao.db.Query(ctx, "SELECT task_id, description, date, status FROM tasks WHERE status != 'removed'")
    if err != nil {
        return nil, fmt.Errorf("failed to fetch tasks: %w", err)
    }
    defer rows.Close()

    var tasks []model.Task
    for rows.Next() {
        var task model.Task
        var date time.Time // Usa time.Time para ler a data do banco de dados

        err := rows.Scan(&task.TaskID, &task.Description, &date, &task.Status)
        if err != nil {
            return nil, fmt.Errorf("failed to scan task: %w", err)
        }

        // Atribuindo a data como string no formato desejado
        task.Date = date.Format("2006-01-02")

        tasks = append(tasks, task)
    }
    if rows.Err() != nil {
        return nil, fmt.Errorf("rows error: %w", rows.Err())
    }

    return tasks, nil
}



func (dao TaskDao) CreateTask(ctx context.Context, task model.Task) error {
	_, err := dao.db.Exec(ctx, "INSERT INTO tasks (description, date, status) VALUES ($1, $2, $3)", task.Description, task.Date, task.Status)
	if err != nil {
		return fmt.Errorf("failed to create task: %w", err)
	}
	return nil
}


func (dao TaskDao) UpdateTask(ctx context.Context, taskID string, task model.Task) error {
	_, err := dao.db.Exec(ctx, "UPDATE tasks SET status = $1 WHERE task_id = $2", task.Status, taskID)
	if err != nil {
		return fmt.Errorf("failed to update task: %w", err)
	}
	return nil
}

// Deletar Tarefa (atualizar status para 'removed')
func (dao TaskDao) DeleteTask(ctx context.Context, taskID string) error {
	_, err := dao.db.Exec(ctx, "UPDATE tasks SET status = 'removed' WHERE task_id = $1", taskID)
	if err != nil {
		return fmt.Errorf("failed to delete task: %w", err)
	}
	return nil
}
