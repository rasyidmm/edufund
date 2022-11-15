package repository

import (
	"edufund/src/adapter/db/model"
	"edufund/src/domain/entity"
	"edufund/src/shared/tracing"
	"edufund/src/shared/util"
	"errors"
	"github.com/opentracing/opentracing-go"
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

func (d *UsersDataHandler) Register(span opentracing.Span, in interface{}) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "Register")
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	reqData := in.(*entity.RegisterRequest)

	data := &model.UsersModel{
		Fullname: reqData.Fullname,
		Username: reqData.Username,
		Password: util.HashSha512(reqData.Password),
	}

	err := d.db.Debug().Create(&data).Error
	if err != nil {
		tracing.LogError(sp, err)
		return nil, err
	}

	res := &entity.RegisterResponse{
		StatusCode: "00",
		StatusDesc: "Register success",
	}

	tracing.LogResponse(sp, res)
	return res, nil
}

func (d *UsersDataHandler) GetUserByUsername(span opentracing.Span, in interface{}) (interface{}, error) {
	sp := tracing.CreateSubChildSpan(span, "GetUserByUsername")
	defer sp.Finish()
	tracing.LogRequest(sp, in)

	reqData := in.(*entity.UserRequest)
	var data *model.UsersModel

	err := d.db.Debug().Where("username = ?", reqData.Username).First(&data).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		tracing.LogError(sp, err)
		return nil, err
	}

	res := &entity.UserResponse{
		Fullname: data.Fullname,
		Username: data.Username,
		Password: data.Password,
	}

	tracing.LogResponse(sp, res)
	return res, nil
}
