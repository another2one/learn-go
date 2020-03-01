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


func TestStore(t *testing.T) {
	m := NewMonster()
	m.Name = "李志"
	res := m.Store("monster.json")
	if !res {
		t.Fatalf("m.Store(\"monster.json\") 结果出错，应为true， 实为：%t \n", res)
	}
	t.Logf("m.Store(\"monster.json\") 结果正确")
}