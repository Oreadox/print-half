package models

import (
	"log"
	"time"
)

type PictureModel struct {
	Id             int `xorm:"pk autoincr"`
	UserId1        int
	UserId2        int
	TopFileName    string    `xorm:"varchar(128)"`
	BottomFileName string    `xorm:"varchar(128)"`
	CreateTime     time.Time `xorm:"datetime created notnull"`
	LikeNum        int       `xorm:"default 0"`
}

func (PictureModel) TableName() string {
	return "pictures"
}

func (model *PictureModel) CreateTable() {
	if has, err := db.IsTableExist(model); err != nil {
		log.Println(err.Error())
	} else if !has {
		db.Charset("utf-8")
		db.Sync2(model)
		db.CreateIndexes(model)
	}
}
