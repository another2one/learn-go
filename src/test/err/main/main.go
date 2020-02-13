package main

import (
	"fmt"
	"errors"
	"time"
	"math"
)

var Start, _ = time.ParseInLocation("2006-01-02", "1990-01-01", time.Local)

// https://studygolang.com/pkgdoc

func test() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err = ", err)
		}
	}()
	// a := 10/0
	a1 := 0
	a2 := 0
	a := a1/a2
	fmt.Println("a = ", a)
	fmt.Println("555 啊!")
}

func test1() (err error){
	return errors.New("test1 error")
}

func checkDay(day time.Time) (isFish bool) {
	time := day.Unix()
	days := int64(math.Ceil(float64((time-Start.Unix())/86400)))
	if days%5 < 3 {
		isFish = true
	}else{
		isFish = false
	}
	return isFish
}


func main() {

	if err := test1(); err != nil {
		fmt.Println("err = ", err)
		// panic(err)
	}

	var needTime, _ = time.ParseInLocation("2006-01-02", "1990-01-09", time.Local)

	isFish := checkDay(needTime)
	fmt.Println("isFish = ", isFish)

	test()
	fmt.Println("666 啊!")
}