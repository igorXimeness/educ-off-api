package api

import (
	"context"

	"net/http"
	"os"
	"path/filepath"

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
	v1.GET("/pdfs/:subject_name", api.FetchPDF) 
}


func (api LessonAPI) FetchLesson(c echo.Context) error {
    moduleName := c.Param("module_name")
    
    lesson, err := api.lessonService.FetchLesson(c.Request().Context(), moduleName)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
    }
    return c.JSON(http.StatusOK, lesson)
}



func (api LessonAPI) FetchPDF(c echo.Context) error {
    subjectName := c.Param("subject_name")
    
    // Caminho para o diretório "assets" a partir do diretório de trabalho
    pdfPath := filepath.Join("assets", subjectName+".pdf")
    
    // Verifica se o arquivo existe
    if _, err := os.Stat(pdfPath); err != nil {
        if os.IsNotExist(err) {
            return c.JSON(http.StatusNotFound, map[string]interface{}{
                "error": "File not found: " + pdfPath,
            })
        }
        // Retornar outros erros, caso ocorra algum problema inesperado
        return c.JSON(http.StatusInternalServerError, map[string]interface{}{
            "error": "Error checking file: " + err.Error(),
        })
    }
    
    // Retornando o PDF como resposta
    return c.File(pdfPath)
}
