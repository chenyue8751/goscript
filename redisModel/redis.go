package redisModel

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

func newPool(addr, password string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", addr)
			if err != nil {
				return nil, err
			}
			if password == "" {
				return c, err
			}
			if _, err := c.Do("AUTH", password); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

var (
	pool *redis.Pool
	conn redis.Conn
)

func InitRedis(addr, password string) *redis.Pool {
	pool = newPool(addr, password)
	return pool
}

func getKeys(pattern string) []string {
	conn = pool.Get()
	defer conn.Close()

	iter := 0

	result := make([]string, 0)
	var keys []string
	for {
		if arr, err := redis.MultiBulk(conn.Do("SCAN", iter, "MATCH", pattern, "COUNT", 10000)); err != nil {
			panic(err)
		} else {
			iter, _ = redis.Int(arr[0], nil)
			keys, _ = redis.Strings(arr[1], nil)
		}
		if keys != nil {
			result = mergeSlice(result, keys)
		}

		if iter == 0 {
			break
		}
	}

	return result
}

func deleteMulti(keys []string) int {
	len := len(keys)
	if len == 0 {
		return 0
	}

	conn = pool.Get()
	defer conn.Close()

	count := 0
	for i, j := 0, 1000; i < len; i, j = i+1000, j+1000 {
		if j > len {
			j = len
		}
		for _, value := range keys[i:j] {
			conn.Send("DEL", value)
		}
		conn.Flush()
		for _, _ = range keys[i:j] {
			result, err := redis.Int(conn.Receive())
			if err == nil {
				count += result
			}
		}
	}
	return count
}

func mergeSlice(s1 []string, s2 []string) []string {
	slice := make([]string, len(s1)+len(s2))
	copy(slice, s1)
	copy(slice[len(s1):], s2)
	return slice
}
