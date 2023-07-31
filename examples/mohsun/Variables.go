package main

import (
	"fmt"
)

//data types
// byte - Typically a single octet(one byte). This is an byte type.
// int - The most natural size of integer for the machine.
// float32 - A single-precision floating point value.
// uint8       the set of all unsigned  8-bit integers (0 to 255)
// uint16      the set of all unsigned 16-bit integers (0 to 65535)
// uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
// uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

// int8        the set of all signed  8-bit integers (-128 to 127)
// int16       the set of all signed 16-bit integers (-32768 to 32767)
// int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
// int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

// float32     the set of all IEEE-754 32-bit floating-point numbers
// float64     the set of all IEEE-754 64-bit floating-point numbers

// complex64   the set of all complex numbers with float32 real and imaginary parts
// complex128  the set of all complex numbers with float64 real and imaginary parts

// byte        alias for uint8
// rune        alias for int32

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
	z = 25.99
	y = 42
	var strVar string
	strVar = "66"
	converedVar := int(z)
	new_var := 100
	_, bb, cc := 33, 66, 77
	fmt.Println("value of z", z)
	fmt.Println("value of y", y)
	fmt.Println("value of x", x)
	fmt.Println("value of new_var", new_var)
	fmt.Println("value of strVar", strVar)
	fmt.Println("value of converedVar", converedVar)
	fmt.Println(bb, cc)

	fmt.Printf("plain string: ")
	fmt.Printf("%s", placeOfInterest)
	fmt.Printf("\n")

	fmt.Println(version)
	fmt.Printf("z is of type %T\n", z)
	fmt.Printf("y is of type %T\n", y)
	fmt.Printf("x is of type %T\n", x)
	fmt.Printf("version is of type %T\n", version)
	fmt.Printf("str_var is of type %T\n", strVar)

	var _, b, c complex64 = 2, 3, 4 + 5i
	compResult := c * b
	fmt.Println("Complex number operation result: ", compResult)
}

// as we see the result of a boolean value is false by default after decleration
