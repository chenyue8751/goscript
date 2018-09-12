package redisModel

import (
    "github.com/garyburd/redigo/redis"
    "time"
)

func newPool(addr, password string) *redis.Pool {
    return &redis.Pool {
        MaxIdle: 3,
        IdleTimeout: 240 * time.Second,
        Dial: func () (redis.Conn, error) {
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

var pool *redis.Pool

func InitRedis(addr, password string) *redis.Pool {
    pool = newPool(addr, password)
    return pool
}
