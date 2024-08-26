package service

import "context"

type SubjectService struct {
	userRepository SubjectRepository
}

type SubjectRepository interface {
	FetchModules(context.Context, string)
}

func NewSubjectService(userRepository SubjectRepository) SubjectService {
	return SubjectService {
		userRepository: userRepository,
	}
}

func FetchModules(context ctx, subject string) error {

}