package connection

import (
	dbConfig "edufund/internal/config/db"
	"errors"
)

var (
	instanceDb map[string]DbDriver = make(map[string]DbDriver)
)

type DbDriver interface {
	Db() interface{}
}

func NewDbDriver(config dbConfig.Database) (DbDriver, error) {
	var (
		err    error
		dnName = config.Dbname
	)
	switch config.Adapter {
	case "mysql":
		dbConn, errMysql := NewDriverMysql(config)
		if errMysql != nil {
			err = errMysql
		}
		instanceDb[dnName] = dbConn
	default:
		err = errors.New("connection Database Failed . ")
	}
	return instanceDb[dnName], err
}
