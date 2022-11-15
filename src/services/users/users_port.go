package users

import "context"

type UsersInputPort interface {
	Register(context.Context, interface{}) (interface{}, error)
	Login(context.Context, interface{}) (interface{}, error)
}

type UsersOutputPort interface {
	RegisterResponse(interface{}) (interface{}, error)
	LoginResponse(interface{}) (interface{}, error)
}
