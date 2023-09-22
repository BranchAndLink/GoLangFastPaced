// When an error occurs, the execution of a program stops completely with a built-in error message.
// it is important to handle those exceptions in our program.
package main

import (
	"errors"
	"fmt"
)

// function that checks if name is Programiz
func checkName(name string) error {

	// create a new error
	newError := errors.New("Invalid Name")

	// return the error if name is not Programiz
	if name != "Programiz" {
		return newError
	}

	// return nil if there is no error
	return nil
}

func divide(num1, num2 int) (int, error) {

	// returns error
	if num2 == 0 {
		return 0, fmt.Errorf("%d / %d\nCannot Divide a Number by zero \n", num1, num2)
	}

	// returns the result of division
	return num1 / num2, nil
}

func main() {

	message := "Hello"

	// create error using New() function
	myError := errors.New("WRONG MESSAGE")

	if message != "Programiz" {
		fmt.Println(myError)
	}

	name := "Salam"

	// call the function
	err := checkName(name)

	fmt.Println(err)

	age := -14

	// create an error using Efforf()
	error := fmt.Errorf("%d is negative\nAge can't be negative", age)
	//error := fmt.Sprintf("%d is negative. Age can't be negative", age)

	if age < 0 {
		fmt.Println(error)
	} else {
		fmt.Printf("Age: %d", age)
	}
	fmt.Println()

	result, diverr := divide(4, 2)

	// error found
	if diverr != nil {
		fmt.Printf("error: %s", diverr)

		// error not found
	} else {
		fmt.Println("Valid Division", "result:", result)
	}

	// defer the execution of Println() function
	defer fmt.Println("Three")
	fmt.Println("One")
	fmt.Println("Two")
	//multiple defer statements
	defer fmt.Println("One")
	defer fmt.Println("Two")
	defer fmt.Println("Three")

	//The lines of code after the panic statement are not executed.
	fmt.Println("Oh! counting for defer is happening now.")

	go display("Process 1")

	display("Process 2")

	division(4, 2)
	division(8, 0)
	division(56, 8)
	panic("Ending the program")
	//unreachable
	fmt.Println("Waiting to execute")

}

func division(num1, num2 int) {

	// if num2 is 0, program is terminated due to panic

	defer handlePanic() //handling
	if num2 == 0 {
		panic("Cannot divide a number by zero")
	} else {
		result := num1 / num2
		fmt.Println("Result: ", result)
	}
}

func handlePanic() {

	// detect if panic occurs or not
	a := recover()

	if a != nil {
		fmt.Println("RECOVER", a)
	}

}

func display(message string) {

	fmt.Println(message)
}
