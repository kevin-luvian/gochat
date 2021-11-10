package redis

import (
	"github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"
)

type RedisConnection struct {
	pool *redis.Pool
}

func MakeRedisDBPool(network string, url string) *RedisConnection {
	return &RedisConnection{
		pool: &redis.Pool{
			MaxIdle:   80,
			MaxActive: 12000,
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial(network, url)
				if err != nil {
					logrus.Panic("Cant establish redis connection. ", err.Error())
				}
				return c, err
			},
		}}
}

func FLUSH(r *RedisConnection) {
	conn := r.pool.Get()
	defer conn.Close()

	_, err := conn.Do("FLUSHALL")
	if err != nil {
		logrus.Panic("Cant flush redis. ", err.Error())
	}
}

func DEL(r *RedisConnection, key string) {
	conn := r.pool.Get()
	defer conn.Close()

	_, err := conn.Do("DEL", key)
	if err != nil {
		logrus.Panic("Cant del redis. ", err.Error())
	}
}

func SET(r *RedisConnection, key string, val string) {
	conn := r.pool.Get()
	defer conn.Close()

	_, err := conn.Do("SET", key, val)
	if err != nil {
		logrus.Panic("Cant set redis. ", err.Error())
	}
}

func SETEX(r *RedisConnection, key string, exp int, val string) {
	conn := r.pool.Get()
	defer conn.Close()

	_, err := conn.Do("SETEX", key, exp, val)
	if err != nil {
		logrus.Panic("Cant set redis. ", err.Error())
	}
}

func GET(r *RedisConnection, key string) (bool, string) {
	conn := r.pool.Get()
	defer conn.Close()

	result, err := redis.String(conn.Do("GET", key))
	if err != nil {
		if err == redis.ErrNil {
			return false, ""
		}
		logrus.Panic("unknown get redis error. ", err.Error())
	}
	return true, result
}

func EXIST(r *RedisConnection, key string) bool {
	conn := r.pool.Get()
	defer conn.Close()

	result, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		logrus.Panic("cant check redis exists. ", err.Error())
	}
	return result
}
