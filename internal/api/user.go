//1) definir os handlers para as rotas de login e cadastro
//2) extrair dados da requisição
//3) chamar os metodos correspondentes a camada de serviço

package api

import (
	"net/http"

	"github.com/igorXimeness/educ-off-api/internal/model"
	"github.com/igorXimeness/educ-off-api/internal/service"
	"github.com/labstack/echo/v4"
	//"net/http"
)


type UserAPI struct {
	userService service.UserService
}

//é como se fosse um construtor de userApi
func newUserAPI(userService service.UserService) *UserAPI {
	return &UserAPI{
		userService: userService,
	}
}

func (api UserAPI) Register(e *echo.Echo){
	e.POST("/signup", api.signup)
	//e.POST("/login", api.login)
}

func (api UserAPI) signup(c echo.Context) error {
	user := model.User{}

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, 
		echo.Map{"error": "invalid data user"})
	}
	

	err := api.userService.signup(&user)

	if err != nil{
		return c.JSON(http.StatusInternalServerError, 
		echo.Map{"error" : "Failed to create user"})
	} 

	return c.JSON(http.StatusCreated, user )

}


/*
func (api UserAPI) login(e echo.Echo) error{

}

*/