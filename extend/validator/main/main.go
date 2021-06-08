package main

// https://www.toutiao.com/i6736832899640738315/
// "github.com/go-playground/validator/v10"
import (
	"fmt"

	"github.com/asaskevich/govalidator"
)

type foo struct {
	A string `valid:"ipv4"`
	B string `valid:"mac"`
	C string `valid:"range(0|100)"` // 也可以使用int类型
}

func main() {
	// 判断字符串值是否为合法的IPv4地址
	ip4 := "192.168.1.1"
	fmt.Println(govalidator.IsIPv4(ip4)) // true

	// 判断字符串值是否为合法的MAC
	mac := "aa:bb:cc:dd:ee:ffffff"
	fmt.Println(govalidator.IsMAC(mac)) // false

	// 判断数字是否在指定范围内
	dig := 101                                    // string类型也可以用
	fmt.Println(govalidator.InRange(dig, 0, 100)) // false

	f := foo{
		A: "192.168.1.1",
		B: "aa:bb:cc:dd:ee:ffffff",
		C: "101",
	}

	result, err := govalidator.ValidateStruct(f)
	if err != nil {
		fmt.Println("error: " + err.Error())
	}
	fmt.Println(result)

	// 驼峰转下划线
	fmt.Println(govalidator.CamelCaseToUnderscore("userFirstName"))
}
