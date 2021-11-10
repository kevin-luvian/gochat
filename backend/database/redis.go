package database

import (
	"fmt"
	"gochat/lib/database/redis"
	"os"
	"sync"
)

var redisonce sync.Once
var redisInstance *redis.RedisConnection

func GetRedisConnection() *redis.RedisConnection {
	if redisInstance == nil {
		redisonce.Do(makeRedisInstance)
	}
	return redisInstance
}

func makeRedisInstance() {
	redisInstance = redis.MakeRedisDBPool("tcp", getRedisAddress())
}

func getRedisAddress() string {
	if os.Getenv("ENV") == "development" {
		return fmt.Sprintf("%s:%s",
			os.Getenv("DB_REDIS_DEV_HOST"),
			os.Getenv("DB_REDIS_DEV_PORT"))
	} else {
		return os.Getenv("REDIS_URL")
	}
}
