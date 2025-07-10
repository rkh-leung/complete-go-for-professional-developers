package main

import (
	"fmt"
)

// package main can import but can't export anything
// go build -o output_file build_file
func main() {
	// Variables
	var name string = "Format specifier"

	fmt.Printf("%s for string is %%s\n", name)

	num := 2
	fmt.Printf("Another way to declare a variable is using inference with walrus operator num = %d\n", num)
	var city string
	city = "Auckland"
	fmt.Printf("%s is my city\n", city)

	var book, cup string = "Golang", "fancy cup"
	fmt.Printf("You can also assign multiple variables with 'var book, cup string = %s, %s'\n", book, cup)

	var (
		isEmployed bool   = true
		salary     int    = 50000
		position   string = "developer"
	)

	fmt.Printf("Am I employed? %t. My position is a %s, and I earn %d\n", isEmployed, position, salary)

	// zero values
	var defaultInt int
	var defaultFloat float64
	var defaultString string
	var defaultBool bool
	fmt.Printf("%d, %f, %s, %t\n", defaultInt, defaultFloat, defaultString, defaultBool)

	// const can be used but variables can't, an issue of memory allocation
	const pi = 3.14

	const (
		Monday    = 1
		Tuesday   = 2
		Wednesday = 3
	)

	const typedNumber int = 25
	const untypedNumber = 25

	fmt.Printf("typed == untyped: %t\n", typedNumber == untypedNumber)

	// declare const using iota
	const (
		Jan = iota + 1
		Feb
		Mar
		Apr
	)

	fmt.Printf("Jan to Apr %d, %d, %d, %d\n", Jan, Feb, Mar, Apr)
}
