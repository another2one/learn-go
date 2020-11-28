package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		Dial: func() (conn redis.Conn, err error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		}, // 初始化连接
		TestOnBorrow:    nil,   // 使用时检查
		MaxIdle:         8,     // 最大空闲连接数
		MaxActive:       0,     // 最大活跃连接数
		IdleTimeout:     100,   // 最大空闲时间
		Wait:            false, // 达到最大连接时，取连接是否等待
		MaxConnLifetime: 100,   // 最大连接时间
	}
}

func main() {

	conn := pool.Get() // 指向连接池
	defer conn.Close()
	_, err := conn.Do("lpush", "userList", "lipan", "shit") // 使用时取连接池里面的连接
	if err != nil {
		fmt.Println("write error: ", err)
	}
	res, err := redis.Strings(conn.Do("lrange", "userList", 0, 10))
	if err != nil {
		fmt.Println("read error: ", err)
	}
	fmt.Printf("%v \n", res)
}
