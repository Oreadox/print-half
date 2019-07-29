package models

import (
	"log"
	"time"
)

type UserModel struct {
	Id         int       `xorm:"pk autoincr"`
	StudentId  string    `xorm:"varchar(11) notnull unique"`
	Name       string    `xorm:"varchar(30) notnull"`
	CreateTime time.Time `xorm:"datetime created notnull"`
}

func (UserModel) TableName() string {
	return "users"
}

func (model *UserModel) CreateTable() {
	if has, err := db.IsTableExist(model); err != nil {
		log.Println(err.Error())
	} else if !has {
		db.Charset("utf8")
		db.Sync2(model)
		db.CreateIndexes(model)
	}
}
