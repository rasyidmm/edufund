package users

import (
	services "edufund/src/services/users"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
	"net/http"
)

type RegisterRequest struct {
	Fullname    string `json:"fullname"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	TryPassword string `json:"try_password"`
}

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type UserController struct {
	service services.UsersInputPort
}

func NewUserController(s services.UsersInputPort) *UserController {
	return &UserController{
		service: s,
	}
}

func (con *UserController) Register(c echo.Context) error {
	reqData := new(RegisterRequest)

	err := c.Bind(reqData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var request *services.RegisterRequest
	err = mapstructure.Decode(reqData, &request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	res, err := con.service.Register(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func (con *UserController) Login(c echo.Context) error {
	reqData := new(LoginRequest)

	err := c.Bind(reqData)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var request *services.LoginRequest
	err = mapstructure.Decode(reqData, &request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	res, err := con.service.Login(request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}
