// In Go, we can also create interfaces without any methods, known as empty interfaces
package main

import "fmt"

// empty interface as function parameter
func displayValue(i interface{}) {
	fmt.Println(i)
}

//We can also use an empty interface to pass any number of arguments to the function definition.

func displayMaltipleValues(i ...interface{}) {
	fmt.Printf("%v \n", i)
}

func main() {

	var empty interface{}
	fmt.Println("Value:", empty)

	//with an empty interface, we can pass parameters of any data type.
	a := "Welcome to Programiz"
	b := 20
	c := false

	// pass string to the function
	displayValue(a)

	// pass integer number to the function
	displayValue(b)

	// pass boolean value to the function
	displayValue(c)

	//passing multiple parameters
	displayMaltipleValues(a, b, c)

	//Go Type Assertions

	// create an empty interface
	var x interface{}

	// store integer to an empty interface
	x = 10

	// type assertion
	interfaceValue := x.(int)

	fmt.Println(interfaceValue)

	x = 12.5

	// type assertion
	interfaceValue, booleanValue := x.(int)

	fmt.Println("Interface Value:", interfaceValue)
	fmt.Println("Boolean Value:", booleanValue)

	switch x.(type) {
	case int:
		fmt.Println(x, "is integer")
	case string:
		fmt.Println(x, "is string")
	case float64:
		fmt.Println(x, "is float64")
	default:
		fmt.Println("unknown type")

	}

}
