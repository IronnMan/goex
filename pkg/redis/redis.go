package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"goex/pkg/logger"
	"sync"
	"time"
)

// RedisClient Redis Serve
type RedisClient struct {
	Client  *redis.Client
	Context context.Context
}

// once make sure that the global Redis object is only instantiated once
var once sync.Once

var Redis *RedisClient

func ConnectRedis(address string, username string, password string, db int) {
	once.Do(func() {
		Redis = NewClient(address, username, password, db)
	})
}

func NewClient(address string, username string, password string, db int) *RedisClient {
	// initialize a custom RedisClient instance
	rds := &RedisClient{}
	// use default context
	rds.Context = context.Background()

	rds.Client = redis.NewClient(&redis.Options{
		Addr:     address,
		Username: username,
		Password: password,
		DB:       db,
	})

	// test connect
	err := rds.Ping()
	logger.LogIf(err)

	return rds
}

// Ping used to test whether the redis connection is normal
func (rds RedisClient) Ping() error {
	_, err := rds.Client.Ping(rds.Context).Result()
	return err
}

// Set store the value corresponding to the key, and set the expiration time
func (rds RedisClient) Set(key string, value interface{}, expiration time.Duration) bool {
	if err := rds.Client.Set(rds.Context, key, value, expiration).Err(); err != nil {
		logger.ErrorString("Redis", "Set", err.Error())
		return false
	}
	return true
}

// Get the value corresponding to the key
func (rds RedisClient) Get(key string) string {
	result, err := rds.Client.Get(rds.Context, key).Result()
	if err != nil {
		if err != redis.Nil {
			logger.ErrorString("Redis", "Get", err.Error())
		}
		return ""
	}
	return result
}

// Has to determine whether a key exists, both internal errors and redis.Nil return false
func (rds RedisClient) Has(key string) bool {
	_, err := rds.Client.Get(rds.Context, key).Result()
	if err != nil {
		if err != redis.Nil {
			logger.ErrorString("Redis", "has", err.Error())
		}
		return false
	}
	return true
}

// Del delete data stored in redis, support multiple key parameters
func (rds RedisClient) Del(keys ...string) bool {
	if err := rds.Client.Del(rds.Context, keys...).Err(); err != nil {
		logger.ErrorString("Redis", "Del", err.Error())
		return false
	}
	return true
}

// FlushDB clear all data in the current redis db
func (rds RedisClient) FlushDB() bool {
	if err := rds.Client.FlushDB(rds.Context).Err(); err != nil {
		logger.ErrorString("Redis", "FlushDB", err.Error())
		return false
	}
	return true
}

// Increment when there is only one parameter, it is the key, and its value is increased by 1.
// when there are 2 parameters, the first parameter is the key, and the second parameter is the int64 type of the value to be added.
func (rds RedisClient) Increment(parameters ...interface{}) bool {
	switch len(parameters) {
	case 1:
		key := parameters[0].(string)
		if err := rds.Client.Incr(rds.Context, key).Err(); err != nil {
			logger.ErrorString("Redis", "Increment", err.Error())
			return false
		}
	case 2:
		key := parameters[0].(string)
		value := parameters[1].(int64)
		if err := rds.Client.IncrBy(rds.Context, key, value).Err(); err != nil {
			logger.ErrorString("Redis", "Increment", err.Error())
			return false
		}
	default:
		logger.ErrorString("Redis", "Increment", "too many parameters")
		return false
	}
	return true
}

func (rds RedisClient) Decrement(parameters ...interface{}) bool {
	switch len(parameters) {
	case 1:
		key := parameters[0].(string)
		if err := rds.Client.Decr(rds.Context, key).Err(); err != nil {
			logger.ErrorString("Redis", "Decrement", err.Error())
			return false
		}
	case 2:
		key := parameters[0].(string)
		value := parameters[1].(int64)
		if err := rds.Client.DecrBy(rds.Context, key, value).Err(); err != nil {
			logger.ErrorString("Redis", "Decrement", err.Error())
			return false
		}
	default:
		logger.ErrorString("Redis", "Decrement", "too many parameters")
		return false
	}
	return true
}
