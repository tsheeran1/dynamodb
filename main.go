package main

import (
	"fmt"

	"os"
)

func main() {

	switch os.Args[1] {
	case "filetest":
		filetest()
	case "httpget":
		httpget()
	default:
		fmt.Println("unknown request")
	}
}
