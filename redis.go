package bidder

import (
 "github.com/garyburd/redigo/redis"
	"time"
)
type RedisPool struct {
	R *redis.Pool
}

func NewRedisPool(config *Config) *RedisPool{
	return &RedisPool{
		R: &redis.Pool{
			MaxIdle:3,
			IdleTimeout:240 * time.Second,
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial("tcp", config.Server+config.Port)
				if err != nil{
					return nil, err
				}
				return c, err
			},
			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				_, err := c.Do("ping")
				return err
			},
		},
	}
}