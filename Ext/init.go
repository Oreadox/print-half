package ext

import (
	"PrintHalf/Config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func init() {
	var err error
	db, err = xorm.NewEngine("mysql", config.DbUri)
	if err != nil {
		println(err.Error())
	}
}
