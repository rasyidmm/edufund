package connection

import (
	"edufund/internal/config"
	"edufund/internal/config/db"
	"edufund/src/adapter/db/connection"
	"fmt"
	"gorm.io/gorm"
)

var Edufund *gorm.DB

func init() {
	var err error
	cfg := config.GetConfig()

	Edufund, err = DbConnection(cfg.Database.Edufund.Mysql)
	if err != nil {
		fmt.Println("Error in connection to database Edufund : ", err)
	}
}
func DbConnection(db db.Database) (*gorm.DB, error) {
	drive, err := connection.NewDbDriver(db)
	if err != nil {
		return nil, err
	}
	return drive.Db().(*gorm.DB), nil
}
