package redis

import (
	"fmt"
	"sync"

	"gochat/env"
)

var redisonce sync.Once
var redisInstance *Redis

func GetRedis() *Redis {
	if redisInstance == nil {
		redisonce.Do(makeRedisInstance)
	}
	return redisInstance
}

func makeRedisInstance() {
	redisDB := MakeRedisDB("tcp", getRedisAddress())
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
