package main

import (
	"log"
	"os"
	"runtime/pprof"
	"strings"
)

func main() {
	cpuf, err := os.Create("cpu_profile")
	if err != nil {
		log.Fatal("cpuf create err: ", err)
	}
	defer cpuf.Close()
	_ = pprof.StartCPUProfile(cpuf)
	defer pprof.StopCPUProfile()

	memf, err := os.Create("mem_profile")
	if err != nil {
		log.Fatal("memf create err: ", err)
	}
	defer memf.Close()
	_ = pprof.WriteHeapProfile(memf)

	num := 1000000
	for num > 0 {
		num--
		shareWith("alice")
	}
}

func shareWith(s string) string {
	if s == "" {
		s = "you"
	}
	str := strings.Join([]string{"one for ", ", one for me"}, s)
	return str
}
