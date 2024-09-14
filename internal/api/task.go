package api

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/igorXimeness/educ-off-api/internal/model"
	"github.com/igorXimeness/educ-off-api/internal/service"
)

type TaskAPI struct {
	taskService service.TaskService
}

func NewTaskAPI(taskService service.TaskService) TaskAPI {
	return TaskAPI{taskService: taskService}
}

func (api TaskAPI) Register(e *echo.Echo) {
	v1 := e.Group("v1")
	v1.GET("/tasks", api.FetchTasks)
	v1.POST("/tasks", api.CreateTask)
	v1.PUT("/tasks/:task_id", api.UpdateTask)
	v1.DELETE("/tasks/:task_id", api.DeleteTask)
}

// Listar Tarefas (funcionando)
func (api TaskAPI) FetchTasks(c echo.Context) error {
	tasks, err := api.taskService.FetchTasks(c.Request().Context())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, tasks)
}

// Criar Tarefa (funcionando)
func (api TaskAPI) CreateTask(c echo.Context) error {
	var task model.Task
	if err := c.Bind(&task); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": err})
	}

	err := api.taskService.CreateTask(c.Request().Context(), task)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, task)
}

// Atualizar Tarefa (funcionando, atualiza somente o status)
func (api TaskAPI) UpdateTask(c echo.Context) error {
	taskID := c.Param("task_id")
	var updatedTask model.Task
	if err := c.Bind(&updatedTask); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "Invalid request payload"})
	}

	err := api.taskService.UpdateTask(c.Request().Context(), taskID, updatedTask)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}

// Deletar Tarefa (atualizar status para "deleted")
func (api TaskAPI) DeleteTask(c echo.Context) error {
	taskID := c.Param("task_id")
	err := api.taskService.DeleteTask(c.Request().Context(), taskID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
	}
	return c.NoContent(http.StatusNoContent)
}
