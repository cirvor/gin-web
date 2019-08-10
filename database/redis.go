package database

import (
	"log"
	"time"

	"github.com/gomodule/redigo/redis"
)

var Redis *redis.Pool

func init() {
	log.Println("++++++++++++++++++ redis connecting +++++++++++++++++++++")
	Redis = &redis.Pool{
		// 从配置文件获取maxidle以及maxactive，取不到则用后面的默认值
		MaxIdle:     10,
		MaxActive:   20,
		IdleTimeout: 5 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "127.0.0.1:6379", redis.DialDatabase(2))
			if err != nil {
				return nil, err
			}
			return c, nil
		},
	}
}
