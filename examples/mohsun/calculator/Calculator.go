package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func main() {
	var input string
	step := 0
	var cavab float64 = 0.0
	var op string
	fmt.Println("eded:")
	for {
		fmt.Scanln(&input)
		if input == "exit" {
			fmt.Println("Tesekkurler sagolun.")
			break
		} else if input == "clear" {
			c := exec.Command("clear")
			c.Stdout = os.Stdout
			c.Run()
			step = 0
			fmt.Println("eded:")
		} else {
			if step == 1 {
				if input == "+" || input == "-" || input == "*" || input == "/" {
					fmt.Println("operator")
					op = input
					step++
					fmt.Println("eded:")
				} else {
					fmt.Println("operatoru yeniden daxil edin (+, -, *, / )")
				}

			} else {
				eded, xeta := strconv.ParseFloat(input, 64)

				if xeta != nil {
					fmt.Println("Xeta. Yeniden daxil edin\neded:")
					continue
				}

				if step == 0 {
					cavab = eded
					step++
				} else {
					switch op {
					case "+":
						cavab += eded
					case "-":
						cavab -= eded
					case "*":
						cavab *= eded
					case "/":
						cavab /= eded
					}
					//
					fmt.Println("cavab:", cavab)
					step = 1

				}
			}
		}
	}
}
