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
	FetchSubjectsWithDoneModules(context.Context) ([]model.Subject, error) 
	FinishModule(context.Context, string) error
	CreateSubject (context.Context, model.Subject) error 
	DeleteSubject (context.Context, string) error
	CreateModule(context.Context, model.Modules) error 
	DeleteModule(context.Context, string) error 
}

func NewSubjectService(subjectRepository SubjectRepository) SubjectService {
	return SubjectService{
		subjectRepository: subjectRepository,
	}
}

func (s SubjectService) CreateSubject(ctx context.Context, subject model.Subject) error {
	return s.subjectRepository.CreateSubject(ctx, subject)
}

func (s SubjectService) DeleteSubject(ctx context.Context, subjectID string) error {
    return s.subjectRepository.DeleteSubject(ctx, subjectID)
}

func (s SubjectService) DeleteModule(ctx context.Context, moduleID string) error {
	return s.subjectRepository.DeleteModule(ctx, moduleID)
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

func (s SubjectService) CreateModule(ctx context.Context, module model.Modules) error {
    return s.subjectRepository.CreateModule(ctx, module)
}


func (s SubjectService) FinishModule(ctx context.Context, moduleID string) error {
	return s.subjectRepository.FinishModule(ctx, moduleID)
}
