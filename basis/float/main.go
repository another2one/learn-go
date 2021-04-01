package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unsafe"
)

var err error

func main() {

	var n1 = 1.1
	var n2 float64 = 1

	// n1,n2都为float64, 和整型不同
	fmt.Printf("n1 type is %T, size is %d, n1 type is %T, int64 size is %d \n", n1, unsafe.Sizeof(n1), n2, unsafe.Sizeof(n2))

	// 精度损失
	var n3 float32 = 123.000090101
	var n4 float64 = 123.00009010010010001
	fmt.Println("n3 =", n3, "n4 =", n4)           // n3 = 123.00009 n4 = 123.0000901001
	fmt.Printf("n3 = %.5f n4 = %.14f \n", n3, n4) // n3 = 123.00009 n4 = 123.00009010010010
	fmt.Printf("n3 = %.6f n4 = %.15f \n", n3, n4) // n3 = 123.000092 n4 = 123.000090100100095

	// 表示
	var n5 float32 = 123.0000901
	var n6 float64 = 9.01e-5
	fmt.Println("n5 =", n5, "n6 =", n6) // n5 = 123.00009 n6 = 9.01e-05

	// float32精度是小数后7位，float64精度是小数点后15位
	var a float32 = 1.00000001
	var b float32 = 1.000000000001
	var c float32 = 1.0000001
	var d float32 = 1.000000000001

	fmt.Println(a == b) //true
	fmt.Println(a > b)  //false
	fmt.Println(c == d) //false
	fmt.Println(c > d)  //true

	var e float64 = 1.0000000000000001
	var f float64 = 1.000000000000000001
	var g float64 = 1.000000000000001
	var h float64 = 1.0000000000000000001

	fmt.Println(e == f) //true
	fmt.Println(e > f)  //false
	fmt.Println(g == h) //false
	fmt.Println(g > h)  //true
}

// 主要逻辑 先乘10的n次方取整再除
func FormatFloat(num float64, decimal int) string {
	// 默认乘1
	d := float64(1)
	if decimal > 0 {
		// 10的N次方
		d = math.Pow10(decimal)
	}
	// math.trunc作用就是返回浮点数的整数部分
	// 再除回去，小数点后无效的0也就不存在了
	return strconv.FormatFloat(math.Trunc(num*d)/d, 'f', -1, 64)
}

// 主要逻辑 字符串截取
func ChangeNumber(f float64, m int) string {
	n := strconv.FormatFloat(f, 'f', -1, 32)
	if n == "" {
		return ""
	}
	if m >= len(n) {
		return n
	}
	newn := strings.Split(n, ".")
	if len(newn) < 2 || m >= len(newn[1]) {
		return n
	}
	return newn[0] + "." + newn[1][:m]
}
