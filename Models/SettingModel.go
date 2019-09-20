package models

import (
	"log"
)

type SettingModel struct {
	Id    int `xorm:"pk autoincr"`
	Desc  string
	Value int
}

func (SettingModel) TableName() string {
	return "settings"
}

func (model *SettingModel) CreateTable() {
	if has, err := db.IsTableExist(model); err != nil {
		log.Println(err.Error())
	} else if !has {
		db.Charset("utf-8")
		db.Sync2(model)
		db.CreateIndexes(model)
		setting := SettingModel{
			Desc:  "NowRound",
			Value: 1,
		}
		db.Insert(setting)
	}
}
