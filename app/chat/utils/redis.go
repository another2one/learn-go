package utils

import (
	"github.com/gomodule/redigo/redis"
)

var (
	redisPool *redis.Pool
)

func init() {
	// 全局变量只在文件第一次引入时初始化一次
	redisPool = &redis.Pool{
		Dial: func() (conn redis.Conn, err error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
		TestOnBorrow:    nil,
		MaxIdle:         8,
		MaxActive:       8,
		IdleTimeout:     100,
		Wait:            true,
		MaxConnLifetime: 3,
	}
	flushOnlineList()
}

func GetConn() redis.Conn {
	return redisPool.Get()
}

// 加入在线列表
func AddOnline(i int) (err error) {
	conn := redisPool.Get()
	defer conn.Close()
	_, err = conn.Do("SADD", "onlineList", i)
	return
}

// 移除在线列表
func RemoveOnline(i int) (err error) {
	conn := redisPool.Get()
	defer conn.Close()
	_, err = conn.Do("SREM", "onlineList", i)
	return
}

// 返回在线列表
func GetOnlineList() (list []int, err error) {
	conn := redisPool.Get()
	defer conn.Close()
	list, err = redis.Ints(conn.Do("SMEMBERS", "onlineList"))
	return
}

// 初始化在线列表
func flushOnlineList() {
	conn := redisPool.Get()
	defer conn.Close()
	conn.Do("del", "onlineList")
}
