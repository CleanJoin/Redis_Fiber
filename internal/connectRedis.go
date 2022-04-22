package internal

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

var pool = newPool()

func ConnectRedis() interface{} {

	client := pool.Get()
	defer client.Close()

	value, err := client.Do("ZRANGE", "hackers", 0, -1, "WITHSCORES")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s \n", value)
	return value
}

func newPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}
