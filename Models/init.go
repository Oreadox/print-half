package models

import ext "PrintHalf/Ext"

var db = ext.GetEngine()

func init() {
	var user UserModel
	user.CreateTable()

	var picture PictureModel
	picture.CreateTable()

	var question QuestionModel
	question.CreateTable()

	var student StudentModel
	student.CreateTable()

	var setting SettingModel
	setting.CreateTable()
}
