---
title: ArrayAndSlice
author:abdelrauf
date: 2023-07-31
tags: ["Array", "Slice", "Go" ]
---

#### Array. 
eyni tipli elementlər ardıcıllığını vermək üçün istifadə edilir. Yaddaşda ardıcıl yer tutur

<pre>
[say]tip{elementlər}   
elan:
var ad [say]tip   
var ad [say]tip = [say]tip{elementlər}   
var ad = [say]tip{elementlər}  
ad := [say]tip{elementlər}  
say Yerinə ... istifadə etsək element sayını özü yazır  

[...]tip{elementlər}  

{ } içindəki Elementləri tam yox spesifikdə vermək olar {index1: qiymet1. index2: qiymet2}. Bu zaman digərləri default olur 

Çox ölçülü :
2d misalında 
[say1][say2] tip { {elementlər}, {elementlər}}
</pre>


```go
%%
var g1 [5]int
var g2 [5]int = [5]int {4,5,6,8,9}
var g3 [5]int = [5]int {2:6, 4:9}
var g4 = [5]int {2:6, 4:9}
var g5 = [...]int {2,5,6,7,8}
var g6 = [...]int {2:6, 4:9}

a1 := [5]int {2:6, 4:9}
a2 := [...]int {2,5,6,7,8}
a3 := [...]int {2:6, 4:9}
multi := [...][3]int{{1,2,3},{5,6,7},{4,5,9},{4,5,9}}
fmt.Println(g1,g2,g3,g4,g5,g6,a1,a2,a3)
fmt.Println(multi)
fmt.Printf("%T %T\n", a2, multi)

```

    [0 0 0 0 0] [4 5 6 8 9] [0 0 6 0 9] [0 0 6 0 9] [2 5 6 7 8] [0 0 6 0 9] [0 0 6 0 9] [2 5 6 7 8] [0 0 6 0 9]
    [[1 2 3] [5 6 7] [4 5 9] [4 5 9]]
    [5]int [4][3]int


Array-in elementini dəyişmək üçün  [indeks] və mənimsətmə ilə edə bilərik


```go
%%
var g1 [5]int
g1[0] = 77
g1[3] = 56
multi := [...][3]int{{1,2,3},{5,6,7},{4,5,9},{4,5,9}}
multi[2][2] = 9898
fmt.Println(g1)
fmt.Println(multi)

```

    [77 0 0 56 0]
    [[1 2 3] [5 6 7] [4 5 9898] [4 5 9]]


**len** builtin Funksiya ilə ölçü və sayı almaq mümkündür. Array üçün də keçərlidir


```go
%%
var g5 = [...]int {2,5,6,7,8}
fmt.Println(len(g5))
```

    5


Qeyd edək ki string tipi də əslində array tipinə oxşardır lakin eyni deyil. Sadəcə burda nəzərə almaq lazımdır ki ASCİİ simvollar olanda indekslər üst üstə düşür. 
 Əks halda rune massivinə çevirmək və ya range istifadə lazımdır


```go
%%
var s= "Men de arrayam"
fmt.Println(len(s), s[0],s[1], s[2])
fmt.Printf("Say %d, %c, %c, %c \n",len(s), s[0],s[1], s[2])
var ss = "😁👍😊al"
fmt.Printf("Say %d, Duzgun olmayan herf indeksleri %c, %c, %c \n",len(ss), ss[0],ss[1], ss[2])
for pos, x := range ss{
	fmt.Printf("%d - %c \n", pos, x)
}
runes := []rune(ss)
fmt.Printf("Say %d, %c, %c, %c \n",len(runes), runes[0],runes[1], runes[2], runes[3])
```

    14 77 101 110
    Say 14, M, e, n 
    Say 14, Duzgun olmayan herf indeksleri ð, ,  
    0 - 😁 
    4 - 👍 
    8 - 😊 
    12 - a 
    13 - l 
    Say 5, 😁, 👍, 😊 
    %!(EXTRA int32=97)

Array-ləri bir birinə mənimsətdikdə kopyalama (sürət alma) baş verir.   
Ona görə funksiyaya verdikdə və s hallarda buna diqqət yetir.   
Qeyd ya slice istifadə elə (hansı ki reference-dir(pointer)) ya pointer


```go
%%
s:=[...]int{5,6,7,8,9}
d:=s
d[0] = 99
fmt.Println(s)
fmt.Println(d)
fmt.Println(s[0], "!=", d[0])

```

    [5 6 7 8 9]
    [99 6 7 8 9]
    5 != 99


**Elementlər sayı eyni olan Arraylər** üzərində müqayisə əməliyyatı olaraq == və != aparmaq olur. string-də isə hamısı olur.  


```go
%%
s:=[...]int{5,6,7,8,9}
d:=s 
d[0] = 99
fmt.Println(s == d )
fmt.Println(s != d )
```

    false
    true


### Slice
Slice arraydən fərqli olaraq **reference tipdir**. Yəni o böyüyə və kiçilə bilən bufferə (array) ə işarə saxlayır və nil-də ola bilər.   
Elanı array-ə oxşayır. Array-dən slice almaq da mümkündür. Sonda fərqliliklərini qeyd edəcəyik      

 <pre>
    var ad []tip   
    var ad []tip = []tip{elementlər}   
    var ad = []tip{elementlər}  
    ad := []tip{elementlər}  
    Çox ölçülü :
    2d misalında 
    [][] tip { {elementlər}, {elementlər}}

 </pre>


```go
%%
var g1 []int
var g2 []int = []int {4,5,6,8,9}
var g3 []int = []int {2:6, 4:9}
var g4 = []int {2:6, 4:9} 

a1 := []int {2:6, 4:9}
a2 := [...]int {2,5,6,7,8}
a3 := [...]int {2:6, 4:9}
multi := [...][3]int{{1,2,3},{5,6,7},{4,5,9},{4,5,9}}
multi[2][2] = 99999
fmt.Println(g1,g2,g3,g4,a1,a2,a3)
fmt.Println(multi)
fmt.Printf("%T %T\n", a2, multi)
//reference tip olmagi
ref_g:=g4
ref_g[3]= 9999
fmt.Println(ref_g[3], "==", g4[3])
```

    [] [4 5 6 8 9] [0 0 6 0 9] [0 0 6 0 9] [0 0 6 0 9] [2 5 6 7 8] [0 0 6 0 9]
    [[1 2 3] [5 6 7] [4 5 99999] [4 5 9]]
    [5]int [4][3]int
    9999 == 9999


Slice-a yeni element əlavə edək və ya pozaq. 
append builtin. 
Qeyd edək ki **yeni bufer yarana bildiyindən** onu yenidən *öz slice-a mənimsətmək lazımdır*.  
 Əks halda yeni bizim slice köhne bufferə, təzəyə isə mənimsətdiyimiz baxır
```Go
   slice1 = append( slice1, elementler,...)
   yeni_slice2 = append( slice1, elementler,...)
   //slice add edəndə slice_adi... istifadə edilir
   yeni_slice3 = append(slice1, yeni_slice2...)
```


```go
%% 
var g3 []int = []int {2:6, 4:9} 

fmt.Println("Say", len(g3), "elementler: ", g3)
g3 = append(g3, 5)
fmt.Println("Say", len(g3), "elementler: ", g3)
g3 = append(g3, 5,88,99)
fmt.Println("Say", len(g3), "elementler: ", g3)
//ozune menimsetmesek 
yeni_slice:=append(g3, 95,88,99)
fmt.Println("kohne slice Say", len(g3), "elementler: ", g3)
fmt.Println("Yeni slice Say", len(yeni_slice), "elementler: ", yeni_slice)
yeni_slice2 := append(yeni_slice, g3...)
fmt.Println("Birleshmis Yeni slice Say", len(yeni_slice2), "elementler: ", yeni_slice2)
```

    Say 5 elementler:  [0 0 6 0 9]
    Say 6 elementler:  [0 0 6 0 9 5]
    Say 9 elementler:  [0 0 6 0 9 5 5 88 99]
    kohne slice Say 9 elementler:  [0 0 6 0 9 5 5 88 99]
    Yeni slice Say 12 elementler:  [0 0 6 0 9 5 5 88 99 95 88 99]
    Birleshmis Yeni slice Say 21 elementler:  [0 0 6 0 9 5 5 88 99 95 88 99 0 0 6 0 9 5 5 88 99]


Array-in slice-in almaq üçün arrayin_adi[başlangıc_index:son_index] istifadə edirik.
Əgər [:] olduqda bütöv arrayə işarə edir.   
Ümumən slice-i dəyişdikdə array-də dəyişilir. Lakin əgər slice-da artırma və s baş veribsə bu yeni bufferə də işarə edə bilər.   
Slice saydan əlavə mümkün tutum həcmə də malik olur. Buna capacity deyilir və **cap** builtin funksiyası vasitəsilə əldə edə bilərik.  
Qeyd: Eynilə slice-dan da slice almaq olur. 


```go
%%
array:=[5]int{1,2,3,4,5}
slice:=array[:]
slice1:=array[0:2]
slice2:=array[2:4]
slice3:=array[1:4]

fmt.Println("say", len(slice), "tutum", cap(slice))
fmt.Println("say", len(slice1), "tutum", cap(slice1))
fmt.Println("say", len(slice2), "tutum", cap(slice2))
fmt.Println("say", len(slice3), "tutum", cap(slice3))

fmt.Println(array)
slice2[0] = 999
//slice2 -de elav artim etmediyimizden o ele array bufferine isahre edir deye array de deyishilecek
fmt.Println(array)


```

    say 5 tutum 5
    say 2 tutum 5
    say 2 tutum 3
    say 3 tutum 4
    [1 2 3 4 5]
    [1 2 999 4 5]


```Go
copy(dest_slice1, slice2)  //slice2 -i dest_slice1-e kopyalayir
```


```go
%%
slice1:=[]int{1,2,5,6}
slice2:=[]int{6,7,8,9,10}
slice3:=[]int{6,7,8,9,10,11,12,13}
copy(slice2,slice1)
fmt.Println(slice2)
//sonuna kopy edek
copy(slice3[len(slice3)-4:],slice1)
fmt.Println(slice3)
```

    [1 2 5 6 10]
    [6 7 8 9 1 2 5 6]


Slice- make  ilə yaratmaq  olur. make vasitəsilə bildiyiniz kimi tutumu əvvəlcədən vermək olur. və make initialize etdiyindən nil olmur. 
```Go
make(T[], say)
make(T[], say, tutum)

```


```go
%%
g:= make([]int, 5, 200)
fmt.Println(g, "say:", len(g), "tutum:", cap(g))
```

    [0 0 0 0 0] say: 5 tutum: 200


Slice sadəcə elan olunduqda nil olur. append nil Slice üzərinde də işləyir


```go
%%
var s0 []int
s1 := []int{}
s2 := make([]int,0)
s3 := make([]int, 50, 200)
fmt.Println("nil? ", s0 == nil)
fmt.Println("nil? ", s1 == nil)
fmt.Println("nil? ", s2 == nil)
fmt.Println("nil? ", s3 == nil)
s0 = append(s0,1,2,3)
fmt.Println(s0)
```

    nil?  true
    nil?  false
    nil?  false
    nil?  false
    [1 2 3]


**Qeyd:** make istifadəsi tutum və sayı əvvəlcədən verməyə imkan verir. hansı ki   
tutumun böyük olması append zamanı hər dəfə yeni buffer  yaranma şansını azaldır və performansa təsir edir.  


append təqribən bu cür koda malikdir   
```Go
func Append(slice, data []byte) []byte {
    l := len(slice)
    if l + len(data) > cap(slice) {  // reallocate
        // Allocate double what's needed, for future growth.
        newSlice := make([]byte, (l+len(data))*2)
        // The copy function is predeclared and works for any slice type.
        copy(newSlice, slice)
        slice = newSlice
    }
    slice = slice[0:l+len(data)]
    copy(slice[l:], data)
    return slice
}
```

Bu koddan həm də görsənir ki nəyə görə dəyişimdən sonra eyni bufferə işarə etməyə bilər.

### Element pozmağa və slice-la işləmə zamanı ehtiyatlı olmağa nümunə.    
##### Performance vs memory -aid nümunə. Həmçinin memory leak ehtimalı

Əgər elementi pozmaq istəsək slice və append -dən, copy-dən bu cür sitifadə edə bilərik:
Lakin bilmək lazımdır ki aşağıdakı üsulda buffer köhnə böyük bufffer olduğundan yaddaşda azalma olmur.
Bu isə bəzən arzuolunmaz haldır və <span style="color:red">memory leak</span> kimi sayıla bilər.   
Ancaq pozma zamanı yeni buffer yaranmadığından sürətli olur.


```go
func SlicedanElementiPoz1(slice []int, index int) []int {
	if index <0 || index > len(slice)-1 {
		return slice
	}
	slice_netice:=append(slice[:index], slice[index+1:]...)
	return slice_netice
}

func SlicedanElementiPoz2(slice []int, index int) []int {
	if index <0 || index > len(slice)-1 {
		return slice
	}
	copy(slice[index:], slice[index+1:])
    return slice[:len(slice)-1]
}

%%
slice2:=[]int{6,7,8,9,10}
slice3:=[]int{6,7,8,9,10,11,12,13} 
fmt.Println(SlicedanElementiPoz1(slice2, 3)) 
fmt.Println(SlicedanElementiPoz2(slice3, 6))
```

    [6 7 8 10]
    [6 7 8 9 10 11 13]


Arzuolunmaz halı göstərək


```go
%%
slice :=make([]int, 500)
fmt.Println("1)\nElement tutumu %v", cap(slice))
for i:=0;i<488;i++{
	slice = SlicedanElementiPoz1(slice, 1)
}
fmt.Printf("Element tutumu Pozduqdan sonra %v, say : %v\n", cap(slice), len(slice))

slice2 :=make([]int, 500)
fmt.Println("2)\nElement tutumu %v", cap(slice))
for i:=0;i<488;i++{
	slice2 = SlicedanElementiPoz2(slice2, 1)
}
fmt.Printf("Element tutumu Pozduqdan sonra %v, say : %v\n", cap(slice2), len(slice2))

```

    1)
    Element tutumu %v 500
    Element tutumu Pozduqdan sonra 500, say : 12
    2)
    Element tutumu %v 500
    Element tutumu Pozduqdan sonra 500, say : 12


Ona görə əgər pozma halları çoxdursa və **yaddaş həcmi önəmlidirsə**. 
- yeni buffer yaratmaq ən münasib variantdır. memory leak halını azaldır
- diapozon pozma. sürət önəmlidirsə


```go
func SlicedanElementiPozNoLeak(slice []int, index int) []int {
	if index <0 || index > len(slice)-1 {
		return slice
	}
	var new = make([]int, len(slice)-1)
	copy(new, slice[:index])
	copy(new[index:], slice[index+1:])
    return new
}

func SlicedanRangePozNoLeak(slice []int, begin_index int, end_index int) []int {
	if begin_index <0 || begin_index > len(slice)-1 || begin_index>end_index {
		return slice
	}
	if end_index>len(slice) {
		end_index = len(slice)
	}
	var new = make([]int, len(slice)-(end_index-begin_index)) 
	copy(new, slice[:begin_index]) 
	copy(new[begin_index:], slice[end_index:])
    return new
}

%%
slice :=make([]int, 500)
for i:=0;i<500;i++{
	slice[i]=i
}
fmt.Println("1)\nElement tutumu %v", cap(slice))
for i:=0;i<488;i++{
	slice = SlicedanElementiPozNoLeak(slice, 3)
}
fmt.Printf("Element tutumu Pozduqdan sonra %v, say : %v, %v\n", cap(slice), len(slice), slice)

slice1 :=make([]int, 500)
for i:=0;i<500;i++{
	slice1[i]=i
}
fmt.Println("2)\nElement tutumu %v", cap(slice1)) 
slice1 = SlicedanRangePozNoLeak(slice1, 3, 3+488) 
fmt.Printf("Element tutumu Diapozon Pozduqdan sonra %v, say : %v, %v\n", cap(slice1), len(slice1), slice1)

```

    1)
    Element tutumu %v 500
    Element tutumu Pozduqdan sonra 12, say : 12, [0 1 2 491 492 493 494 495 496 497 498 499]
    2)
    Element tutumu %v 500
    Element tutumu Diapozon Pozduqdan sonra 12, say : 12, [0 1 2 491 492 493 494 495 496 497 498 499]


### Array və Slice arasında fərqlər. xülasə

- Array ölçüsü statikdir. Slice artırmaq olur. Əgər tutum imkan verirsə eyni bufferə işarə edecək
- Array value tipdir. Slice reference tipdir. Mənimsətmə və funksiyaya vermə zamanı dəyişim slice-da həmin buffer üzərində olur.   
Array-də isə pointer istifadə etməsən kopyası üzərində olur. (<span style="color:red">Funksiyalar bölməsində nümunə var</span>)
