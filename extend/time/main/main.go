package main

import (
	"fmt"
	"time"
)

// https://studygolang.com/pkgdoc

func main() {

	now := time.Now()

	// 格式化
	// formatTime := now.Format("2006-01-02 15:04:05") 记忆法：12345
	formatTime := now.Format("15") // 小时
	fmt.Printf("value is %v, type is %T \n", formatTime, formatTime)

	// 时间戳
	// timestamp := now.Unix()
	timestamp := now.UnixNano()
	fmt.Printf("value is %v, type is %T \n", timestamp, timestamp)

	// 解析字符串
	timeFormat, err := time.ParseInLocation("2006-01-02 15:04-05", "2020-03-04 12:00-00", time.Local)
	if err != nil {
		fmt.Printf("parse error: %v \n", err)
	}
	fmt.Printf("value is %v, type is %T \n", timeFormat, timeFormat)

	// 时间差
	// diff := time.Since(timeFormat) // until 相反，到当前时间的距离
	diff := now.Sub(timeFormat)
	fmt.Printf("value is %v, type is %T \n", diff, diff)

	// 比较
	fmt.Printf("value is %v, type is %T \n", now.After(timeFormat), now.After(timeFormat))

	// 加减时间
	timeBefore1hour := now.Add(time.Hour * 1) // 1小时前
	fmt.Printf("value is %v, type is %T \n", timeBefore1hour, timeBefore1hour)
	// ParseDuration parses a duration string.
	// A duration string is a possibly signed sequence of decimal numbers,
	// each with optional fraction and a unit suffix,
	// such as "300ms", "-1.5h" or "2h45m".
	//  Valid time units are "ns", "us" (or "µs"), "ms", "s", "m", "h".
	// 10分钟前
	m, _ := time.ParseDuration("-1m")
	m1 := now.Add(m)
	fmt.Println(m1)

	// 月初
	timeBeginMonth := time.Now().AddDate(0, -1, -1*now.Day()+1)
	timeBeginMonth = time.Date(timeBeginMonth.Year(), timeBeginMonth.Month(), timeBeginMonth.Day(), 0, 0, 0, 0, time.Local)
	fmt.Printf("value is %v, type is %T \n", timeBeginMonth, timeBeginMonth)

	// 时间常量
	fmt.Printf("value is %v, type is %T \n", time.Millisecond, time.Millisecond) // Second Minute Hour Microsecond Nanosecond

}
