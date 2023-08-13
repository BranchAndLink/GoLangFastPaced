---
title: " Basic Interface Type Assertions"
author: "abdelrauf"
date: 2023-08-13
tags: ["Basic", "Interface", "Type", "Assertions", "Go" ]
---

### Basic Interface (Sadə interfeys)

Sadə interfeys dedikdə  metodlar siyahısı (boş siyahı daxilolmaqla) nəzərdə tutulur.   
Bu metodları əgər tip üçün müəyyən edilibsə onda həmin tip interfeysin çoxluğunda olur.  
Məhz bu cəhətdən boş interfeysin tiplər çoxluğu bütün tiplərdir.  
<pre>
[type AD] interface {
    [metod1..]
    [metod2...]
    ...
}
Burada:
[ ] - məcburi olmayan yazı
metod - funksiyalar (func açar sözsüz).   
Qeyd: Funksiya adı _ (blank) və təkrar olmamalıdır

</pre>

```Go
interface {
	Read([]byte) (int, error)
	Write([]byte) (int, error)
	Close() error
}

interface {
	String() string
	String() string  // qadağandır: String adı unikal deyil
	_(x int)         // qadağandır: metod adı _ olmamalıdır
}
```


Bildirdiyimiz kimi, Hər hansı tip üçün interfeysdə əks olunan metodları bəyan edilibsə o həmin interfeysin   
çoxluğuna daxil olur.   

```Go
type Locker interface {
	Lock()
	Unlock()
}
// T tipi

func (p T) Lock() { … }
func (p T) Unlock() { … }

```

Burada T tipi Locker interfeysinə tam uyğun gəlir.  

#### Boş interfeys  

Yuxarıda qeyd etdiyimiz kimi heç bir metodu olmayan boş sade interfeysdir.  
Bütün tiplər bu interfeysin çoxluğunda olur. 
```Go
interface{}
```
**any** boş interfeys tip üçün alias (sinonim) kimi istifadə edilir.

#### Interfeysin istifadəsi.   

Yuxarıda qeyd etdik ki, interfeysin metodlarının hamısını bəyan edən tip onun çoxluğunda olur.   
Bu o deməkdir ki, biz interfeys tipində olan dəyişənə həmin tipdə olan qiyməti mənimsədə bilərik.  
Bu tiplərə interfeysin **dinamik tip**ləri deyilir.
İnterfeysin çoxluğunda bir neçə tip olduğundan onlar üçün polimorfizm yaratmış oluruq.  
Yəni bu tiplə bir neçə tipi eyni qaydada işlədə bilirik.    
Sadə halda interfeysin qiyməti nil olur və həmçinin qeyd edək ki, interfeys reference tip-dir.   

```Go

var x interface{}  // x qiyməti nil-dir və statik tipi interface{}
var v *T           // v qiyməti nil, statik tipi *T
x = 42             // x qiyməti 42, dinamik tipi int
x = v              // x qiyməti (*T)(nil), dinamik tipi *T

```

Dediklərimizə nümunələr göstərək.  


```go
type Hendese interface {
    perimetr() float32
    sahe() float32 
    fiqurun_adı() string
}


type Kvadrat struct {
    teref float32
}

type Duzbucaqli struct {
    uzunluq float32
    en float32
}

const (
    Pi float32 = 3.14
)

type Cevre struct {
    radius float32
}

func (p Kvadrat) perimetr() float32 {
    return 4.0 * p.teref
}

func (p Kvadrat) sahe() float32 {
    return p.teref * p.teref
}

func (p Kvadrat) fiqurun_adı() string {
    return "Kvadrat"
}

func (p Duzbucaqli) perimetr() float32 {
    return p.uzunluq * p.en
}

func (p Duzbucaqli) sahe() float32 {
    return p.uzunluq * p.en
}

func (p Duzbucaqli) fiqurun_adı() string {
    return "Düzbucaqlı"
}

func (p Cevre) perimetr() float32 {
    return 2.0 * Pi * p.radius
}

func (p Cevre) sahe() float32 {
    return Pi * p.radius * p.radius
}

func (p Cevre) fiqurun_adı() string {
    return "Çevrə Dairə"
}


%%

kv := Kvadrat{teref:44}
dz := Duzbucaqli{en:45, uzunluq:55}
cv := Cevre{radius: 23}

var x Hendese 
fmt.Println(x)
x = kv
fmt.Println(x.fiqurun_adı(), "sahə:", x.sahe(), "perimetr:",x.perimetr() )
x = dz
fmt.Println(x.fiqurun_adı(), "sahə:", x.sahe(), "perimetr:",x.perimetr() )
x = cv
fmt.Println(x.fiqurun_adı(), "sahə:", x.sahe(), "perimetr:",x.perimetr() )
```

    <nil>
    Kvadrat sahə: 1936 perimetr: 176
    Düzbucaqlı sahə: 2475 perimetr: 2475
    Çevrə Dairə sahə: 1661.06 perimetr: 144.44


#### Type Assertions (Tipin təsdiqi), Type Switch 

Bu mövzuya bəsit də olsa flow controls-da toxunmuşduq.  
Bəzən bizə interfeys dəyişəninin saxladığı qiymətin hansı dinamik tip olduğunu müəyyənləşdirib  
ona uyğun hərəkət etmək lazım gəlir.   
Bunun üçün Type assertions və ya Type switch-dən istifadə edirik.   
 
Type Assertions üçün bu cür yazılışdan istifadə edirik.
```Go
//T tipi, x isə interfeysdir
x.(T)

```
Burada x-in saxladığı qiymətin T dinamik tipində olması yoxlanılır.  
Yoxlamanın nəticəsini 2 cür (qısa və xəta ilə birlikdə tam) mənimsədə bilərik.   
```Go
v := x.(T)

```
Yuxarıdakı halda x-in qiyməti T tipinə malik olmazsa xəta baş verəcək.    
Ona görə çox zaman aşağıdakı mənimsətmədən istifadə edirik.  
```Go
v, ok := x.(T) // ok bool tipidir

```
Burada v uğurlu halda qiyməti, ok isə halın uğurlu olub olmamağını saxlayacaq.   

Hər iki halda T tipi x-in dinamik tipi olmalıdır, yoxsa **kompilyasiya** xətası baş verəcək.

Dediklərimizi nümunə ilə göstərək.  


```go
%%
kv := Kvadrat{teref:44}
var x Hendese = kv

iii := x.(Kvadrat)
fmt.Printf("%T %v %v \n", iii, iii, iii.teref)
```

    main.Kvadrat {44} 44 



```go
%%
kv := Kvadrat{teref:44}
var x Hendese = kv

iii := x.(Cevre)
fmt.Println(iii)
```

    panic: interface conversion: main.Hendese is main.Kvadrat, not main.Cevre
    
    goroutine 1 [running]:
    main.main()
    	 [7m[[ Cell [18] Line 5 ]][0m /tmp/gonb_0e480a67/main.go:68 +0x6f
    exit status 2



```go
%%
kv := Kvadrat{teref:44}
var x Hendese = kv

v, ok := x.(Cevre)
if ok {
	fmt.Println(v)
} else {
	fmt.Println( "Cevre tipinde olani saxlamir")
}

```

    Cevre tipinde olani saxlamir



```go
type HerHansiTip int 
%%
var x Hendese
_, ok := x.(HerHansiTip)
if  ok {
   fmt.Println("ok")
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


<span style="white-space: pre;"> # gonb_0e480a67</span>

<br/>

<span class="gonb-cell-line-info">Cell[20]: Line 4</span>
<span class="gonb-error-location">./main.go:67:11: </span> impossible type assertion: x.(HerHansiTip)
<div class="gonb-error-context">
func main() {
	flag.Parse()
	var x Hendese
<div class="gonb-error-line">	_, ok := x.(HerHansiTip)
</div>	if  ok {
	   fmt.Println(&#34;ok&#34;)

</div>

<br/>


<span style="white-space: pre;"> 	HerHansiTip does not implement Hendese (missing method fiqurun_adı)</span>

<br/>


<span style="white-space: pre;"> </span>

<br/>

</div>




    failed to run "/usr/local/go/bin/go build -o /tmp/gonb_0e480a67/gonb_0e480a67": exit status 1



```go
func  goster(x Hendese) {
	
	if x == nil {
		return
	}
    if kv, ok:= x.(Kvadrat); ok {
		//kvadratdır
		fmt.Println("kvadratin terefi:", kv.teref )
	} else if cv,ok:=x.(Cevre); ok {
		//cevre
		fmt.Println("Cevrenin radiusu:", cv.radius )
	} else if v,ok:=x.(Duzbucaqli); ok {
		fmt.Println("Duzbucaqli en:", v.en, "uzunluq:", v.uzunluq )
	}
	fmt.Println(x.fiqurun_adı(), "sahə:", x.sahe(), "perimetr:",x.perimetr() )
}

%%

kv := Kvadrat{teref:44}
dz := Duzbucaqli{en:45, uzunluq:55}
cv := Cevre{radius: 23}
goster(kv)
goster(dz)
goster(cv)
```

    kvadratin terefi: 44
    Kvadrat sahə: 1936 perimetr: 176
    Duzbucaqli en: 45 uzunluq: 55
    Düzbucaqlı sahə: 2475 perimetr: 2475
    Cevrenin radiusu: 23
    Çevrə Dairə sahə: 1661.06 perimetr: 144.44


##### Type Switch (Tip seçim)

```Go
// deyishen ifadəsinin saxladığı qiymətin nil və ya hansı dinamik tip  
// olmasının müəyyən edilməsi
switch deyishen.(type) {
// cases
}

Burada:
switch də yoxlanılan deyishen.(type) kimi yazılır
deyishen interfeys tip olmalıdır
case-lər isə:
nil,
deyishen interfeysinin dinamik tipləri
və təkrarlanmayan olmalıdır. 
```
Type switch-ə aid nümunəni boş interfeys üçün göstərək.


```go
//qeyd any ele interface{} deməkdir
func gosterUmumi( x any){

switch i := x.(type) {
case nil:
	fmt.Println("x nil-dir", i)               //  i-in tipi x (interface{}) esl tipi olacaq
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
}

%%
x:="hello"
y:=33
z:=55.5
gosterUmumi(x)
gosterUmumi(y)
gosterUmumi(z)
gosterUmumi(nil)
```

    type  bool ve ya string -dir
    33
    float64 55.5
    x nil-dir <nil>


#### Embeded Interface  

Strukt-da olduğu kimi hər hansı bir T interfeysin daxilində digər bir neçə interfeysi embed etmək olur.   
T interfeysi həm öz metodlarının həm də embed etdiyi interfeys metodlarını bəyan edən tiplər çoxluğuna malik olur.   
Bu zaman fikir vermək lazımdır ki embed olan interfeyslərdə olan metodlarda oxşar funksiyalar eyni quruluşda olsunlar.  
Əks halda bu **kompilyasiya** xətasına səbəb olacaq. 


```go
type Reader interface {
	Read(p []byte) (n int, err error)
	Close() error
}

type Writer interface {
	Write(p []byte) (n int, err error)
	Close() error
}

// ReadWriter metodları Read, Write, Close.
type ReadWriter interface {
	Reader  // Reader-in metodları
	Writer  // Üriter-in metodları
}
```

Məsələn eyniadlı metodda əgər quruluş fərqi olsa


```go
type ReadCloser interface {
	Reader   // Reader-in metodları
	Close()  //  Reader.Close və Close eyniadlıdır, lakin quruluş fərqi var deyə yol verilməzdir
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


<span style="white-space: pre;"> # gonb_0e480a67</span>

<br/>

<span class="gonb-cell-line-info">Cell[21]: Line 2</span>
<span class="gonb-error-location">./main.go:23:2: </span> undefined: Reader
<div class="gonb-error-context">
    teref float32
}
type ReadCloser interface {
<div class="gonb-error-line">	Reader   // Reader-in metodları
</div>	Close()  //  Reader.Close və Close eyniadlıdır, lakin quruluş fərqi var deyə yol verilməzdir
}

</div>

<br/>


<span style="white-space: pre;"> </span>

<br/>

</div>




    failed to run "/usr/local/go/bin/go build -o /tmp/gonb_0e480a67/gonb_0e480a67": exit status 1


<span style="color:red">Qeyd edək ki,</span> interfeys tiplər üçün məhdudiyyət vermək üçün də istifadə edilir.  
Lakin bu sadə interfeysdən fərqlənir və Generics mövzusuna aid oluğundan onu o mövzuda öyrənəcəyik. 
