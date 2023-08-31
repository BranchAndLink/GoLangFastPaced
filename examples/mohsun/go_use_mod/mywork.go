package main

import (
	"fmt"

	"github.com/mohsun/mymodule/package1"
)

// /package2 i de sitifade ele
// create add function
func main() {

	number1 := 9
	number2 := 5

	// use the add function of calculator package
	fmt.Println(package1.Add(number1, number2))

	// use the subtract function of calculator package
	fmt.Println(package1.Subtract(number1, number2))

	package1.Salam2("Mohsun")
}
