package service
import (
	"context"
	"errors"       
	"github.com/igorXimeness/educ-off-api/internal/model"

)
type SubjectService struct {
	userRepository SubjectRepository
}

type SubjectRepository interface {
	FetchModules(context.Context, string) error 
	FetchSubjects(context.Context) error 
}

func NewSubjectService(userRepository SubjectRepository) SubjectService {
	return SubjectService {
		userRepository: userRepository,
	}
}


func FetchSubjects(context ctx) error {

}

func FetchModules(context ctx, subject string) error {

}

