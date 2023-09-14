package main

import "fmt"

// function that returns an anonymous function
func displayNumber() func() int {

	number := 10
	return func() int {
		number--
		return number //Inside the displayNumber() function, we have created an anonymous function.
	}
}

// regular function to calculate square of numbers
func findSquare(num int) int {
	square := num * num
	return square
}

//CURRING examples

type CR1 func(int) int
type CR2 func(int, int) int
type CR3 func(int, int, int) int
type CR4 func(int, int, int, int) int

func (f CR1) Curry(arg int) func() int {
	return func() int {
		return f(arg)
	}
}

func (f CR2) Curry(arg int) CR1 {
	return func(h int) int {
		return f(arg, h)
	}
}

func (f CR3) Curry(arg int) CR2 {
	return func(x, y int) int {
		return f(arg, x, y)
	}
}

func (f CR4) Curry(arg int) CR3 {
	return func(x, y, z int) int {
		return f(arg, x, y, z)
	}
}

func main() {
	var greet = func() {
		fmt.Println("Function without name")
	}
	greet()

	a := displayNumber()

	fmt.Println(a())

	// anonymous function with arguments
	var sum = func(n1, n2 int) {
		sum := n1 + n2
		fmt.Println("Sum is:", sum)
	}

	// function call
	sum(5, 3)

	//Program to return value from an anonymous function
	var product = func(n1, n2 int) int {
		product := n1 * n2
		return product
	}
	result := product(5, 3)

	fmt.Println("Product is:", result)

	//Anonymous Function as Arguments to Other Functions
	// anonymous function that returns sum of numbers
	sum_of_numbers := func(number1 int, number2 int) int {
		return number1 + number2
	}

	// function call
	squareOfnumber := findSquare(sum_of_numbers(6, 9))
	fmt.Println("Square of the sum is:", squareOfnumber)

	//Running currying
	var fn CR4 = func(x, y, z, d int) int {
		return (x + y + z) * d
	}
	fn10 := fn.Curry(10)
	fn10_11 := fn10.Curry(11)
	fn10_11_12 := fn10_11.Curry(12)
	fn10_11_12_2 := fn10_11_12.Curry(2)
	fmt.Printf("%T %v\n", fn, fn(10, 11, 12, 2))
	fmt.Printf("%T %v\n", fn10, fn10(11, 12, 2))
	fmt.Printf("%T %v\n", fn10_11, fn10_11(12, 2))
	fmt.Printf("%T %v\n", fn10_11_12, fn10_11_12(2))
	fmt.Printf("%T %v\n", fn10_11_12_2, fn10_11_12_2())
	fn882 := fn.Curry(8).Curry(8).Curry(2)
	fmt.Printf("(8+8+2) * 55 = %v\n", fn882(55))
}
