package main

import (
	"fmt"
	"reflect"
)

// 注意
// 1. 地址需要使用Elem()进行操作 reflect.ValueOf(i).Field(1).SetInt(11)
// 2. 结构体指针调用方法只能地址放射才能取到，普通的都可以
// 3. 只有地址(reflect.ValueOf(i).Elem())才能进行值修改， CanSet()判断

type Student struct {
	Name  string `json:"name"`
	Age   int
	score float64
}

func (this *Student) Add(a, b int) int {
	return a + b
}

func (this Student) Add1(a, b int) int {
	return a + b
}

func (this Student) GetName() string {
	return this.Name
}

func testInt(i interface{}) {
	// 将 i 转为 reflect.Value
	ri := reflect.ValueOf(i)
	fmt.Printf("ri type is %v, kind is %v, value is %+v \n", reflect.TypeOf(i), ri.Kind(), ri)
	// 将 reflect.Value 转回 interface{}
	i = ri.Interface()
	fmt.Printf("i type is %T \n", i)
	// 将 i 转为 int
	stu := i.(int)
	fmt.Printf("stu type is %T \n", stu)
}

func testStruct(i interface{}) {

	// 获取i的值、类型、类别
	rvi := reflect.ValueOf(i)
	rti := reflect.TypeOf(i)
	ki := rvi.Kind()

	if ki == reflect.Ptr {
		rvi = rvi.Elem()
		rti = rti.Elem()
	}
	fmt.Printf("i type = %v, kind = %v, value = %v \n", rti, ki, rvi)

	// 修改字段
	if rvi.CanSet() {
		rvi.FieldByName("Age").SetInt(11)
	}

	// 获取结构体字段及标签
	num := rti.NumField()
	fmt.Printf("stu has %d field : \n", num)
	for i := 0; i < num; i++ {
		fname := rti.Field(i).Name
		fjsonTag := rti.Field(i).Tag.Get("json")
		fvalue := rvi.Field(i)
		if fjsonTag != "" {
			fjsonTag = "json:" + fjsonTag
		}
		fmt.Printf("%v : %v %v \n", fname, fvalue, fjsonTag)
	}

	// 获取方法并调用 (注意： ASCALL 排序)
	rvi = reflect.ValueOf(i)
	rti = reflect.TypeOf(i)
	num = rvi.NumMethod()
	fmt.Printf("stu has %d method : \n\n", num)
	params := []reflect.Value{reflect.ValueOf(10), reflect.ValueOf(20)}
	for i := 0; i < num; i++ {
		mname := rti.Method(i).Name
		mtype := rti.Method(i).Type

		fmt.Printf("%v : %v \n", mname, mtype)

		numParams := mtype.NumIn() - 1
		fmt.Println("num params:", numParams)
		res := rvi.Method(i).Call(params[:numParams])
		numRes := mtype.NumOut()
		for i := 0; i < numRes; i++ {
			out := mtype.Out(i)
			switch out.Kind() {
			case reflect.String:
				result := res[i].String()
				fmt.Printf("%v %d 个返回参数为：%v \n", mname, i+1, result)
			case reflect.Int:
				result := res[i].Int()
				fmt.Printf("%v %d 个返回参数为：%v \n", mname, i+1, result)
			}
		}
		fmt.Println()
	}

}

func main() {

	testInt(3)

	var ss Student
	ss.Name = "lizhi"
	testStruct(ss)
	testStruct(&ss)
	fmt.Println(ss)
}
