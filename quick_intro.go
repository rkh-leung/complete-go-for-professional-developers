package main

import (
	"fmt"
)

type Person struct { // custom data type
	Name string
	Age  int
}

// package main can import but can't export anything
// go build -o output_file build_file
func quick_intro() {
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

	// Arrays and slices
	numbers := [5]int{10, 20, 30, 40, 50} // can only have one type
	fmt.Printf("This is out array %v\n", numbers)
	fmt.Printf("This is the last value %v\n", numbers[len(numbers)-1])
	numbersAtInit := [...]int{20, 20, 20} // set array capacity based on initializtion
	fmt.Printf("This is out array %v\n", numbersAtInit)
	fmt.Printf("This is the last value %v\n", numbersAtInit[len(numbersAtInit)-1])
	matrix := [2][3]int{ //[row][column]
		{1, 2, 3},
		{4, 5, 6},
	}
	fmt.Println("this is matrix", matrix)

	// slice is a dynamic array or a portion of an array
	allNumbers := numbers[:] // slice copy of numbers array
	fmt.Println(allNumbers)
	firstThree := numbers[0:3]
	fmt.Println(firstThree)

	fruits := []string{"apple", "banana", "strawberry"}
	fmt.Printf("these are my fruits %v\n", fruits)
	fruits = append(fruits, "kiwi")
	fmt.Printf("these are my fruits with kiwi %v\n", fruits)

	moreFruits := []string{"blueberries", "durian"}
	fruits = append(fruits, moreFruits...)
	fmt.Printf("these are my fruits with more fruits %v\n", fruits)

	for index, value := range numbers {
		fmt.Printf("index %d and value %d\n", index, value)
	}

	// Maps (key value store, hash map)
	capitalCities := map[string]string{
		"USA":   "Washington D.C.",
		"India": "New Delhi",
		"UK":    "London",
	}
	fmt.Println(capitalCities["UK"])

	capital, exists := capitalCities["Germany"] // exists returns boolean
	if exists {
		fmt.Println("This is the capital", capital)
	} else {
		fmt.Println("Does not exist", exists) // prints false
	}
	delete(capitalCities, "UK")
	fmt.Printf("This is new map %v\n", capitalCities)

	// Struct - data type that can hold data and pass around
	person := Person{Name: "John", Age: 20}
	fmt.Printf("This is person struct: %v\n", person)
	fmt.Printf("This is person struct with fields: %+v\n", person)

	employee := struct { // {type}{data}
		name string
		id   int
	}{
		name: "alice",
		id:   123,
	}
	fmt.Println("this is employee", employee)

	type Address struct {
		Street string
		City   string
	}
	type Contact struct {
		Name    string
		Address Address
		Phone   string
	}

	contact := Contact{
		Name: "Marc",
		Address: Address{
			Street: "123 Main street",
			City:   "Anytown",
		},
		Phone: "12345", // doesn't have to have all fields in Contact
	}
	fmt.Println("this is nested contact", contact)

	// Pointers & Struct methods
	fmt.Println("name before: ", person.Name) // John
	modifyPersonName(person)                  // pass by copied value
	fmt.Println("name after: ", person.Name)  // No change
	new_name := modifyPersonName(person)
	fmt.Println("name after: ", new_name)

	x := 20
	ptr := &x // set ptr to the reference of x
	fmt.Printf("value of x: %d and address of x %p\n", x, ptr)
	*ptr = 30 // dereference ptr to get the value and set to 30
	fmt.Printf("value of new x: %d and address of x %p\n", x, ptr)

	person.modifyPersonNameMethod("new method name")
	fmt.Println("name after method: ", person.Name) // name modified
}

func modifyPersonName(person Person) string {
	person.Name = "K" // modifying within scope
	fmt.Println("inside scope new name: ", person.Name)
	return person.Name // returning modified value
}

// func (method receiver) nameOfMethod()
func (person *Person) modifyPersonNameMethod(name string) {
	person.Name = name
	fmt.Println("inside scope new name: ", person.Name)
}

// Add (capitalized means it is exportable)
// Every parameter has to specify a type, as well as the return fucntion type
func multiple_return_add_and_mul(a, b int) (int, int) {
	return a + b, a * b
}

func add(a int, b int) int {
	return a + b
}
