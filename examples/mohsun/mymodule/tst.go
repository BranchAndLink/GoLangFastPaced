package main

import (
	"fmt"
	"mymodule/package1"
)

// create add function
func main() {

	number1 := 9
	number2 := 5

	// use the add function of calculator package
	fmt.Println(package1.Add(number1, number2))

	// use the subtract function of calculator package
	fmt.Println(package1.Subtract(number1, number2))
}
