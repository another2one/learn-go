package calc

import (
	// "fmt"
	"testing"
)

// 注意：
// 1. 以 _test.go 结尾
// 2. 测试函数以 Test 开头, 驼峰法命名
// 3. go test -v -v用来显示详细信息
// 4. 测试单个文件： go test -v 文件名 go test -v sub_test.go
// 5. 测试单个方法：go test -v -run 文件名


func TestCalc(t *testing.T) {
	res := Calc(10)
	if res != 55 {
		t.Fatalf("calc(10) 结果出错，应为55， 实为：%d \n", res)
	}
	t.Logf("calc(10) 结果正确")
}

func TestCalc1(t *testing.T) {
	res := Calc(10)
	if res != 55 {
		t.Fatalf("calc(10) 结果出错，应为55， 实为：%d \n", res)
	}
	t.Logf("calc(10) 结果正确")
}

func TestSub(t *testing.T) {
	res := Sub(10, 3)
	if res != 7 {
		t.Fatalf("Sub(10, 3) 结果出错，应为7， 实为：%d \n", res)
	}
	t.Logf("Sub(10, 3) 结果正确")
}