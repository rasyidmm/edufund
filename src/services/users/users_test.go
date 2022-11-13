package users

import (
	"edufund/src/domain/entity"
	"edufund/src/mock/dto"
	"edufund/src/mock/repository"
	"edufund/src/shared/util"
	"errors"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestUserService_Register(t *testing.T) {
	mockDto := &dto.MockUserDto{}
	mockRepo := &repository.MockUserRepository{}
	reqEntity := &entity.RegisterRequest{
		Fullname: "ratata",
		Username: "ratata@gmail.com",
		Password: "123412341234",
	}
	reqS := &RegisterRequest{
		Fullname:    "ratata",
		Username:    "ratata@gmail.com",
		Password:    "123412341234",
		TryPassword: "123412341234",
	}
	resS := &RegisterResponse{
		StatusCode: "00",
		StatusDesc: "Transaction success",
	}
	resEntity := &entity.RegisterResponse{
		StatusCode: "00",
		StatusDesc: "Transaction success",
	}

	// negative Fullname less than 2
	reqSF2 := &RegisterRequest{
		Fullname:    "r",
		Username:    "ratata@gmail.com",
		Password:    "123412341234",
		TryPassword: "123412341234",
	}
	reqEntityf2 := &entity.RegisterRequest{
		Fullname: "r",
		Username: "ratata@gmail.com",
		Password: "123412341234",
	}
	// negative Please provide a valid email address
	reqSUV2 := &RegisterRequest{
		Fullname:    "ratata",
		Username:    "ratatagmail.com",
		Password:    "123412341234",
		TryPassword: "123412341234",
	}
	reqEntityUV2 := &entity.RegisterRequest{
		Fullname: "ratata",
		Username: "ratatagmail.com",
		Password: "123412341234",
	}
	// negative Password should be at least 12 characters long
	reqEntityP12 := &entity.RegisterRequest{
		Fullname: "ratata",
		Username: "ratata@gmail.com",
		Password: "1234",
	}
	reqSP12 := &RegisterRequest{
		Fullname:    "ratata",
		Username:    "ratata@gmail.com",
		Password:    "1234",
		TryPassword: "123412341234",
	}
	// negative Password Not Match A
	reqEntityPNMA := &entity.RegisterRequest{
		Fullname: "ratata",
		Username: "ratata@gmail.com",
		Password: "123412341234",
	}
	reqSPNMA := &RegisterRequest{
		Fullname:    "ratata",
		Username:    "ratata@gmail.com",
		Password:    "123412341244",
		TryPassword: "123412341234",
	}
	// negative Password Not Match B
	reqEntityPNMB := &entity.RegisterRequest{
		Fullname: "ratata",
		Username: "ratata@gmail.com",
		Password: "123412341234ckkk",
	}
	reqSPNMB := &RegisterRequest{
		Fullname:    "ratata",
		Username:    "ratata@gmail.com",
		Password:    "123412341244sss",
		TryPassword: "123412341234",
	}

	Convey("Test Service Register", t, func() {
		Convey("Positive Scenario", func() {
			Convey("When Register Success, should return data", func() {
				mockRepo.On("Register", reqEntity).Return(resEntity, nil).Once()
				mockDto.On("RegisterResponse", resS).Return(resS, nil).Once()
				sv := NewUserService(mockRepo, mockDto)
				res, err := sv.Register(reqS)
				So(err, ShouldBeNil)
				So(res, ShouldNotBeNil)
			})
		})
		Convey("Negative Scenario", func() {
			Convey("When Register Failed, should return error", func() {
				mockRepo.On("Register", reqEntity).Return(nil, errors.New("error")).Once()
				mockDto.On("RegisterResponse", resS).Return(resS, nil).Once()
				sv := NewUserService(mockRepo, mockDto)
				res, err := sv.Register(reqS)
				So(res, ShouldBeNil)
				So(err, ShouldNotBeNil)
			})
			Convey("When Confirmation password does not match B, should return error", func() {
				mockRepo.On("Register", reqEntityPNMB).Return(nil, errors.New("Confirmation password does not match")).Once()
				mockDto.On("RegisterResponse", resS).Return(resS, nil).Once()
				sv := NewUserService(mockRepo, mockDto)
				res, err := sv.Register(reqSPNMB)
				So(res, ShouldBeNil)
				So(err, ShouldNotBeNil)
			})
			Convey("When Confirmation password does not match A, should return error", func() {
				mockRepo.On("Register", reqEntityPNMA).Return(nil, errors.New("Confirmation password does not match")).Once()
				mockDto.On("RegisterResponse", resS).Return(resS, nil).Once()
				sv := NewUserService(mockRepo, mockDto)
				res, err := sv.Register(reqSPNMA)
				So(res, ShouldBeNil)
				So(err, ShouldNotBeNil)
			})
			Convey("When Password should be at least 12 characters long, should return error", func() {
				mockRepo.On("Register", reqEntityP12).Return(nil, errors.New("Please provide a valid email address")).Once()
				mockDto.On("RegisterResponse", resS).Return(resS, nil).Once()
				sv := NewUserService(mockRepo, mockDto)
				res, err := sv.Register(reqSP12)
				So(res, ShouldBeNil)
				So(err, ShouldNotBeNil)
			})
			Convey("When username not valid, should return error", func() {
				mockRepo.On("Register", reqEntityUV2).Return(nil, errors.New("Please provide a valid email address")).Once()
				mockDto.On("RegisterResponse", resS).Return(resS, nil).Once()
				sv := NewUserService(mockRepo, mockDto)
				res, err := sv.Register(reqSUV2)
				So(res, ShouldBeNil)
				So(err, ShouldNotBeNil)
			})
			Convey("When Fullname less than 2, should return error", func() {
				mockRepo.On("Register", reqEntityf2).Return(nil, errors.New("Name should be 2 characters or more")).Once()
				mockDto.On("RegisterResponse", resS).Return(resS, nil).Once()
				sv := NewUserService(mockRepo, mockDto)
				res, err := sv.Register(reqSF2)
				So(res, ShouldBeNil)
				So(err, ShouldNotBeNil)
			})

		})
	})
}

func TestUserService_Login(t *testing.T) {
	mockDto := &dto.MockUserDto{}
	mockRepo := &repository.MockUserRepository{}
	reqEntity := &entity.UserRequest{
		Username: "ratata@gmail.com",
	}
	reqS := &LoginRequest{
		Username: "ratata@gmail.com",
		Password: "123412341234",
	}
	resS := &LoginResponse{
		Fullname: "",
		Username: "",
		Token:    "",
		Expired:  "",
	}
	resEntity := &entity.UserResponse{
		Fullname: "ratata",
		Username: "ratata@gmail.com",
		Password: util.HashSha512("123412341234"),
	}
	//Invalid password / username
	reqSIPU := &LoginRequest{
		Username: "ratata@gmail.com",
		Password: "1234123412344",
	}

	resEntityIPU := &entity.UserResponse{
		Fullname: "ratata",
		Username: "ratata@gmail.com",
		Password: util.HashSha512("123412341234"),
	}

	//Password should be at least 12 characters long
	reqSP12 := &LoginRequest{
		Username: "ratata@gmail.com",
		Password: "1234123412",
	}

	resEntityP12 := &entity.UserResponse{
		Fullname: "ratata",
		Username: "ratata@gmail.com",
		Password: util.HashSha512("123412341234"),
	}

	//Please provide a valid email address
	reqSPiV := &LoginRequest{
		Username: "ratatagmail.com",
		Password: "1234123412",
	}

	Convey("Test Service Register", t, func() {
		Convey("Positive Scenario", func() {
			Convey("When Login Success, should return data", func() {
				mockRepo.On("GetUserByUsername", reqEntity).Return(resEntity, nil).Once()
				mockDto.On("LoginResponse", resS).Return(resS, nil).Once()
				sv := NewUserService(mockRepo, mockDto)
				res, err := sv.Login(reqS)
				So(err, ShouldBeNil)
				So(res, ShouldNotBeNil)
			})
		})
		Convey("Negative Scenario", func() {
			Convey("When Invalid password / username, should return error", func() {
				mockRepo.On("GetUserByUsername", reqEntity).Return(resEntityIPU, nil).Once()
				mockDto.On("LoginResponse", resS).Return(resS, nil).Once()
				sv := NewUserService(mockRepo, mockDto)
				res, err := sv.Login(reqSIPU)
				So(res, ShouldBeNil)
				So(err, ShouldNotBeNil)
			})
			Convey("When GetUserByUsername error, should return error", func() {
				mockRepo.On("GetUserByUsername", reqEntity).Return(nil, errors.New("error")).Once()
				mockDto.On("LoginResponse", resS).Return(resS, nil).Once()
				sv := NewUserService(mockRepo, mockDto)
				res, err := sv.Login(reqS)
				So(res, ShouldBeNil)
				So(err, ShouldNotBeNil)
			})
			Convey("When Password should be at least 12 characters long, should return error", func() {
				mockRepo.On("GetUserByUsername", reqEntity).Return(resEntityP12, nil).Once()
				mockDto.On("LoginResponse", resS).Return(resS, nil).Once()
				sv := NewUserService(mockRepo, mockDto)
				res, err := sv.Login(reqSP12)
				So(res, ShouldBeNil)
				So(err, ShouldNotBeNil)
			})
			Convey("When Please provide a valid email address, should return error", func() {
				mockRepo.On("GetUserByUsername", reqEntity).Return(resEntityP12, nil).Once()
				mockDto.On("LoginResponse", resS).Return(resS, nil).Once()
				sv := NewUserService(mockRepo, mockDto)
				res, err := sv.Login(reqSPiV)
				So(res, ShouldBeNil)
				So(err, ShouldNotBeNil)
			})
		})
	})
}
