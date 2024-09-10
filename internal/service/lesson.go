package service

import (
	"context"
	"github.com/igorXimeness/educ-off-api/internal/model"
)

type LessonService struct {
	lessonRepository LessonRepository
}

type LessonRepository interface {
	FetchLesson(context.Context, string) (model.Lesson, error) 
}

func NewLessonService(lessonRepository LessonRepository) LessonService {
	return LessonService {
		lessonRepository: lessonRepository,
	}
}

func (l LessonService) FetchLesson(ctx context.Context, module string) (model.Lesson, error) {
	return l.lessonRepository.FetchLesson(ctx, module)
}