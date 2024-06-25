package main

import "fmt"

type IntegerAlias = int
type NewIntegerType int

func newTypeInAction( integerType NewIntegerType ) NewIntegerType {
	return integerType * integerType
}

type Person struct {
	name string
	age int
	friends map[int]*Person
}

func (p *Person) unfriendAll() {
	p.friends = map[int]*Person{}
}

func main() {
	var i IntegerAlias = 10

	fmt.Println( newTypeInAction( NewIntegerType(i) ) )

	simon := Person{ name: "Simon", age: 25, friends: nil }

	adam := Person{
		"Adam",
		21,
		map[int]*Person{ 1: &simon },
	}

	fmt.Println( adam.friends )
	adam.unfriendAll()

	fmt.Println( adam.friends )
}

