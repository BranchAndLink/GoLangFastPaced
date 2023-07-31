package main

import (
	"fmt"
	"os"
)

// functions to print output messages in Go programming.
func main() {

	fmt.Print("sen", "men", "o", "\n")
	fmt.Print("biz", "siz", "onlar", 5, "\n")

	//fmt.Println() prints a new line at the end by default.
	//If we print multiple values and variables at once, a space is added between the values by default.
	fmt.Println("sen", "men", "o")
	fmt.Println("biz", "siz", "onlar", 5)

	//The fmt.Printf() function formats the strings and sends them to the screen. Let's see an example.
	currentAge := 21
	fmt.Printf("Age = %d. ", currentAge)
	fmt.Println("Here, the fmt.Printf() function replaces the d with the value of currentAge")
	//%d is a format specifier that replaces integer variables with their values.
	//In Go, every data type has a unique format specifier. Data Type	Format Specifier
	// integer	%d
	// float	%g
	// string	%s
	// bool	    %t

	var annualSalary float32 = 65000.5

	fmt.Printf("Annual Salary: %g", annualSalary)
	fmt.Println("")
	//It's also possible to print output without using the fmt package. For that, we use print() and println()
	print("Using print instead of fmt.Print")
	println("")
	println("Using println instead of fmt.Println")

	fmt.Fprintln(os.Stdout, "Hello stdout")
	fmt.Fprintln(os.Stderr, "Hello stderr")
}
