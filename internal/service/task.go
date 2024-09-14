package service

import (
	"context"
	"github.com/igorXimeness/educ-off-api/internal/model"
)

type TaskService struct {
	taskRepository TaskRepository
}

type TaskRepository interface {
	FetchTasks(ctx context.Context) ([]model.Task, error)
	CreateTask(ctx context.Context, task model.Task) error
	UpdateTask(ctx context.Context, taskID string, task model.Task) error
	DeleteTask(ctx context.Context, taskID string) error
}

func NewTaskService(taskRepository TaskRepository) TaskService {
	return TaskService{taskRepository: taskRepository}
}

func (s TaskService) FetchTasks(ctx context.Context) ([]model.Task, error) {
	return s.taskRepository.FetchTasks(ctx)
}

func (s TaskService) CreateTask(ctx context.Context, task model.Task) error {
	return s.taskRepository.CreateTask(ctx, task)
}

func (s TaskService) UpdateTask(ctx context.Context, taskID string, task model.Task) error {
	return s.taskRepository.UpdateTask(ctx, taskID, task)
}

func (s TaskService) DeleteTask(ctx context.Context, taskID string) error {
	return s.taskRepository.DeleteTask(ctx, taskID)
}
