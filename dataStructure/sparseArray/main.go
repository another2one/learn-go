package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Node struct {
	row int
	col int
	val int
}

// 稀疏数据
// 第一行 记录数组总的行列及默认元素
// 第二行 记录不同元素的行列及值

func main() {

	// 定义一个数组
	arr1 := [3][4]int{
		{0, 1, 0, 1},
		{0, 2, 0, 1},
		{0, 1, 0, 0},
	}

	fmt.Println(arr1)

	// 创建稀疏数组
	sparseSlice := make([]Node, 1, 10)
	sparseSlice[0] = Node{3, 4, 0}
	for row, cols := range arr1 {
		for col, val := range cols {
			if val != 0 {
				sparseSlice = append(sparseSlice, Node{row, col, val})
			}
		}
	}
	fmt.Println(sparseSlice)

	// 还原数组
	// TODO: 二维切片不好处理，二维数组不能以变量来初始化，使用一维切片处理，
	var position = 0
	arrSlice := make([]int, sparseSlice[0].row*sparseSlice[0].col)
	var col = sparseSlice[0].col
	for index, node := range sparseSlice {
		if index > 0 {
			position = node.row*col + node.col // 转换为一维数组中的位置
			arrSlice[position] = node.val
		}
	}
	fmt.Println(arrSlice)

	// 存盘
	var str string
	for _, node := range sparseSlice {
		str += fmt.Sprintf("%d %d %d \n", node.row, node.col, node.val)
	}
	err := ioutil.WriteFile("D:/go/learn/src/dataStructure/sparseArray/sparse.data", []byte(str), 0666)
	if err != nil {
		fmt.Println("write error: ", err)
	}

	//复盘
	strBS, _ := ioutil.ReadFile("D:/go/learn/src/dataStructure/sparseArray/sparse.data")
	strSlice := strings.Split(string(strBS), "\n")
	node1, _ := getSlice(strSlice[0])
	s1 := make([]int, node1[0]*node1[1])
	col = node1[1]
	for index, _ := range strSlice {
		if index > 0 {
			node1, err := getSlice(strSlice[index])
			if err != nil {
				continue
			}
			position = node1[0]*col + node1[1]
			s1[position] = node1[2]
		}
	}
	var realStr = ""
	for index, val := range s1 {
		if index > 0 && index%node1[0] == 0 {
			realStr += "\n" + strconv.Itoa(val)
		} else if index == 0 {
			realStr += strconv.Itoa(val)
		} else {
			realStr += " " + strconv.Itoa(val)
		}
	}
	ioutil.WriteFile("D:/go/learn/src/dataStructure/sparseArray/array.data", []byte(realStr), 0666)
}

func getSlice(str string) ([]int, error) {
	str = strings.Trim(str, " \n")
	s1 := make([]int, 3)
	if len(str) > 1 {
		ss := strings.Split(str, " ")
		for i, val := range ss {
			s1[i], _ = strconv.Atoi(val)
		}
	} else {
		return s1, errors.New("empty string")
	}
	return s1, nil
}
