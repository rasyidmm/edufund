package model

type UsersModel struct {
	Id       int64  `gorm:"primaryKey;unique;autoIncrement"`
	Fullname string `gorm:"column:fullname"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
}

func (UsersModel) TableName() string {
	return "users"
}
