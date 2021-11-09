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
	conn := r.pool.Get()
	defer conn.Close()

	_, err := conn.Do("FLUSHALL")
	if err != nil {
		logrus.Panic("Cant flush redis. ", err.Error())
	}
}

func (r *Redis) SET(key string, val string) {
	conn := r.pool.Get()
	defer conn.Close()

	_, err := conn.Do("SET", key, val)
	if err != nil {
		logrus.Panic("Cant set redis. ", err.Error())
	}
}

func (r *Redis) GET(key string) (bool, string) {
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

func (r *Redis) EXIST(key string) bool {
	conn := r.pool.Get()
	defer conn.Close()

	result, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		logrus.Panic("cant check redis exists. ", err.Error())
	}
	return result
}
