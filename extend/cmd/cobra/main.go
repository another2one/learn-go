package main

import (
	"learn-go/extend/cmd/cobra/cmd"
)

// go build -o main.exe .\main.go
// .\main.exe test1 -h
func main() {
	cmd.Execute()
}
