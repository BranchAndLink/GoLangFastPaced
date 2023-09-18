package main

import "fmt"

//Generics in Go are implemented using type parameters, which allow for the creation of generic functions and data structures
//that can operate on different types without the need for explicit type conversions.

func generic_circumference[r int | float32](radius r) {
	c := 2 * 3 * radius
	fmt.Println("The circumference is: ", c)
}

func Print[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func Sum[T []int](slice T) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum
}

func is64Bit[T float64|float32|int](v T) bool {
    switch (interface {})(v).(type) {
        case float32:
            return false
        case float64:
            return true
    }
	return false
}

func main() {
	var radius1 int = 8
	var radius2 float32 = 9.5

	generic_circumference(radius1)
	generic_circumference(radius2)

	Print([]string{"a", "b", "c"})
	Print([]int{1, 2, 3})
	fmt.Println("Sum of the list elem:", Sum([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
	fmt.Println(is64Bit(8.9))

}
