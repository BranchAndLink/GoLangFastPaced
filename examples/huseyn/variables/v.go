package main

import "fmt"

func main() {
	var helloWorldString string = "Hello World"
	var temp1 float64 = 34.2131
	var temp2 float32 = 23.12
	var num1 uint64 = 242
	var num2 = 312
	var st = "Hi !"
	a := true

	fmt.Println(helloWorldString, temp1, temp2, num1, num2, st, a)

	var p1, p2, p3 = 23, "salam", true

	fmt.Println(p2, p3, p1)

	l1, ö2, ğ3 := 34.001, "ğ", 5+4i

	fmt.Println(l1, ö2, ğ3)

	var (
		word1 string = "User1 ❤"
		word2 rune   = '❤'
		word3        = "❤ User3 ❤"
	)

	word1 = "asd"

	fmt.Println(word1, word2, word3)
}
