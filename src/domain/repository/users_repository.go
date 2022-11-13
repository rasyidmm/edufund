package repository

type UsersRepository interface {
	Register(interface{}) (interface{}, error)
	GetUserByUsername(interface{}) (interface{}, error)
}
