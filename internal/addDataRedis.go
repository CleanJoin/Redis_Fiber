package internal

import "github.com/go-redis/redis"

func AddDataInRedis(client redis.Conn) {
	err := client.Do("ZADD", "hackers", 1953, "Richard Stallman")
	if err != nil {
		panic(err)
	}
	err = client.Do("ZADD", "hackers", 1940, "Alan Kay")
	if err != nil {
		panic(err)
	}
	err = client.Do("ZADD", "hackers", 1965, "Yukihiro Matsumoto")
	if err != nil {
		panic(err)
	}
	err = client.Do("ZADD", "hackers", 1916, "Claude Shannon")
	if err != nil {
		panic(err)
	}

	err = client.Do("ZADD", "hackers", 1969, "Linus Torvalds")
	if err != nil {
		panic(err)
	}
	err = client.Do("ZADD", "hackers", 1912, "Alan Turing")
	if err != nil {
		panic(err)
	}
}
