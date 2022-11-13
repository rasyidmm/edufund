package dto

import "github.com/stretchr/testify/mock"

type MockUserDto struct {
	mock.Mock
}

func (m *MockUserDto) RegisterResponse(in interface{}) (interface{}, error) {
	call := m.Called(in)
	res := call.Get(0)
	if res == nil {
		return nil, call.Error(1)
	}
	return call.Get(0), nil
}
func (m *MockUserDto) LoginResponse(in interface{}) (interface{}, error) {
	call := m.Called(in)
	res := call.Get(0)
	if res == nil {
		return nil, call.Error(1)
	}
	return call.Get(0), nil
}
