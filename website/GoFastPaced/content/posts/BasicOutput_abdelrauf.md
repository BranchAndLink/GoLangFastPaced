---
title: " Basic Output "
author: "abdelrauf"
date: 2023-07-31
tags: ["Basic", "Output", "Go" ]
---

### BASIC INPUT OUTPUT

Sade konsola giris çıxış əmirlərini öyərənək. Ümumən isə stdin stderr stdout söhbəti qısa şərh edək

Proses yarananda 3 standard fayl stream-i ona bağlanir
- 0 - stdin  
- 1 - stdout   
- 2 - stderr   



<pre>
             |         |    -> stdout
stdin  --->  | proqram |
             |         |    -> stderr

</pre>

Sadə halda 

stdin adi halda keyboarda    
stdout - konsol ekranına   
stderr - konsol ekranına  bağlanır


Shell-de program run edildikde Bunu ferqli yerlere de baglamaq olur.

Meselen:   
```bash
./bizim_program  <input_fayl.txt   1>output.txt  2> error.txt   
```
Burada 

stdin - input_fayl.txt   
stdout - output.txt  
stderr - error.txt faylina bağlanır


#### Golang -de output

Bunun üçün "fmt" standard kitabxanadan istifade edirik

```Go
func Print(a ...any) (n int, err error)
func Println(a ...any) (n int, err error)
func Printf(format string, a ...any) (n int, err error)

Burada Arqumentlər:
    a istenilen tip və sayda dəyişən
    format - format mətni
Funksiya qaytarır:
  n - yazılmış bayt sayı
  err - xəta. Baş vermədikdə nil olur
```



```go
import "fmt"
```

Indi ise ekrana nese verek
fmt.Println - ekrana deyisenleri eks etdirir. Qeyd edek ki o deyishenlerin arasina " " boshluq elave edecek. Ve hemchinin ln -de yeni setir elave edir


```go
%%
fmt.Println("Hello","WORLD", "Pi", 3.14)
fmt.Println("Hello","WORLD", "Pi", 3.14)
```

    Hello WORLD Pi 3.14
    Hello WORLD Pi 3.14


Print - de eyni ishi gorur amma sona yeni setir ve bosluqlari elave etmir elave etmir


```go
%%
fmt.Print("Hello","WORLD", "Pi", 3.14)
fmt.Print("Hello","WORLD", "Pi", 3.14)
```

    HelloWORLDPi3.14HelloWORLDPi3.14

Eyni sheyi etmek uchun ozumuz etmeliyik. "\n" - yeni setirdir " " - boshluqdur
bundan elave:  
- \\  - arxsa \ slash, \ isharesi xususi rol oynadigindan
- \t  - tab 
- \n  - yeni setir
- \n\r  -yeni setir windows uchun

Strings bolmesinde bu movzuya baxacagiq


```go
%%
fmt.Print("Hello ","WORLD ", "Pi ", 3.14, "\n")
fmt.Print("Hello ","WORLD ", "Pi ", 3.14, "\n")
```

    Hello WORLD Pi 3.14
    Hello WORLD Pi 3.14


##### Formatli print uchun fmt.Printf

Bezi Formatlara baxaq:

|  format |  menasi |
|---|---|
| %b  |  ikili sayda eks etdirmek |
| %o  |  8-lik say sisteminde |
|  %O | 8-lik say sistemi o prefixi ile  |
| %d  |  10-lu sayda  |
|  %x |  16-liq sayda eks etdirir 0-9A-F|
| %X  | 16-liq say  |
|  %g |  onluq kesr ededleri kimi eks etdirmek |
|  %s |  metn ve simvollar ardicilligi |
|  %t |  bul üçün true false eks etdirir |
|%e| elmi riyazi formatda|
|%v | default tipin formatina uygun|
|	%#v | Go sintaksis formatina uygun eks|
|%T | arqumentin Tipi|
|%c | simvol |
|%q|  dirnaqda simvol|
|%u| unikod|
|%p| pointer eks etdirir|





```go
%%
h:=3.85
metn:="Metn"

fmt.Printf("Format li eks %d, %g, %e, %s, %v, %v , %#v, %T,\n", int(h), h, h, metn, h, metn, h, h)
```

    Format li eks 3, 3.85, 3.850000e+00, Metn, 3.85, Metn , 3.85, float64,


Formatda da ededler uchun %+d qoysaq isareni de gosterecek
BUndan elave asagidaki deyishme usullarini da formatlara tetbiq etmek olar. Qeyd (_ isahresi boshlugu ve sixmani gostermek uchundur)

|  Goruntu | FormatNumune  | Izah   |
|---|---|---|
| +24  |  %+d |  isareli eded ,yalniz edede aiddir |
| __20  | %{n}d , %4d | 4 sayda olub saga sixir yazini   |
| 15__  |  %-{n}d |  n sayda sola sixir ()  |
|  0025 | %0{n}d  | n sayda 0 elave edir (eded formatina aiddir) |
|122.15   | %.{n}f ,%.2f  |  n sayda onluqdan sonraki kesir hisse sayi, (kesir edede aiddir) |
| __122.46  |   %{n}.{m}f| n sayda umumi uzunluq sayi, m sayda   kesir hisse sayi (kesir edede aiddir) |



```go
%%
fmt.Printf("%5s | %5s | %7.5f \n","Al", "sat", 3.45)
fmt.Printf("%5s | %5s | %7.5f \n","Vur", "yix", 3.45)
```

       Al |   sat | 3.45000 
      Vur |   yix | 3.45000 



```go
%%
//sola sixilmis, saga bosluq elave etmekle
fmt.Printf("%-5s | %-5s | %7.5f \n","Al", "sat", 3.45)
fmt.Printf("%-5s | %-5s | %7.5f \n","Vur", "yix", 3.45)
```

    Al    | sat   | 3.45000 
    Vur   | yix   | 3.45000 


##### Bes nece edek ki stdout, stderr-i ozumuz sechek (Qeyd bu hemçinin öz açdığın fayla da aiddir)
Bunun üçün Fprintf -dən istifadə edirik. Ve stdout stdin i vermek üçün    
os paketini import edek
Əslən Printf ele Fprintf stdout -dur


```go
import "os"
```


```go
%%
fmt.Fprintln(os.Stdout, "Hello stdout")
fmt.Fprintln(os.Stderr, "Hello stderr")
```

    Hello stderr


    Hello stdout


Eger komandamizi icra etsek 1> ile stdout-a yazilana basqa fayla 2> ile ise stderr-i basqa fayla yaz bileceyik.

Qeyd linux, windows komanda cagirmalari, redirection bolmesini oxuyun

Qeyd: GO-da print funksiyaları özündə də var.     
print, println  builtin funksiyalar


```go
%%
print("builtin print  ")
println(99)
println(88)
```

    builtin print  99
    88

