package models

import "log"

// 学生模型，储存学号和姓名供首次登录验证
type StudentModel struct {
	StudentId string `xorm:"pk varchar(11) notnull unique"`
	Name      string `xorm:"varchar(30) notnull"`
}

func (StudentModel) TableName() string {
	return "students"
}

func (model *StudentModel) CreateTable() {
	if has, err := db.IsTableExist(model); err != nil {
		log.Println(err.Error())
	} else if !has {
		db.Charset("utf-8")
		db.Sync2(model)
		db.CreateIndexes(model)
	}
}
