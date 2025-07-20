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

	result := add(5, 3)
	fmt.Printf("Here's the result: %d\n", result)

	// sum, product destructuring is hanlded as a first class citizen
	// a, _ or _, b are both valid
	sum, product := multiple_return_add_and_mul(5, 3)
	fmt.Printf("Here's the result of sum and product: %d, %d\n", sum, product)
	_, new_res := multiple_return_add_and_mul(5, 2)
	fmt.Println(new_res)

	// conditional statements

	age := 30
	if age >= 18 {
		fmt.Println("You're an adult")
	} else if age >= 13 {
		fmt.Println("you are a teenager")
	} else {
		fmt.Println("you are a child")
	}

	day := "Tuesday"

	switch day {
	case "Monday":
		fmt.Println("Start of the week")
	case "Tuesday", "Wednesday", "Thursday":
		fmt.Println("Midweek")
	case "Friday":
		fmt.Println("TGIF")
	default:
		fmt.Println("its the weekend")
	}

	// loops
	for i := 0; i < 5; i++ {
		fmt.Println("this is i", i)
	}

	counter := 0
	for counter < 3 { // equivalent to while loop
		counter++
		fmt.Println("this is counter", counter)
	}

	iterations := 0
	for { // infinite loop, require break
		if iterations > 3 {
			break
		}
		iterations++
	}
}

// Add (capitalized means it is exportable)
// Every parameter has to specify a type, as well as the return fucntion type
func multiple_return_add_and_mul(a, b int) (int, int) {
	return a + b, a * b
}

func add(a int, b int) int {
	return a + b
}
