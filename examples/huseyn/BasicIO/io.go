package main

import "fmt"

func main() {
	fmt.Println("Hello Admin")
	fmt.Println("Hello", "Huseyn")
	fmt.Println(100, 10, 1)

	fmt.Print("No", "Space") // eyni isi gorur lakin bosluq ve yeni setir olmur
	fmt.Print("No", "Tab")

	fmt.Printf("idk") // bunu basa dusdum sadece neye lazimdir anlamadim
	fmt.Printf("👍✔❤")

	// stdin - Standart Input
	// stdout - Standart Output
	// stderr - Standart Error
	// Bunlari stream'de adlandira bilerik

	// Axinlar adi halda bele baglanir
	// stdin - keyboard
	// stdout - console
	// stderr - console

	//Niye adi halda? Cunki Shell'de command ile bunu ferqli
	//yere baglamaq mumkundur

	// ./BasicIO <stdin.txt 1>stdout.txt 2>stderr.txt

	// Go dilinde ise fmt standart kitabxanasi ile output mumkundur
	// yuxarda bunu numune var

	//Yeni setir ve bosluq elave etmek istedikde
	/*
		' ' - bosluq
		\n  - yeni setir
		\\  - geri slesh
		\t  - tab
		\n\r  - yeni setir (windows)
		\" - dirnaq
	*/

	//BASIC INPUT

	// Bunun ucunde fmt standart kitabxanadan istifade edeceyik

	var a int
	var b int
	var c int
	var d int

	input := "10 \n 20"

	fmt.Sscan(input, &a, &b)
	// burada inputu deyisenle gonderdik cunki goda input islemediyinden bele edirik
	fmt.Println("a:",a, "b:",b)

	fmt.Sscanln(input, &c, &d) 
	fmt.Println("c:",c, "d:",d)
	// buradad eynidir lakin bu sefer \n den sonrani yeni
	// yeni setiri oxumayacaq
	//output: c: 10 d: 0

	fmt.Sscanf(input, "%d %f", &c, &d)
	//Buda birinci yazdiqimiz sscan kimidir
	//lakin burada daha deqiq olaraq format vermek olur
	
}
