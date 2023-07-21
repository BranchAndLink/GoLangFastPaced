package main

import "fmt"

//data types

// byte - Typically a single octet(one byte). This is an byte type.

// int - The most natural size of integer for the machine.

// float32 - A single-precision floating point value.

//A variable definition tells the compiler where and how much storage to create for the variable. A variable definition specifies a data type and contains a list of one or more variables of that type as follows −

// printing variable and variable data type
func main() {
	var x int
	x = 20
	fmt.Println(x)
	fmt.Printf("type of x is %T\n", x)

	var z float64 = 20.0
	z = 25
	y := 42
	fmt.Println(z)
	fmt.Println(y)
	fmt.Printf("z is of type %T\n", z)
	fmt.Printf("y is of type %T\n", y)
}
