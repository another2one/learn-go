package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string `json:"name"`
	score int
	hobby []string
	grade map[string]interface{}
	Animal
}

type Animal struct {
	Name  string
	Wight float64
}

// 遍历结构体
func main() {
	grade := make(map[string]interface{}, 6)
	grade["语文"] = 105
	grade["数学"] = 96
	grade["英语"] = 112
	student := Student{
		Name:  "lizhi",
		score: 80,
		hobby: []string{"music", "football", "draw"},
		grade: grade,
		Animal: Animal{
			Name:  "小白",
			Wight: 11,
		},
	}

	rvi := reflect.ValueOf(student)
	rti := reflect.TypeOf(student)
	for i := 0; i < rti.NumField(); i++ {
		fname := rti.Field(i).Name
		fvalue := rvi.Field(i)
		if fvalue.IsValid() && fvalue.Kind() == reflect.Struct {
			// 如果是结构体就再次遍历
			fmt.Printf("%v : \n", fname)
			fType := fvalue.Type()
			for j := 0; j < fvalue.NumField(); j++ {
				fname1 := fType.Field(j).Name
				fvalue1 := fvalue.Field(j)
				fmt.Printf("\t%v : %v \n", fname1, fvalue1)
			}
		} else {
			fmt.Printf("%v : %v \n", fname, fvalue)
		}
	}
}
