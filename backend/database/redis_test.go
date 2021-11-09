package database

import (
	"gochat/lib/database/redis"
	"testing"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
)

func init() {
	logrus.Info("Starting userDao test")
	err := godotenv.Load("../.env")

	if err != nil {
		logrus.Fatal("Error loading .env file")
	}
}

func makeTestRedisDB() *redis.Redis {
	redisdb := GetRedis()
	redisdb.FLUSH()
	return redisdb
}

func TestRedisGetSet(t *testing.T) {
	redisdb := makeTestRedisDB()

	key := "state"
	value := "wabalasladk"

	redisdb.SET(key, value)
	logrus.Info("set redis key: ", key, ", value: ", value)

	ok, res := redisdb.GET(key)
	if !ok {
		t.Fatalf(`can't get key.`)
	} else {
		logrus.Info("get redis value: ", res)
	}

	if value != res {
		t.Fatalf(`result and value doesn't match`)
	} else {
		logrus.Info("result and value match")
	}

	if ok, _ = redisdb.GET("somekey"); ok {
		t.Fatalf(`non existent key found`)
	} else {
		logrus.Info("non existent key not found")
	}
}

func TestRedisUnique(t *testing.T) {
	redisdb := makeTestRedisDB()

	key := "state"
	value := "value"

	redisdb.SET(key, value)
	logrus.Info("set redis key")

	if exist := redisdb.EXIST(key); !exist {
		t.Fatalf(`key not found`)
	} else {
		logrus.Info("key found")
	}

	if exist := redisdb.EXIST("somekey"); exist {
		t.Fatalf(`non existent key found`)
	} else {
		logrus.Info("non existent key not found")
	}
}
