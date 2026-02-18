---
title: "Flow Controls"
weight: 5
tags: ["Flow", "Controls", "Go"]
---

### Şərt və dövr blokları

## Blok {  } və dəyişənin görünməsi (scope)

Go dilində blok fiqurlu mötərizə ilə bildirilir. Bu həmçinin dəyişənin görünmə dəirəsinə da təsir edir.   
Blok içində ondan üst blokda elan olunmuş dəyişən də istifadə edilir.   
Əgər dəyişən blokun içində təkrar elan edilibsə o blokun xaricindəki dəyişəni kölgəyə salacaq və blok üçün lokal sayılacaq.  



```go
%%
var uzunluq float32 = 4.5
var adamlarin_sayi = 40
var azər = 88 
{
  var uzunluq float32 = 3.5 
  var azər = 45 
  metn := "Sade yazi 1)"
  fmt.Println(metn, uzunluq, adamlarin_sayi, azər) 
}
{
  var uzunluq float32 = 9.5
  var adamlarin_sayi = 44
  var azər = 85 
  metn := "Sade yazi 2) "
  fmt.Println(metn, uzunluq, adamlarin_sayi, azər) 
} 
fmt.Println("Sade yazi Ana blok) ", uzunluq, adamlarin_sayi, azər)
```

    Sade yazi 1) 3.5 40 45
    Sade yazi 2)  9.5 44 85
    Sade yazi Ana blok)  4.5 40 88


Təkrar Qeyd edək ki funksiya xaricində := ilə yox var ilə dəyişən elan edilməlidir


```go
var h = 45
//h:=5 //xeta
func main(){
	fmt.Println(h)
}
```

    45


### If Şərt bloku

<pre>
Bizim burda [ ] işarəsi daxilindəkilər buraxıla bilər

if [Sadə elan kod ";"] müqayisə və bul ifadəsi {
	//şərt ödəndikdə
    kodlar...
} [ else if müqayisə və bul ifadəsi {
	// ikinci şərt ödənildikdə
    kodlar...
} else if müqayisə və bul ifadəsi {
	// 3-cu şərt ödənildikdə
    kodlar...
....
} else {
    // digər hallarda
	kodlar...
} ]

Sadə kod elanı (Simple Statement) yerinə bunlar ola bilər:  
     <b>Boşluq, İfadə, <- Kanal göndərməsi , ++ və ya -- operatoru,  Mənimsətmə, := ilə dəyişən elanı</b>
</pre>

<span style= "color:red">
<pre>
Go dilində fiqurlu mötərizələrin qoyulma qaydası pozula bilməz
Məsələn
if 5==5 {
    fmt.Println("---")
}

aşağıdakı kimi yazsaq xəta olacaq
if 5==5 
{
    fmt.Println("---")
}
</pre>
</span>


```go
%%
if 5==5 {
	fmt.Println("--------")
}
```

    --------



```go
//xeta
if 5==5 
{
	fmt.Println("--------")
}

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

<span class="gonb-cell-line-info">Cell[4]: Line 2</span>
<span class="gonb-error-location">/tmp/gonb_07a99419/main.go:4:1: </span> expected declaration, found 'if'
<div class="gonb-error-context">
package main

//xeta
<div class="gonb-error-line">if 5==5 
</div>{
	fmt.Println(&#34;--------&#34;)

</div>

<br/>

</div>




    in goexec.ExecuteCell(): parsing go files in TempDir "/tmp/gonb_07a99419": /tmp/gonb_07a99419/main.go:4:1: expected declaration, found 'if'


if Nümunələr


```go
%%
if g:=7; g<12 {
	fmt.Println(g,"<", 12)
}
fmt.Println("---Blokdan sonra----")
```

    7 < 12
    ---Blokdan sonra----



```go
%%
if g:=7; g>12 {
	fmt.Println(g,">", 12)
}
fmt.Println("---Blokdan sonra----")
```

    ---Blokdan sonra----



```go
%%
var g = 7
if g>12 {
	fmt.Println(g,">", 12)
} else {
	fmt.Println(g,">", 12, "deyil")
}
fmt.Println("---Blokdan sonra----")
```

    7 > 12 deyil
    ---Blokdan sonra----



```go
%%
var g = 7
if g>12 {
	fmt.Println(g,">", 12)
} else if g>6{
	fmt.Println(g,">", 6)
} 
fmt.Println("---Blokdan sonra----")
```

    7 > 6
    ---Blokdan sonra----



```go
%% 
if g:=7;g>12 {
	fmt.Println(g,">", 12)
} else if g>9{
	fmt.Println(g,">", 9)
} else {
	fmt.Println(g,">", 9, "və ya", g, ">", 12, "deyil")
}
fmt.Println("---Blokdan sonra----")
```

    7 > 9 və ya 7 > 12 deyil
    ---Blokdan sonra----



```go
%%
if gun:=5; gun==1{
	fmt.Println("Bazar ertesi")
} else if(gun==2){
	fmt.Println("Cersenbe axsami")
} else if(gun==3){
	fmt.Println("Cersenbe")
} else if(gun==4){
	fmt.Println("Cume axsami")
} else if(gun==5){
	fmt.Println("Cume")
} else if(gun==6){
	fmt.Println("Senbe")
} else if(gun==7){
	fmt.Println("Bazar")
} else  {
	fmt.Println("Hefte 7 gundur")
}
```

    Cume


#### Switch seçim və şərt bloku

<pre>

switch [Sadə elan ";"] [ifadə] {  //Əgər ifadə yoxdursa bu true kimi qəbul olunur
case ifadə listi:  
 kod...
 [fallthrough]
[
    case ifadə listi:
      kod ...
    ...
]
default: 
  kod...
}

ifadə listi müqayisə və bul operatorları və ya   
Əgər yuxarıda switch-də dəyişən  varsa ona bərabər olanlar listi ola bilər.   
Bu zaman fikir vermək lazımdır ki duplikatlar olmasın, çünki kompilyator tərəfindən qəbul edilməyəcək.

default yəni heç biri olmadıqda həmin hissə yerinə yetirilir  

fallthrough - Əgər bu qeyd edilsə, ondan sonra gələn case -də avtmomatik icra edilir


</pre>

Nümunələrlə misal gətirək


```go
// if g:=7;g>12 {
// 	fmt.Println(g,">", 12)
// } else if g>9{
// 	fmt.Println(g,">", 9)
// } else {
// 	fmt.Println(g,">", 9, "və ya", g, ">", 12, "deyil")
// }

%%
switch g:=7; {
case g > 12:
	fmt.Println(g,">", 12)
case g > 9 :
	fmt.Println(g,">", 9)
default:
	fmt.Println(g,">", 9, "və ya", g, ">", 12, "deyil")
}
```

    7 > 9 və ya 7 > 12 deyil



```go
%%
var g = 10
switch {
case g > 12:
	fmt.Println(g,">", 12)
case g > 9 :
	fmt.Println(g,">", 9)
case g == 7 :
	fmt.Println(g, "7 dir")
default:
	fmt.Println(g,">", 9, "və ya", g, ">", 12, "deyil")
}
```

    10 > 9



```go
%%
// if gun:=5; gun==1{
// 	fmt.Println("Bazar ertesi")
// } else if(gun==2){
// 	fmt.Println("Cersenbe axsami")
// } else if(gun==3){
// 	fmt.Println("Cersenbe")
// } else if(gun==4){
// 	fmt.Println("Cume axsami")
// } else if(gun==5){
// 	fmt.Println("Cume")
// } else if(gun==6){
// 	fmt.Println("Senbe")
// } else if(gun==7){
// 	fmt.Println("Bazar")
// } else  {
// 	fmt.Println("Hefte 7 gundur")
// }
var gun = 5
//switch gun:=5; gun {
switch gun {
case 1:
 	fmt.Println("Bazar ertesi")
case 2:
    fmt.Println("Cersenbe axsami")
case 3:
    fmt.Println("Cersenbe")
case 4:
    fmt.Println("Cume axsami")
case 5:
    fmt.Println("Cume")
case 6:
    fmt.Println("Senbe")
case 7:
    fmt.Println("Bazar")
default:  
 	fmt.Println("Hefte 7 gundur") 
}
```

    Cume



```go
%%
switch gun:=4;gun {
case 1,2,3,4,5:
	fmt.Println("is gunleri")
case 6:
    fmt.Println("Senbe")
case 7:
    fmt.Println("Bazar")
default:  
 	fmt.Println("Hefte 7 gundur") 
}
```

    is gunleri



```go
%%
switch gun:=5;gun {
case 1,2,3,4 :
	fmt.Println("1,2,3,4")
case 5:
	fmt.Println("Cume")
	fallthrough
case 6:
    fmt.Println("Senbe")
	fallthrough
case 7:
    fmt.Println("Bazar")
default:  
 	fmt.Println("Hefte 7 gundur") 
}
```

    Cume
    Senbe
    Bazar


**(Type switch) Tip Switch bloku** və tip təyininə aid nümunə. 
<span style="color:red"> Qeyd interface ve type assertions mövzusunda daha aydın olacaq.</span>

<pre>
Xüsusi Type switch
</pre>
```Go
x interface{} = menimsedilen dəyişən
// adətən funskiya arqumenti kimi de vermek olur
// func tipe_esasen_et(x interface{}) { }

switch v := x.(type) {
case T:
    //  v T tipindedir
case S:
    //  v S tipindedir
default:
    // diger
}
```
<pre>
<b>Burada adi switch-dən fərqli olaraq case x interfeysinə mənimsədilənin tipine uyğun  baş verir</b>
  := ilə həmin tipdə olan dəyişənə mənimsətmək olar
  .(type)  bu switch-də  type assertion x.(T) siktaksisi ilə eyni olsa da x.(type) yəni <b>type sözünün işlətmək lazımdır</b>
  case -də nil olub olmadığını yoxlamaq olur
</pre>

Əvvəlcə if else bloku və type Assertions ilə misala baxaq. Sonra onu Type Switch-lə daha qısa yazaq


```go
%%
var x interface{} = 3.56
v := x  // x in qiymətini bir dəfə alırıq
if v == nil {
	i := v                                 //  i-in tipi x (interface{}) esl tipi olacaq
	fmt.Println("x is nil", i)
} else if i, isInt := v.(int); isInt {
	fmt.Println(i)                            // i-in tipi: int
} else if i, isFloat64 := v.(float64); isFloat64 {
	fmt.Println("float64", i)                        // i-in tipi: float64
} else if i, isFunc := v.(func(int) float64); isFunc {
	fmt.Printf("func(int) float64 : %T\n",i)                       // i-in tipi: func(int) float64
} else {
	_, isBool := v.(bool)
	_, isString := v.(string)
	if isBool || isString {
		i := v                         // i-in tipi: x (interface{}) esl tipi olacaq
		fmt.Println("type  bool ve ya string -dir", i)
	} else {
		i := v                         // i-in tipi: x (interface{}) esl tipi olacaq
		fmt.Println("bilinmeyen tip", i)
	}
}
```

    float64 3.56



```go
%%
var x interface{}= 3.56
//Burada i:= deyiseni olsa da case tipe esasen bas verir
switch i := x.(type) {
case nil:
	fmt.Println("x is nil", i)               //  i-in tipi x (interface{}) esl tipi olacaq
case int:
	fmt.Println(i)                            // i-in tipi: int
case float64:
	fmt.Println("float64", i)                        // // i-in tipi: float64
case func(int) float64:
	fmt.Printf("func(int) float64 : %T\n",i)  // i-in tipi: func(int) float64
case bool, string:
	fmt.Println("type  bool ve ya string -dir")  //  i-in tipi x (interface{}) esl tipi olacaq
default:
	fmt.Println("bilinmeyen tip")     //  i-in tipi x (interface{}) esl tipi olacaq
}
```

    float64 3.56



```go
%%
var i interface{} = "salam"
switch i.(type) {
case int:
	fmt.Printf("Tam eded")
case string:
	fmt.Printf("Metn")
default:
	fmt.Printf("diger tip")
}
```

    Metn

### For dövrlər
<pre>

for [muqayise] { 
     //muqayise buraxildiqda true olduğu üçün dövr daimi olur for true {}. dövrdən çıxma daxildə müəyyən edilə bilinər
    kod...
}
//SadeKod: Sadə kod elanı (Simple Statement) 
for [dövrün əvvəlində icra edilən SadeKod] ; [muqayise]; [blok sonunda olan hər dəfə icra edilən SadeKod]{
    //burada da hər hansı hissə buraxıla bilər. for ;; { } elə for {} for true{} kimi olacaq
    kod...
}
//range ifadesi array string map channel ola bilər
for [deyisenler, := ] range ifade { 
}

</pre>
```Go
//
//string olduqda bu metni emele getiren simvollar olacaq
var metn string
for pozisiya, simvol := range metn { 
}
//pozisiya ASCII-də <b>indeksə</b> uyğun gəlsə də, unikodlu simvollarda <b>byte indeksinə</b> uyğun gələcək
```
<pre>
<span style="color:red">Digər tiplərə  nümunə həmin tipdə göstəriləcək. Diqqət edin ki range-də elementlər kopyalanır(sureti alınır)</span>

break [label]  bu ifadə blok daxilindən çıxmaq üçün istifadə olunur. 
Əgər label olsa çıxma həmin labelə edilir
label bu cür verilir:
// Ad ve iki nöqtə :
Label_adi:
continue [label]  bu ifadə blok daxilindən ondan sonra gələnləri adlayaraq davam etmək üçün istifadə olunur. 
Əgər label olsa həmin labelə edilir


</pre>


```go
%%
i:=0
for i<5 {
	fmt.Print(i, ", ")
	i+=1
}
fmt.Println("\n-------")
```

    0, 1, 2, 3, 4, 
    -------



```go
%%
i:=0
for  {
	fmt.Print(i, ", ")
	i+=1
	if i>4 {
		break;
	}
}
fmt.Println("\n-------")
```

    0, 1, 2, 3, 4, 
    -------



```go
%% 
for i:=0;i<5; i++{
	fmt.Print(i, ", ")
}
fmt.Println("\n-------")
```

    0, 1, 2, 3, 4, 
    -------



```go
%%
for i, x := range "Salam Go" {
	fmt.Printf("%d: %c ",i, x)
}
fmt.Println("\n-------")
```

    0: S 1: a 2: l 3: a 4: m 5:   6: G 7: o 
    -------



```go
%%
for i, x := range "👍❤️❤️alma👍😊" {
	fmt.Printf("%d: %c ",i, x)
}
fmt.Println("\n-------")
```

    0: 👍 4: ❤ 7: ️ 10: ❤ 13: ️ 16: a 17: l 18: m 19: a 20: 👍 24: 😊 
    -------



```go
%% 
for i:=0;i<5; i++{
	fmt.Print(i, ", ")
    if i %2 == 0 {
		continue
	}
	fmt.Print(i, " tek ededdir ")

}
fmt.Println("\n-------")
```

    0, 1, 1 tek ededdir 2, 3, 3 tek ededdir 4, 
    -------



```go
%%
for i:=0;i<2; i++{
	for j:=0;j<5; j++{
		fmt.Print(i*10 + j, " ") 
	}
	fmt.Println("\n-------")
}
fmt.Println("--~~~---")
```

    0 1 2 3 4 
    -------
    10 11 12 13 14 
    -------
    --~~~---



```go
%%
for i:=0;i<2; i++{
	for j:=0;j<5; j++{
		fmt.Print(i*10 + j, " ") 
		if (j>2) {
			break; //j for-undan cixir
		}
	}
	fmt.Println("\n-------")
}
fmt.Println("--~~~---")
```

    0 1 2 3 
    -------
    10 11 12 13 
    -------
    --~~~---



```go
%%
Label1:
for i:=0;i<2; i++{
	for j:=0;j<5; j++{
		fmt.Print(i*10 + j, " ") 
		if (j>2) {
			break Label1; //tam for-dan chixir
			//Qeyd baxmayaraq Label1 for-dan üstdədir   bir daha for-un içinə girmir
		}
	}
	fmt.Println("\n-------")
}
fmt.Println("--~~~---")
```

    0 1 2 3 --~~~---



```go
%%
for i:=0;i<2; i++{
	for j:=0;j<5; j++{
		if (j>2) {
			continue; 
		}
		fmt.Print(i*10 + j, " ") 

	}
	fmt.Println("\n-------")
}
fmt.Println("--~~~---")
```

    0 1 2 
    -------
    10 11 12 
    -------
    --~~~---



```go
%%
Label1:
for i:=0;i<2; i++{
	for j:=0;j<5; j++{
		if (j>2) {
			continue Label1; //ust for-ile davam edir 
			//neticede fmt.Println("\n-------") hec vaxt gorunmur
		}
		fmt.Print(i*10 + j, " ")

	}
	fmt.Println("\n-------")
}
fmt.Println("--~~~---")
```

    0 1 2 10 11 12 --~~~---



```go
%%
for i, j := 0, 10; i < 10; i, j = i+1, j-1 {
	fmt.Print(i , " + ", j, " == 10, ")
}
fmt.Println("\n---------")
```

    0 + 10 == 10, 1 + 9 == 10, 2 + 8 == 10, 3 + 7 == 10, 4 + 6 == 10, 5 + 5 == 10, 6 + 4 == 10, 7 + 3 == 10, 8 + 2 == 10, 9 + 1 == 10, 
    ---------


Goto Şərtsiz keçid.   

goto label  

Fikir vermək lazımdır ki bu zaman hansısa dəyişən elanı kənarda qalmasın   
```Go
goto L  // Xəta
	v := 3
L:
```
Həmçinin bir blokun içindən digər blokun( qonşu blokdursa) içinə necə gəldi atlanma mümkün deyil.
```Go
{
    got L
}
{
    L:
    fmt.Println("__")
}
```

Baxmayaraq goto şərtsiz keçiddir onu şərtlə birləşdirə edə bilərik  



```go
%%
goto L  // Xəta
	v := 3
L:
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


<span style="white-space: pre;"> # gonb_5bb6b08a</span>

<br/>

<span class="gonb-cell-line-info">Cell[249]: Line 2</span>
<span class="gonb-error-location">./main.go:14:7: </span> goto L jumps over declaration of v at ./main.go:15:5
<div class="gonb-error-context">

func main() {
	flag.Parse()
<div class="gonb-error-line">	goto L  // Xəta
</div>		v := 3
	L:

</div>

<br/>

<span class="gonb-cell-line-info">Cell[249]: Line 3</span>
<span class="gonb-error-location">./main.go:15:3: </span> v declared and not used
<div class="gonb-error-context">
func main() {
	flag.Parse()
	goto L  // Xəta
<div class="gonb-error-line">		v := 3
</div>	L:


</div>

<br/>


<span style="white-space: pre;"> </span>

<br/>

</div>




    failed to run "/usr/local/go/bin/go build -o /tmp/gonb_5bb6b08a/gonb_5bb6b08a": exit status 1



```go
%%
{
    goto L
}
{
    L:
    fmt.Println("__")
}
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


<span style="white-space: pre;"> # gonb_5bb6b08a</span>

<br/>

<span class="gonb-cell-line-info">Cell[251]: Line 3</span>
<span class="gonb-error-location">./main.go:16:11: </span> goto L jumps into block starting at ./main.go:18:2
<div class="gonb-error-context">
func main() {
	flag.Parse()
	{
<div class="gonb-error-line">	    goto L
</div>	}
	{

</div>

<br/>


<span style="white-space: pre;"> </span>

<br/>

</div>




    failed to run "/usr/local/go/bin/go build -o /tmp/gonb_5bb6b08a/gonb_5bb6b08a": exit status 1



```go
%%
fmt.Println("a")
    goto SON
    fmt.Println("b")
SON:
    fmt.Println("c")
```

    a
    c


For dövrünü goto və if-lə edək


```go
%%
// for i:=0;i<10;i++{
// 	fmt.Print(i, ", ")
// }

i:=0
start:
	fmt.Print(i, ", ")
	i++
	if i<10 {
		goto start
	}
fmt.Println("\n-----")
```

    0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 
    -----

