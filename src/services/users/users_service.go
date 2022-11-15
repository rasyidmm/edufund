package users

import (
	"context"
	"edufund/src/domain/entity"
	"edufund/src/domain/repository"
	"edufund/src/shared/enum"
	"edufund/src/shared/jwt_builder"
	"edufund/src/shared/tracing"
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

func (s *UserService) Register(ctx context.Context, in interface{}) (interface{}, error) {
	sp := tracing.CreateChildSpan(ctx, string(enum.StartService))
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	reqData := in.(*RegisterRequest)

	var (
		request *entity.RegisterRequest
		out     *RegisterResponse
	)

	//TODO:Checker len FullName
	if len(reqData.Fullname) < 2 {
		tracing.LogError(sp, errors.New("Name should be 2 characters or more"))
		return nil, errors.New("Name should be 2 characters or more")
	}

	//TODO: UsernameValidation
	if !util.IsEmailValid(reqData.Username) {
		tracing.LogError(sp, errors.New("Please provide a valid email address"))
		return nil, errors.New("Please provide a valid email address")
	}

	//TODO:Checker len Password
	if len(reqData.Password) < 12 {
		tracing.LogError(sp, errors.New("Password should be at least 12 characters long"))
		return nil, errors.New("Password should be at least 12 characters long")
	}

	//TODO: Checker Password not match
	if len(reqData.Password) == len(reqData.TryPassword) {
		for i := 0; i < len(reqData.Password); i++ {
			if reqData.Password[i] != reqData.TryPassword[i] {
				tracing.LogError(sp, errors.New("Confirmation password does not match"))
				return nil, errors.New("Confirmation password does not match")
			}
		}
	} else {
		tracing.LogError(sp, errors.New("Confirmation password does not match"))
		return nil, errors.New("Confirmation password does not match")
	}

	err := mapstructure.Decode(reqData, &request)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	reqGetUser := &entity.UserRequest{Username: reqData.Username}
	resUser, err := s.repo.GetUserByUsername(sp, reqGetUser)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}
	outUser := resUser.(*entity.UserResponse)
	if outUser.Username != "" {
		tracing.LogError(sp, errors.New("User Existing"))
		return nil, errors.New("User Existing")
	}

	res, err := s.repo.Register(sp, request)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	err = mapstructure.Decode(res, &out)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	tracing.LogResponse(sp, out)
	return s.out.RegisterResponse(out)
}

func (s *UserService) Login(ctx context.Context, in interface{}) (interface{}, error) {
	sp := tracing.CreateChildSpan(ctx, string(enum.StartService))
	defer sp.Finish()
	tracing.LogRequest(sp, in)
	reqData := in.(*LoginRequest)

	//TODO: UsernameValidation
	if !util.IsEmailValid(reqData.Username) {
		tracing.LogError(sp, errors.New("Please provide a valid email address"))
		return nil, errors.New("Please provide a valid email address")
	}

	//TODO:Checker len Password
	if len(reqData.Password) < 12 {
		tracing.LogError(sp, errors.New("Password should be at least 12 characters long"))
		return nil, errors.New("Password should be at least 12 characters long")
	}

	reqGetUser := &entity.UserRequest{Username: reqData.Username}
	res, err := s.repo.GetUserByUsername(sp, reqGetUser)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}
	outUser := res.(*entity.UserResponse)

	//TODO: Password Validation
	if util.HashSha512(reqData.Password) != outUser.Password {
		tracing.LogError(sp, errors.New("Invalid username / password"))
		return nil, errors.New("Invalid username / password")
	}

	//TODO: JwtToken Builder
	reqJwt := &jwt_builder.TokenRequest{
		Username: reqData.Username,
		Fullname: outUser.Fullname,
	}

	resJwt, err := jwt_builder.CreateJwtToken(sp, reqJwt)
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	outJwt := resJwt.(*jwt_builder.TokenResponse)

	out := LoginResponse{
		Fullname: outUser.Fullname,
		Username: outUser.Username,
		Token:    outJwt.Token,
		Expired:  outJwt.Expired,
	}

	tracing.LogResponse(sp, out)
	return out, nil
}
