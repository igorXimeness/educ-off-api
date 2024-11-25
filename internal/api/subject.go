package api

import (
	"net/http"

	"github.com/igorXimeness/educ-off-api/internal/model"
	"github.com/igorXimeness/educ-off-api/internal/service"
	"github.com/labstack/echo/v4"
)

type SubjectAPI struct {
	subjectService service.SubjectService
}

func NewSubjectAPI(subjectService service.SubjectService) SubjectAPI {
	return SubjectAPI{
		subjectService: subjectService,
	}
}

func (api SubjectAPI) Register(e *echo.Echo) {
	v1 := e.Group("v1")
	v1.GET("/modules/:subject_name", api.FetchModules)
	v1.GET("/subjects", api.FetchSubjects)
	v1.GET("/subjects-with-done-modules", api.FetchSubjectsWithDoneModules) 
	v1.PUT("/modules/:module_id/finish", api.FinishModule)
	v1.POST("/subjects", api.CreateSubject) 
	v1.DELETE("/subjects/:subject_id", api.DeleteSubject)
	v1.POST("/modules", api.CreateModule)
	v1.DELETE("/modules/:module_id", api.DeleteModule)
}


func (api SubjectAPI) DeleteModule(c echo.Context) error {
	moduleID := c.Param("module_id")

	err := api.subjectService.DeleteModule(c.Request().Context(), moduleID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}

func (api SubjectAPI) CreateModule(c echo.Context) error {
    var module model.Modules
    if err := c.Bind(&module); err != nil { // Faz o bind do JSON para o modelo
        return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid payload"})
    }

    err := api.subjectService.CreateModule(c.Request().Context(), module)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
    }

    return c.JSON(http.StatusCreated, map[string]interface{}{"message": "Module created successfully!"})
}

func (api SubjectAPI) DeleteSubject(c echo.Context) error {
    subjectID := c.Param("subject_id") // Obtém o ID da matéria a partir da rota
    err := api.subjectService.DeleteSubject(c.Request().Context(), subjectID)

    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
    }
    return c.NoContent(http.StatusNoContent) // Retorna 204 No Content
}


func (api SubjectAPI) CreateSubject(c echo.Context) error {
	var subject model.Subject

	if err := c.Bind(&subject); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "invalid request body"})
	}

	// Chama o serviço para criar o novo assunto
	err := api.subjectService.CreateSubject(c.Request().Context(), subject)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{"message": "subject created successfully"})
}


func (api SubjectAPI) FetchSubjects(c echo.Context) error {
	subjects, err := api.subjectService.FetchSubjects(c.Request().Context())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, subjects)
}

func (api SubjectAPI) FetchModules(c echo.Context) error {
	subjectName := c.Param("subject_name") // Captura o parâmetro de rota

	modules, err := api.subjectService.FetchModules(c.Request().Context(), subjectName)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, modules)
}

func (api SubjectAPI) FetchSubjectsWithDoneModules(c echo.Context) error {
	subjects, err := api.subjectService.FetchSubjectsWithDoneModules(c.Request().Context())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, subjects)
}

func (api SubjectAPI) FinishModule(c echo.Context) error {
	moduleID := c.Param("module_id")
	err := api.subjectService.FinishModule(c.Request().Context(), moduleID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}
