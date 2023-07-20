package main

import "fmt"

//data types

// byte - Typically a single octet(one byte). This is an byte type.

// int - The most natural size of integer for the machine.

// float32 - A single-precision floating point value.

//A variable definition tells the compiler where and how much storage to create for the variable. A variable definition specifies a data type and contains a list of one or more variables of that type as follows −

//Warning : short variable declaration cannot be used outside functions !
//here we can not use := for defining variable
// A constant is a value in your program that will stand firm, that will not change during the execution. A variable can change during the runtime; a constant will not change; it will stay constant.

// printing variable and variable data type
func main() {

	var (
		y int
		x bool
		z float64
	)
	const placeOfInterest = `⌘` //Unicode character
	const version int = 5
	z = 25
	y = 42
	fmt.Println("value of z", z)
	fmt.Println("value of y", y)
	fmt.Println("value of x", x)

	fmt.Printf("plain string: ")
	fmt.Printf("%s", placeOfInterest)
	fmt.Printf("\n")

	fmt.Println(version)
	fmt.Printf("z is of type %T\n", z)
	fmt.Printf("y is of type %T\n", y)
	fmt.Printf("x is of type %T\n", x)
	fmt.Printf("version is of type %T\n", version)
}

// as we see the result of a boolean value is false by default after decleration
