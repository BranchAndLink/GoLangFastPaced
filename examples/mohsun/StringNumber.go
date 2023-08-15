package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	strVar := "100"
	//We can use the strconv package's Atoi() function to convert the string into an integer value.
	//Atoi stands for ASCII to integer. The Atoi() function returns two values:
	//the result of the conversion, and the error (if any)
	MyintVar, err := strconv.Atoi(strVar)
	fmt.Println(MyintVar, err, reflect.TypeOf(MyintVar))

	//ParseInt interprets a string s in the given base (0, 2 to 36) and bit size (0 to 64) and returns the corresponding value i.
	//This function accepts a string parameter, convert it into a corresponding int type based on a base parameter. By default, it returns Int64 value.

	intVar, err := strconv.ParseInt(strVar, 0, 8)
	fmt.Println(intVar, err, reflect.TypeOf(intVar))

	intVar, err = strconv.ParseInt(strVar, 0, 16)
	fmt.Println(intVar, err, reflect.TypeOf(intVar))

	intVar, err = strconv.ParseInt(strVar, 0, 32)
	fmt.Println(intVar, err, reflect.TypeOf(intVar))

	intVar, err = strconv.ParseInt(strVar, 0, 64)
	fmt.Println(intVar, err, reflect.TypeOf(intVar))

	//The fmt package provides sscan() function which scans string argument and store into variables.
	//This function read the string with spaces and assign into consecutive Integer variables.
	strVar = "200   armud"
	intValue := 4
	_, error_ := fmt.Sscan(strVar, &intValue)
	fmt.Println(intValue, error_, reflect.TypeOf(intValue))
	//conversion from string to float
	floatNum := "55.66"
	num, err1 := strconv.ParseFloat(floatNum, 64)
	fmt.Println(num, err1, reflect.TypeOf(num))

	//conversion number to string
	numb := 776
	var mystring string
	mystring = fmt.Sprintf("%d", numb)
	fmt.Println(mystring, reflect.TypeOf(mystring))

	var nomre = 109
	var metnliNomre string
	metnliNomre = strconv.Itoa(nomre)
	fmt.Printf("%T metnliNomre type : %T %v \n", nomre, metnliNomre, metnliNomre)

	//converting string number to integer by using ASCII table
	var NumberWithSpace = "99   9   "
	var number int
	for _, char := range NumberWithSpace {
		if char != 32 {
			number = number*10 + (int(char) - 48)
		}
	}
	fmt.Println("number without space: ", number, ", its type: ", reflect.TypeOf(number))

}
