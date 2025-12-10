package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		//Password: "", // no password
		DB:       3, // use default DB
		Protocol: 2,
	})

	ctx := context.Background()

	bfInfo := rdb.BFReserve(ctx, "testBloom", 0.001, 30000000)
	if bfInfo.Err() != nil {
		fmt.Printf("bfInfo == %+v \n", bfInfo.Err())
		return
	}
}
