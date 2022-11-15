package main

import (
	cfg "edufund/internal/config"
	usercontroller "edufund/src/infrastructure/restful/controllers/users"
	"edufund/src/infrastructure/router"
	"edufund/src/services/users"
	container "edufund/src/shared/di"
	"edufund/src/shared/tracing"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
	"log"
)

func main() {
	log.Println("Starting Restful Server")

	config := cfg.GetConfig()

	fmt.Println(config)
	e := echo.New()
	ctn := container.NewContainer()

	tracer, closer := tracing.Init("edufund")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	Apply(e, ctn)
	svcPort := config.Server.Rest.Port

	e.Logger.Fatal(e.Start(":" + svcPort))
}

func Apply(e *echo.Echo, ctn *container.Container) {
	router.NewUsersRouter(e, usercontroller.NewUserController(ctn.Resolve("usersService").(*users.UserService)))
}
