package calc

import (
	"os"
	"strconv"
	"testing"
)

// 注意：
// 1. 以 _test.go 结尾
// 2. 测试函数以 Test 开头, 驼峰法命名
// 3. go test -v -v用来显示详细信息
// 		go clean -cache
// 		测试指定文件： go test -v [文件名/夹]   eg: go test -v calc1_test.go calc.go
// 		测试指定方法：go test -v -run 方法	eg: go test -v -run TestCalc\d (支持正则)
// 		并行测试 go test -v -parallel n
// 		性能测试 go test -v -bench . -run . -cpu 1,2,4 -benchmem -count 10 -coverprofile=calc.out -covermode=count
//		查看覆盖率 go tool cover -html=calc.out

func TestMe(t *testing.T) {
	t.Log("TestCalc2")
	res := Calc(10)
	if res != 55 {
		// t.fail 记录错误继续执行
		// t.Fatal 打印错误并停止
		t.Fatalf("calc(10) 结果出错，应为55， 实为：%d \n", res)
	}
	t.Logf("calc(10) 结果正确")
}

func TestCalc3(t *testing.T) {
	if os.Getenv("SOME_ACCESS_TOKEN") == "" {
		t.Skip("skipping test; $SOME_ACCESS_TOKEN not set") // 跳过测试
	}
	t.Log("TestCalc3")
	res := Calc(10)
	if res != 55 {
		t.Fatalf("calc(10) 结果出错，应为55， 实为：%d \n", res)
	}
	t.Logf("calc(10) 结果正确")
}

func BenchmarkCalc(b *testing.B) {
	s := []string{"0"}
	for i := 1; i < b.N; i++ {
		s = append(s, strconv.Itoa(i))
	}
}
