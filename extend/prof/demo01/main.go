package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"reflect"
	"runtime/pprof"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")
var memprofile = flag.String("memprofile", "", "write memory profile to `file`")

// 非 web 程序 cpu 及 内存分析
// go build -o main.exe main.go
// ./main.exe -cpuprofile=cpu -memprofile=mem
// go tool pprof cpu
// top10
// list comb
// web comd (需安装 graphviz)
func main() {
	flag.Parse()

	fmt.Printf("cpuprofile = %s, memprofile = %s \n", *cpuprofile, *memprofile)

	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close()
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close()
		pprof.WriteHeapProfile(f)
	}

	fmt.Println("comb(30) = ", comb(30))
	fmt.Println("comb(40) = ", comb(40))
	fmt.Println("fiber(15) = ", fiber(15))
	fmt.Println("fiber(20) = ", fiber(20))

	s1 := getS1()
	s2 := getS2()
	fmt.Println(reflect.DeepEqual(s1, s2))
}

func getS1() []int {
	s1 := make([]int, 10000000)
	for i := 0; i < 10000000; i++ {
		s1[i] = i
	}
	return s1
}

func getS2() []int {
	s2 := make([]int, 1)
	for i := 1; i < 10000000; i++ {
		s2 = append(s2, i)
	}
	return s2
}

func fiber(n int) int {
	if n <= 2 {
		return n
	}
	return n * fiber(n-1)
}

// 分析后这里耗时太久了
func comb(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	if n < 1 {
		return 0
	}
	return comb(n-1) + comb(n-2)
}

// 优化
// func comb(n int) int {
// 	s1 := make([]int, n)
// 	for i := 0; i < n; i++ {
// 		if i <= 2 {
// 			s1[i] = i
// 		} else {
// 			s1[i] = s1[i-1] + s1[i-2]
// 		}

// 	}
// 	return s1[n-1]
// }
