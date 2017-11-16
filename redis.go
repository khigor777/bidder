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

func (rp *RedisPool) IncrIfa(st *StatisticsForIfa, lifeTime int)(int64, error) {
	c := rp.Get()
	defer c.Close()
	r, e:= redis.Int64(c.Do("INCR", st.Device.Ifa))
	c.Do("EXPIRE", st.Device.Ifa, lifeTime)
	if e != nil{
		return 0, e
	}
	return r, nil
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
		log.Fatal(err)
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