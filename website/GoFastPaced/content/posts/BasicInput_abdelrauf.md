---
title: BasicInput
author:abdelrauf
date: 2023-07-31
tags: ["Basic", "Input", "Go" ]
---

#### Basic Input

Bunun üçün "fmt" standard kitabxanadan istifade edirik   
Burada **bosluqlarla format olunmus** stdin (klaviatura) daxilolmalarını dəyişənlərə mənimsətməni nəzərdən keçirəcəyik.    
Scan və bacıları Scanln Scanf
```Go
func Scan(a ...any) (n int, err error)
func Scanf(format string, a ...any) (n int, err error)
func Scanln(a ...any) (n int, err error)
Burada Arqumentlər:
    format - format mətni 
    a istenilen tip və sayda dəyişən. Dəyişənin əvvəlinə & qoyuruq
Funksiya qaytarır:
  n - uğurlu Scan olunmuş element sayı
  err - xəta. Baş vermədikdə nil olur

```
Bildiyimiz kimi stdin klaviaturaya bağlı olur, əgər istifadəçi proqramı çağırdığı zaman fayla bağlamayıbsa

Əlavə oxu: 
- BasicOutput.ipynb nəzərdən keçirin
- (ingiliscə) Command line stdin, stdout, stderr, Redirections, Pipe 



```go
import "fmt"
```

Scan - funksiyasi " \n\t" boşluqlarla ayrılmış daxiletməni dəyişənlərə mənimsədəcək. 

fmt.Scan(&deyishen1, &deyishen2)

Fikir verin ki  **& işarəsini dəyişənin əvvəlinə** əlavə edirik. Bu işarə funksiya və digər bölmələrdə aydın olacaq


```go
 //Gonb -de input işləmədiyindən Sscan la əvəzləyib daxiletməmizi input-da göstərəcəyik
func main(){
	var a int
	var b int

	//fmt.Scan(&a,&b)  //Daxil edilən mətn: 456 enter düyməsi 546 enter
	 input:="456 \n 546 "
	 fmt.Sscan(input, &a, &b)
	fmt.Println("a:", a, "b:", b)
	}
```

    a: 456 b: 546


fmt.Scanln -de eyni ile işləyir lakin yeni sətirə '\n' kimi oxuyur

Qeyd: Gonb -de input işləmədiyindən Sscanln la əvəzləyib daxiletməmizi input-da göstərəcəyik


```go
 //Gonb -de input işləmədiyindən Sscanln la əvəzləyib daxiletməmizi input-da göstərəcəyik
 func main(){
	var a int
	var b int

	//fmt.Scanln(&a,&b)  //Daxil edilən mətn: 456 enter düyməsi 546 enter
	 input:="456 \n 546 "
	 fmt.Sscanln(input, &a, &b)

	fmt.Println("a:", a, "b:", b) 
	// b budəfə 0 oldu çünki \n yeni sətirə qədər oxuyub dayanır
	}
```

    a: 456 b: 0


Bu qebilden olanlara Scanf -de aiddir. Lakin evvelkilerden ferqli olaraq hansi tipde saxlayacagimizi formatla daha deqiq vermek olur



```go
 //Gonb -de input işləmədiyindən Sscanf la əvəzləyib daxiletməmizi input-da göstərəcəyik
 func main(){
	var a int
	var b float32

	//fmt.Scanf("%d %f", &a,&b)  //Daxil edilən mətn: 456 boşluq 546 enter
	 input:="456 546 "
	 fmt.Sscanf(input, "%d %f", &a, &b)

	fmt.Println("a:", a, "b:", b)  
	}
```

    a: 456 b: 546


Qeyd edək ki fmt.Scan elə **fmt.Fscan (os.Stdin, ...)** dır.   
Sadəcə burada **os.Stdin in yerinə özümüz açdığımız faylı** da vermək olur.   
Sadə halda fmt.Scan istifadə edib İstifadəçidən istədiyi  **fayla bağlamağını** seçim kimi verə bilərik

```bash
./proqram  **< stdin_yazilari_olan_fayl.txt**
```

<span style="color:red">Ümumən **Scan ve bacıları yalnız boşluqlarla ayrılmış formatlarla yaxşı işləyir**. </span>  
Başqa hallar üçün isə ümumi sətiri oxuyub parse etmək lazımdır.  
Bunu daha sonra keçəcəyik  
<pre>
import (
    "bufio"
    "fmt"
    "os"
)

scanner := bufio.NewScanner(os.Stdin)
</pre>


