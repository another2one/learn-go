package main

import (
	"fmt"
	"time"
)

// https://studygolang.com/pkgdoc

func main() {

	now := time.Now()

	// formatTime := now.Format("2006-01-02 15:04:05")
	formatTime := now.Format("15") // 小时
	fmt.Printf("value is %v, type is %T \n", formatTime, formatTime)

	// timestamp := now.Unix()
	timestamp := now.UnixNano()
	fmt.Printf("value is %v, type is %T \n", timestamp, timestamp)

	// 时间常量
	fmt.Printf("value is %v, type is %T \n", time.Millisecond, time.Millisecond) // Second Minute Hour Microsecond Nanosecond

}
