package bidder

import (
 "github.com/garyburd/redigo/redis"
)
type RedisPool struct {
	R *redis.Pool
}

//func NewRedisPool() *RedisPool{
//
//}