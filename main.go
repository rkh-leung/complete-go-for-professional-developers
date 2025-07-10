package main

import (
	"fmt"
)

// package main can import but can't export anything
// go build -o output_file build_file
func main() {
	fmt.Println("hello world")
}
