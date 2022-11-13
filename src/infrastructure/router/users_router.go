package router

import (
	controller "edufund/src/infrastructure/restful/controllers/users"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type UsersRouter struct {
}

func NewUsersRouter(e *echo.Echo, userController *controller.UserController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	r := e.Group("/user")
	r.POST("/register", userController.Register)
	r.POST("/login", userController.Login)
	return e
}
