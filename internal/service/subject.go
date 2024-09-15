package service

import (
	"context"

	"github.com/igorXimeness/educ-off-api/internal/model"
)

type SubjectService struct {
	subjectRepository SubjectRepository
}

type SubjectRepository interface {
	FetchModules(context.Context, string) ([]model.Modules, error)
	FetchSubjects(context.Context) ([]model.Subject, error)
	FetchSubjectsWithDoneModules(context.Context) ([]model.Subject, error) // Novo método
	FinishModule(context.Context, string) error
}

func NewSubjectService(subjectRepository SubjectRepository) SubjectService {
	return SubjectService{
		subjectRepository: subjectRepository,
	}
}

func (s SubjectService) FetchSubjects(ctx context.Context) ([]model.Subject, error) {
	return s.subjectRepository.FetchSubjects(ctx)
}

func (s SubjectService) FetchModules(ctx context.Context, subject string) ([]model.Modules, error) {
	return s.subjectRepository.FetchModules(ctx, subject)
}

func (s SubjectService) FetchSubjectsWithDoneModules(ctx context.Context) ([]model.Subject, error) {
	return s.subjectRepository.FetchSubjectsWithDoneModules(ctx) // Chama o novo método no repositório
}

func (s SubjectService) FinishModule(ctx context.Context, moduleID string) error {
	return s.subjectRepository.FinishModule(ctx, moduleID)
}
