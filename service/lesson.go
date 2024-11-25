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

func (s SubjectService) CreateLesson(ctx context.Context, lesson model.Lesson) error {
    module, err := s.moduleRepo.FindByNameAndSubject(ctx, lesson.ModuleName, lesson.SubjectName)
    if err != nil {
        return err
    }

    if module == nil {
        return errors.New("module not found for the given subject")
    }

    lesson.ModuleID = module.ID

    return s.lessonRepo.Save(ctx, lesson)
}