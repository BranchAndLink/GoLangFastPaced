---
title: " Functions Recursive Anonymous Closure"
author: "abdelrauf"
date: 2023-08-15
tags: ["Functions", "Recursive", "Anonymous", "Closure", "Go" ]
---

#### First Class Function (1-ci sinif funksiyalar)   

Go dilində funksiyalar 1-ci sinifdir. Bu o deməkdir ki, funksiyalar dəyişənə mənimsədilə, digər funksiyaya ötürülə, nəticə kimi qaytarıla bilinir.  
Dediklərimizi misallarla göstərək   


Funksiya dəyişən kimi


```go
func Go(h int) int {
    return h + 5
}

%%
funksiya := Go
fmt.Printf("%T\n", funksiya)
fmt.Println(funksiya(55))
```

    func(int) int
    60


Funksiya arqument kimi


```go
func DoCalc( operator func(int, int) int, initial int,  args ...int) (netice int) {

    netice = initial

    for _, x:= range args {
        netice = operator(netice, x)
    }
    return
}

func Toplama (x, y int) int {
    return x+y
}

func Vurma (x, y int) int {
    return x*y
}

%%

ededler := [...] int {4,5,6,7,8,9,7,5,3,2}

fmt.Println( "Toplama", DoCalc(Toplama, 0, ededler[:]...))
fmt.Println( "Vurma",  DoCalc(Vurma, 1, ededler[:]...))

```

    Toplama 56
    Vurma 12700800


Funksiya nəticə kimi


```go
func Funksiya ( op string )  func(int, int) int  {

    switch op {
        case "+":
        return Toplama
        case "*":
        return Vurma
        default:
        return Toplama
    }

}

%%
ededler := [...] int {4,5,6,7,8,9,7,5,3,2}

fmt.Println( "Toplama", DoCalc(Funksiya("+"), 0, ededler[:]...))
fmt.Println( "Vurma",  DoCalc(Funksiya("*"), 1, ededler[:]...))
```

    Toplama 56
    Vurma 12700800


##### Anonymous Function (Anonim funksiya) 

Bəzən müvəqqəti adsız funksiya düzəltmək istəyirik. Bu zaman anonim funksiyadan (və ya lambda-dan) istifadə edə bilərik.   
Bu funksiyanı istənilən yerdə elan edə bilməmimiz işimizi asanlaşdırır



```go
%%
funksiya := func (h int) int {
	return h + 5
}
fmt.Printf("%T\n", funksiya)
fmt.Println(funksiya(55))
```

    func(int) int
    60



```go


%%
sechim := func ( op string )  func(int, int) int  {

    switch op {
        case "+":
        return func(x,y int) int {
			return x + y
		}
        case "*":
        return  func(x,y int) int {
			return x * y
		}
        default:
        return  func(x,y int) int {
			return x + y
		}
    }

}
ededler := [...] int {4,5,6,7,8,9,7,5,3,2}

fmt.Println( "Toplama", DoCalc(sechim("+"), 0, ededler[:]...))
fmt.Println( "Vurma",  DoCalc(sechim("*"), 1, ededler[:]...))
```

    Toplama 56
    Vurma 12700800



```go
%%
func(name string) {
	fmt.Println("Hello",name)
}("GoLang")
```

    Hello GoLang


##### Closure  (klojr)
Closure dedikdə funksiya xaricində olan dəyişənləri funksiya daxilində (arqument kimi verilmədən) işlətmək imkanı nəzərdə tutulur.  
Bu özəllik bizə bir çox məsələləri daha sadə həll etməyə imkan  verəcək.  



```go
%%
name := "Golang"
func() {
	fmt.Println("Hello",name)
}()
```

    Hello Golang


Funksiya əgər digər funksiyaları arqument kimi götürüb və ya nəticə kimi qaytarırsa   
ona **High Order functions** deyirik.

Yuxarıda öyrəndiklərimizi tətbiq etməklə bəzi maraqlı funksiyalar yığaq


```go
func MurekkebFunksiya( f func(x int) float64, g  func (x float64) string )   func(int) string {
	return func(x int) string {
		return g(f(x))
	}
}

%%
f:= func(x int) float64 {
	return float64(x)/2.0
}
g:= func(x float64) string {
	metn := fmt.Sprintf("~~%g~~", x) 
	return metn
}
mmm := MurekkebFunksiya(f,g)
fmt.Println(mmm(456))
```

    ~~228~~



```go
type Filter  func(int) bool

func SliceFilterInPlace(a []int, filterFunk Filter) (len int) {
    for _, x := range a {
       if filterFunk(x) {
			a[len] = x
			len++
	   }
	}
	
	return
}

func SliceFilter(a []int, filterFunk Filter) [] int {
	slice := make([]int, 0, len(a)/2)
    for _, x := range a {
       if filterFunk(x) {
			slice = append(slice, x)
	   }
	}
	
	return slice
}
%%

vals := []int{-2, -8, 8, 5, 99, 0, 1, 9, 7, 55, 66, 7, -3, -5, 6, 111}
lenx := SliceFilterInPlace(vals, func (x int) bool {
	return x<77
})

filtered := SliceFilter(vals[:lenx], func (x int ) bool {
	return x>0
})
filtered2 := SliceFilter(filtered, func (x int ) bool {
	return x>10
})
fmt.Println(vals[:lenx], filtered, filtered2)
```

    [-2 -8 8 5 0 1 9 7 55 66 7 -3 -5 6] [8 5 1 9 7 55 66 7 6] [55 66]


#### Rekursion (Rekursiya) 
Funksiya əgər daxildə öz funksiyasını çağırırsa buna rekursiya deyilir. Həmçinin Əgər funksiya çağırdığı funksiyalardan biri də təzədən onu çağırırsa  
rekursiya baş verir. Rekursiv funksiyalarla biz bəzi alqoritmik məsələləri həll edə bilirik.   
Lakin nəzərə almaq lazımdır ki rekursiya (call stack) stek yaddaşın dolmasına səbəb olur.   
(<span style="color:red">Qeyd GO-da hələki rekursiv funksiyalar (tail recursion optimization) <b>optimallaş-ma-dığından</b> alqoritmik həlləri for iterasiyası ilə yığmaq məsləhətdir</span>)


```go
func faktorial(n int) int {
    if n == 0 {
        return 1
    }
    return n * faktorial(n-1)
}

%%
print(faktorial(6))
```

    720


```go
func cem( a [] int) int {
	if len(a) <2 {
		return a[0]
	}
	return a[0] + cem(a[1:])
}

%%
print(cem([]int{1,2,3,4,5,6,7,8,9,100,1000,10000}))
```

    11145

Aşağıdakı misalda  anonim closure funksiyasında rekursiyanı vermək üçün fibonacci dəyişənini var ilə **əvvəlcədən** bəyan etməli oluruq.  
Əks halda closure-mız funksiyanı təyin edə bilmir.  


```go
%%
var fibonacci func(n int) int
fibonacci = func(n int) int {
	if n < 2 {
		return n
	} 
	return fibonacci(n-1) + fibonacci(n-2)
}
print(fibonacci(15))
```

    610

Fibonnaci misalında təkrar hesablamalar üçün memoization(yaddaş) texnikasını əlavə edə bilərik


```go
%%
var fibonacci func(n int) int
var yaddash = make(map[int]int)

fibonacci = func(n int) int  {
 if n < 2 {
  return n
 }
 //eger hesablama varsa yaddash-dan gotur
 if v, ok := yaddash[n]; ok {
  return v
 }
 result := fibonacci(n-1) + fibonacci(n-2)
 yaddash[n] = result
 return result
}
print(fibonacci(15))

```

    610
