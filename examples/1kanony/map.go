package main

import "fmt"

func main() {
	var xerite map[string]int = map[string]int{ "Baki": 1, "Sumqayit": 2 }

	xerite2 := make( map[string][]int, 10 )
	xerite2["Gence"] = append(xerite2["Gence"], 10, 20, 30)

	fmt.Println(xerite)
	fmt.Println(xerite2)

	value, exists := xerite2["Gence"]
	fmt.Println( value, exists )

	_, exists2 := xerite2["Quba"]

	if !exists2 {
		xerite2["Quba"] = []int{10, 33, 213, 44}
	}
	fmt.Println(xerite2)

	delete(xerite2, "Gence")
	fmt.Println(xerite2)
}

