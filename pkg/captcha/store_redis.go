package captcha

import (
	"errors"
	"goex/pkg/app"
	"goex/pkg/config"
	"goex/pkg/redis"
	"time"
)

// RedisStore implement base64Captcha.Store interface
type RedisStore struct {
	RedisClient *redis.RedisClient
	KeyPrefix   string
}

// Set implement method Set of base64Captcha.Store interface
func (s *RedisStore) Set(key string, value string) error {
	ExpireTime := time.Minute * time.Duration(config.GetInt64("captcha.expire_time"))

	if app.IsLocal() {
		ExpireTime = time.Minute * time.Duration(config.GetInt64("captcha.expire_time"))
	}

	if ok := s.RedisClient.Set(s.KeyPrefix+key, value, ExpireTime); !ok {
		return errors.New("unable to store image captcha answer")
	}
	return nil
}

// Get implement method Get of base64Captcha.Store interface
func (s *RedisStore) Get(key string, clear bool) string {
	key = s.KeyPrefix + key
	val := s.RedisClient.Get(key)
	if clear {
		s.RedisClient.Del(key)
	}
	return val
}

// Verify implement method Verify of base64Captcha.Store interface
func (s *RedisStore) Verify(key, answer string, clear bool) bool {
	v := s.Get(key, clear)
	return v == answer
}
