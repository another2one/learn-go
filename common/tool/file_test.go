package tool

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

var file = ProjectPath + "common/funcs/test.xlsx"

func setup() {

}

// TestMain 测试的主入口点
func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	//teardown()
	os.Exit(code)
}

func TestWriteExcel(t *testing.T) {
	var b strings.Builder
	b.Grow(100)
	os.Remove(file)
	err := WriteExcel([][]int{{1, 2, 3}, {4, 5, 6}}, file)
	if err != nil {
		t.Errorf("TestWriteExcelWithArray error: %s \n", err)
	}
}

func TestReadExcel(t *testing.T) {
	data, err := ReadExcel(file)
	if err != nil {
		t.Errorf("read err: %s", err)
	}
	fmt.Printf("read data : \n %v \n", data)
}
