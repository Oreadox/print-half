package models

import "log"

type QuestionModel struct {
	Id   int    `xorm:"pk autoincr"`
	Name string `xorm:"varchar(30) notnull"`
}

func (QuestionModel) TableName() string {
	return "questions"
}

func (model *QuestionModel) CreateTable() {
	if has, err := db.IsTableExist(model); err != nil {
		log.Println(err.Error())
	} else if !has {
		db.Charset("utf-8")
		db.Sync2(model)
		db.CreateIndexes(model)
	}
}
