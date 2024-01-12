package main

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func storeAuthTokenInRedis(authToken string) {
	conn := getRedisInstance().Get()
	defer conn.Close()
	conn, err := redis.Dial("tcp", redisAddress)
	if err != nil {
		fmt.Println("Error connecting to Redis:", err)
		return
	}
	defer conn.Close()

	_, err = conn.Do("SET", fmt.Sprintf("global:%s:limit", authToken), 100)
	if err != nil {
		fmt.Println("Error storing limit in Redis:", err)
	}

	_, err = conn.Do("SET", fmt.Sprintf("global:%s:usage", authToken), 0)
	if err != nil {
		fmt.Println("Error storing usage in Redis:", err)
	}
}

func getRedisValues(limitKey, usageKey string) (int, int, error) {
	conn := getRedisInstance().Get()
	defer conn.Close()
	conn, err := redis.Dial("tcp", redisAddress)
	if err != nil {
		return 0, 0, err
	}
	defer conn.Close()

	limit, err := redis.Int(conn.Do("GET", limitKey))
	if err != nil {
		return 0, 0, err
	}

	usage, err := redis.Int(conn.Do("GET", usageKey))
	if err != nil {
		return 0, 0, err
	}

	return limit, usage, nil
}

func incrementUsage(usageKey string) {
	conn := getRedisInstance().Get()
	defer conn.Close()
	conn, err := redis.Dial("tcp", redisAddress)
	if err != nil {
		fmt.Println("Error connecting to Redis:", err)
		return
	}
	defer conn.Close()

	_, err = conn.Do("INCR", usageKey)
	if err != nil {
		fmt.Println("Error incrementing usage in Redis:", err)
	}
}
