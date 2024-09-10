package api

import (
	"context"
	"net/http"

	"github.com/igorXimeness/educ-off-api/internal/model"
	"github.com/igorXimeness/educ-off-api/internal/service"
	"github.com/labstack/echo/v4"
)

type Lesson interface {
	FetchLesson(ctx context.Context, subject_name string) ([]model.Modules, error)
}


type LessonAPI struct {
	lessonService service.LessonService
}

func NewLessonAPI(lessonService service.LessonService) LessonAPI {
    return LessonAPI {
        lessonService: lessonService,
    }
}


func (api LessonAPI) Register(e *echo.Echo) {
	v1 := e.Group("v1")
	v1.GET("/modules/:module_name/lesson", api.FetchLesson)
}


func (api LessonAPI) FetchLesson(c echo.Context) error {
    moduleName := c.Param("module_name")
    
    lesson, err := api.lessonService.FetchLesson(c.Request().Context(), moduleName)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, lesson)
}
