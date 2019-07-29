package config

var config ConfigModel

func init() {
	config.DBUri = DBUri
	config.Debug = Debug
	config.SECRET_KEY = SECRET_KEY
	config.ExpiresTime = ExpiresTime
}
