package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var randomnum = rand.Intn(100)

	var guessNumber int
	fmt.Println("fikrimdeki ededi tapin:")
	for step := 0; step < 10; step++ {
		fmt.Scanln(&guessNumber)

		if guessNumber < randomnum {
			fmt.Println("Siz kicik eded daxil etmisiniz")
		} else if guessNumber > randomnum {
			fmt.Println("Siz boyuk eded daxil etmisiniz")
		} else {
			fmt.Println("Tebrikler, ededi tapdiniz :)")
			break
		}
	}
	fmt.Println("Oyun bitdi")
}
