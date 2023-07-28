package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {

	fmt.Println(reflect.TypeOf(10))       // int
	fmt.Println(reflect.TypeOf(10.1))     // float64
	fmt.Println(reflect.TypeOf(true))     // bool
	fmt.Println(reflect.TypeOf('A'))      // int32 (a rune)
	fmt.Println(reflect.TypeOf("Mansur")) // string

	/*
	 * variable declaration and default values
	 */
	// using var keyword
	var number int
	println(number) // prints 0
	var temp float64
	println(temp) //prints +0.000000e+000
	var name string
	println(name) // prints ''
	var isLogical bool
	println(isLogical) // prints false
	//short variable declaration
	shortNumber := 999
	println(shortNumber)
	shortSize := 10.5
	println(shortSize)
	shortName := "Metosh"
	println(shortName)
	isShort := true
	println(isShort)

	/*
	 * Integer data types
	 */
	//integers
	var i1 int8 = 127
	fmt.Println(i1)
	var i2 int16 = 32767
	fmt.Println(i2)
	var i3 int32 = 2147483647
	fmt.Println(i3)
	var i4 int64 = 9223372036854775807
	fmt.Println(i4)
	//It can be omitted, at least 32bit but not an alias to int32
	var i5 int = 9223372036854775807
	fmt.Println(i5)
	//unsigned integers
	var ui1 uint8 = 255
	fmt.Println(ui1)
	var ui2 uint16 = 65535
	fmt.Println(ui2)
	var ui3 uint32 = 4294967295
	fmt.Println(ui3)
	var ui4 uint64 = 18446744073709551615
	fmt.Println(ui4)
	//at least 32 bit
	var ui5 uint = 9223372036854775807
	fmt.Println(ui5)

	// byte type, it is alias to uint8
	var d byte
	d = 34
	fmt.Println(d)

	n := 15
	// Print hexadecimal representation of number
	fmt.Printf("%X\n", n) //prints 'F'
	fmt.Printf("%x\n", n) //prints 'f'

	// represent hex number in code
	n = 0xF
	fmt.Printf("%X\n", n) //prints 'F'

	// Get binary representation of decimal number
	fmt.Printf("%b", n) // prints 1111

	//Strings
	raw := `spring rain:
			browsing under an umbrella
			at the picture-book store`
	fmt.Println(raw)

	interpreted := "i love spring"
	fmt.Println(interpreted)

	//Rune: rune represent a Unicode code point
	var aRune rune = 'Z'
	fmt.Printf("%c  %U\n", aRune, aRune) //prints Z U+005A
	fmt.Println(aRune)                   // prints 90

	/*
	* Type casting
	 */
	intNum := 23
	intToFloat64 := float64(intNum)
	fmt.Println(reflect.TypeOf(intToFloat64)) //float64

	intToFloat32 := float32(intNum)
	fmt.Println(reflect.TypeOf(intToFloat32)) //float32

	floatNum := 34.0
	floatToToInt := int(floatNum)
	fmt.Println(reflect.TypeOf(floatToToInt)) // int

	toString := strconv.Itoa(intNum)
	fmt.Println(toString)                 // 23
	fmt.Println(reflect.TypeOf(toString)) // string
}
