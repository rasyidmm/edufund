package repository

import "github.com/opentracing/opentracing-go"

type UsersRepository interface {
	Register(opentracing.Span, interface{}) (interface{}, error)
	GetUserByUsername(opentracing.Span, interface{}) (interface{}, error)
}
