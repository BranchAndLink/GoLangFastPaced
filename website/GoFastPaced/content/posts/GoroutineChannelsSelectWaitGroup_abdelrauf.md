---
title: " Goroutine Channels Select WaitGroup"
author: "abdelrauf"
date: 2023-09-06
tags: ["Goroutine", "Channels", "Select", "WaitGroup", "Go" ]
---

### Goroutine

Goroutine parallel eyni anda bir neçə funksiyanı icra etmək üçündür.  



Goroutine-ləri istifadə etmək üçün funksiya çağırışlarının əvvəlinə ```go``` sözünü əlavə edirik.  
defer kimi bunda da funksiyaya verilən arqumentlərin qiymətləndirilməsi çağırış anında baş verir.  
Funksiyanın icrası isə yeni goroutine-də baş verir.  
```

go funksiyamız(a, b, ...)

```



```go
import "time"

func MyRoutine(söz string, times int) {
	for i:=0; i< times; i++ {
		println(i, ") ", söz)
		//fasilə
		time.Sleep(50 * time.Millisecond)
	}
   
}

func main() {
	go MyRoutine("Salam", 7)
	go MyRoutine("Dünya", 3)
	MyRoutine("++++++", 2)
	time.Sleep(10 * time.Millisecond)
	MyRoutine("++++++", 1)
	time.Sleep(10 * time.Millisecond)
	MyRoutine("++++++", 2)
	time.Sleep(200 * time.Millisecond)
}
```

    0 )  ++++++
    0 )  Dünya
    0 )  Salam
    1 )  Salam
    1 )  ++++++
    1 )  Dünya
    2 )  Dünya
    2 )  Salam
    0 )  ++++++
    3 )  Salam
    0 )  ++++++
    4 )  Salam
    1 )  ++++++
    5 )  Salam
    6 )  Salam


##### Go planlaşdırıcı(scheduler)  
Go planlaşdırıcı çoxlu sayda goroutinlərin icrasını nizamlayır. Go planlaşdırıcı Go runtime-ın tərkib hissəsidir.  
Bu əməliyyat sisteminin verdiyi thread-lərə (parallel icra axınları) nisbətdə daha yüngüldür.  
Həmçinin goroutine-lər arası keçidlər də daha tez və sürətlidir.  



```go
func MyRoutine(söz string, times int) {
	for i:=0; i< times; i++ {
		fmt.Println(i, ")", söz)
		//fasilə
		time.Sleep(500 * time.Millisecond)
	}
   
}

func main() {
	var sözlər = [...] string {"Salam", "Baku", "Mars", "Yupiter"}
	for _, s := range sözlər {
		go MyRoutine(s, 2) 
	}
	fmt.Println("NumGoroutine", runtime.NumGoroutine())
	fmt.Println("GOMAXPROCS", runtime.GOMAXPROCS(-1))
	fmt.Println("NumCPU", runtime.NumCPU())
	time.Sleep(2000 * time.Millisecond)
}
```

    NumGoroutine 5
    GOMAXPROCS 2
    NumCPU 2
    0 ) Yupiter
    0 ) Salam
    0 ) Baku
    0 ) Mars
    1 ) Salam
    1 ) Mars
    1 ) Baku
    1 ) Yupiter


Qeyd edək ki, runtime.GOMAXPROCS(-1) eynianda icra edə biləcək maximum sayı göstərir.  
Göründüyü kimi bu adi halda sistemdə olan məntiqi nüvə sayına bərabərdir.  
Ümumilikdə qeyd edək ki, Go planlaşdırıcı icra edilən goroutine-lərin hamısı üçün vaxt ayırmasına şərait yaratmağa çalışacaq.  
Həmçinin Goroutine daxilində primitiv sinxronlaşma blokları, sistem, io çağırışları olduqda digər goroutine icrasına keçid ola bilər. Bu məsələlər adətən parallel məntiqi nüvə sayı az, goroutine sayı isə həddindən artıq olanda vacib olur.   
Bu mövzuya xüsusi baxış lazım olduğundan gələcəkdə dərindən nəzərdən keçirəcəyik.

### Goroutine-lərlə işləmə

Yuxarıdakı misallarda biz gözləmə ilə fasilə verirdik. 
Fasilə vermədiyimiz halda nə baş verəcəyini nəzərdən keçirək.


```go
func MyRoutine(söz string, times int) {
	for i:=0; i< times; i++ {
		fmt.Println(i, ")", söz) 
	} 
   
}

func main() {
	go MyRoutine("Salam", 7)
	go MyRoutine("Dünya", 3)
	MyRoutine("++++++", 2) 
	MyRoutine("++++++", 1) 
	MyRoutine("++++++", 2) 
}
```

    0 ) ++++++
    1 ) ++++++
    0 ) ++++++
    0 ) ++++++
    1 ) ++++++


Göründüyü kimi əsas (main) goroutine-i bitdikdən sonra proqram çıxır. 
Proqram digər goroutine-lərə vaxt ayırmağa macal tapmır.  
Aydın məsələdir ki, proqram çıxmamışdan əvvəl goroutine-lərin icrasının bitməsini gözləmək 
daha düzgün olardı.  

#### Channels (kanallar)

Channel dedikdə tipə sahib kanal nəzərdə tutulur. Adi halda bu kanal iki istiqamətlidir. Yəni ona həm yaza 
həmdə ondan oxuya bilərik. 
Channel elanı aşağıdakı kimi olur. 
```Go
// kanal := make(chan Tip)
// var kanal2 chan Tip //inisializasiya olunmadığından nil-dir
// kanal2 =  make(chan Tip)

kanal := make(chan string)
var kanal2 chan string
kanal2 = make(chan string)

```
Channel-ə yazmaq və oxumaq üçün <- operatorundan istifadə edirik.  
Oxun istiqaməti (channel dəyişəninə nisbətdə) ona yazı və ya ondan oxuma olduğunu göstərir
```Go
kanal <- v    // kanal Channel-inə v məlumatını göndər 
z := <- kanal  // kanal Channel-dən məlumatı oxu və z dəyişəninə mənimsət 

```

Adi halda Channel-ə oxu və yazı əməliyyatı digər tərəfin istifadə edəcəyi anadək bloka düşür.  
Bu səbəbdən Channel vasitəsilə biz goroutine-ləri nəinki bir birilə əlaqələndirə həm də sinxronlaşdıra bilərik.  



```go
%%
//Qeyd: sadə elan etdikdə kanal inisializasiya olunmur və nil-dir
var kanal2 chan string
fmt.Println(kanal2)
kanal2 = make(chan string)
fmt.Println(kanal2)

```

    <nil>
    0xc000082060



```go
%%
var kanal2 chan string
kanal2 <- "blok deadlock"
```

    fatal error: all goroutines are asleep - deadlock!
    
    goroutine 1 [chan send (nil chan)]:
    main.main()
    	/tmp/gonb_e784b3b9/main.go:11 +0x65
    exit status 2



Channel-i ```close``` builtin funksiyası ilə bağlamaq olur. Bu zaman yadda saxlamaq lazımdır ki, 
bağlanmış kanala yazmaq xəta ilə nəticələnir.  Adətən channel əgər for dövrəsi içində oxuyursa bağlama 
zamanı o dövrdən çıxacaq. Ümumən isə channel-i bağlamaq vacib deyil, sadəcə for range loop üçün əhəmiyyət kəsb edir. 
channel-in bağlandığını başqa cür isə belə yoxlamaq olar. 
```Go

v, işləkdir := <-kanal 

if !işləkdir {
    print("kanal bağlıdır" )
}

```


```go
func İcracıAsanXidmət(növbə chan string) { 
	for adam := range növbə {
        fmt.Println(adam)
		//sanki iş görürük
		time.Sleep(500 * time.Millisecond)
    }
   
}

func main() {
	kanal := make(chan string)
	go İcracıAsanXidmət(kanal)
	var t = time.Now().UnixMilli()
	kanal <- "Akif"
	kanal <- "Zakir"
	kanal <- "Samir"
	kanal <- "Sona"
	kanal <- "Vahid"
	kanal <- "Əli"
	kanal <- "Səid"
	t = time.Now().UnixMilli() - t
	close(kanal)
	print("göndərmələrə sərf olunan vaxt ", t, "ms")
	time.Sleep(100 * time.Millisecond)
}
```

    Akif
    Zakir
    Samir
    Sona
    Vahid
    Əli


    göndərmələrə sərf olunan vaxt 3003ms

    Səid


Qedy edək ki, kanallar həmçinin **bufferləşmiş** də ola bilər. Bu zaman kanala yazma əməliyyatları buffer tam dolduqdan sonra bloka düşür. 


```Go
kanal := make(chan string, 100)

```
Aşağıdakı nümunədə göndərmələrə sərf olunan vaxtı əvvəlki ilə müqayisə edin. 


```go
func main() {
	kanal := make(chan string, 100)
	go İcracıAsanXidmət(kanal)
	var t = time.Now().UnixMilli()
	kanal <- "Akif"
	kanal <- "Zakir"
	kanal <- "Samir"
	kanal <- "Sona"
	kanal <- "Vahid"
	kanal <- "Əli"
	kanal <- "Səid"
	t = time.Now().UnixMilli() - t
	close(kanal)
	print("göndərmələrə sərf olunan vaxt ", t, "ms")
	time.Sleep(5 * time.Second)
}
```

    göndərmələrə sərf olunan vaxt 0ms

    Akif
    Zakir
    Samir
    Sona
    Vahid
    Əli
    Səid


Yuxarıdakı misalda channel ikitərəfli olduğundan hər iki tərəf həm göndərə, həm də 
oxuya bilər. Bu çox zaman arzuolunmazdır. Bəzən biz istəyirik ki, bir tərəf yalnız oxuya bilsin. 
və bir tərəf yalnız yaza bilsin. Həmçinin, bu zaman müvafiq əməliyyatlar bir tərəfin məsuliyyətində olur.  
```Go

var adiKanal chan string // kanala hər kəs yaza, ondan oxuya və ya onu bağlaya bilər.
var oxuKanalı <-chan string // kanaldan oxuya bilər lakin  yaza bağlaya bilməz.
var yazıKanalı chan<- string  // kanala yaza və ya onu bağlaya bilər. Lakin ondan oxuya bilməz.

readOnly := make(<-chan string) // oxu
sendOnly := make(chan<- string) // yazı


```


```go
func İcracıAsanXidmət(növbəOxuKanalı <-chan string) { 
	for adam := range növbəOxuKanalı {
        fmt.Println(adam)
		//sanki iş görürük
		time.Sleep(500 * time.Millisecond)
    }
   
}

func main() {
	kanal := make(chan string)
	go İcracıAsanXidmət(kanal)
	var t = time.Now().UnixMilli()
	kanal <- "Akif"
	kanal <- "Zakir"
	kanal <- "Samir"
	kanal <- "Sona"
	kanal <- "Vahid"
	kanal <- "Əli"
	kanal <- "Səid"
	t = time.Now().UnixMilli() - t
	close(kanal)
	print("göndərmələrdə keçən vaxt ", t, "ms")
	time.Sleep(100 * time.Millisecond)
}
```

    Akif
    Zakir
    Samir
    Sona
    Vahid
    Əli
    Səid


    göndərmələrdə keçən vaxt 3004ms

Baxmayaraq ki, kanal channel-i ikitərəfli elan olunub, goroutine funksiyasına verilən arqument daxildə
yalnız birtərəfli oxu kanalı olur. channel reference olduğundan aydın məsələ hər ikisinin alt kanalı eynidir. 

```Go

var kanal chan string
var növbəOxuKanalı <-chan string

kanal = make(chan string)

növbəOxuKanalı = kanal //hər ikisi altda eyni kanalı saxlasa da növbəOxuKanalı 
                       //vasitəsilə yalnız oxuma əməliyyatları etmək olar

```

Kanal üzərində əməliyyatlar və nəticələri:
 

"Əməliyyat" | Kanalın vəziyyəti      | Nəticə
----------|--------------------|-------------
Oxumaq      | nil                | Bloka Deadlock-a düşür
_         |Açıq və boş deyil  | qiyməti oxuyur
_         | Açıq və boş| Bloka düşür
_         | Bağlı              | (default qiymət, false) qaytarır
_         | Yalnız Yazı kanalı        | Kompilyasiya xətası
Yazmaq     | nil                | Bloka Deadlock-a düşür
_         | Açıq və doludur      | Bloka düşür
_         | Açıq və dolu deyil  | Kanala yazır
_         | Bağlı              | panic
_         | Yalnız Oxu kanalı       | Kompilyasiya xətası
Bağlamaq (Close)     | nil                | panic 
_         | Açıq və boş deyil | kanalı bağlayır; oxuma kanal boşalana kimi uğurla davam edir, sonra default qiymətlə nəticələnir
_         | Açıq və boş     | kanalı bağlayır; oxuma default qiymətlə nəticələnir
_         | Bağlı              | panic

#### Select

Select vasitəsilə biz bir neçə kanallarda baş verən kommunikasiya əməliyyatlarını izləyə və gözləyə bilərik.  
Həmçinin, default halını da əlavə etsək, bu zaman default hissə digər əməliyyatlarda bloka düşmə olduğu halda işə düşəcək.  
 

```Go
select {
    case s <- v:
        fmt.Println("Göndər:", v)
    case vr := <-r:
        fmt.Println("Oxu:", vr) 
    default:
    // yuxarıdakılar blok olduğu halda icra et
}


```



<span style="color:red"> Diqqət: select bloku heçnəsiz və ya nil kanallarla icra edilsə tam bloka və deadlock-a düşür </span>



```go
%%
select {
	
}
```

    fatal error: all goroutines are asleep - deadlock!
    
    goroutine 1 [select (no cases)]:
    main.main()
    	/tmp/gonb_bb7e8f5b/main.go:21 +0x5a
    exit status 2


Select-ə aid nümunəyə baxaq: 


```go

func İcracıAsanXidmət(növbəOxuKanalı <-chan string) {
	for {
		select {
		case adam, ok := <-növbəOxuKanalı:
			if !ok {
				break
			}
			fmt.Println("~~~", adam, "~~~")
			time.Sleep(50 * time.Millisecond)
			fmt.Println("~~~", adam, "~~~", " işi bitib, yola salınır")
		default:
			fmt.Println("icraçı istirahet edir")
			time.Sleep(100 * time.Millisecond)

		}

	}

}

func main() {
	kanal := make(chan string)
	go İcracıAsanXidmət(kanal)
	var arr = [...]string{"Akif", "Zakir", "Samir"}
	var arr1 = [...]string{"Sona", "Vahid", "Əli"}
	go func() {
		for _, el := range arr {
			fmt.Println(el, " növbəyə durur")
			kanal <- el
			time.Sleep(100 * time.Millisecond)
		}
	}()

	for _, el := range arr1 {
		fmt.Println(el, " növbəyə durur")
		kanal <- el
		time.Sleep(100 * time.Millisecond)
	}
	close(kanal)
	time.Sleep(1000 * time.Millisecond)
}

```

    Sona  növbəyə durur
    Akif  növbəyə durur
    ~~~ Sona ~~~
    ~~~ Sona ~~~  işi bitib, yola salınır
    ~~~ Akif ~~~
    Vahid  növbəyə durur
    ~~~ Akif ~~~  işi bitib, yola salınır
    ~~~ Vahid ~~~
    Zakir  növbəyə durur
    ~~~ Vahid ~~~  işi bitib, yola salınır
    ~~~ Zakir ~~~
    ~~~ Zakir ~~~  işi bitib, yola salınır
    icraçı istirahet edir
    Əli  növbəyə durur
    Samir  növbəyə durur
    ~~~ Əli ~~~
    ~~~ Əli ~~~  işi bitib, yola salınır
    ~~~ Samir ~~~
    ~~~ Samir ~~~  işi bitib, yola salınır


Ümumən Select və channel-dən ehtiyatla istifadə etmək lazımdır.  

Qeyd edək ki, biz əlavə channel-dən goroutine-in bitmə siqnalı kimi istifadə edə və main-dən o halda çıxa bilərik.  
Lakin bunun üçün sync primitivlərdən WaitGroup istifadə etmək olar.

#### sync.WaitGroup (gözləmə qrupu) 
sync paketində olan WaitGroup la tanış olaq. 

Biz indiyə kimi gözləmə kimi fasilə istifadə edirdik (time.Sleep).  
İndi isə daha düzgün üsulla gözləməni həyata keçirək.

 | metod | izah |
 |---|---|
|Add(n int) |WaitGroup sayını n sayda artırır
|Done()	 |WaitGroup sayını 1 vahid azaldır, Bunu goroutine-in bitdiyini göstərmək üçün istifadə edəcəyik
|Wait()	| WaitGroup sayı 0 olana qədər bloka düşür.


```go
import "sync"

func MyRoutine(söz string, times int, wg *sync.WaitGroup) {
	if wg!=nil {
		defer wg.Done()
	}
	for i:=0; i< times; i++ {
		println(i, ") ", söz)
		time.Sleep(70 * time.Millisecond)
	} 
   
}

func main() {
	wg := new(sync.WaitGroup)
    wg.Add(2)
	go MyRoutine("Salam", 4, wg)
	go MyRoutine("Dünya", 4, wg)
	MyRoutine("++++++", 3, nil)  

	wg.Wait()
}
```

    0 )  ++++++
    0 )  Dünya
    0 )  Salam
    1 )  Salam
    1 )  ++++++
    1 )  Dünya
    2 )  Dünya
    2 )  ++++++
    2 )  Salam
    3 )  Salam
    3 )  Dünya


sync paketində digər sinxronlaşdırma primitivləri ilə daha sonra tanış olacağıq.  
Belə ki, goroutine-lər eyni yaddaş fəzasında icra edildiyindən dəyişənlərə müraciət sinxronlaşdırılmalıdır.  
əks halda data race (məlumatı dəyişmək üzrə yarış) hadisəsi baş verir və dəyişənin qiyməti arzuolunmaz pozğunluqla nəticələnə bilər. 
