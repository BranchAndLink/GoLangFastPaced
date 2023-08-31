package main

import "fmt"

/*
  ----DATA TYPES----
	int => tam ededler
	float32 => kesr ededler 4 baytliq
	float64 => kesr ededler 8 baytliq
	complex64, complex128 => kompleks ededler
	string => "" ve ya `` daxilinde yazilanlar
	bool => true ve false deyerlerinden ibaretdir
	byte => 1 baytliq ededler araliq ise [0,255] 
	rune => `` daxilinde yazilan simvol, unikod
	int8 => 1baytliq musbet ve menfi ededler [-128,127]
	uint8 => byte type ile eynidir [0,255]
	int16, int32, int64 => 2,4,8, baytliq tam ededler
	uint16,uint32,uint64 => 2,4,8, baytliq musbet ededler

*/

func main() {

	// Variable adlari a-z A-Z 0-9 _ ve herf olan unikodlar ola biler
	// Variable adlari herfle ve ya _ ile basliya biler
	// Variable adlari herflerin kicik ve boyuk olmasi ile ferqlenir
	// Variable adlari Goland dilinin ozel sozleri ola bilmez
	// Variable var variable_adi variable_type seklinde yazilir
	// Variable varsa mutleq istifade edilmelidir
	// := istifade etmeklede variable yaratmaq olar
	// Ferqleri bunlardi
	// var => funksiya daxilinde ve xaricinde istifade etmek olur. Deyeri sonrada menimsede bilerik
	// := => Yalniz funksiya daxilinde istifade etmek olur. Yarandiqi zaman deyer menimsedilmelidir
	// Menimsetme emeliyyati bunun vasitesi "=" ile olur
	// Riyazi beraberlikle eyni anlami dasimasi ucun variable const olmalidir

	var helloWorldString /*Camel case*/ string = "Hello World"
	var temp1 float64 = 34.2131
	var temp2 float32 = 23.12
	var num1 uint64 = 242
	var num2 = 312
	var StDiv /*Pascal case*/ = "Hi !"
	a_variable /*Snake case*/  := true

	fmt.Println(helloWorldString, temp1, temp2, num1, num2, StDiv, a_variable)

	var p1, p2, p3 = 23, "salam", true

	fmt.Println(p2, p3, p1)

	l1, ö2, ğ3 := 34.001, "ğ", 5+4i

	fmt.Println(l1, ö2, ğ3)

	var (
		word1 string = "User1 ❤"
		word2 rune   = '❤'
		word3        = "❤ User3 ❤"
	)

	word1 = "asd"

	fmt.Println(word1, word2, word3)

	const lorem = 20;
	const ipsum string = "salam";

	const(
		first= "salam"
		second //deyeri verilmediyinden evvelkinin deyerin menimseyir
	)

	fmt.Println(first, second)

	const (
		Bazar  = iota   // 0
		BazarErtesi     // 1
		CersenbeA       // 2
		Cersenbe        // 3
		CumeA           // 4
		Cume            // 5
		Senbe           // 6
	)

	//iota keywordunu elave olaraq Cume = ioata etseydik yenede Cume = 5 olacaqdi
	//iota + her hansisa bir eded yazdiqda bu hesablama gerceklesir

	const (
		Abb = iota + 90   //0 + 90
		Acc = iota + 10   // 1 + 10
	)

	 const zero = iota // zero == 0 beraber olacaq
	fmt.Println(zero, Bazar, BazarErtesi, Cume, Abb, Acc)

	// type (deyisenin_adi) yazaraq variablenin type deyismek mumkundur

	standard := 20.5
	var standard2 int = int(standard)
	
	fmt.Println(standard2)


	fmt.Println(number1,number2,number3,number4)
}

//Burada ardicilliq variableye gore yox asililliga gore olur
// Qeyd funksiya xaricindinde yaradilan variableden gedir sohbet
var(
	number1 int = number2
	number3 int = number4
	number2 int = Hesabla()
	number4 int = 20
)

func Hesabla() int {
	number4 ++
	return number4
}