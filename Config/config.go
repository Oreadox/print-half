package config

import (
	"os"
	"time"
)

var (
	DBUri                     = "user:user@/db1?charset=utf8"
	Debug                     = false
	SECRET_KEY                = GetEnviron("SECRET_KEY", "viWxab60cncwhkpetBlwmmpi")
	ExpiresTime time.Duration = 24 * 7 // token过期时间，单位为小时
)

type ConfigModel struct {
	DBUri       string
	Debug       bool
	SECRET_KEY  string
	ExpiresTime time.Duration
}

func GetConfig() ConfigModel {
	return config
}

func GetEnviron(key, define string) string {
	value := os.Getenv(key)
	if value == "" {
		value = define
	}
	return value
}
