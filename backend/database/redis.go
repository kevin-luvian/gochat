package database

import (
	"fmt"
	libredis "gochat/lib/database/redis"
	"os"
	"sync"
)

var redisonce sync.Once
var redisInstance *libredis.Redis

func GetRedis() *libredis.Redis {
	if redisInstance == nil {
		redisonce.Do(makeRedisInstance)
	}
	return redisInstance
}

func makeRedisInstance() {
	redisDB := libredis.MakeRedisDB("tcp", getRedisAddress())
	redisInstance = &redisDB
}

func getRedisAddress() string {
	if os.Getenv("ENV") == "development" {
		return fmt.Sprintf("%s:%s",
			os.Getenv("DB_REDIS_DEV_HOST"),
			os.Getenv("DB_REDIS_DEV_PORT"))
	}
	return os.Getenv("REDIS_URL")
}
