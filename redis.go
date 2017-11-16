package bidder

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"time"
	"log"
	"strings"

)
const StatisticsKey = "request:stats"

type StringInt struct {
	I int
	S string
 }

type RedisPool struct {
	*redis.Pool
}

func NewRedisPool(config *Config) *RedisPool {
	return &RedisPool{
		&redis.Pool{
			MaxIdle:     3,
			IdleTimeout: 240 * time.Second,
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", config.Server, config.Port))
				if err != nil {
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

func (rp *RedisPool) IncrIfa(st *StatisticsForIfa, lifeTime int) (s string, e error){
	c := rp.Get()
	defer c.Close()
	c.Send("MULTI")
	c.Send("INCR", st.Device.Ifa)
	c.Flush()

	c.Send("EXPIRE", st.Device.Ifa, lifeTime)

	r, err := redis.Bytes(c.Receive())
	fmt.Println(r)
	if err != nil{
		return "", err
	}
	return string(r), nil

}

func (rp *RedisPool) AddStat(st *Statistics)  error {
	con := rp.Get()
	defer con.Close()
	_, err := con.Do("zincrby", StatisticsKey, 1, st.GetFormattedStatString())
	if err != nil {
		return err
	}
	return nil
}

func (rp *RedisPool) GetStat()(result []ReturnStatistics)   {
	return rp.getStat("ZRANGE")
}

func (rp *RedisPool) GetSortedStat()(result []ReturnStatistics)   {
	return rp.getStat("ZREVRANGE")
}


func (rp *RedisPool) getStat(cmd string) (result []ReturnStatistics) {
	con := rp.Get()
	defer con.Close()
	values, err := redis.Values(con.Do(cmd, StatisticsKey, 0, -1, "WITHSCORES"))
	if err != nil {
		log.Println(err)
	}

	pairs := make([]StringInt, len(values)/2)

	for i := range pairs {
		values, err = redis.Scan(values, &pairs[i].S, &pairs[i].I)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, stringToStruct(pairs[i].S, pairs[i].I))
	}

	return
}


func stringToStruct(str string, count int) ReturnStatistics{
	r := strings.Split(str, ":")
	if len(r) >= 2 {
		return ReturnStatistics{Country:r[0], App:r[1], Platform:r[2], Count:count}
	}
	return ReturnStatistics{}

}
/*
127.0.0.1:6379> zincrby most_popular 1 rus:com.err:android
127.0.0.1:6379> zincrby most_popular 1 rus:302324249:android
ZREVRANGE most_popular 0 -1 WITHSCORES // all sorted keys
zrange key 0 -1 // all keys
*/