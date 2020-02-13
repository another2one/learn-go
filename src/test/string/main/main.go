package main

import (
	"fmt"
	"strconv"
	"strings"
)

// https://studygolang.com/pkgdoc

func main() {

	str := "11我是李志"
	fmt.Println("str = ", str)

	fmt.Printf("str length is %d \n", len(str)) // 字符串长度
	
	strSlice := []rune(str)  // 转字符切片 []byte同理  string(strSlice) 转回
	fmt.Printf(" strSlice = %v, strSlice length is %d, type is %T \n", strSlice, len(strSlice), strSlice)
	fmt.Printf(" string(strSlice) = %v \n", string(strSlice))

	i, err := strconv.Atoi("11") // 字符串转数字 Itoa(数字转字符串)
	if err != nil {
		fmt.Println("字符转数字转换错误：", err)
	}else{
		fmt.Println("字符转数字转换成功：", i)
	}

	i2 := strconv.FormatInt(123, 8) // 10进转其他进制
	fmt.Printf("i2 = %v, i2 type is %T \n", i2, i2)

	fmt.Printf("66.jpg contain jpg is %t \n", strings.Contains("66.jpg", "jpg")) // 是否包含某字符串
	fmt.Printf("66.jpg has %d  个 6 \n", strings.Count("66.jpg", "6")) // 包含多少个某字符串
	fmt.Printf(". first find in 66.jpg is %d \n", strings.Index("66.jpg", "v6")) // 字符串第一次出现位置 LastIndex (最后一次)
	fmt.Printf("不区分大小写后 66.jpg 等于 66.JPG 为 %t \n", strings.EqualFold("66.jpg", "66.JPG")) // 不区分大小写比较字符串
	fmt.Printf("将66.jpg 中的 6 替换为8, 而且只替换一次 %q \n", strings.Replace("66.jpg", "6", "8", 1)) // 不区分大小写比较字符串 (-1表示全部替换)
	fmt.Printf("将66.jpg 按 . 切割 %v \n", strings.Split("66.jpg", ".")) // 切割
	fmt.Printf("将66.Jpg  转大写：%q, 转小写 %q \n", strings.ToUpper("66.Jpg"), strings.ToLower("66.Jpg")) // 大小写转换
	fmt.Printf("将 %q 去空格 %q \n", " 66.jpg ", strings.Trim(" 66.Jpg ", " 6g")) // 去掉左右两边指定字符 TrimLeft TrimRight
	fmt.Printf("%q 以.jpg结尾: %t \n", "66.jpg", strings.HasSuffix(" 66.jpg", ".jpg")) // 去掉左右两边指定字符 TrimLeft TrimRight
}