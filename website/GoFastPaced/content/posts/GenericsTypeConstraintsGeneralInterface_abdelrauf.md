---
title: " Generics Type Constraints General Interface"
author: "abdelrauf"
date: 2023-08-23
tags: ["Generics", "Type", "Constraints", "General", "Interface", "Go" ]
---

### Generics

Bəzən biz funksiya və yeni mürəkkəb tip quranda onun istifadə etdiyi parametrlərin və   
dəyişənlərin tipini istifadə zamanı müəyyən edilməsini istəyirik.  
Aydın məsələdir ki bu boş interfeysdən fərqli olaraq daha sürətli olacaq.  
Buna Go dilində bir çox dillərdə olduğu kimi generics type deyilən yanaşma ilə nail olunur.  
Go dilinin yeni versiyalarında (>=1.18) artıq bunu istifadə etmək mümkündür. 

Tip elan edilməsinə yenidən baxaq:  
<pre>

 
type ad [TipParametri] tip 

Burada TipParametri:
[ ] arasında ad və Məhdudiyyət cütlüklərindən ibarət olan siyahı olur
Məsələn [T Constraint], [T Məhdudiyyət1, V Məhdudiyyət2, ...]

(Tip Məhdudiyyətlərini aşağıda xüsusi bölmədə izah edəcəyik.  )

Yeni tipimizə metod vermək üçün  

func (x * ad[T]) metod_adı(parametrlər... ) nəticə... { ...}

</pre>

Həmçinin generic funksiya elan edilməsinə baxaq:
<pre>

func G[T Constraint](p T) { ... }

</pre>

Hələki tip məhdudiyyəti kimi **any** götürək hansı ki bütün tipləri götürür.  
Onda Generics-ə nümunə olaraq bunları göstərə bilərik  
```Go
//generic List
type List[T any] struct {
	next  *List[T]
	value T
}

// List[T] biz üçün metod vermə qaydasına nümunə
func (l *List[T]) Len() int  { ... }

//generic Print funksiyamız
func Print[T any](s []T) {
    for _, v := range s {
        fmt.Println(v)
    }
}

```


```go

func Print[T any](s []T) {
    for _, v := range s {
        fmt.Println(v)
    }
}

%%
Print[int]([]int{1, 2, 3})
Print[string]([]string{"alma","banan", "kivi"})
```

    1
    2
    3
    alma
    banan
    kivi


Yuxarıdakı çağırış tərzini sadələşdirmək də olar.  
Go bizim yerimizə parametrdən tipi təyin edəcək.  



```go
%%
Print([]string{"alma","banan", "kivi", "armud"})
```

    alma
    banan
    kivi
    armud


#### Type Contraints  (Tip məhdudiyyətləri)  

Tip məhdudiyyəti müvafiq tip parametri üçün icazə verilən tip arqumentlər dəstini təyin edir.  
Bu dəstəklənən əməliyyatlara nəzarət edən interfeysdir.  

Tərifdən də aydın olur ki biz niyə any istifadə etmişik. Çünki any boş interfeysdir.  

Aşağıdakı tip məhdudiyyətlər elanına fikir verək  
<pre>
[T any]  //[T interface{}]
[T interface{[]P}]
[T interface{~int}]
[T interface{int|string}]
</pre>

Fikir verdinizsə digər interfeys yazıları bizə tanış deyil.  
Biz interfeys daxilində ya boşluq ya da funksiya (func-suz) elanı istifadə edirdik.  
Bu tip interfeys sadə interfeysdən fərqlidir və buna ümumi interfeys deyirik. (General İnterface)

##### General Interface (Ümumi interfeys)  
Ümumi interfeysin elanı.  
Asan olsun deyə qaydalar şəklində yox göstərməklə təyin nümunlərinə baxaq:  
```Go
// Interfeys yalnız "tip" tipini əhatə edir.
interface {
	tip
}
// Məsələn:
interface {
	int
}

// Interfeys alt tipi "tip" olan bütün tipləri əhatə edir.
// yəni alt tipi "tip" olan bütün tiplər interfeysin tip çoxluğunda olur.
interface {
	~tip
}

// Məsələn:
interface {
	~int
}
// MyInt alt tipi int olduğundan interfeys tərəfindən qəbul olunacaq
type MyInt int 


// Interfeys alt tipi "tip" olan və həmçinin "metod_adı" metodu elan edilmiş tipləri əhatə edir.  
// yəni interfeysin tip çoxluğunda olur.
interface {
	~tip
	metod_adı(parameterlər) nəticələr
}

// Məsələn:
interface {
	~int
	String() string
}
type MyInt2 int

func (v MyInt2) String() string {
    ...
}

// Aşağıdakı misalda isə heç bir tip əhatə olunmayacaq. Çünki bir tipin həm int, həm string tipi olması mümkün deyil.
interface {
	int
	string
}

//Qeyd edək ki | məntiqi və ya bağlayıcısı ilə də məhdudiyyət interfeysi hazırlaya bilərik

// Interfeys "tip1" və ya "tip2" və ya "tip3" tipini əhatə edir.
interface {
	tip1 | tip2 | tip3
}

// Məsələn:

//alt tipi int və ya string olan tiplər
interface {
	~int | string
}


// https://pkg.go.dev/golang.org/x/exp/constraints 
type Signed interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Unsigned interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr
}

type Integer interface {
	Signed | Unsigned
}

```


Daha dəqiq və ətraflı qaydalar və istifadə toplumu üçün dilə baxa bilərsiniz.  
https://go.dev/ref/spec#Interface_types

Qeyd edək ki, bəzi tip məhdudiyyətlərini sadə üsulla da elan edə bilərik.   
Məsələn:  
<pre>
[T []P]                      // = [T interface{[]P}]
[T int]                     // = [T interface{int}]
[T ~int]                     // = [T interface{~int}]
[T int|string]               // = [T interface{int|string}]
</pre>

(Qeyd: bu halda [T int] -də int-i tip yox, tip məhdudiyyəti kimi başa düşmək daha düzgündür) 

Amma  nəzərə almaq lazımdır ki bu sadələşdirmə 
qarışıqlığa səbəb olsa onda onu xüsusi üsulla qeyd etmək lazımdır.  
Məsələn aşağıdakı C məhdudiyyəti üçün qarşılıq yarana bilər.  

<pre>
type X1[P *C] ..
type X2[P (C)] ...
type X3[P *C|Q] ...
</pre>

Bu qarşılıqlığı aradan qaldırmaq üçün ya təzədən məhdudiyyətləri interface{} almaq, ya da "," işarəsini sonuna əlavə etmək lazımdır  
<pre>
type X1[P interface{*C}] ...
type X1[P *C,] ...
</pre>

Dediyimiz qarşılıqlığa aid real nümunə verək və çıxan səhvə baxaq:


```go
type Lx [ T *int] struct {
	val T
}
```



<style>
.gonb-err-location {
	background: var(--jp-err-color2);  
	border-radius: 3px;
	border-style: dotted;
	border-width: 1px;
	border-color: var(--jp-border-color2);
}
.gonb-err-location:hover {
	border-width: 2px;
	border-style: solid;
	border-color: var(--jp-border-color2);
}
.gonb-err-context {
	display: none;
}
.gonb-err-location:hover + .gonb-err-context {
	background: var(--jp-dialog-background);  
	border-radius: 3px;
	border-style: solid;
	border-width: 1px;
	border-color: var(--jp-border-color2);
	display: block;
	white-space: pre;
	font-family: monospace;
}
.gonb-err-line {
	border-radius: 3px;
	border-style: dotted;
	border-width: 1px;	
	border-color: var(--jp-border-color2);
	background-color: var(--jp-rendermime-err-background);
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


<span style="white-space: pre;"> # gonb_93df4279</span>

<br/>

<span class="gonb-cell-line-info">Cell[24]: Line 1</span>
<span class="gonb-err-location">./main.go:11:11: </span> undefined: T
<div class="gonb-err-context">
type List[T any] struct {
    head, tail *element[T]
}
<div class="gonb-err-line">type Lx [ T *int] struct {
</div>	val T
}

</div>

<br/>

<span class="gonb-cell-line-info">Cell[24]: Line 1</span>
<span class="gonb-err-location">./main.go:11:14: </span> int (type) is not an expression
<div class="gonb-err-context">
type List[T any] struct {
    head, tail *element[T]
}
<div class="gonb-err-line">type Lx [ T *int] struct {
</div>	val T
}

</div>

<br/>

<span class="gonb-cell-line-info">Cell[24]: Line 2</span>
<span class="gonb-err-location">./main.go:12:6: </span> undefined: T
<div class="gonb-err-context">
    head, tail *element[T]
}
type Lx [ T *int] struct {
<div class="gonb-err-line">	val T
</div>}
type element[T any] struct {

</div>

<br/>


<span style="white-space: pre;"> </span>

<br/>

</div>




    failed to run "/usr/local/go/bin/go build -o /tmp/gonb_93df4279/gonb_93df4279": exit status 1


Gördüyünüz kimi Go bunu düzgün parse edə bilmədi. Bu adətən o zaman baş verirki  
yazılış ifadəyə oxşayır. İndi isə həmin problemi aradan qaldıran yazılışa baxaq:


```go
// Normal interface{ } içində elanla
type Lx [ T interface{*int}] struct {
	val T
}
```


```go
// Əlavə vergül yazmaqla
type Lx [ T *int,] struct {
	val T
}
```

##### comparable
Go dilində any interfeysi ilə yanaşı comparable interfeys məhdudiyyəti var.  
comparable məhdudiyyəti müqayisə əməliyyatlarından(==, !=) yerinə yetirən tipləri əhatə edir.  
(Go 1.20-dən any də comparable tipin interfeys çoxlundadır)  


```go
func beraber[T comparable](x, y T) bool {
    if x == y {
        return true
    }
    return false
} 

%%
var a,b int
a=5
b=8
fmt.Println(beraber[int](a,b))
fmt.Println(beraber(b, b))

```

    false
    true


Məsələn bilirik ki funksiya tipləri müqayisə olunmur. 


```go
 
%%
a:= func() int {return 5}
b:= func() int {return 8}
fmt.Println(beraber(a,b))

```



<style>
.gonb-err-location {
	background: var(--jp-err-color2);  
	border-radius: 3px;
	border-style: dotted;
	border-width: 1px;
	border-color: var(--jp-border-color2);
}
.gonb-err-location:hover {
	border-width: 2px;
	border-style: solid;
	border-color: var(--jp-border-color2);
}
.gonb-err-context {
	display: none;
}
.gonb-err-location:hover + .gonb-err-context {
	background: var(--jp-dialog-background);  
	border-radius: 3px;
	border-style: solid;
	border-width: 1px;
	border-color: var(--jp-border-color2);
	display: block;
	white-space: pre;
	font-family: monospace;
}
.gonb-err-line {
	border-radius: 3px;
	border-style: dotted;
	border-width: 1px;	
	border-color: var(--jp-border-color2);
	background-color: var(--jp-rendermime-err-background);
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


<span style="white-space: pre;"> # gonb_93df4279</span>

<br/>

<span class="gonb-cell-line-info">Cell[32]: Line 5</span>
<span class="gonb-err-location">./main.go:56:21: </span> func() int does not satisfy comparable
<div class="gonb-err-context">
	flag.Parse()
	a:= func() int {return 5}
	b:= func() int {return 8}
<div class="gonb-err-line">	fmt.Println(beraber(a,b))
</div>


</div>

<br/>


<span style="white-space: pre;"> </span>

<br/>

</div>




    failed to run "/usr/local/go/bin/go build -o /tmp/gonb_93df4279/gonb_93df4279": exit status 1


Təəssüf ki Go dilində öz tipimizi comparable çoxluğuna əlavə etmək üçün elan  
etdiyimiz tipə xüsusi == != metodlarını təyin edə bilmirik.  Bu xüsusi qaydalara əsasən təyin edilir.  
Məsələn struct tipimizin bütün eleməntləri müqayisə edilən olduqda bizim struct da comparable olur.  
Tam qaydaları   
https://go.dev/ref/spec#Comparison_operators bölməsində tapmaq olar


Generics-ə aid List nümunəsinə baxaq.  
Bu nümunə https://gobyexample.com/generics götürülüb


```go
type List[T any] struct {
    head, tail *element[T]
}
type element[T any] struct {
    next *element[T]
    val  T
} 
func (lst *List[T]) Push(v T) {
    if lst.tail == nil {
        lst.head = &element[T]{val: v}
        lst.tail = lst.head
    } else {
        lst.tail.next = &element[T]{val: v}
        lst.tail = lst.tail.next
    }
}
func (lst *List[T]) GetAll() []T {
    var elems []T
    for e := lst.head; e != nil; e = e.next {
        elems = append(elems, e.val)
    }
    return elems
}

%%
lst := List[int]{}
lst.Push(10)
lst.Push(13)
lst.Push(23)
fmt.Println("list:", lst.GetAll())
```

    list: [10 13 23]

