---
title: " Basic Output"
author: "abdelrauf"
date: 2023-07-31
tags: ["Basic", "Output", "Go" ]
---

### BASIC INPUT OUTPUT

Sadə konsola giriş və çıxış əmrlərini öyərənək.   
Ümumən, ilk öncə stdin stderr stdout anlamlarını qısa şərh edək.

- stdin - Standart giriş (Standard Input) proqramın daxil olan məlumatı oxuduğu kommunikasiya kanalıdır.  
- stdout - Standart çıxış (Standard Output) proqramın xaric olan məlumatı yazdığı kommunikasiya kanalıdır.  
- stderr - Standart xəta (Standard Error) proqramın xaric olan xəta məlumatlarını yazdığı kommunikasiya kanalıdır.  


Bunları həmçinin axın (stream) də adlandıra bilərik.  
Linux-da ümumən isə fayl sözünü də işlədə bilərik.  
Bu zaman stdin, stdout, stderr -in fayl deskriptorları da olacaq. Fayl deskriptoru deyəndə açılan fayla xas ədəd nəzərdə tutulur.  

Proses yarananda 3 standard fayl axını yaranır.  

|Fayl desktiptoru| Standard fayl axını|
|:---:|:---:|
| 0 | stdin   |
| 1 | stdout   |
| 2 | stderr |  



<pre>
             |         |    -> stdout
stdin  --->  | proqram |
             |         |    -> stderr

</pre>

Adi halda axınlar belə bağlanır. 

- stdin  - klaviaturaya (keyboard)  
- stdout - konsol ekranına   
- stderr - konsol ekranına  




Shell-də program icra etdikdə fərqli yerlərə də bağlamaq olar.

Meselen:   
```bash
./bizim_program  <input_fayl.txt   1>output.txt  2> error.txt   
```
Burada 

stdin - input_fayl.txt   
stdout - output.txt  
stderr - error.txt faylına bağlanır.


#### Golang -de output

Bunun üçün "fmt" standard kitabxanasından istifade edirik.

```Go
func Print(a ...any) (n int, err error)
func Println(a ...any) (n int, err error)
func Printf(format string, a ...any) (n int, err error)

Burada Arqumentlər:
  a istənilen tip və sayda dəyişən
  format - format mətni
Funksiya qaytarır:
  n - yazılmış bayt sayı
  err - xəta. Baş vermədikdə nil olur
```



```go
import "fmt"
```

İndi ise ekrana yazı verək. 
```Go 
fmt.Println 
```
Qeyd edək ki, o dəyişənlərin qiymətlərini stdout-a yazanda arasına " " boşluq əlavə edəcək.   
Sonuna isə yeni sətir əlavə edəcək.  


```go
%%
fmt.Println("Salam","Dünya")
fmt.Println("Salam","Azərbaycan")
fmt.Println(44, 2023, 5.789)
```

    Salam Dünya
    Salam Azərbaycan
    44 2023 5.789


```Go 
fmt.Print 
``` 
Bu da eyni işi görür. Ancaq string dəyişənlərin sonuna boşluq, həmçinin ümumi sonuna yeni sətir əlavə etmir. 


```go
%%
fmt.Print("Salam","Dünya")
fmt.Print("Salam","Azərbaycan")
fmt.Print(44, 2023, 5.789)
```

    SalamDünyaSalamAzərbaycan44 2023 5.789

Yeni sətir və boşluqları özümüz əlavə etmək istədikdə:
<pre> 
' ' - boşluq 
\n  - yeni sətir 
\\  - geri sləş \ 
\t  - tab 
\n\r  - yeni sətir (windows üçün)
\" - dırnaq
</pre>

Burada \ geri sləş vasitəsilə həmin yazıla bilinməyən simvolları veririk.  
Həmçinin geri sləş-in özünü vermək üçün iki dəfə istifadə etmişik.  
Həmçinin string tipimizin qiyməti " içərisində olduğundan onu da \-lə veririk. 


```go
%%
fmt.Print("Salam ","Dünya\n")
fmt.Print("Salam ","\"Azərbaycan\"\n")
fmt.Print(44, 2023, 5.789)
```

    Salam Dünya
    Salam "Azərbaycan"
    44 2023 5.789

##### Formatlı print 
```Go
fmt.Printf
```

Bezi Formatlara baxaq:

|  format |  əks etdirmə şəkli |
|---|---|
| %b  |  ikili sayda |
| %o  |  8-lik say sistemində |
|  %O | 8-lik say sistemi o prefiksi ilə  |
| %d  |  10-luq sayda  |
|  %x |  16-lıq sayda 1 baytı iki simvolla  [0-9a-f]|
| %X  | 16-lıq say  1 baytı iki simvolla, lakin  böyük hərflə [0-9A-F]|
|  %g |  onluq kəsr ədədləri kimi |
|  %s |  mətn ve simvollar ardıcıllığı |
|  %t |  bul üçün true false kimi|
|%e| elmi riyazi formatda|
|%v | təyin edilən tipin formatına uyğun|
|	%#v | Go sintaksis formatına uygun|
|%T | arqumentin Tipi|
|%c | simvol |
|%q|  simvol və ya simvol ardıcıllığı dırnaq daxilində, bəzi görülməyən simvolları \ ilə əks etdirir|
|%+q|  simvol və ya simvol ardıcıllığı dırnaq daxilində, ASCII-dən fərqli simvolları \ ilə əks etdirir|
|%U| unikod|
|%#U| unikod + simvol |
|%p| pointer|



```go
%%
h := 3.85
mətn := "👍Mətn"
simvol := '👍'

fmt.Printf("Formatlı əks %d, %g, %e, %s, %v, %v , %#v, %T, %q, %+q, %q, %U, %#U\n", int(h), h, h, mətn, h, mətn, h, h, mətn, mətn, simvol, simvol, simvol)
```

    Formatlı əks 3, 3.85, 3.850000e+00, 👍Mətn, 3.85, 👍Mətn , 3.85, float64, "👍Mətn", "\U0001f44dM\u0259tn", '👍', U+1F44D, U+1F44D '👍'



Bndan elave asagidaki deyishme usullarini da formatlara tetbiq etmek olar. 

|  Görüntü | Format  | izah   |
|---|---|---|
| ' 24' | % d|
| '+24'  |  %+d |  isareli eded ,yalniz edede aiddir |
| '  20'  | %{n}d , %4d | 4 sayda olub sağa sıxılmış yazı   |
| '15  '  |  %-{n}d |  n sayda sola sıxılmış yazı  |
|  '0025' | %0{n}d  | n sayda 0 əlavə |
|'122.15'   | %.{n}f ,%.2f  |  n sayda onluqdan sonrakı kəsr ədəd |
| '  122.46'  |   %{n}.{m}f| n sayda ümumi uzunluq sayı, m sayda kesir hissə sayı|
| '41 6C 6D 61 6C C4 B1'  |   % X | 16-lıq sayda bir bayt iki simvolla və aralarında boşluqla |



```go
%%
fmt.Printf("%5s | %5s | %7.5f | A% d \n","Al", "sat", 3.45, 44)
fmt.Printf("%X \n","Almalı")  
fmt.Printf("% X \n","Almalı") 
```

       Al |   sat | 3.45000 | A 44 
    416C6D616CC4B1 
    41 6C 6D 61 6C C4 B1 



```go
%%
fmt.Printf("%-5s | %-5s | %7.5f \n","Al", "sat", 3.45)
fmt.Printf("%-5s | %-5s | %7.5f \n","Vur", "yix", 3.45)
```

    Al    | sat   | 3.45000 
    Vur   | yix   | 3.45000 


Bundan başqa həmin format dəyişimlərində olan ədədləri də dəyişən kimi verə bilərik.  
Buna dəyişim ədədlərini [arqumentin_indeksi + 1]*, formatı isə [arqumentin_indeksi + 1]format verməklə nail ola bilərik.  


```go
%%
fmt.Printf("%6.4f\n", 12.0)
fmt.Printf("%[3]*.[2]*[1]f\n", 12.0, 4, 6) //[1]f 12.0, [2]* 4, [3]* isə 6

fmt.Printf("%-10s |\n", "Al")
fmt.Printf("%-[2]*[1]s |\n", "Al", 10) //[1]s AL, [2]* isə 10 
```

    12.0000
    12.0000
    Al         |
    Al         |


##### stdout, stderr-i özümüz yönləndirək  
(Qeyd: bununla həmçinin digər açdığımız fayla da yönləndirə bilərik)   
Bunun üçün Fprintf -dən istifadə edirik.  
Əslən Printf elə Fprintf stdout-dur.   


```go
// os paketini import edək.  
import "os"
```


```go
%%
fmt.Fprintln(os.Stdout, "Hello stdout")
fmt.Fprintln(os.Stderr, "Hello stderr")
```

    Hello stderr


    Hello stdout


Qeyd: 
Go-da print funksiyaları özündə builtin funksiya kimi də var.     
print, println  builtin funksiyalar


```go
%%
print("builtin print  ")
println(99)
println(88)
```

    builtin print  99
    88

