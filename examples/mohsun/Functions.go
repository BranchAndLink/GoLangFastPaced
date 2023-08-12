package main

import (
	"fmt"
	"math"
)

//we use the func keyword to create a function.

func greet() {
	fmt.Println("Good Morning")
}

// creating functions with function parameters
func product(a int, b int, c int) int {
	return a * b * c
}

// function to add two numbers
func addNumbers(n1 int, n2 int) int {
	sum := n1 + n2
	return sum
}

//The function parameter data type should always match the data passed during the function call
// calling functions.

// once function returns value to the place from where it is called,  we can store the returned value in a variable
var result = product(4, 5, 6)

// we can also return multiple values from a function
func calculate(n1 int, n2 int) (int, int) {
	sum := n1 + n2
	difference := n1 - n2
	return sum, difference
}

// We can reuse the same function multiple times in our program.
func getSquare(num int) {
	square := num * num
	fmt.Printf("Square of %d is %d\n", num, square)
}

// assigning dunction to a variable
var adding func(int, int) int = addNumbers

// variablese are copied once using in function. To change the variable value via function we should use pointer.
func myFunction(number *int) int {
	*number *= 2
	return *number
}

//named return values.
//Go's return values may be named. If so, they are treated as variables defined at the top of the function.

func getCoordinates(locationValue float64) (x float64, y float64) {
	x = math.Sin(locationValue)
	y = math.Cos(locationValue)
	return
}

func main() {
	greet()
	fmt.Println(product(1, 2, 3))
	fmt.Println(addNumbers(1, 9))
	//result := product(4,5,6)
	fmt.Println("product result:", result)

	sum, difference := calculate(21, 13)
	fmt.Println("Sum:", sum, "Difference:", difference)

	getSquare(3)
	getSquare(5)
	getSquare(10)

	fmt.Println(adding(1, 3))

	var a = 6
	fmt.Println(a)
	fmt.Println(myFunction(&a))
	fmt.Println(a)
	x_axis, y_axis := getCoordinates(0.5)
	fmt.Println("Coordinates:", x_axis, y_axis)

}

//The return statement should be the last statement of the function.
//When a return statement is encountered, the function terminates.
//Functions help us break our code into chunks to make our program readable and easy to understand.
//It also makes the code easier to maintain and debug.
