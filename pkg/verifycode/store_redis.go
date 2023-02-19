package verifycode

import (
	"goex/pkg/app"
	"goex/pkg/config"
	"goex/pkg/redis"
	"time"
)

// RedisStore implement verifycode.Store interface
type RedisStore struct {
	RedisClient *redis.RedisClient
	KeyPrefix   string
}

func (s *RedisStore) Set(key string, value string) bool {
	ExpireTime := time.Minute * time.Duration(config.GetInt64("verifycode.expire_time"))

	if app.IsLocal() {
		ExpireTime = time.Minute * time.Duration(config.GetInt64("verifycode.debug_expire_time"))
	}

	return s.RedisClient.Set(s.KeyPrefix+key, value, ExpireTime)
}

func (s *RedisStore) Get(key string, clear bool) (value string) {
	key = s.KeyPrefix + key
	val := s.RedisClient.Get(key)
	if clear {
		s.RedisClient.Del(key)
	}
	return val
}

func (s *RedisStore) Verify(key, answer string, clear bool) bool {
	v := s.Get(key, clear)
	return v == answer
}
