package redis

import (
	"fmt"
	"log/slog"

	"github.com/Mubasshira-Tazrin/rate-limiting-api/internal/constants"
	"github.com/gomodule/redigo/redis"
)

// StoreAuthTokenInRedis supposed to do something.
// Unused? - Verify the purpose or utilization of this function.
func StoreAuthTokenInRedis(authToken string, l *slog.Logger) {
	conn, err := redis.Dial("tcp", constants.RedisAddress)
	if err != nil {
		l.Error("error connecting to Redis", "err", err)
		return
	}

	defer conn.Close()

	if _, err = conn.Do(constants.RedisCmdSet, fmt.Sprintf("global:%s:limit", authToken), 100); err != nil {
		l.Error("error storing limit in Redis", "err", err)
	}

	if _, err = conn.Do(constants.RedisCmdSet, fmt.Sprintf("global:%s:usage", authToken), 0); err != nil {
		l.Error("error storing usage in Redis", "err", err)
	}
}

func GetRedisValues(limitKey, usageKey string) (int, int, error) {
	conn, err := redis.Dial("tcp", constants.RedisAddress)

	if err != nil {
		return 0, 0, fmt.Errorf("unable to connect redis server:%w", err)
	}

	defer conn.Close()

	limit, err := redis.Int(conn.Do(constants.RedisCmdGet, limitKey))
	if err != nil {
		return 0, 0, err
	}

	usage, err := redis.Int(conn.Do(constants.RedisCmdGet, usageKey))
	if err != nil {
		return 0, 0, fmt.Errorf("unable to decode redis cmd reply:%w", err)
	}

	return limit, usage, nil
}

func IncrementUsage(usageKey string, l *slog.Logger) {
	conn, err := redis.Dial("tcp", constants.RedisAddress)
	if err != nil {
		l.Error("error connecting to Redis", "err", err)
		return
	}

	defer conn.Close()

	if _, err = conn.Do(constants.RedisCmdIncrement, usageKey); err != nil {
		l.Error("error incrementing usage in Redis", "err", err)
	}
}
