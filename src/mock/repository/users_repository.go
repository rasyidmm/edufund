package repository

import (
	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Register(in interface{}) (interface{}, error) {
	call := m.Called(in)
	res := call.Get(0)
	if res == nil {
		return nil, call.Error(1)
	}
	return res, nil

}
func (m *MockUserRepository) GetUserByUsername(in interface{}) (interface{}, error) {
	call := m.Called(in)
	res := call.Get(0)
	if res == nil {
		return nil, call.Error(1)
	}
	return res, nil
}
