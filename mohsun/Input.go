package main

import "fmt"

//In Go, we use the scan() function to take input from the user

func main() {
	var name string
	var age int

	// takes input value for name
	fmt.Println("Enter your name and age: ")
	fmt.Scan(&name, &age)

	fmt.Printf("Name: %s\n Age: %d\n", name, age)

	//We use the Scanln() function to get input values up to the new line. When it encounters a new line, it stops taking input values. For example,
	var newName string
	var newAge int
	fmt.Println("Enter your new ame and new age:")
	fmt.Scanln(&newName, &newAge)
	fmt.Printf("New Name: %s\n New Age: %d\n", newName, newAge)
	// in Scanln if we want to take input values for age, we must provide values separated by space.
	fmt.Scanf("%s %d", &name, &age)

	var temperature float32
	var sunny bool

	// take float input
	fmt.Println("Enter the current temperature:")
	fmt.Scanf("%g", &temperature)

	// take boolean input
	fmt.Println("Is the day sunny?")
	fmt.Scanf("%t", &sunny)

	fmt.Printf("Current temperature: %g\nIs the day Sunny? %t", temperature, sunny)
}
