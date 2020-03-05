package main
import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func main() {
	
	// 创建连接
	c, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	// 设置值
	res, _ := c.Do("set", "Inum", "lizhi")
	c.Do("hmset", "user1", "name", "lizhi", "age", 8)
	fmt.Printf("res type %T, %v \n", res, res)

	// 第一种取值方法 类型断言 （不推荐）
	r1, _ := c.Do("get", "Inum")
	r2, ok := r1.([]byte)
	if !ok {
		fmt.Println("not string")
	}
	fmt.Println("r2 = ", string(r2))

	// 第二种取值方法 redis 转换值类型
	r, _ := redis.String(c.Do("get", "Inum"))
	rs, _ := redis.Strings(c.Do("hgetall", "user1"))
	fmt.Printf("r type %T, %v \n", r, r)
	fmt.Printf("rs type %T, %+v \n", rs, rs)
}