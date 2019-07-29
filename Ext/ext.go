package ext

import "github.com/go-xorm/xorm"

var (
	db *xorm.Engine
)

func GetEngine() *xorm.Engine {
	return db
}
