package main

import (
	"sync"

	"github.com/gomodule/redigo/redis"
)

var (
	once          sync.Once
	singleRedis   *redis.Pool
	redisPoolSize = 10
)

func getRedisInstance() *redis.Pool {
	once.Do(func() {
		singleRedis = createRedisPool()
	})
	return singleRedis
}

func createRedisPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:     redisPoolSize,
		IdleTimeout: 240,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", redisAddress)
		},
	}
}
