package main

import (
    "fmt"
    "net/http"

    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"

    "github.com/igorXimeness/educ-off-api/internal/api"
    "github.com/igorXimeness/educ-off-api/internal/dao"
    "github.com/igorXimeness/educ-off-api/internal/service"
)
func main() {

    db, err := dao.ConnectDB()
    if err != nil{
        fmt.Println("Failed to connect to database", err)
    }
    defer db.Close()

    //inicializando DAO e service 

    userDao     := dao.NewUserDAO(db)
    userService := service.NewUserService(*userDao) 

    //inicializando ECHO
    e := echo.New()

    //Middleware
    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    //defininado api(rotas)
    userApi := api.NewUserAPI(*userService)
    userApi.Register(e)

    port := ":8080"

    fmt.Println("Server is running at port ", port)
    e.Start(port)



}
