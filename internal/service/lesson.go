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
	CreateLesson(context.Context, model.Lesson) (int, error)
	DeleteLesson(context.Context, string) error 
	CreateQuestion(context.Context, model.Question) (int, error)
	FetchQuestionsByLessonID(ctx context.Context, lessonID int) ([]model.Question, error)

}

func NewLessonService(lessonRepository LessonRepository) LessonService {
	return LessonService {
		lessonRepository: lessonRepository,
	}
}

func (l LessonService) FetchQuestionsByLessonID(ctx context.Context, lessonID int) ([]model.Question, error) {
    return l.lessonRepository.FetchQuestionsByLessonID(ctx, lessonID)
}


func (l LessonService) DeleteLesson(ctx context.Context, lessonID string) error {
    return l.lessonRepository.DeleteLesson(ctx, lessonID)
}

func (l LessonService) CreateQuestion(ctx context.Context, question model.Question) (int, error) {
    return l.lessonRepository.CreateQuestion(ctx, question)
}

func (l LessonService) CreateLesson(ctx context.Context, lesson model.Lesson) (int, error) {
	return l.lessonRepository.CreateLesson(ctx, lesson)
}

func (l LessonService) FetchLesson(ctx context.Context, module string) (model.Lesson, error) {
	return l.lessonRepository.FetchLesson(ctx, module)
}