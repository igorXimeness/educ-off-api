package main

import (
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"fmt"

	"github.com/igorXimeness/educ-off-api/internal/api"
	"github.com/igorXimeness/educ-off-api/internal/dao"
	"github.com/igorXimeness/educ-off-api/internal/service"
)

func main() {
    // Cria um contexto raiz com cancelamento
 

    // Conexão com o banco de dados
    db, err := dao.ConnectDB()
    if err != nil {
        fmt.Printf("Failed to connect to database: %v\n", err)
        return
    }
    defer db.Close()

    // Inicializando DAO e Service com o contexto
    userDao := dao.NewUserDAO(db)
    userService := service.NewUserService(userDao) 

    // Inicializando Echo
    server := echo.New()

    // Middlewares
    server.Use(middleware.Logger())
    server.Use(middleware.Recover())

    // Definindo API (rotas)
    userApi := api.NewUserAPI(userService)
    userApi.Register(*server)

    // Definindo porta (pode ser configurável via variável de ambiente)
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    fmt.Printf("Server is running at port %s\n", port)
    if err := server.Start(":" + port); err != nil {
        fmt.Printf("Failed to start server: %v\n", err)
    }
}
