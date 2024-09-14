package main

import (
	"fmt"
	"os"

	"github.com/igorXimeness/educ-off-api/internal/api"
	"github.com/igorXimeness/educ-off-api/internal/dao"
	"github.com/igorXimeness/educ-off-api/internal/service"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Conexão com o banco de dados
	db, err := dao.ConnectDB()
	if err != nil {
		fmt.Printf("Failed to connect to database: %v\n", err)
		return
	}
	defer db.Close()

	// Inicializando DAO e Service
	userDao := dao.NewUserDAO(db)
	subjectDAO := dao.NewSubjectDAO(db)
	lessonDAO := dao.NewLessonDAO(db)
	taskDAO := dao.NewTaskDAO(db) // Novo TaskDAO

	userService := service.NewUserService(userDao)
	lessonService := service.NewLessonService(lessonDAO)
	subjectService := service.NewSubjectService(subjectDAO)
	taskService := service.NewTaskService(taskDAO) // Novo TaskService

	// Inicializando Echo
	server := echo.New()

	// Middlewares
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())

	// Definindo APIs (rotas)
	userApi := api.NewUserAPI(userService)
	userApi.Register(*&server)

	subjectApi := api.NewSubjectAPI(subjectService)
	subjectApi.Register(*&server)

	lessonApi := api.NewLessonAPI(lessonService)
	lessonApi.Register(*&server)

	taskApi := api.NewTaskAPI(taskService) // Novo TaskAPI
	taskApi.Register(*&server)             // Registrando rotas de tarefas

	// Configurando porta do servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server is running at port %s\n", port)
	if err := server.Start(":" + port); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}
