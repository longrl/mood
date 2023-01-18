package captcha

import (
	"errors"
	"github.com/longrl/mood/config"
	"time"
)

// RedisStore 实现base64Captcha的store接口
type RedisStore struct {
	RedisClient *config.RedisClient
	KeyPrefix   string
}

func (s *RedisStore) Set(key string, value string) error {
	t := 5 * time.Minute
	ok := s.RedisClient.Set(s.KeyPrefix+key, value, t)
	if !ok {
		return errors.New("无法存储验证码答案")
	}
	return nil
}

func (s *RedisStore) Get(key string, clear bool) string {
	key = s.KeyPrefix + key
	val := s.RedisClient.Get(key)
	if clear {
		s.RedisClient.Del(key)
	}
	return val
}

// Verify 实现 base64Captcha.Store interface 的 Verify 方法
func (s *RedisStore) Verify(key, answer string, clear bool) bool {
	v := s.Get(key, clear)
	return v == answer
}
