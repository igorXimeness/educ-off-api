package api

import (
	"context"
	"log"
	"net/http"

	"github.com/igorXimeness/educ-off-api/internal/model"
	"github.com/igorXimeness/educ-off-api/internal/service"
	"github.com/labstack/echo/v4"
)

type UserService interface {
    Signup(ctx context.Context, user model.User) error
    Login(ctx context.Context, email, password string) (model.User, error)
}
//INJETAR service na camada de api 

type UserAPI struct {
    userService service.UserService
}

func NewUserAPI(userService service.UserService) UserAPI {
    return UserAPI {
        userService: userService,
    }
}

func (api UserAPI) Register(e *echo.Echo) {
    v1 := e.Group("v1")

    v1.POST("/signup", api.signup)
    v1.POST("/login", api.login)
}

// 1) preencher os campos para cadastrar 
// 2) apertar botao para cadastrar
// 3) educ-off/signup  {nome: "", "email": ""}


func (api UserAPI) signup(c echo.Context) error {
    user := model.User{}
    if err := c.Bind(&user); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "invalid request"})
    }

    if user.Email == "" || user.Password == "" || user.FirstName == "" || user.LastName == "" {
        return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "missing required fields"})
    }

    if len(user.FirstName) < 2 || len(user.LastName) < 2 {
        return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "invalid name"})
    }

    err := api.userService.Signup(c.Request().Context(), user)
    if err != nil {
        log.Printf("Error during user signup: %v", err)
        
        return c.JSON(http.StatusInternalServerError, map[string]interface{}{"error": err.Error()})
    }

    return c.JSON(http.StatusCreated, user)
}


func (api UserAPI) login(c echo.Context) error {

	form := model.LoginForm{}
    if err := c.Bind(&form); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]interface{}{"error": "invalid request"})
    }

    user, err := api.userService.Login(c.Request().Context(), form.Email, form.Password)
    if err != nil {
        return c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": err.Error()})
    }

    return c.JSON(http.StatusOK, user)
}
