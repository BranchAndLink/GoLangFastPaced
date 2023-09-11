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


	fmt.Scan(&a, &b)
	// burada inputu deyisenle gonderdik cunki goda input islemediyinden bele edirik
	fmt.Println("a:",a, "b:",b)

	fmt.Scanln(&c, &d) 
	fmt.Println("c:",c, "d:",d)
	// buradad eynidir lakin bu sefer \n den sonrani yeni
	// yeni setiri oxumayacaq
	//output: c: 10 d: 0

	fmt.Scanf("%d %f", &c, &d)
	//Buda birinci yazdiqimiz sscan kimidir
	//lakin burada daha deqiq olaraq format vermek olur
	

	//Userlerin maasi
	fmt.Print("Huseynin maashi: ")
	fmt.Scan(&a)
	fmt.Print("Abdulkerimin maashi: ")
	fmt.Scan(&b)
	fmt.Print("Rufet maashi: ")
	fmt.Scan(&c)
	fmt.Printf("%-15s | %-15s | %-15s |  %-10s \n","Ad", "Soyad", "Shirket", "Maash")
	fmt.Println("────────────────────────────────────────────────────────────")
	fmt.Printf("%-15s | %-15s | %-15s |  %5d \n","Huseyn", "Huseynli", "Smart Scoring", a)
	fmt.Println("────────────────────────────────────────────────────────────")
	fmt.Printf("%-15s | %-15s | %-15s |  %5d  \n","Abdulkerim", "Eliyev", "Asan Yol", b)
	fmt.Println("────────────────────────────────────────────────────────────")
	fmt.Printf("%-15s | %-15s | %-15s |  %5d  \n","Rufet", "Recebov", "SOCAR", c)
	fmt.Println("────────────────────────────────────────────────────────────")
}
