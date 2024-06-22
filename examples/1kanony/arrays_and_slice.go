package main

import "fmt"

func main() {
	////////////////////////////////////////////////////////
	// Array
	////////////////////////////////////////////////////////

	a0 := [3]int{10, 1:100, 200}
	var a1 [3]int = [3]int{1, 2, 3}
	var a2 = [3]int{100, 200, 300}
	var a3 = [...]int{1000, 2000, 3000}
	var aa1 = [3][2]int{ {11, 12}, {21,22}, {31,32} }
	var aa2 = [...][2]int{ {1,1}, {2,2}, {3,3} }

	aa2[ 1 ] = [2]int{22, 22}

	fmt.Println(a0, a1, a2, a3, aa1, aa2)
	fmt.Printf( "%T\n", aa2 )

	fmt.Println(len(aa2))
	fmt.Println(len(aa2[0]))

	if [...]int{1,2,3} == [...]int{1,2,3} {
		fmt.Println("Correct")
	}


	////////////////////////////////////////////////////////
	// Slice
	////////////////////////////////////////////////////////

	var s1 []int= []int{0xfa,0xfb,0xfc}
	var s2 = []int{0xfa,0xfb,0xfc}
	s3 := []int{0xfa,0xfb,0xfc}
	fmt.Println(s1, s2, s3)

	s1 = append( s1, 100, 200, 300 )
	fmt.Printf("type: %T, array: %v, len: %d, cap: %d\n", s1, s1, len(s1), cap(s1))

	tempArray := [...]int{1, 2, 3, 4, 5} // getting slice from array
	sliceFromArray := tempArray[2:4]
	fmt.Println(sliceFromArray)

	ss := s1[1:3] // slice from slice
	fmt.Println(ss)

	copy(ss, s1) // note due to len of ss is 2 it will just copy first 2 elements from s1
	fmt.Println(ss)

	copy(ss, s1[2:6])
	fmt.Println(ss)

	copy(ss[0:], s1[1:3])
	fmt.Println(ss)

	sliceFromMake := make([]int, 10, 200)
	fmt.Println(sliceFromMake)

	var nilSlice []int
	fmt.Println( nilSlice == nil )
	nilSlice = append(nilSlice, 10, 20, 30)
	fmt.Println( nilSlice == nil )
}

