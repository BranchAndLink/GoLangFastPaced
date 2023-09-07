// A map is a data structure that provides you with an unordered collection of key/value pairs (maps are also sometimes called
// associative arrays in Php, hash tables in Java, or dictionaries in Python). Maps are used to look up a value by its associated key.
package main

import (
	"fmt"
	"sort"
)

//Map initialization

var employee = map[string]int{"Mark": 10, "Sandy": 20}

func main() {
	fmt.Println(employee)
	//Empty map declaration
	var employees = map[string]int{}
	fmt.Println(employees)        // map[]
	fmt.Printf("%T\n", employees) // map[string]int

	//Map declaration using make function
	employeeList := make(map[string]int)
	employeeList["Mark"] = 10
	employeeList["Sandy"] = 20
	fmt.Println(employeeList)
	//len of map
	fmt.Println(len(employee))
	fmt.Println(len(employeeList))
	//we can access the items of a map by referring to its key name, inside square brackets.
	fmt.Println(employee["Mark"])
	//adding new elements
	employee["Rocky"] = 30
	employee["Josef"] = 40
	fmt.Println(employee)
	employee["Mark"] = 50 // Edit item
	fmt.Println(employee)
	//built-in delete function deletes an item from a given map associated with the provided key.
	delete(employee, "Mark")
	fmt.Println(employee)
	//Iterate over a Map
	for key, element := range employee {
		fmt.Println("Key:", key, "=>", "Element:", element) //it is printing unordered way
	}

	//There are two methods to clear all items from a Map.
	// Method 1
	for k := range employee {
		delete(employee, k)
	}
	// Method 2
	employee = make(map[string]int)

	//Sort Map Keys
	unSortedMap := map[string]int{"India": 20, "Canada": 70, "Germany": 15}

	keys := make([]string, 0, len(unSortedMap))

	for k := range unSortedMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k, unSortedMap[k])
	}

	//sort map values
	values := make([]int, 0, len(unSortedMap))

	for _, v := range unSortedMap {
		values = append(values, v)
	}

	// Sort slice values.
	sort.Ints(values)
	fmt.Println(values)
	//merging maps

	first := map[string]int{"a": 1, "b": 2, "c": 3}
	second := map[string]int{"a": 1, "e": 5, "c": 3, "d": 4}

	for k, v := range second {
		first[k] = v
	}

	fmt.Println(first)

}
