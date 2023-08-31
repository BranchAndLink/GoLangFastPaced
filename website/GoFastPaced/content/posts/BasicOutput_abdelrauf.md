---
title: " Basic Output"
author: "abdelrauf"
date: 2023-07-26
tags: ["Basic", "Output", "Go" ]
---

### BASIC INPUT OUTPUT

Sadə proqramlarda giriş və əks etdirmə üsullarına baxaq.  
  
Ümumən, ilk öncə stdin stderr stdout anlamlarını qısa şərh edək.

- stdin - Standart giriş (Standard Input) proqramın daxil olan məlumatı oxuduğu kommunikasiya kanalıdır.  
- stdout - Standart çıxış (Standard Output) proqramın xaric olan məlumatı yazdığı kommunikasiya kanalıdır.  
- stderr - Standart xəta (Standard Error) proqramın xaric olan xəta məlumatlarını yazdığı kommunikasiya kanalıdır.  


Bunları həmçinin axın (stream) adlandıra bilərik.  
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
Sonuna isə yeni sətir əlavə olunur.  


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
Bu da eyni işi görür. Ancaq string qiymətlərinin sonuna boşluq, və ümumi sonluğa yeni sətir əlavə etmir. 


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
&#92;n  - yeni sətir 
&#92;&#92;  - geri sləş &#92; 
&#92;t  - tab 
&#92;n&#92;r  - yeni sətir (windows üçün)
&#92;" - dırnaq
</pre>

Burada &#92; geri sləş vasitəsilə həmin yazıla bilinməyən simvolları veririk.  
Həmçinin geri sləş-in özünü vermək üçün iki dəfə &#92;&#92; istifadə etmişik.  
Həmçinin string tipimizin qiyməti " içərisində olduğundan onu da &#92;-lə veririk. 


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

Bəzi Formatlara baxaq:

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
|%q|  simvol və ya simvol ardıcıllığı dırnaq daxilində|
|%+q|  simvol və ya simvol ardıcıllığı dırnaq daxilində, ASCII-dən fərqli simvolları \ ilə əks etdirir|
|%U| unikod|
|%#U| unikod + simvol |
|%p| pointer|



```go
%%
h := 3.85
mətn := "👍Mətn"
simvol := '👍'

fmt.Printf("Formatlı əks %d, %g, %e, %s, %v, %v\n", int(h), h, h, mətn, h, mətn)
fmt.Printf("Formatlı əks %#v, %T, %q, %+q, %q, %U, %#U\n", h, h, mətn, mətn, simvol, simvol, simvol)
```

    Formatlı əks 3, 3.85, 3.850000e+00, 👍Mətn, 3.85, 👍Mətn
    Formatlı əks 3.85, float64, "👍Mətn", "\U0001f44dM\u0259tn", '👍', U+1F44D, U+1F44D '👍'



Bundan əlavə aşağıdakı format dəyişimlərindən də istifadə etmək olar. 

|  Görüntü | Format dəyişimi  | izah   |
|---|---|---|
|  |'#'	| alternativ format yaradır:  Məsələn: 0b (%#b) ikilik, 0x 16-lıq %#x  və s
| '&nbsp;24' | '%&nbsp;', '%&nbsp;d'| boşluq əlavə edir. formatlar: '% d', '% x', '% X'
| '41 6C 6D 61 6C C4 B1'  |   '%&nbsp;', '%&nbsp;X' | 16-lıq sayda bir bayt iki simvolla və aralarında boşluqla |
| '+24'  | '%+', '%+d' |  işarəli ədəd ,yalnız (%q və %d ilə işlənir) |
| '\U0001f44dM\u0259tn'  | '%+', '%+q' |  %q ilə işlənib ASCII simvol  |
| '&nbsp;&nbsp;&nbsp;&nbsp;20'  | '%{n}d' , '%4d' | n sayda olub sağa sıxılmış yazı (nümunədə 4 sayda)   |
| '15&nbsp;&nbsp;&nbsp;&nbsp;'  |  '%-{n}d', '%-4d' |  n sayda sola sıxılmış yazı (nümunədə 4 sayda)|
|  '0025' | '%0{n}d', '%02d' | n sayda 0 əlavə (nümunədə 2 sayda)|
|'122.15'   | '%.{m}f', '%.2f'  |  m sayda onluqdan sonrakı kəsr hissə sayı |
| '  122.46'  |   '%{n}.{m}f'| n sayda ümumi uzunluq sayı, m sayda kəsr hissə sayı|



```go
%%
mətn := "👍Mətn"
fmt.Printf("%X və ya %#X \n", "alma", "alma")
fmt.Printf("%U və ya %#U \n", '👍', '👍')
fmt.Printf("%q və ya %+q \n", mətn, mətn)
fmt.Printf("%5s | %5s | %7.5f | A% d \n","Al", "sat", 3.45, 44)
fmt.Printf("%X \n","Almalı")  
fmt.Printf("% X \n","Almalı") 
```

    616C6D61 və ya 0X616C6D61 
    U+1F44D və ya U+1F44D '👍' 
    "👍Mətn" və ya "\U0001f44dM\u0259tn" 
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


%v Vasitəsilə Go dili tipə uyğun formatı təyin edir. 
Həmçinin bununla mürəkkəb tipləri də çıxara bilirik.   
%#v isə Go sintaksisinə uyğun əks etdirir.



```go
%%
fmt.Printf("%v, %v, %v, %v\n", 55, 55.789, 5 + 4i, [...]int{4, 5, 6})
fmt.Printf("%#v, %#v, %#v, %#v\n", 55, 55.789, 5 + 4i, [...]int{4, 5, 6})
```

    55, 55.789, (5+4i), [4 5 6]
    55, 55.789, (5+4i), [3]int{4, 5, 6}


Əgər format-ı səhv versək xəta baş verir və o yazıda da ( %! işarəsi ilə) görsənir.  
Məsələn: ədəd formatına string versək 


```go
%%
fmt.Printf("%d", "hi")
```

    %!d(string=hi)

##### stdout, stderr-i özümüz yönləndirək  
(Qeyd: digər açdığımız fayla da yönləndirə bilərik)   
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


Go-da fmt paketi yerinə **print, println** builtin funksiyalarını da istifadə etmək olar.


```go
%%
print("builtin print  ")
println(99)
println(88)
```

    builtin print  99
    88

