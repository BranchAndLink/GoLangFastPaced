---
title: " String Number "
author: "abdelrauf"
date: 2023-07-31
tags: ["String", "Number", "Go" ]
---

#### Go-da string və ədəd tipləri bir birinə çevirmək

Bunun üçün biz **strconv** və ya **fmt** paketini istifadə edəcəyik



```go
import "fmt"
import "strconv"
```

``` Go
func Atoi(s string) (int, error)   
``` 
ədəd mətnini tam ədədə çevirir  və həmin tipdəki ədəd və xəta qaytarır. Xəta baş vermədikdə xəta nil olur   

``` Go
func ParseInt(s string, base int, bitSize int) (i int64, err error) 
````
Bu funksiya vasitəsilə isə istənilən əsaslı ədədi tam ədəd çevirmək olur 



```go
%%
a_metn:= "4566"
b_metn:= "45AB"
a, xeta1:=strconv.Atoi(a_metn)
if xeta1 != nil {
	fmt.Println(xeta1)
}
b, xeta2:=strconv.ParseInt(b_metn, 16, 32)
if xeta2 != nil {
	fmt.Println(xeta2)
}
fmt.Printf("%T %T %v\n", a_metn, a, a )
fmt.Printf("%T %T %x %v\n", b_metn, b, b, b)


```

    string int 4566
    string int64 45ab 17835


```Go
func ParseFloat(s string, bitSize int) (float64, error)
```
Mətn ədədini float tipə çevirmək  
bitSize 32 olduqda nəticə yenə float64 olsa da onu float32-ə dəyişilmədən çevirmək olur  


```go
%%
b_metn:= "456.455" 
b, xeta:=strconv.ParseFloat(b_metn, 64)
if xeta != nil {
	fmt.Println(xeta)
} 
fmt.Printf("%T %T %v\n", b_metn, b, b )


```

    string float64 456.455


Digər konversiyalar üçün strconv paketini incələyin

**fmt** paketi və formatlarla mətn ədədini ədəd tipə çevirək  

Bunun üçün Sscanf dən istifadə edəcəyik. Bunun vasitəsilə dəqiq bildiyimiz mətn içərisindəki ədədi də rahat çevirmək olar
```Go
func Sscanf(str string, format string, a ...any) (n int, err error)
```


```go
%%
b_metn:= "456.455" 
c_metn:= "armud:56kq"
var b float64
var c int
_, xeta:=fmt.Sscanf(b_metn, "%g", &b)
if xeta != nil {
	fmt.Println(xeta)
} 
fmt.Printf("%T %T %v\n", b_metn, b, b)
_, xeta1:=fmt.Sscanf(c_metn, "armud:%dkq", &c)
if xeta1 != nil {
	fmt.Println(xeta1)
} 
fmt.Printf("%T %T %v\n", c_metn, c , c)
```

    string float64 456.455
    string int 56


Ədədi mətnə çevirmək

```Go
func Sprintf(format string, a ...any) string
```


```go
%%
var g=456.45
var h string
//BasicOutput-dan bildiyimiz kimi v tipe uygunlugu ozu teyin edir
h = fmt.Sprintf("%v", g)
fmt.Printf("%T %T %v \n",g, h, h)
```

    float64 string 456.45 



```go
%%
var g=456.45
var h string
//%g float tipi
h = fmt.Sprintf("%g", g)
fmt.Printf("%T %T %v \n",g, h, h)
```

    float64 string 456.45 


```Go
//tam eded uchun
func Itoa(i int) string
//istenilen esasli
func FormatInt(i int64, base int) string
//float
func FormatFloat(f float64, fmt byte, prec, bitSize int) string
//complex
func FormatComplex(c complex128, fmt byte, prec, bitSize int) string

Burada:
base - tam ədəddə əsas
fmt - format-dan bildiklərimiz %  işarəsiz. 'e', 'E', 'f', 'g', 'G', 'x', and 'X' 
prec - (-1) verməklə float-u mətnə çevirəndə ən az rəqəm sayısına nail oluruq
bitSize - bit ölçüsü
```


```go
%%
var g=456.45
var h string 
h = strconv.FormatFloat(g, 'g', -1 ,64)
fmt.Printf("%T %T %v \n",g, h, h)
```

    float64 string 456.45 



```go
%%
var g=45
var h string 
h = strconv.Itoa(g)
fmt.Printf("%T %T %v \n",g, h, h)
```

    int string 45 

