package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

func main() {
	var daxilolma1 string
	var daxilolma2 string
	fmt.Println("1-ci ededi daxil edin:")
	fmt.Scanf("%s", &daxilolma1)
	eded1, xeta1 := strconv.ParseFloat(daxilolma1, 64)
	if xeta1 != nil {
		if daxilolma1 == "exit" {
			fmt.Println("Tesekkurler sagolun.")
		} else if daxilolma1 == "clear" {
			c := exec.Command("clear")
			c.Stdout = os.Stdout
			c.Run()
		} else {
			fmt.Println("Zehmet olmasa duzgun ifade daxil edin")
		}
	} else {
		fmt.Println("2-ci ededi daxil edin:")
		fmt.Scanf("%s", &daxilolma2)
		eded2, xeta2 := strconv.ParseFloat(daxilolma2, 64)
		if xeta2 != nil {
			if daxilolma2 == "exit" {
				fmt.Println("Tesekkurler sagolun.")
			} else if daxilolma2 == "clear" {
				c := exec.Command("clear")
				c.Stdout = os.Stdout
				c.Run()
			} else {
				fmt.Println("Zehmet olmasa duzgun ifade daxil edin")
			}
		} else {
			var operator string
			fmt.Println("operator daxil edin:")
			fmt.Scanf("%s", &operator)
			if operator == "+" {
				cavab := eded1 + eded2
				fmt.Println(cavab)
			} else if operator == "-" {
				cavab := eded1 - eded2
				fmt.Println(cavab)
			} else if operator == "*" {
				cavab := eded1 * eded2
				fmt.Println(cavab)
			} else if operator == "/" {
				cavab := eded1 / eded2
				fmt.Println(cavab)
			}
		}
	}

}
