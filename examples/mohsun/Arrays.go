package main

import "fmt"

//An array is a collection of similar types of data

func main() {

	// declare array variable of type integer
	// defined size [size=5]
	var arrayOfInteger = [5]int{1, 5, 8, 0, 3}

	fmt.Println(arrayOfInteger)
	// undefined size
	var arrayOfString = [...]string{"Hello", "Student"}

	fmt.Println(arrayOfString)
	// using shorthand notation
	newArray := [...]string{"Hello", "Shagird"}

	fmt.Println(newArray)
	//how to access the array elements
	languages := [3]string{"Go", "Java", "C++"}

	// access element at index 0
	fmt.Println(languages[0])

	// access element at index 2
	fmt.Println(languages[2])

	//We can also use index numbers to initialize an array
	var arrayOfIntegers [3]int

	// elements are assigned using index
	arrayOfIntegers[0] = 5
	arrayOfIntegers[1] = 10
	arrayOfIntegers[2] = 15
	arrayOfIntegers[2] = 20 // changing the value at index 2
	fmt.Println(arrayOfIntegers)

	// initialize the elements of index 0 and 3 only
	arrayOfSpecificnumbes := [5]int{0: 7, 3: 9}

	fmt.Println(arrayOfSpecificnumbes)

	arrayOfNumbers := [...]int{1, 5, 8, 0, 3, 10}
	// find the length of array using len()
	length := len(arrayOfNumbers)

	fmt.Println("The length of array is", length)
	//Looping through an array
	var age = [...]int{14, 15, 16}
	for i := 0; i < len(age); i++ {
		fmt.Println(age[i])
	}
	//Multidimensional array
	// create a 2 dimensional array
	arrayInteger := [2][2]int{{1, 2}, {3, 4}}

	// access the values of 2d array
	for i := 0; i < len(arrayInteger); i++ {
		for j := 0; j < len(arrayInteger[i]); j++ {
			fmt.Println(arrayInteger[i][j])
		}
	}
}
