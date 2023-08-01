---
title: " Functions Intro"
author: "abdelrauf"
date: 2023-07-31
tags: ["Functions", "Intro", "Go" ]
---

### Funksiyalar ( Functions)  Sadə Giriş
Funksiyalar proqrami sadələşdirir və eyni işi gördüyündən bir neçə dəfə çağırmağa imkan verir.  


```Go
func FunksiyaninAdı ( [parametr listi]) [nəticə listi] {
    kodlar...
    return [nəticə listi]
}
```

Burada:
 - parametr listi arqument adları və tipləridir   
 - nəitcə listi də funksiyanın qaytardığı tiplər listidir   
 nəticədə həmçinin qaytarılan adları qeyd edilə bilər.   
 Bu zaman həmin dəyişənlər mənimsədildikdə return-də qeyd edilmir.   
 - return funksiyadan çağrılan bloka qayıtmaq və nəticəni qaytarmaq üçündür.   
 - { } içində funksiyanın görəcəyi kodlar olur və buna funksiyanın özəyi bədəni deyilir (Function body)


Parametr listi 
  - boşluq - bu zaman funksiya arqumentsiz olur   
  - arg1 tip1, arg2 tip2, ..    
  - arg1, arg2, arg3,.. tipi // hamısı eyni tip   
  - args ...tip  //variadic arqument ... 3 nöqtə ilə verilir , istənilən sayda eyni tipli arqument. Slice kimi baxıla bilər   

Nəticə listi
  - boşluq - bu zaman funskiya heç nə qaytarmır     
  - tip1  -  funksiya tip1 tipində nəticə qaytarır. return tip1_tipində_nəticə     
  - (tip1, tip2, ..) -  funksiya çoxlu tipdə (tuple) nəticə qaytarır return tip1_tipində_nəticə, tip2_tipində_nəticə, ..    
  - (ad1 tip1, ad2 tip2, ..) - funksiya adlı nəticəsi olduğundan həmin dəyişənlərə mənimsətmək və return ilə geri qaytarmaq olur   
  - (ad1 , ad2, ad3, ... tip) - yuxarıdakı kimi, sadəcə eyni tip olduğundan qısa yazılış    



Funskiya Çağrıldıqda 

[neticeler, .. = ]FunksiyaninAdı ( [arqumentlər listi])

Nümunələrə baxaq


```go
import "fmt"
func salamDe(){
	fmt.Println("Salam")
}

%%
salamDe()
salamDe()
salamDe()
```

    Salam
    Salam
    Salam



```go
func salamDe( ad string) {
	fmt.Println("Salam", ad)
}

%%
salamDe("Dunya")
salamDe("Go")
```

    Salam Dunya
    Salam Go



```go
func salamDe( ad string) string{
	fmt.Println("Salam", ad)

	return "Wa aleykum assalam "
}

%%
netice := salamDe("Dunya")
fmt.Println(netice)
```

    Salam Dunya
    Wa aleykum assalam 


Go dilində **funksiya overlaoding yoxdur**. Ona görə eyni adlı fərqli arqumentlər götürən funksiyalar mümkün deyil.   
Ya onları birləşdirməliyik ya da adlarını dəyişdirməliyik


```go
func salamDe(){
	fmt.Println("Salam")
}
func salamDe( ad string) string{
	fmt.Println("Salam", ad)

	return "Wa aleykum assalam "
}
%%
netice := salamDe("Dunya")
fmt.Println(netice)
salamDe()
```



<style>
.gonb-error-location {
	background: var(--jp-error-color2);  
	border-radius: 3px;
	border-style: dotted;
	border-width: 1px;
	border-color: var(--jp-border-color2);
}
.gonb-error-location:hover {
	border-width: 2px;
	border-style: solid;
	border-color: var(--jp-border-color2);
}
.gonb-error-context {
	display: none;
}
.gonb-error-location:hover + .gonb-error-context {
	background: var(--jp-dialog-background);  
	border-radius: 3px;
	border-style: solid;
	border-width: 1px;
	border-color: var(--jp-border-color2);
	display: block;
	white-space: pre;
	font-family: monospace;
}
.gonb-error-line {
	border-radius: 3px;
	border-style: dotted;
	border-width: 1px;	
	border-color: var(--jp-border-color2);
	background-color: var(--jp-rendermime-error-background);
	font-weight: bold;
}
.gonb-cell-line-info {
	background: var(--jp-layout-color2);
	color: #999;
	margin: 0.1em;
	border: 1px solid var(--jp-border-color1);
	padding-left: 0.2em;
	padding-right: 0.2em;
}
</style>
<div class="lm-Widget p-Widget lm-Panel p-Panel jp-OutputArea-child">
<div class="lm-Widget p-Widget jp-RenderedText jp-mod-trusted jp-OutputArea-output" data-mime-type="application/vnd.jupyter.stderr" style="font-family: monospace;">


<span style="white-space: pre;"> # gonb_ca8a2206</span>

<br/>

<span class="gonb-cell-line-info">Cell[24]: Line 12</span>
<span class="gonb-error-location">./main.go:49:2: </span> not enough arguments in call to salamDe
<div class="gonb-error-context">
	flag.Parse()
	netice := salamDe(&#34;Dunya&#34;)
	fmt.Println(netice)
<div class="gonb-error-line">	salamDe()
</div>
}

</div>

<br/>


<span style="white-space: pre;"> 	have ()</span>

<br/>


<span style="white-space: pre;"> 	want (string)</span>

<br/>


<span style="white-space: pre;"> </span>

<br/>

</div>




    failed to run "/usr/local/go/bin/go build -o /tmp/gonb_ca8a2206/gonb_ca8a2206": exit status 1



```go
func salamDe1(){
	fmt.Println("Salam")
}
func salamDe2( ad string) string{
	fmt.Println("Salam", ad)

	return "Wa aleykum assalam "
}
%%
netice := salamDe2("Dunya")
fmt.Println(netice)
salamDe1()
```

    Salam Dunya
    Wa aleykum assalam 
    Salam



```go
func topla( a int,b int,c int) int {
	return a +b +c
}

%%
fmt.Println(topla(5,6,7))
```

    18



```go
func topla( a,b,c int) int {
	return a +b +c
}

%%
fmt.Println(topla(5,6,7))
```

    18



```go
func salamDeVeTopla(ad string, a,b,c int) (string,  int){
    fmt.Println("salam", ad)
	return "Wa aleykum assalam. \nNeticemiz: ",  a +b +c
}

%%
fmt.Println (salamDeVeTopla("Go", 5, 6, 7) )
```

    salam Go
    Wa aleykum assalam. 
    Neticemiz:  18



```go
func toplaVeVurAyriliqda( a, b int) (int, int){
	return a+b, a*b
}

%%
toplama, vurma := toplaVeVurAyriliqda(8,9)
fmt.Println("Toplama:",toplama,"Vurma:", vurma)
```

    Toplama: 17 Vurma: 72



```go
func toplaVeVurAyriliqda( a, b int) (toplama_netice int, vurma_netice int){
	toplama_netice = a+b
	vurma_netice = a*b
	return
}

%%
toplama, vurma := toplaVeVurAyriliqda(8,9)
fmt.Println("Toplama:",toplama,"Vurma:", vurma)
```

    Toplama: 17 Vurma: 72



```go
func toplaVeVurAyriliqda( a, b int) (toplama_netice, vurma_netice int){
	toplama_netice = a+b
	vurma_netice = a*b
	return
}

%%
toplama, vurma := toplaVeVurAyriliqda(8,9)
fmt.Println("Toplama:",toplama,"Vurma:", vurma)
```

    Toplama: 17 Vurma: 72


Variadic arqumentlə. Burada array və slice-ı vermək qaydasını da göstərəcəyik. Array və slice o biri dərslərdə olacaq.


```go
func topla( ededler ... int) int{
	netice:=0
	for _, k := range ededler {
		netice += k
	}
	return netice
}

%%
fmt.Println(topla())
fmt.Println(topla(3))
fmt.Println(topla(3,4))
fmt.Println(topla(3,4,5,6,7,8,9))

//Array Ve Slice Variadic function-a vermek qaydasi.
//Array ad[:]...
var arr [5]int = [5]int{5,6,7,8,9}
fmt.Println(topla(arr[:]...))
//Slice   ad...
var slice []int = []int{5,6,7,8,9}
fmt.Println(topla(slice...))
```

    0
    3
    7
    42
    35
    35


Funksiyanı digər funksiya dəyişəninə mənimsədikdə bir növ onun alias-ını sinonimini almış oluruq


```go
func topla( a int,b int,c int) int {
	return a +b +c
}

%%
ustegel:=topla
yaz:=fmt.Println
//var ile elan etmek
var addition func(int,int,int)int = topla
yaz(ustegel(5,6,8))
yaz(addition(5,6,8))
```

    19
    19


Go dilində dillə hazır gələn funksiyalar da var. Builtin functions.  
Onları müvafiq bölmələrdə öyrənəcəyik <span style="color:red">len, close, make, new, append, copy, cap, delete, real, complex, imag, panic, recover, print, println  ve s</span>


```go
%%
print("builtin print")
```

    builtin print


```go
%%
var g = 4 + 9i
print(imag(g))
print("\n")
println(real(g))
println("-----")
```

    +9.000000e+000
    +4.000000e+000
    -----


Qeyd. Funksiyada arqumentler by copy  (kopyası sürətilə) ötürülür. Yəni arqument dəyişildikdə ona heçnə olmur.   
Ona görə funksiyadan kənarda olan dəyişəni dəyişmək üçün onun addresini Pointer istifadə etməklə ötürməliyik.   
Amma yenə bu zaman da address kopya ilə ötürülür. Qaydada həminki qaydadır. Biz pointeri dəyişmirik.


```go
func Funksiya(a int){
	a+=8
	fmt.Println(a)
}

%%
var g = 8
fmt.Println(g)
Funksiya(g)
fmt.Println(g)
```

    8
    16
    8



```go
//Pointer gələcəkdə keçəcəyik
func Funksiya(a *int){
	*a+=8
	fmt.Println(*a)
}

%%
var g = 8
fmt.Println(g)
Funksiya(&g)
fmt.Println(g)
```

    8
    16
    16



```go
func Funksiya( a [3]int){
	a[0] = a[0] + a[1] + a[2]
	fmt.Println(a)
}

%%
g:=[3]int{5,6,7}
fmt.Println(g)
Funksiya(g)
fmt.Println(g)


```

    [5 6 7]
    [18 6 7]
    [5 6 7]



```go
func Funksiya( a *[3]int){
	//array de (*a)[0] yazmayib qisa yazmaqla da Go basha dushur. a[0]
	a[0] = a[0] + a[1] + (*a)[2]
	fmt.Println(*a)
}

%%
g:=[3]int{5,6,7}
fmt.Println(g)
Funksiya(&g)
fmt.Println(g)
```

    [5 6 7]
    [18 6 7]
    [18 6 7]


ve ya pointer to array əvəzinə slice işlətsək


```go
func Funksiya( a_slice []int){ 
	a_slice[0] = a_slice[0] + a_slice[1] + a_slice[2]
	fmt.Println(a_slice)
}

%%
g:=[3]int{5,6,7}
fmt.Println(g)
//chagirdiqda [:] ile slice
Funksiya(g[:])
fmt.Println(g)
```

    [5 6 7]
    [18 6 7]
    [18 6 7]

