package db

import (
	"fmt"
	"sync"

	"gochat/env"
	"gochat/pkg/db/redis"
)

var redisonce sync.Once
var redisInstance *redis.Redis

func GetRedis() *redis.Redis {
	if redisInstance == nil {
		redisonce.Do(makeRedisInstance)
	}
	return redisInstance
}

func makeRedisInstance() {
	redisDB := redis.MakeRedisDB("tcp", getRedisAddress())
	redisInstance = &redisDB
}

func getRedisAddress() string {
	if env.IsDevEnv() {
		return fmt.Sprintf("%s:%d",
			env.GetStr(env.DB_REDIS_DEV_HOST),
			env.GetInt(env.DB_REDIS_DEV_PORT))
	}
	return env.GetStr(env.DB_REDIS_URL)
}
