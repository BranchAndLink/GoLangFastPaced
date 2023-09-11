// we use interfaces to store a set of methods without implementation
// To implement an interface, a struct should provide implementations for all methods of an interface
package main

import "fmt"

// interface
type Shape interface {
	area() float32
}

// struct to implement interface
type Rectangle struct {
	length, breadth float32
}

// use struct to implement area() of interface. Here, we are trying to implement this interface by the Rectangle struct
func (r Rectangle) area() float32 {
	return r.length * r.breadth
}

//In Go, more than 1 struct can also implement a single interface.

// Triangle struct implements the interface
type Triangle struct {
	base, height float32
}

// Triangle provides implementation for area()
func (t Triangle) area() float32 {
	return 0.5 * t.base * t.height
}

// access method of the interface
func calculate(s Shape) float32 {
	return s.area()
}

// main function
func main() {

	// assigns value to struct members
	rect := Rectangle{7, 4}

	// call calculate() with struct variable rect
	fmt.Println("with rectangle struct method:", rect.area())

	tri := Triangle{8, 12}

	// call calculate() with struct variable rect
	rectangleArea := calculate(rect)
	fmt.Println("Area of Rectangle:", rectangleArea)

	triangleArea := calculate(tri)
	fmt.Println("Area of Triangle:", triangleArea)

	//both structs provide implementation for the area() method

}

//When a struct implements an interface, it should provide an implementation for all the methods of the interface.
//If it fails to implement any method, we will get an error.
