package api

import (
	"context"
	"github.com/igorXimeness/educ-off-api/internal/model"
    "github.com/igorXimeness/educ-off-api/internal/service"
	"github.com/labstack/echo/v4"
)

type Subject interface {
	fetchModules(ctx context.Context, subject_name string) ([]model.Modules, error)
}


type SubjectAPI struct {
	subjectService service.SubjectService
}

func (api SubjectAPI) Register(e *echo.Echo) {
	v1 := e.Group("v1")
    v1.GET("/fetch-modules", api.FetchModules)
}

func FetchModules(c echo.Context) error {
	return error 

}
