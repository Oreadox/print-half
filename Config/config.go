package config

import (
	"os"
	"time"
)

var (
	DbUri                     = GetEnviron("dBUri", "user:user@/db1?charset=utf8")
	Debug                     = false
	SecretKey                 = GetEnviron("SecretKey", "viWxab60cncwhkpetBlwmmpi")
	ExpiresTime time.Duration = 24 * 7 // token过期时间，单位为小时
	// OSS相关
	AccessKeyId     = GetEnviron("AccessKeyId", "")
	AccessKeySecret = GetEnviron("AccessKeySecret", "")
	Endpoint        = GetEnviron("Endpoint", "")
	BucketName      = GetEnviron("BucketName", "")
)

func GetEnviron(key, define string) string {
	value := os.Getenv(key)
	if value == "" {
		value = define
	}
	return value
}
