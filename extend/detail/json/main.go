package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string `json:"name"`
	age  int
}

func main() {

	// map
	// 序列化
	testMap := make(map[string]interface{})
	testMap["name"] = "lizhi"
	testMap["age"] = 12
	jsonMap, _ := json.Marshal(testMap)
	fmt.Println(string(jsonMap))

	testSturct := Student{"lizhi", 66}
	fmt.Printf("testSturct = %+v \n", testSturct)
	jsonStruct, err := json.Marshal(testSturct)
	if err != nil {
		fmt.Println("json marshal error: ", err)
	}
	fmt.Println(string(jsonStruct))

	// 反序列化
	err = json.Unmarshal(jsonStruct, &testSturct)
	if err != nil {
		fmt.Println("json Unmarshal error: ", err)
	}
	fmt.Printf("testSturct = %+v \n", testSturct)

	// slice
	// 序列化
	intArray := [5]int{1, 2, 3, 4, 5}
	byteArray, err := json.Marshal(intArray)
	if err != nil {
		fmt.Println("json marshal error: ", err)
	}
	fmt.Printf("intArray Marshal  = %q \n", string(byteArray))

	// 反序列化
	var intSlice []int
	err = json.Unmarshal(byteArray, &intSlice)
	if err != nil {
		fmt.Println("json Unmarshal error: ", err)
	}
	fmt.Printf("intSlice = %+v, type is %T \n", intSlice, intSlice)
}
