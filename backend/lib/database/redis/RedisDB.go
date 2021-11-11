package redis

import (
	"github.com/gomodule/redigo/redis"
	"github.com/sirupsen/logrus"
)

type Redis struct {
	network string
	url     string
	pool    *redis.Pool
}

func MakeRedisDB(network string, url string) Redis {
	rPool := redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(network, url)
			if err != nil {
				logrus.Panic("Cant establish redis connection. ", err.Error())
			}
			return c, err
		},
	}
	return Redis{
		network: network,
		url:     url,
		pool:    &rPool,
	}
}

func (r *Redis) FLUSH() {
	_, err := doOnce(r, "FLUSHALL")
	if err != nil {
		logrus.Panic("Cant flush redis. ", err.Error())
	}
}

func (r *Redis) DEL(key string) {
	_, err := doOnce(r, "DEL", key)
	if err != nil {
		logrus.Panic("Cant del redis. ", err.Error())
	}
}

func (r *Redis) SET(key string, val string) {
	_, err := doOnce(r, "SET", key, val)
	if err != nil {
		logrus.Panic("Cant set redis. ", err.Error())
	}
}

func (r *Redis) SETEX(key string, exp int, val string) {
	_, err := doOnce(r, "SETEX", key, exp, val)
	if err != nil {
		logrus.Panic("Cant set redis. ", err.Error())
	}
}

func (r *Redis) GET(key string) (bool, string) {
	result, err := redis.String(doOnce(r, "GET", key))
	if err != nil {
		if err == redis.ErrNil {
			return false, ""
		}
		logrus.Panic("unknown get redis error. ", err.Error())
	}
	return true, result
}

func (r *Redis) EXIST(key string) bool {
	result, err := redis.Bool(doOnce(r, "EXISTS", key))
	if err != nil {
		logrus.Panic("cant check redis exists. ", err.Error())
	}
	return result
}

func doOnce(r *Redis, cmd string, args ...interface{}) (interface{}, error) {
	conn := r.pool.Get()
	defer conn.Close()
	return conn.Do(cmd, args...)
}
