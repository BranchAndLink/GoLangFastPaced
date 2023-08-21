---
title: " Error"
author: "abdelrauf"
date: 2023-08-21
tags: ["Error", "defer", "recover", "panic", "Go" ]
---

### Error (Xəta) 

Ümumən, proqramlaşdırmada xətaların idarə edilməsi önəmlidir.  
Bu səbəbdən Go dilində universal blokda təyin edilən error builtin tipi var.   
Bu error tipi interfeysdir.  
```Go
type error interface {
    Error() string
}
``` 

Biz bu interfeyslə dolayı olsa da Basic Input Output fəslində qarşılaşmışdıq.  
Məsələn: 
```Go
func Scan(a ...any) (n int, err error)
```

Məsələn burada ədəd yerinə fərqli bir şey yazsaq  
"expected integer" xətası baş verəcək.


```go
import (
	"fmt"
	"github.com/janpfeifer/gonb/gonbui"
)
%%
gonbui.RequestInput("Eded daxil et: ", false)
var x int
_, err := fmt.Scan(&x) 
if err !=nil {
	fmt.Println(err)
}
```

    expected integer


İnterfeysimiz sadədir. Aydın məsələdir ki, bunu implementasiya etmək rahatdır.  
Lakin çox zaman bizə elə mətn qaytaran error və onunla əlaqəli funksiyalar bəs edir, hansı ki errors və fmt package-də var. 
```Go
// errorString   error-un sadə implementasiyası .
type errorString struct {
    s string
}

func (e *errorString) Error() string {
    return e.s
}
```
Onun istifadəsinə baxaq



```go
import (
	"fmt"
	"errors"
)

%%

err1 := errors.New("Xeta#1")
fmt.Printf("%T %v\n", err1, err1)
err2 := fmt.Errorf("Xeta#%d", 2)
fmt.Printf("%T %v\n", err2, err2)
```

    *errors.errorString Xeta#1
    *errors.errorString Xeta#2


errors və fmt package üzrə digər metodlara aid nümunəyə baxaq.  
Xüsusilə  real kodlama zamanı səhvləri yuxarı ötürməyimizdə kömək edən aşağıdakı  
qayda ilə bağlamaya almaq və açmaq məsələsinə baxaq.  


```go
%%
err1 := errors.New("Xeta#1")  
//fmt.Errorf %w ilə error bağlaması düzəldə bilərik

err_wrapped := fmt.Errorf("Xeta#2 (%w )" , err1 ) 
fmt.Println("{\n",err_wrapped,"\n}")
err_wrapped2 := fmt.Errorf("Xeta#3 (%w )" , err_wrapped ) 

fmt.Println("{\n",err_wrapped2,"\n}")

//bağlamanı açaq
err_unwrapped2 := errors.Unwrap(err_wrapped2)
fmt.Println(err_unwrapped2 )
err_unwrapped1 := errors.Unwrap(err_unwrapped2)
fmt.Println(err_unwrapped1 )

```

    {
     Xeta#2 (Xeta#1 ) 
    }
    {
     Xeta#3 (Xeta#2 (Xeta#1 ) ) 
    }
    Xeta#2 (Xeta#1 )
    Xeta#1


Və son olaraq öz error tipimizə aid nümunəyə baxaq.  
(Sadəlik üçün burda bütün hallarda xəta qaytarırıq.)


```go
type BizimXeta struct {
    xeta_mesaji string
	xeta_kodu int
}

func (bx *BizimXeta) Error() string {
	return fmt.Sprintf("msg:{%v}, kod:{%d}", bx.xeta_mesaji, bx.xeta_kodu)
}

func hemishe_xetali() error {
	return &BizimXeta{"xeta mesajimiz", 44}
}

func hemishe_xetali_ikinci_funksiya() error {
	ret := hemishe_xetali()
   return fmt.Errorf("hemishe_xetali_ikinci_funksiya xeta [%w]", ret)
}

%%
err := hemishe_xetali_ikinci_funksiya()

fmt.Println(err) 

```

    hemishe_xetali_ikinci_funksiya xeta [msg:{xeta mesajimiz}, kod:{44}]


Errors və fmt paketini, həmçinin bu mövzunu daha ətraflı və geniş nəzərdən keçirmək üçün   
https://go.dev/blog/error-handling-and-go   
https://pkg.go.dev/errors   
https://pkg.go.dev/fmt#Errorf  

### Defer  (Təxirə salma)
**defer** ifadəsi ilə funksiya bloku daxilində olan funksiya çağırışlarını siyahıya alıb icrasını təxirə salmaq olur.  
Həmin siyahıdakı funksiya çağırışlarının icrası funksiya blokundan return ilə çıxdıqdan sonra baş verir.  
Defer adətən müxtəlif resurslarla işlədiyimiz zaman onların düzgün azad edilməsində işimizə yarayır. (Məsələn: File-larla işlədikdə)  
İndi isə defer-in işləmə xüsusiyyətlərini sadalayaq:  


1.  defer edilmiş funksiya çağırışlarının arqumentlərinin qiymətləndirilməsi defer çağrıldığı anda həyata keçir
2.  siyahıya alınmış funksiya çağırışları lifo(last in first out) sıralaması ilə icra edilir. Yəni sonuncu ilk icra edilir. 
3.  defer edilmiş funksiya çağırışları içərisində adlı nəticə dəyişənlərini həm istifadə edə həm də onlara oxuyub mənimsədə bilərik.  
( return-dən sonra icra edildiyini unutma).

1-ciyə aid nümunəyə baxaq


```go
func a() {
    i := 0
    defer fmt.Println("defer icra", i)
    i++
	fmt.Println("---", i)
    return
}
%%
a()
```

    --- 1
    defer icra 0


2-ciyə aid nümunə


```go
func b() {
    for i := 0; i < 4; i++ {
        defer fmt.Println("deferred funksiya", i)
    }
}
%%
b()
```

    deferred funksiya 3
    deferred funksiya 2
    deferred funksiya 1
    deferred funksiya 0


3-cüyə aid nümunə:


```go
func c() (i int) {
    defer func() { i++ }()
    return 1
}

%%
ret:= c()
fmt.Println(ret)
```

    2


#### Panic və Recover
Go dilində **panic** builtin funksiyası vasitəsilə proqramın normal icrası dayanır.   
Nəticə etibarilə panic baş verən funksiyada defer edilmiş funksiya çağırışları icra edilir və  
geri qayıdır. Geri qayıdılan funksiyada yenə həmin iş təkrarlanır.  Bu proses  proqramdan  
çıxışa səbəb olana kimi davam edir.   
Qeyd edək ki panic kodda istifadə edilməsilə yanaşı, həmçinin məşhur runtime(icra prosesində) xətalar zamanı da baş verir. 


```go
%%
h:=456
z:=0
//Qeyd h/0 birbaşa yazsaq ağıllı GO kompilyasiya zamanı problemdən xəbər verəcək
h=h/z 
```

    panic: runtime error: integer divide by zero
    
    goroutine 1 [running]:
    main.main()
    	 [7m[[ Cell [2] Line 4 ]][0m /tmp/gonb_f39407f7/main.go:12 +0x5a
    exit status 2



```go
%%
var h =   [...] int {1,2,3,4,5}
fmt.Println(h[2]) 
fmt.Println(h[h[4]]) 
```

    3


    panic: runtime error: index out of range [5] with length 5
    
    goroutine 1 [running]:
    main.main()
    	 [7m[[ Cell [8] Line 4 ]][0m /tmp/gonb_f39407f7/main.go:13 +0x12b
    exit status 2



```go
func hemishe_panic() {
	defer fmt.Println("defer#hemishe_panic")
	panic("gözlənilməz xəta")
}

func caller() {
	defer fmt.Println("defer#caller")
	hemishe_panic()
}

%%
caller()
//görsənməyəcək
fmt.Println("davam")
```

    defer#hemishe_panic
    defer#caller


    panic: gözlənilməz xəta
    
    goroutine 1 [running]:
    main.hemishe_panic()
    	 [7m[[ Cell [27] Line 3 ]][0m /tmp/gonb_f39407f7/main.go:42 +0x73
    main.caller()
    	 [7m[[ Cell [27] Line 8 ]][0m /tmp/gonb_f39407f7/main.go:33 +0x70
    main.main()
    	 [7m[[ Cell [27] Line 12 ]][0m /tmp/gonb_f39407f7/main.go:57 +0x5e
    exit status 2


##### recover    
Bəzən bəzi runtime xətalarını həll etməklə proqramın icrasını dayandırma-maq daha münasib sayılır.  
Bu hallar üçün biz **recover** funksiyasından istifadə edə bilərik.  
Qeyd edək ki, recover defer edilmiş funksiya çağırışları içində əhəmiyyət kəsb edir.  
Əgər proqramın icrası əsasında panic baş vermişsə recover həmin panic kəmiyyətini götürür və icra prosesi normal qaydada davam edir.  
Həmçinin heç bir şey baş verməyibsə bu halda recover sadəcə nil qaytarır və əks təsirə səbəb olmur.   



```go
func division (a,b int) int {
	return a/b
}

func caller(a,b int) int{
	defer func(){
		if ret:=recover(); ret !=nil {
			fmt.Println("Xəta", ret)
		}
	}()
	return division(a,b)
}

%%
fmt.Println(caller(44,11))
fmt.Println(caller(44,0))
fmt.Println("davam")
```

    4
    Xəta runtime error: integer divide by zero
    0
    davam



```go
func hemishe_panic() {
	defer fmt.Println("defer#hemishe_panic")
	panic("gözlənilməz xəta")
}

func caller() {
	defer func(){
		fmt.Println("defer#caller")
		ret:=recover()
		if ret!=nil {
			fmt.Println(ret)
		}
	}()
	hemishe_panic()
}

%%
caller()
//görsənəcək
fmt.Println("davam")
```

    defer#hemishe_panic
    defer#caller
    gözlənilməz xəta
    davam


Qeyd:
Proqramda normal gözlənilən xətalar üçün error-dan istifadə etmək lazımdır.  
Məsələn, istifadəçi tərəfindən daxil edilən məlumatın düzgünlüyü və s.   
Lakin gözlənilməyən və ya proqramın dayandırılmasını labüd edən xətalar üçün panic istifadə etmək daha məsləhətdir.
