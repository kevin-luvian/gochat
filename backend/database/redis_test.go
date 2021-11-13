package database

import (
	"gochat/env"
	"gochat/lib/database/redis"
	"testing"
	"time"

	"github.com/sirupsen/logrus"
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
		t.Fatalf(`non existent key found using get`)
	} else {
		logrus.Info("non existent key not found using get")
	}
}

func TestExist(t *testing.T) {
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

func TestDelete(t *testing.T) {
	redisdb := makeTestRedisDB()

	key := "state"
	value := "value"

	redisdb.SET(key, value)
	logrus.Info("set redis key")

	redisdb.DEL(key)
	logrus.Info("key deleted")

	if exist := redisdb.EXIST(key); exist {
		t.Fatalf(`key still exist`)
	} else {
		logrus.Info("key not found")
	}
}

func TestSetFor(t *testing.T) {
	redisdb := makeTestRedisDB()

	key := "state"
	value := "value"

	redisdb.SETEX(key, 1, value)
	logrus.Info("set redis key for 1 second")

	time.Sleep(time.Second / 4)
	if exist := redisdb.EXIST(key); !exist {
		t.Fatalf(`key not found`)
	} else {
		logrus.Info("key found")
	}

	time.Sleep(time.Second)
	if exist := redisdb.EXIST(key); exist {
		t.Fatalf(`key still exist`)
	} else {
		logrus.Info("key not found")
	}
}
