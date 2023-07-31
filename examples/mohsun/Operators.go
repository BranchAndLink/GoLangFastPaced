package main

import "fmt"

/* Go programming provides wide range of operators that are categorized into following major categories:

Arithmetic operators
Assignment operator
Relational operators
Logical operators */

func main() {
	num1 := 9
	num2 := 12
	// + adds two variables
	sum := num1 + num2
	fmt.Printf("%d + %d = %d\n", num1, num2, sum)
	// - subtract two variables
	difference := num1 - num2
	fmt.Printf("%d - %d = %d\n", num1, num2, difference)
	// * multiply two variables
	product := num1 * num2
	fmt.Printf("%d * %d = %d\n", num1, num2, product)
	num3 := 11
	num4 := 4

	// / divide two integer variables
	quotient := num3 / num4
	fmt.Printf(" %d / %d = %d\n", num3, num4, quotient)
	fmt.Printf("If we want the actual result we should always use the / operator with floating point numbers. For example: \n")
	{
		num3 := 11.0
		num4 := 4.0
		quotient := num3 / num4
		fmt.Printf(" %g / %g = %g\n", num3, num4, quotient)
	}

	result := num3 + num4*num2/num1
	fmt.Println("result of mixed operation:", result)

	num := 5
	// increment of num by 1
	num++
	fmt.Println(num) // 6

	// decrement of num by 1
	num--
	fmt.Println(num) // 5

	fmt.Println("assinment operations")
	num += 2
	fmt.Println(num)
	num -= 2
	fmt.Println(num)
	num *= 2
	fmt.Println(num)
	num /= 2
	fmt.Println(num)
	num %= 2
	fmt.Println(num)
	num <<= 2
	fmt.Println(num)
	num >>= 2
	fmt.Println(num)

	fmt.Println("Relational and Logigal Operators")
	number1, number2, number3 := 12, 20, 6
	var boolResult bool

	// equal to operator
	boolResult = (number1 == number2)

	fmt.Printf("%d == %d returns %t \n", number1, number2, boolResult)

	// not equal to operator
	boolResult = (number1 != number2)

	fmt.Printf("%d != %d returns %t \n", number1, number2, boolResult)

	// greater than operator
	boolResult = (number1 > number2)

	fmt.Printf("%d > %d returns %t \n", number1, number2, boolResult)

	// less than operator
	boolResult = (number1 < number2)

	fmt.Printf("%d < %d returns %t \n", number1, number2, boolResult)

	// returns false because number1 > number2 is false
	boolResult = (number1 > number2) && (number1 == number3)

	fmt.Printf("Result of AND operator is %t \n", boolResult)

	// returns true because number1 == number3 is true
	boolResult = (number1 > number2) || (number1 == number3)

	fmt.Printf("Result of OR operator is %t \n", boolResult)

	// returns false because number1 == number3 is true
	boolResult = !(number1 == number3)

	fmt.Printf("Result of NOT operator is %t \n", boolResult)

	shiftNum, shiftingNum := 212, 2

	fmt.Printf("Right Shift by %d: %d\n", shiftingNum, shiftNum>>shiftingNum)

}
