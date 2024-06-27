package main

import (
	"fmt"
	"math"
)

type Computable interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

type Circle[T Computable] struct {
	radius T
}

type Rectangle[T Computable] struct {
	width, height T
}

func (c Circle[T]) Area() T {
	return T( math.Pi * float64( c.radius ) * float64( c.radius ) )
}

func (r *Rectangle[T]) Area() T {
	return r.width * r.height
}

type Figure[T Computable] interface {
	Area() T
}

func getArea[T Computable](f Figure[T]) T {
	return f.Area()
}

func main() {
	rect := Rectangle[float64]{ 15.0, 22.5 }

	figures := []Figure[float64]{ Circle[float64]{ 10.0 }, &rect }

	for _, figure := range figures {
		fmt.Println( getArea( figure ) )
	}

	var i interface{} = "hello"
	s := i.(string)
	fmt.Println(s)
}

