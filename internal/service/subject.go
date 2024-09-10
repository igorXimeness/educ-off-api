package service

import (
	"context"
	"github.com/igorXimeness/educ-off-api/internal/model"
)

type SubjectService struct {
	userRepository SubjectRepository
}

type SubjectRepository interface {
	FetchModules(context.Context, string) ([]model.Modules, error)  
	FetchSubjects(context.Context) ([]model.Subject, error) 
	FinishModule(context.Context, string) error 
}

func NewSubjectService(userRepository SubjectRepository) SubjectService {
	return SubjectService {
		userRepository: userRepository,
	}
}


func(s SubjectService) FetchSubjects(ctx context.Context) ([]model.Subject, error) {
	return s.userRepository.FetchSubjects(ctx)
}

func (s SubjectService) FetchModules(ctx context.Context, subject string) ([]model.Modules, error) {
    return s.userRepository.FetchModules(ctx, subject)
}


func (s SubjectService) FinishModule(ctx context.Context, moduleID string) error {
	return s.userRepository.FinishModule(ctx, moduleID)
}