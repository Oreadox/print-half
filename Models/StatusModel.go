package models

// 服务器统计信息
type StatusModel struct {
	Id    int    `xorm:"pk autoincr"`
	Desc  string `xorm:"notnull"`
	Value int
}

func (StatusModel) TableName() string {
	return "status"
}

func (model *StatusModel) CreateTable() {
	sql := "TRUNCATE TABLE status"
	db.Exec(sql)
	//db.DropTables(model)
	//db.Charset("utf8")
	//db.Sync2(model)
	//db.CreateIndexes(model)
	onlineNum := StatusModel{
		Desc:  "在线人数",
		Value: 0,
	}
	playingNum := StatusModel{
		Desc:  "游戏人数",
		Value: 0,
	}
	db.Insert(onlineNum, playingNum)
}
