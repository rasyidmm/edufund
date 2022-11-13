package users

import (
	"edufund/src/domain/entity"
	"edufund/src/domain/repository"
	"edufund/src/shared/jwt_builder"
	"edufund/src/shared/util"
	"errors"
	"github.com/mitchellh/mapstructure"
)

type RegisterRequest struct {
	Fullname    string
	Username    string
	Password    string
	TryPassword string
}

type RegisterResponse struct {
	StatusCode string
	StatusDesc string
}

type LoginRequest struct {
	Username string
	Password string
}

type LoginResponse struct {
	Fullname string
	Username string
	Token    string
	Expired  string
}
type UserService struct {
	repo repository.UsersRepository
	out  UsersOutputPort
}

func NewUserService(r repository.UsersRepository, o UsersOutputPort) *UserService {
	return &UserService{
		repo: r,
		out:  o,
	}
}

func (s *UserService) Register(in interface{}) (interface{}, error) {

	reqData := in.(*RegisterRequest)

	var (
		request *entity.RegisterRequest
		out     *RegisterResponse
	)

	//TODO:Checker len FullName
	if len(reqData.Fullname) < 2 {
		return nil, errors.New("Name should be 2 characters or more")
	}

	//TODO: UsernameValidation
	if !util.IsEmailValid(reqData.Username) {

		return nil, errors.New("Please provide a valid email address")
	}

	//TODO:Checker len Password
	if len(reqData.Password) < 12 {

		return nil, errors.New("Password should be at least 12 characters long")
	}

	//TODO: Checker Password not match
	if len(reqData.Password) == len(reqData.TryPassword) {
		for i := 0; i < len(reqData.Password); i++ {
			if reqData.Password[i] != reqData.TryPassword[i] {
				return nil, errors.New("Confirmation password does not match")
			}
		}
	} else {
		return nil, errors.New("Confirmation password does not match")
	}

	err := mapstructure.Decode(reqData, &request)
	if err != nil {
		return nil, err
	}

	res, err := s.repo.Register(request)
	if err != nil {
		return nil, err
	}

	err = mapstructure.Decode(res, &out)
	if err != nil {
		return nil, err
	}

	return s.out.RegisterResponse(out)
}

func (s *UserService) Login(in interface{}) (interface{}, error) {

	reqData := in.(*LoginRequest)

	//TODO: UsernameValidation
	if !util.IsEmailValid(reqData.Username) {
		return nil, errors.New("Please provide a valid email address")
	}

	//TODO:Checker len Password
	if len(reqData.Password) < 12 {
		return nil, errors.New("Password should be at least 12 characters long")
	}

	reqGetUser := &entity.UserRequest{Username: reqData.Username}
	res, err := s.repo.GetUserByUsername(reqGetUser)
	if err != nil {
		return nil, err
	}
	outUser := res.(*entity.UserResponse)

	//TODO: Password Validation
	if util.HashSha512(reqData.Password) != outUser.Password {
		return nil, errors.New("Invalid username / password")
	}

	//TODO: JwtToken Builder
	reqJwt := &jwt_builder.TokenRequest{
		Username: reqData.Username,
		Fullname: reqData.Username,
	}

	resJwt, err := jwt_builder.CreateJwtToken(reqJwt)
	if err != nil {
		return nil, err
	}

	outJwt := resJwt.(*jwt_builder.TokenResponse)

	out := LoginResponse{
		Fullname: outUser.Fullname,
		Username: outUser.Username,
		Token:    outJwt.Token,
		Expired:  outJwt.Expired,
	}

	return out, nil
}
