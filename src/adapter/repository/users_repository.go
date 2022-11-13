package repository

import (
	"edufund/src/adapter/db/model"
	"edufund/src/domain/entity"
	"edufund/src/shared/util"
	"errors"
	"gorm.io/gorm"
)

type UsersDataHandler struct {
	db *gorm.DB
}

func NewUserDataHandler(db *gorm.DB) *UsersDataHandler {
	return &UsersDataHandler{
		db: db,
	}
}

func (d *UsersDataHandler) Register(in interface{}) (interface{}, error) {

	reqData := in.(*entity.RegisterRequest)

	data := &model.UsersModel{
		Fullname: reqData.Fullname,
		Username: reqData.Username,
		Password: util.HashSha512(reqData.Password),
	}

	err := d.db.Debug().Create(&data).Error
	if err != nil {
		return nil, err
	}

	res := &entity.RegisterResponse{
		StatusCode: "00",
		StatusDesc: "Register success",
	}

	return res, nil
}

func (d *UsersDataHandler) GetUserByUsername(in interface{}) (interface{}, error) {

	reqData := in.(*entity.UserRequest)
	var data *model.UsersModel

	err := d.db.Debug().Where("username = ?", reqData.Username).First(&data).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	res := &entity.UserResponse{
		Fullname: data.Fullname,
		Username: data.Username,
		Password: data.Password,
	}

	return res, nil
}
