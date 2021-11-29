package db

import (
	"gochat/env"
	"gochat/pkg/db/redis"
	"testing"
	"time"
)

func init() {
	env.LoadDotEnvForTest()
}

func makeTestRedisDB() *redis.Redis {
	redisdb := GetRedis()
	redisdb.FLUSH()
	return redisdb
}

func TestGetSet(t *testing.T) {
	redisdb := makeTestRedisDB()

	key := "state"
	value := "wabalasladk"

	redisdb.SET(key, value)

	ok, res := redisdb.GET(key)
	if !ok {
		t.Fatalf(`can't get key.`)
	}

	if value != res {
		t.Fatalf(`result and value doesn't match`)
	}

	if ok, _ = redisdb.GET("somekey"); ok {
		t.Fatalf(`non existent key found using get`)
	}
}

func TestExist(t *testing.T) {
	redisdb := makeTestRedisDB()

	key := "state"
	value := "value"

	redisdb.SET(key, value)

	if exist := redisdb.EXIST(key); !exist {
		t.Fatalf(`key not found`)
	}

	if exist := redisdb.EXIST("somekey"); exist {
		t.Fatalf(`non existent key found`)
	}
}

func TestDelete(t *testing.T) {
	redisdb := makeTestRedisDB()

	key := "state"
	value := "value"

	redisdb.SET(key, value)
	redisdb.DEL(key)

	if exist := redisdb.EXIST(key); exist {
		t.Fatalf(`key still exist`)
	}
}

func TestSetFor(t *testing.T) {
	redisdb := makeTestRedisDB()

	key := "state"
	value := "value"

	redisdb.SETEX(key, 1, value)

	time.Sleep(time.Second / 4)
	if exist := redisdb.EXIST(key); !exist {
		t.Fatalf(`key not found`)
	}

	time.Sleep(time.Second)
	if exist := redisdb.EXIST(key); exist {
		t.Fatalf(`key still exist`)
	}
}
