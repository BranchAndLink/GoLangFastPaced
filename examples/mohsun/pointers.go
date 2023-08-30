//Pointers in Go programming is a variable that is used to store the memory address of another variable.
//Pointers in Golang is also termed as the special variables. The variables are used to store some data at a particular memory
// address in the system. The memory address is always found in hexadecimal format(starting with 0x like 0xFFAAF etc.).

package main

import "fmt"

func main() {

	// storing the hexadecimal
	// values in variables
	x := 0xFF
	y := 0x9C

	// Displaying the values
	fmt.Printf("Type of variable x is %T\n", x)
	fmt.Printf("Value of x in hexadecimal is %X\n", x)
	fmt.Printf("Value of x in decimal is %v\n", x)

	fmt.Printf("Type of variable y is %T\n", y)
	fmt.Printf("Value of y in hexadecimal is %X\n", y)
	fmt.Printf("Value of y in decimal is %v\n", y)

	//* Operator also termed as the dereferencing operator used to declare pointer variable and access the value stored in the address.
	//& operator termed as address operator used to returns the address of a variable or to access the address of a variable to a pointer

	// taking a normal variable
	var v int = 5748

	// declaration of pointer
	var p *int

	// initialization of pointer
	p = &v

	// displaying the result
	fmt.Println("Value stored in x = ", v)
	fmt.Println("Address of x = ", &v)
	fmt.Println("Value stored in variable p = ", p)

	//The default value or zero-value of a pointer is always nil. Or you can say that an uninitialized pointer will always have a nil value.
	var s *int

	// displaying the result
	fmt.Println("s = ", s)
	m := 45
	point := &m

	fmt.Println("Value stored in m = ", m)
	fmt.Println("Address of m = ", &m)
	fmt.Println("Value stored in pointer variable point = ", point)

	//for better understanding * operator

	// using var keyword
	// we are not defining
	// any type with variable
	var i = 458

	// taking a pointer variable using
	// var keyword without specifying
	// the type
	var newpoin = &i

	fmt.Println("Value stored in i = ", i)
	fmt.Println("Address of i = ", &i)
	fmt.Println("Value stored in pointer variable p = ", newpoin)

	// this is dereferencing a pointer
	// using * operator before a pointer
	// variable to access the value stored
	// at the variable at which it is pointing
	fmt.Println("Value stored in i(*newpoin) = ", *newpoin)

	// using var keyword
	// we are not defining
	// any type with variable

	// taking a pointer variable using
	// var keyword without specifying
	// the type

	fmt.Println("Value stored in v before changing = ", v)
	fmt.Println("Address of v = ", &v)
	fmt.Println("Value stored in pointer variable p = ", p)

	// this is dereferencing a pointer
	// using * operator before a pointer
	// variable to access the value stored
	// at the variable at which it is pointing
	fmt.Println("Value stored in v(*p) Before Changing = ", *p)

	// changing the value of y by assigning
	// the new value to the pointer
	*p = 500

	fmt.Println("Value stored in v(*p) after Changing = ", v, "pointer:", p)

	//comparing pointers

	val1 := 2345
	val2 := 567

	// Creating and initializing pointers
	var p1 *int
	p1 = &val1
	p2 := &val2
	p3 := &val1

	// Comparing pointers
	// with each other
	// Using == operator
	res1 := &p1 == &p2
	fmt.Println("Is p1 pointer is equal to p2 pointer: ", res1)

	res2 := p1 == p2
	fmt.Println("Is p1 pointer is equal to p2 pointer: ", res2)

	res3 := p1 == p3
	fmt.Println("Is p1 pointer is equal to p3 pointer: ", res3)

	res4 := p2 == p3
	fmt.Println("Is p2 pointer is equal to p3 pointer: ", res4)

	res5 := &p3 == &p1
	fmt.Println("Is p3 pointer is equal to p1 pointer: ", res5)

	//changing value of vairable with function

	a := 11
	modify(&a) // *a + 10 * (*a)
	fmt.Println(a)

	x = 100

	fmt.Printf("The value of x before function call is: %d\n", x)

	// taking a pointer variable
	// and assigning the address
	// of x to it
	var pa *int = &x

	// calling the function by
	// passing pointer to function
	ptf(pa)

	fmt.Printf("The value of x after function call is: %d\n", x)

}

func modify(a *int) {
	*a = *a + 10**a
}

func ptf(a *int) {

	// dereferencing
	*a = 748
}
