// 暂时测试
package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type CookieSameSite string
type CookiePriority string

// CookieSourceScheme represents the source scheme of the origin that
// originally set the cookie. A value of "Unset" allows protocol clients to
// emulate legacy cookie scope for the scheme. This is a temporary ability and
// it will be removed in the future.
//
// See: https://chromedevtools.github.io/devtools-protocol/tot/Network#type-CookieSourceScheme
type CookieSourceScheme string

// String returns the CookieSourceScheme as string value.
func (t CookieSourceScheme) String() string {
	return string(t)
}

// CookieSourceScheme values.
const (
	CookieSourceSchemeUnset     CookieSourceScheme = "Unset"
	CookieSourceSchemeNonSecure CookieSourceScheme = "NonSecure"
	CookieSourceSchemeSecure    CookieSourceScheme = "Secure"
)

var v = 4

func main() {

	// dialSocksProxy, err := proxy.SOCKS5("tcp", "127.0.0.1:10808", nil, proxy.Direct)
	// if err != nil {
	// 	fmt.Println("Error connecting to proxy:", err)
	// }
	// tr := &http.Transport{Dial: dialSocksProxy.Dial}

	// // Create client
	// myClient := &http.Client{
	// 	Transport: tr,
	// }
	// response, err := myClient.Get("https://etherscan.io/")
	// if err != nil {
	// 	fmt.Println("Error get url:", err)
	// }
	// defer response.Body.Close()
	// var by []byte
	// by, _ = ioutil.ReadAll(response.Body)
	err := writeFile("et.html", []byte("我是666\n"))
	fmt.Println(err)
	v := 1 //声明
	if v == 1 {
		v = 3 // 赋值
		fmt.Println(v)
	}
	fmt.Println(v) //依然可以访问

	ch := 'b'
	fmt.Println(ch / 2.0)

	//s1 := []int{1, 2, 4}
	s1 := make(map[int]int, 6)
	s1[1] = 1
	printAddress(s1)
	for i, v := range s1 {
		printAddress(s1)
		s1[i+1] = v + 1
	}
	fmt.Println(s1)

	for _, v := range []int{1, 2, 4} {
		fmt.Printf("v address is %v \n", &v)
		go func(v int) {
			fmt.Println(v)
		}(v)
	}
	time.Sleep(time.Microsecond * 100)
}

func printAddress(v interface{}) {
	fmt.Printf("v address is %p \n", &v)
}

// writeFile 实现追加写入
func writeFile(filename string, data []byte) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	n, err := file.Write(data)
	if err != nil && n < len(data) {
		err = io.ErrShortWrite
	}
	return err
}
