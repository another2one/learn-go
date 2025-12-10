package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"learn-go/common/tool"
)

var key = "testBloom"

func main() {

	// 创建连接
	c, err := redis.Dial("tcp", "127.0.0.1:6379", redis.DialDatabase(3))
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
	defer c.Close()

	//base(c)
	err = bloomFilter(c)
	if err != nil {
		fmt.Printf("err == %+v \n", err)
	}
}

func base(c redis.Conn) {
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

// bloomFilter 布隆过滤器 测试三千万数据失败率及速度
func bloomFilter(c redis.Conn) error {
	timeSpend := tool.NewTimeSpend()

	exists, err := redis.Bool(c.Do("EXISTS", key))
	if err != nil {
		fmt.Println("判断key是否存在失败", err)
	}
	timeSpend.Record("EXISTS")
	if !exists {
		// 创建布隆过滤器
		if _, err := c.Do("BF.RESERVE", key, 0.001, 30000000); err != nil {
			fmt.Println("创建布隆过滤器失败", err)
			return err
		}
		timeSpend.Record("RESERV")
		step := 1000
		addIntSlices := make([]any, step+1)
		addIntSlices[0] = key
		for i := 0; i < 30000000; i += step {
			for j := i; j < i+step; j++ {
				addIntSlices[j-i+1] = fmt.Sprintf("%d-lizhi", j)
			}
			// 添加数据
			if _, err := c.Do("BF.MADD", addIntSlices...); err != nil {
				fmt.Println("添加数据失败", err)
				return err
			}
		}
		timeSpend.Record("MADD")
	}

	res, err := c.Do("BF.INFO", key)
	if err != nil {
		fmt.Println("获取信息失败", err)
	}
	res1, _ := res.([]any)
	infoData := make([][]string, len(res1)/2)
	for i := 0; i < len(res1); i += 2 {
		infoData[i/2] = []string{res1[i].(string), fmt.Sprintf("%v", res1[i+1])}
	}
	tool.PrettyPrint("BF.INFO", infoData)
	timeSpend.Record("INFO")

	for i := 0; i < 30000000; i += 10101 {
		item := fmt.Sprintf("%d-lizhi", i)
		res, err := redis.Bool(c.Do("BF.EXISTS", key, item))
		if err != nil {
			fmt.Println("redis操作 BF.EXISTS 失败", err)
		}
		if !res {
			fmt.Printf("数据出错 %s 出错了 \n", item)
		}
	}

	timeSpend.EndPrint()

	return nil
}
