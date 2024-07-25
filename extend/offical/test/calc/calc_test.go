package calc

import (
	"testing"
)

// 注意：
// 1. 以 _test.go 结尾
// 2. 测试函数以 Test 开头, 驼峰法命名
// 3. go test -v -v用来显示详细信息
// 4. 测试指定文件： go test -v 文件名
// 5. 测试指定方法：go test -v -run 方法

func TestCalc(t *testing.T) {
	t.Log("TestCalc")
	res := Calc(10)
	if res != 55 {
		// t.fail 记录错误继续执行
		// t.Fatal 打印错误并停止
		t.Fatalf("calc(10) 结果出错，应为55， 实为：%d \n", res)
	}
	t.Logf("calc(10) 结果正确")
}

func TestCalc1(t *testing.T) {
	t.Log("TestCalc1")
	res := Calc(10)
	if res != 55 {
		t.Fatalf("calc(10) 结果出错，应为55， 实为：%d \n", res)
	}
	t.Logf("calc(10) 结果正确")
}
