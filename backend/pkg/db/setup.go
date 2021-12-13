package db

import (
	"gochat/pkg/db/redis"
)

func GetRedis() *redis.Redis {
	return redis.GetRedis()
}

func Setup() {
	r := GetRedis()
	r.TESTCONN()
}
