package main

import (
	"fmt"
)

// The below one is a good blog regarding interface.
// https://blog.knoldus.com/how-to-use-interfaces-in-golang/#:~:text=An%20interface%20in%20Go%20is,that%20enables%20polymorphism%20in%20Go.

func main() {

	number := 15

	// true if number is less than 0
	if number > 0 {
		fmt.Printf("%d is a positive number\n", number)
	}

	fmt.Println("Out of the block")

	number = -6

	if number > 0 {
		fmt.Printf("%d is a positive number\n", number)
	} else {
		fmt.Printf("%d is a negative number\n", number)
	}

	// if we need to make a choice between more than two alternatives, then we use the else if statement.

	number1 := 12
	number2 := 20

	if number1 == number2 {
		fmt.Printf("Result: %d == %d \n", number1, number2)
	} else if number1 > number2 {
		fmt.Printf("Result: %d > %d \n", number1, number2)
	} else {
		fmt.Printf("Result: %d < %d \n", number1, number2)
		// inner if statement
		if number1*number2 > 100 {
			fmt.Println("product of the numbers is:", number2*number1)
		}
	}

	//In Go, the switch statement allows us to execute one code block among many alternatives

	dayOfWeek := 3

	switch dayOfWeek {
	case 1:
		fmt.Println("Sunday")
	case 2:
		fmt.Println("Monday")
	case 3:
		fmt.Println("Tuesday")
		fallthrough
	case 4:
		fmt.Println("Wednesday")
	case 5:
		fmt.Println("Thursday")
	case 6:
		fmt.Println("Friday")
	case 7:
		fmt.Println("Saturday")

	default:
		fmt.Println("Invalid day")
	}

	nameOfDay := "Friday"

	switch nameOfDay {
	case "Saturday", "Sunday":
		fmt.Println("Weekend")

	case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
		fmt.Println("Weekday")

	default:
		fmt.Println("Invalid day")
	}
	//we can also use an optional statement along with the expression
	switch day := 4; day {

	case 1:
		fmt.Println("Sunday")
	case 2:
		fmt.Println("Monday")
	case 3:
		fmt.Println("Tuesday")
	case 4:
		fmt.Println("Wednesday")
	case 5:
		fmt.Println("Thursday")
	case 6:
		fmt.Println("Friday")
	case 7:
		fmt.Println("Saturday")
	default:
		fmt.Println("Invalid Day!")
	}
	// for loop terminates when i becomes 6
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	var n, sum = 10, 0

	for i := 1; i <= n; i++ {
		sum += i
	}

	fmt.Println("sum =", sum)

	//Go while loop
	number = 1

	for number <= 5 {
		fmt.Println(number)
		number++
	}

	multiplier := 1

	// run while loop for 10 times
	for multiplier <= 10 {

		// find the product
		product := 5 * multiplier

		// print the multiplication table in format 5 * 1 = 5
		fmt.Printf("5 * %d = %d\n", multiplier, product)
		multiplier++
	}

	number = 1

	// loop that runs infinitely
	for {

		// condition to terminate the loop
		if number > 5 {
			break
		} // This statement acts as the while clause inside a do...while loop and is used to terminate the loop

		fmt.Printf("%d\n", number)
		number++

	}

	var x interface{} = 3.56
	//Burada i:= deyiseni olsa da case tipe esasen bas verir
	switch i := x.(type) {
	case nil:
		fmt.Println("x is nil", i) //  i-in tipi x (interface{}) esl tipi olacaq
	case int:
		fmt.Println(i) // i-in tipi: int
	case float64:
		fmt.Println("float64", i) // // i-in tipi: float64
	case func(int) float64:
		fmt.Printf("func(int) float64 : %T\n", i) // i-in tipi: func(int) float64
	case bool, string:
		fmt.Println("type  bool ve ya string -dir") //  i-in tipi x (interface{}) esl tipi olacaq
	default:
		fmt.Println("bilinmeyen tip") //  i-in tipi x (interface{}) esl tipi olacaq
	}

	var i interface{} = "world"
	switch i.(type) {
	case int:
		fmt.Println("Type is integer")
	case string:
		fmt.Println("Type is string")
	default:
		fmt.Println("Type is other")
	}

Label1:
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			fmt.Print(i*10+j, " ")
			if j > 2 {
				break Label1 //tam for-dan chixir
				//Qeyd baxmayaraq Label1 for-dan üstdədir   bir daha for-un içinə girmir
			}
		}
		fmt.Println("\n-------")
	}
	fmt.Println("Label 1 --~~~---")

Label2:
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			if j > 2 {
				continue Label2 //ust for-ile davam edir
				//neticede fmt.Println("\n-------") hec vaxt gorunmur
			}
			fmt.Print(i*10+j, " ")

		}
		fmt.Println("\n-------")
	}
	fmt.Println("Label 2 --~~~---")

	anumber := 0
start:
	fmt.Print(anumber, ", ")
	anumber++
	if anumber < 10 {
		goto start
	}
	fmt.Println("\n-----")

	string := "Golang"
	fmt.Println("Index: Character")

	// i access index of each character
	// item access each character
	for index, character := range string {
		fmt.Printf("%d= %c \n", index, character)
	}

}
