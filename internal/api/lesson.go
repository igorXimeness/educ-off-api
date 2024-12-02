package api

import (
	"context"
	"strconv"

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
    v1.POST("/modules/:module_name/lesson", api.CreateLesson) 
    v1.DELETE("/lessons/:lesson_id", api.DeleteLesson)        
    v1.POST("/lesson/:lesson_id/question", api.CreateQuestion) 
    v1.GET("/lessons/:lesson_id/questions", api.FetchQuestionsByLessonID)
    v1.POST("/pdfs", api.UploadPDF)

}

func (api LessonAPI) FetchQuestionsByLessonID(c echo.Context) error {
    lessonIDStr := c.Param("lesson_id")

    lessonID, err := strconv.Atoi(lessonIDStr)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid lesson ID"})
    }

    questions, err := api.lessonService.FetchQuestionsByLessonID(c.Request().Context(), lessonID)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
    }

    return c.JSON(http.StatusOK, questions)
}


func (api LessonAPI) DeleteLesson(c echo.Context) error {
    lessonID := c.Param("lesson_id")

    err := api.lessonService.DeleteLesson(c.Request().Context(), lessonID)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
    }

    return c.JSON(http.StatusOK, map[string]interface{}{"message": "Lesson deleted successfully"})
}

func (api LessonAPI) CreateLesson(c echo.Context) error {
    moduleName := c.Param("module_name")
    var newLesson model.Lesson

    // Bind do corpo da requisição para a estrutura da lição
    if err := c.Bind(&newLesson); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid request body"})
    }

    // Buscar o ModuleID a partir do moduleName
    module, err := api.lessonService.FetchLesson(c.Request().Context(), moduleName)
    if err != nil {
        return c.JSON(http.StatusNotFound, map[string]interface{}{"error": "Module not found"})
    }

    // Agora associamos o ModuleID da lição com o ModuleID do módulo
    newLesson.ModuleID = module.ModuleID

    // Cria a lição
    lessonID, err := api.lessonService.CreateLesson(c.Request().Context(), newLesson)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
    }

    return c.JSON(http.StatusCreated, map[string]interface{}{"lesson_id": lessonID})
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
    
    pdfPath := filepath.Join("assets", subjectName+".pdf")
    
    if _, err := os.Stat(pdfPath); err != nil {
        if os.IsNotExist(err) {
            return c.JSON(http.StatusNotFound, map[string]interface{}{
                "error": "File not found: " + pdfPath,
            })
        }
        return c.JSON(http.StatusInternalServerError, map[string]interface{}{
            "error": "Error checking file: " + err.Error(),
        })
    }
    
    // Retornando o PDF como resposta
    return c.File(pdfPath)
}


func (api LessonAPI) UploadPDF(c echo.Context) error {
	file, err := c.FormFile("pdf")
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Failed to read file: " + err.Error(),
		})
	}

	if filepath.Ext(file.Filename) != ".pdf" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": "Only PDF files are allowed.",
		})
	}

	pdfPath := filepath.Join("assets", file.Filename)

	src, err := file.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to open file: " + err.Error(),
		})
	}
	defer src.Close()

	dst, err := os.Create(pdfPath)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to save file: " + err.Error(),
		})
	}
	defer dst.Close()

	if _, err = dst.ReadFrom(src); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Failed to write file: " + err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "File uploaded successfully",
		"path":    pdfPath,
	})
}


func (api LessonAPI) CreateQuestion(c echo.Context) error {
    lessonIDStr := c.Param("lesson_id")
    var newQuestion model.Question

    // Converte lessonID de string para int
    lessonID, err := strconv.Atoi(lessonIDStr)
    if err != nil {
        return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid lesson ID"})
    }

    // Bind do corpo da requisição para a estrutura da questão
    if err := c.Bind(&newQuestion); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid request body"})
    }

    newQuestion.LessonID = lessonID // Agora lessonID é do tipo int

    questionID, err := api.lessonService.CreateQuestion(c.Request().Context(), newQuestion)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
    }

    return c.JSON(http.StatusCreated, map[string]interface{}{"question_id": questionID})
}
