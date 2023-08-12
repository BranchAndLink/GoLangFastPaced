---
title: " Pointers"
author: "abdelrauf"
date: 2023-08-12
tags: ["Pointers", "Go" ]
---

### Pointer

Pointer yaddaşda hər hansı bir tip üçün ayrılan yerin ünvanını saxlayır. Həmin tipə pointer-in əsas tipi deyilir.   

Pointer tipi:  * Esas_Tip   

Göründüyü kimi Tip-in əvvəlinə * artırmaqla pointer tipini təyin edirik.   
Adi elan edildikə pointer tipi nil saxlayır. nil o demekdir ki pointer heç bir address saxlamır və boşluğa işarə edir.   
(Qeyd: C dilində ədəd olaraq null ( Go nil ) dedikdə 0 addressi nəzərdə tutulur. Çünki bir çox sistemlərdə 0 virtual addressi fiziki yaddaşa map olmur və ona müraciət xətaya səbəb olur.   
 Lakin məntiq olaraq nil 0 deyil. )   




Bir neçə pointer tipinə  nümunə yazaq   
```Go
var ptrToStr  *string
var ptrToInt *int
var ptrToStr *struct {
    m int
    k float32
}


```


```go
%%
var ptrToStr  *string
var ptrToInt *int
var ptrToStruct *struct {
    m int
    k float32
}
fmt.Println(ptrToStr, ptrToInt, ptrToStruct)

```

    <nil> <nil> <nil>


Pointer tipini nəyəsə işarə etməsi üçün ona hər hansı real dəyişənin addresini mənimsədə bilərik.   
Bunun üçün & işarəsindən istifadə edirik.  


```go
%%
var ptrToInt *int
k := 999
ptrToInt = &k
fmt.Println("k-in ünvanı ", ptrToInt)
```

    k-in ünvanı  0xc0000a2000


Pointer tipinin işarə etdiyi əsas tipdə olanın qiymətini dəyişmək üçün isə * istifadə edirik. Buna dereference deyilir. 



```go
%%
var ptrToInt *int
k := 999
fmt.Println( k)
ptrToInt = &k
*ptrToInt = 888
fmt.Println( *ptrToInt, k)
```

    999
    888 888


Qeyd edəki ki əgər pointer boşluğa işarə edirsə onu dereference etmək xətaya səbəb olacaq 


```go
%%
var ptrToInt *int 
fmt.Println( *ptrToInt) 
```

    panic: runtime error: invalid memory address or nil pointer dereference
    [signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x48c681]
    
    goroutine 1 [running]:
    main.main()
    	 [7m[[ Cell [298] Line 3 ]][0m /tmp/gonb_7a8e7f4d/main.go:22 +0x61
    exit status 2


Go dilində C dilindən fərqli olaraq pointer məhdud imkanlara malikdir. Pointer dəyişəni üzərində **riyazi əməliyyatlar aparmaq olmur**.  
Go dilində bu özəlliyi əldə etmək üçün **unsafe** paketindən istifadə etməli oluruq. Onu növbəti dərslərdə nəzərdən keçirəcəyik.  

Biz bilirik Go dilində funksiyaya arqumentlər sürəti ilə (kopyası ilə) ötürülür.   
Lakin bəzən dəyişəni funksiyaya həm arqument kimi ötürüb həm də onu daxildə dəyişmək ehtiyacımız yaranır.  
Bu zaman pointer-dən istifadə edə bilərik.  


```go
func modify(a *int) {
	*a = *a + 10 * *a // *a + 10 * (*a) 
}

%%
a:=11
modify(&a)
modify(&a)
modify(&a)
fmt.Println(a)
```

    14641


Yuxarıdakı misaldan sonra Scan və onun bacılarında nə üçün & istifadə etdiyimiz anlaşılır.  

Go dilində həmçinin qeyd etmişdik ki bəzi tiplər (slice, map..) reference tip-dir.   
Bu o deməkdir ki Onlar da həmçinin pointer-dirlər və real data struktura göstərici saxlayırlar.   
Reference və pointer sinonim sözlərdir. Bəzi dillərdə (C++) reference daha ciddi məhdud xarakter daşıyır. 

Əvvəl qeyd etdik ki, pointer dəyişəni adi  halda nil olur. Ona görə istifadə üçün onu real dəyişənə işarə edirik.   
Lakin çox zaman biz istəyirik ki əsas tipdə   yaddaş ayrılsın və bizim pointer onun ünvanını saxlasın.   
Bunun üçün built-in funksiya olan **new** dan istifadə edirik.  
```Go
  new(Type)
```


```go
%%
var a *int = new(int)
*a = 889
*a += 10
fmt.Println(a, *a)
```

    0xc0000a2000 899


Bu new-dan əslən bir çox məsələlərdə ekvivalent bəyanlar üçün də istifadə edə bilərik.  
Məsələn biz make-lə slice hazırlamaq əvəzinə new ilə array alıb onun slice-ın ala bilərik. 



```go
%%
var slice_len50_cap100 = make([]int, 50, 100)
var slice_len50_cap100_2ci_usul =  new([100]int)[0:50]

fmt.Println(len(slice_len50_cap100), cap(slice_len50_cap100))
fmt.Println(len(slice_len50_cap100_2ci_usul), cap(slice_len50_cap100_2ci_usul))
```

    50 100
    50 100


Go dilində funksiya daxilində yaradılmış dəyişənin addresini pointer olaraq qaytarmaq normal hal sayılır.  
(Qeyd edək ki Bu C/C++ dilində stack lokal dəyişəninə işarə olduğundan kobud xətadır )


```go
func itsOk() *int {
	a:=500

	return &a
}


%%
ptrToH := itsOk()
*ptrToH += 50
fmt.Println(ptrToH, *ptrToH)
```

    0xc0000220c8 550


O üzdən bəzən new-suz belə lokal dəyişən yaradıb onun addresini pointerə mənimsədib işləmək normaldır.   
Həmçinin nəzərə almaq lazımdır ki Go kompilyatoru bizə əlavə * dereference əməliyyatı aparmadan belə işləməyə imkan verir.  


```go
%%
//lokal anonim yaratdığımız array-in addresini birbaşa mənimsədək
ptrToArray := & [10]int{}
(*ptrToArray)[0] = 44
// Go dili * dereference etmədən belə bizi başa düşəcək
ptrToArray[5] = 5
fmt.Println(*ptrToArray)


```

    [44 0 0 0 0 5 0 0 0 0]


Qeyd etdiyimiz struktur üçün də keçərlidir. Və onda da üzv dəyişən üçün əlavə * əməliyyat aparmaq lazım deyil.  



```go
%%
//lokal anonim yaratdığımız strukt-un addresini birbaşa mənimsədək
ptrToStruct := & struct {
    m int
    n float32
}{}
(*ptrToStruct).m = 44
// Go dili * dereference etmədən belə bizi başa düşəcək
ptrToStruct.n = 5.5
fmt.Println(*ptrToStruct)
```

    {44 5.5}


lakin bilmək lazımdır ki **new** ilə **reference tip** elan etdikdə lokal anonim dəyişənin addresini götürməyimizdən fərqli olaraq inisializasiya olunmamış yaddaşla işləyirik


```go
%%
p1 := &[]int{}    // p1  [5]int{} inisializasiya olun-muş slice-a işarə edir.
p2 := new([]int)  // p2  inisializasiya olun-ma-mış slice-a işarə edir. nil olur.
fmt.Println(*p1 == nil,  *p2 == nil)
fmt.Println(len(*p1), len(*p2))

```

    false true
    0 0


Pointerdən Pointerə işarələr


```go
%%
var h int
var ptr_h *int = &h
var ptr_ptr_h **int = &ptr_h
var ptr_ptr_ptr_h ***int = &ptr_ptr_h

***ptr_ptr_ptr_h = 555
fmt.Println(h)

```

    555



```go
%% 
var h int = 898
var ptr_ptr_X **int = new( *int)
*ptr_ptr_X = &h
fmt.Println( **ptr_ptr_X)

```

    898


& və * ardıcıllığı.  
**Qeyd bu cür kod yazılışı düzgün deyil.**     
Sadəcə bu iki operatorun bir-birinə əks olduğuna nümunə üçün qeyd edirik.   


```go
%%
var h = 444
fmt.Println(*&h)
fmt.Println(*&*&h)
fmt.Println(*&*&*&h)
fmt.Println(*&*&*&*&h)
```

    444
    444
    444
    444

