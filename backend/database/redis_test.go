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

func makeTestRedisPool() *redis.RedisConnection {
	conn := GetRedisConnection()
	redis.FLUSH(conn)
	return conn
}

func TestGetSet(t *testing.T) {
	rpool := makeTestRedisPool()

	key := "state"
	value := "wabalasladk"

	redis.SET(rpool, key, value)
	logrus.Info("set redis key: ", key, ", value: ", value)

	ok, res := redis.GET(rpool, key)
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

	if ok, _ = redis.GET(rpool, "somekey"); ok {
		t.Fatalf(`non existent key found using get`)
	} else {
		logrus.Info("non existent key not found using get")
	}
}

func TestExist(t *testing.T) {
	rpool := makeTestRedisPool()

	key := "state"
	value := "value"

	redis.SET(rpool, key, value)
	logrus.Info("set redis key")

	if exist := redis.EXIST(rpool, key); !exist {
		t.Fatalf(`key not found`)
	} else {
		logrus.Info("key found")
	}

	if exist := redis.EXIST(rpool, "somekey"); exist {
		t.Fatalf(`non existent key found`)
	} else {
		logrus.Info("non existent key not found")
	}
}

func TestDelete(t *testing.T) {
	rpool := makeTestRedisPool()

	key := "state"
	value := "value"

	redis.SET(rpool, key, value)
	logrus.Info("set redis key")

	redis.DEL(rpool, key)
	logrus.Info("key deleted")

	if exist := redis.EXIST(rpool, key); exist {
		t.Fatalf(`key still exist`)
	} else {
		logrus.Info("key not found")
	}
}

func TestSetFor(t *testing.T) {
	rpool := makeTestRedisPool()

	key := "state"
	value := "value"

	redis.SETEX(rpool, key, 1, value)
	logrus.Info("set redis key for 1 second")

	time.Sleep(time.Second / 2)
	if exist := redis.EXIST(rpool, key); !exist {
		t.Fatalf(`key not found`)
	} else {
		logrus.Info("key found")
	}

	time.Sleep(time.Second / 2)
	if exist := redis.EXIST(rpool, key); exist {
		t.Fatalf(`key still exist`)
	} else {
		logrus.Info("key not found")
	}
}
