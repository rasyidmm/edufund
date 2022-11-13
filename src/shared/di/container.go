package di

import (
	"edufund/src/adapter/repository"
	"edufund/src/infrastructure/connection"
	"edufund/src/infrastructure/restful/controllers/dto"
	"edufund/src/services/users"
	"github.com/sarulabs/di/v2"
)

type Container struct {
	ctn di.Container
}

func NewContainer() *Container {
	builder, _ := di.NewBuilder()
	_ = builder.Add([]di.Def{
		{Name: "usersService", Build: usersService},
	}...)
	return &Container{
		ctn: builder.Build(),
	}
}

func (c *Container) Resolve(name string) interface{} {
	return c.ctn.Get(name)
}
func usersService(_ di.Container) (interface{}, error) {
	repo := repository.NewUserDataHandler(connection.Edufund)
	out := &dto.UsersBuilder{}
	return users.NewUserService(repo, out), nil
}
