package api

import (
	"context"
	"net/http"

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

func NewSubjectAPI(subjectService service.SubjectService) SubjectAPI {
    return SubjectAPI {
        subjectService: subjectService,
    }
}


func (api SubjectAPI) Register(e *echo.Echo) {
	v1 := e.Group("v1")
    v1.GET("/modules/:subject_name", api.FetchModules)
	v1.GET("/subjects", api.FetchSubjects)
    v1.PUT("/modules/:module_id/finish", api.FinishModule) 
}

func (api SubjectAPI) FetchSubjects(c echo.Context) error {
    subjects, err := api.subjectService.FetchSubjects(c.Request().Context())

    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
    }
	
    return c.JSON(http.StatusOK, subjects)
}

func (api SubjectAPI) FetchModules(c echo.Context) error {
    subjectName := c.Param("subject_name")  // Corrigindo a captura do parâmetro de rota

    modules, err := api.subjectService.FetchModules(c.Request().Context(), subjectName)

    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
    }

    return c.JSON(http.StatusOK, modules)
}


func (api SubjectAPI) FinishModule(c echo.Context) error {
	moduleID := c.Param("module_id")
	err := api.subjectService.FinishModule(c.Request().Context(), moduleID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent) // Retorna apenas um status 204 No Content
}