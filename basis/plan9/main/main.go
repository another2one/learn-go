package main

import (
	"runtime"
)

// go tool compile -S main.go
//  go build -GCflags -S main.go
//
// go build -o main main.go
// go tool objdump -s main.main main
func main() {
	println(runtime.GOOS)
}
