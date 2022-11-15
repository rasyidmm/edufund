package users

import (
	services "edufund/src/services/users"
	"edufund/src/shared/tracing"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
	"github.com/opentracing/opentracing-go"
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
	sp, ctx := opentracing.StartSpanFromContext(c.Request().Context(), "Register")
	defer sp.Finish()
	reqData := new(RegisterRequest)

	err := c.Bind(reqData)
	if err != nil {
		tracing.LogError(sp, err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	tracing.LogRequest(sp, reqData)

	var request *services.RegisterRequest
	err = mapstructure.Decode(reqData, &request)
	if err != nil {
		tracing.LogError(sp, err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	res, err := con.service.Register(ctx, request)
	if err != nil {
		tracing.LogError(sp, err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	tracing.LogResponse(sp, res)
	return c.JSON(http.StatusOK, res)
}

func (con *UserController) Login(c echo.Context) error {
	sp, ctx := opentracing.StartSpanFromContext(c.Request().Context(), "Login")
	defer sp.Finish()
	reqData := new(LoginRequest)

	err := c.Bind(reqData)
	if err != nil {
		tracing.LogError(sp, err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	tracing.LogRequest(sp, reqData)

	var request *services.LoginRequest
	err = mapstructure.Decode(reqData, &request)
	if err != nil {
		tracing.LogError(sp, err)
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	res, err := con.service.Login(ctx, request)
	if err != nil {
		tracing.LogError(sp, err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	tracing.LogResponse(sp, res)
	return c.JSON(http.StatusOK, res)
}
