package dto

import "github.com/mitchellh/mapstructure"

type RegisterResponse struct {
	StatusCode string `json:"status_code"`
	StatusDesc string `json:"status_desc"`
}
type LoginResponse struct {
	Fullname string `json:"fullname"`
	Username string `json:"username"`
	Token    string `json:"token"`
	Expired  string `json:"expired"`
}
type UsersBuilder struct {
}

func (b *UsersBuilder) RegisterResponse(in interface{}) (interface{}, error) {
	var out *RegisterResponse
	err := mapstructure.Decode(in, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
func (b *UsersBuilder) LoginResponse(in interface{}) (interface{}, error) {
	var out *LoginResponse
	err := mapstructure.Decode(in, &out)
	if err != nil {
		return nil, err
	}
	return out, nil
}
