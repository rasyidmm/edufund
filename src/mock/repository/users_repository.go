package repository

import (
	"github.com/opentracing/opentracing-go"
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Register(span opentracing.Span, in interface{}) (interface{}, error) {
	call := m.Called(in)
	res := call.Get(0)
	if res == nil {
		return nil, call.Error(1)
	}
	return res, nil

}
func (m *MockUserRepository) GetUserByUsername(span opentracing.Span, in interface{}) (interface{}, error) {
	call := m.Called(in)
	res := call.Get(0)
	if res == nil {
		return nil, call.Error(1)
	}
	return res, nil
}
