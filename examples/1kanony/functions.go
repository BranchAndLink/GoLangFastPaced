package main

import "fmt"

func greet() string {
	return "salam"
}

func greetExtended(isEng bool) string {
	if isEng{
		return "hello"
	}

	return greet()
}

func tupleReturner() (int, float32){
	return 10, 10.09
}

func tupleReturnerWithLocalVars() (localInt int, localFloat32 float32){
	localInt = 10
	localFloat32 = 10.09

	return
}

func shorthandForArg(a, b int, c float32) float64{
	return float64(float32(a * b) - c)
}

func variadic(list...int) int{
	sum := 0
	for _, e := range list{
		sum += e
	}

	return sum
}

func pointerArged(a *int){
	*a = *a * *a
}

func f2() func(int)int{
	var state int = 1

	return func(n int)int{
		state *= n
		return state
	}
}

func rec(n int)int{
	if n > 1{
		return n + rec(n - 1)
	} else {
		return n
	}
}

func main() {
	fmt.Println(greet())
	fmt.Println(greetExtended(true))
	fmt.Println(tupleReturner())
	fmt.Println(tupleReturnerWithLocalVars())
	fmt.Println(shorthandForArg(100, 200, 5.4))

	fmt.Println(variadic([]int{10, 20, 30}...))

	array := [...]int{100, 200, 300}
	fmt.Println(variadic(array[:]...))

	variadicAliased1 := variadic
	var variadicAliased2 func(...int)int = variadic

	fmt.Println(variadicAliased1([]int{ 5, 21 }...))
	fmt.Println(variadicAliased2([]int{ 5, 22 }...))

	a := 100
	pointerArged(&a)

	fmt.Println(a)

	f := func (a int) int {
		return a * a
	}

	fmt.Println(f(10))

	state := f2()

	fmt.Println(state(100))
	fmt.Println(state(200))

	fmt.Println(rec(5))
}

