package users

type UsersInputPort interface {
	Register(interface{}) (interface{}, error)
	Login(interface{}) (interface{}, error)
}

type UsersOutputPort interface {
	RegisterResponse(interface{}) (interface{}, error)
	LoginResponse(interface{}) (interface{}, error)
}
