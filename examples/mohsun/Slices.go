package main

import "fmt"

func main() {
	//how we create a slice
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(numbers)
	//If we provide size inside the [] notation, it becomes an array.
	arr := [3]int{1, 2, 3} //array
	slc := []int{1, 2, 3}  //slice
	fmt.Println(arr, slc)
	number_arr := [8]int{10, 20, 30, 40, 50, 60, 70, 80}
	sliceNumbers := number_arr[4:7]
	fmt.Println(sliceNumbers)
	new_slice := []int{1, 2, 3}
	//copying slices
	copy(sliceNumbers, new_slice)
	fmt.Println("newslice:", new_slice, "old slice:", sliceNumbers)
	primeNumbers := []int{2, 3}

	// add elements 5, 7 to the slice
	primeNumbers = append(primeNumbers, 5, 7, 11, 13)
	fmt.Println("Prime Numbers:", primeNumbers)
	evenNumbers := []int{2, 4}
	oddNumbers := []int{1, 3}

	// add elements of oddNumbers to evenNumbers
	evenNumbers = append(evenNumbers, oddNumbers...)
	fmt.Println("evenNumbers:", evenNumbers)

	//length and capacity of slice
	fmt.Println("len:", len(evenNumbers), "cap:", cap(evenNumbers))
	//lets see what happens when the number of slice items exceeds its capacity.

	var s []int
	for i := 0; i < 10; i++ {
		fmt.Printf("length: %d, capacity: %d, address: %p\n", len(s), cap(s), s)
		s = append(s, i)
	}
	//As you can see in the output, every time the length of the slice is beyond its
	//capacity (the length of the underlying array), the append() function expands the slice by allocating a new underlying
	//array of twice its size and copying all of its elements there.
	//Notice that the pointer to the underlying array changes with each change in capacity.

	//Create a slice using make() function
	newnumbers := make([]int, 5, 7)
	//5 is the length of the slice (number of elements present in the slice)
	//7 is the capacity of the slice (maximum size up to which a slice can be extended)
	fmt.Println("Numbers:", newnumbers)
	// add elements to numbers
	newnumbers[0] = 13
	newnumbers[1] = 23
	newnumbers[2] = 33
	newnumbers = append(newnumbers, 40, 41, 42, 43, 44, 45, 46, 47)
	fmt.Println("Numbers:", newnumbers)
	fmt.Printf("length: %d, capacity: %d\n", len(newnumbers), cap(newnumbers))
	//here we see that once we increase slice length the capacity has increased twice (from 7 to 14).
	//here we can also see that append function adds new elements beginning from 6th index.

	// Program that loops over a slice using for range loop

	floatNum := []float64{2.2, 4.8, 6.0, 8.9, 10}

	// for range loop that iterates through a slice
	for index, value := range floatNum {
		fmt.Println("index:", index, "value:", value)
	}

}
