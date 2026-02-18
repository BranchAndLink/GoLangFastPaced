---
title: "Goroutine Channels Select Wait Group"
weight: 1
tags: ["Goroutine", "Channels", "Select", "Wait", "Group", "Go"]
---

### Goroutine

Goroutine — eyni zaman daxilində bir neçə funksiyanı bölüşdürülmüş şəkildə(concurrent) icra etmək üçündür.

Qeyd: Eyni vaxtda dedikdə hər ikisinin eyni anda tam paralel icrası nəzərdə tutulmur. Bu baş verə bilər, baş verməyə də bilər. Əsasən nəzərdə tutulur ki, bu funksiyaların icrası üçün ayrılan prosess bölüşdürülür. Bu bölüşdürmə isə nizamlayıcı və əməliyyat sisteminin öz nizamlayıcısı ilə baş verir. Bütün proqramlaşdırma məsələlərini real həyatdan nümunə çəkərək izah edə bilərik. Məsələn: siz yeməyi su ilə birlikdə yeyərkən, birinci yeməyi bitirib sonra su içə bilərsiniz və ya əksinə. Eyni vaxtlı dedikdə isə nəzərdə tutulur ki, bir neçə qaşıq yeməkdən yeyir, bir neçə udum sudan içirsiniz. Və bunun necə və hansı vaxt bölgüsündə baş verəcəyini nizamlayan isə özünüzsünüz.



Goroutine-lərdən istifadə etmək üçün funksiya çağırışlarının əvvəlinə `go` sözünü əlavə edirik.  
`defer` kimi, bunda da funksiyaya verilən arqumentlərin qiymətləndirilməsi çağırış anında baş verir.  
Funksiyanın icrası isə yeni goroutine-də həyata keçirilir.
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
    1 )  Dünya
    1 )  ++++++
    1 )  Salam
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
Go planlaşdırıcı çoxlu sayda goroutinlərin icrasını nizamlayır. Go planlaşdırıcı, Go runtime tərkib hissəsidir.
Bu, əməliyyat sisteminin verdiyi thread-lərə (concurrent icra axınlarına) nisbətdə daha yüngüldür.
Həmçinin goroutinlər arası keçidlər də daha tez və sürətlidir.




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
    1 ) Mars
    1 ) Salam
    1 ) Yupiter
    1 ) Baku


Qeyd edək ki, runtime.GOMAXPROCS(-1) eyni anda (parallel) icra edə biləcək maksimum məntiqi CPU nüvələrini göstərir.
Göründüyü kimi, bu adi halda sistemdə olan nüvə sayına bərabərdir.
Ümumilikdə qeyd edək ki, Go planlaşdırıcı icra edilən goroutinlərin hamısı üçün vaxt ayırmasına şərait yaratmağa çalışacaq.
Həmçinin goroutine daxilində primitiv sinxronlaşma blokları, sistem və IO çağırışları olduqda digər goroutine icrasına keçid ola bilər.
Bu məsələlər adətən tam paralel məntiqi nüvə sayı 1 və ya az olduqda, goroutine sayı isə həddindən artıq olanda vacib olur.
Bu mövzuya xüsusi baxış lazım olduğundan gələcəkdə dərindən nəzərdən keçirəcəyik.


### Goroutine-lərlə işləmə

Yuxarıdakı misallarda biz gözləmə ilə fasilə verirdik. 
Fasilə vermədiyimiz halda nə baş verəcəyini nəzərdən keçirək


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


Fikir verdiksə, əsas main goroutine-i bitdikdən sonra proqram çıxır.
Proqram digər goroutinlərə vaxt ayırmağa macal tapmır.
Aydın məsələdir ki, main çıxmamışdan əvvəl goroutinlərin icrasının bitməsini də gözləmək daha düzgün olardı.
Bundan başqa, goroutinlər bir-birilə necə əlaqə saxlaya və eyni dəyişənə problemsiz müraciət edə bilərlər — bu barədə də qısa danışaq.
Belə ki, goroutine-də icra edilən funksiyalar eyni yaddaş fəzasını əhatə edir.
Ona görə ortaq dəyişənlərin dəyişdirilməsi məlumat pozğunluğuna səbəb ola bilər (data race — məlumat dəyişimi üçün yarış).
O səbəbdən həmin məlumatın dəyişdirilməsi sinxronlaşdırılmalıdır.


#### Channels (kanallar)

Channel dedikdə tipə sahib kanal nəzərdə tutulur. Adi halda bu kanal iki istiqamətlidir. Yəni ona həm yaza,
həm də ondan oxuya bilərik.
Channel elanı aşağıdakı kimi olur:

// kanal := make(chan Tip)
// var kanal2 chan Tip // inisializasiya olunmadığından nil-dir
// kanal2 = make(chan Tip)

kanal := make(chan string)
var kanal2 chan string
kanal2 = make(chan string)

Channel-ə yazmaq və oxumaq üçün <- operatorundan istifadə edirik.
Oxun istiqaməti yazı və ya oxuma olduğunu göstərir:

kanal <- v     // kanal-a v məlumatını göndər
z := <- kanal  // kanal-dan məlumatı oxu və z dəyişəninə mənimsət

Adi halda Channel-ə oxu və yazı əməliyyatı digər tərəfin istifadə edəcəyi anadək blok olur.
Bu səbəbdən Channel vasitəsilə biz goroutinləri nəinki bir-birilə əlaqələndirə, həm də sinxronlaşdıra bilərik.



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
bağlanmış kanala göndərmək xəta ilə nəticələnir. Adətən channel əgər for dövrəsi içində oxuyursa,
bağlama zamanı o dövrdən çıxacaq. Ümumən isə channel-i bağlama vacib deyil, sadəcə for range loop üçün əhəmiyyət kəsb edir. Channel-in bağlandığını başqa cür isə belə yoxlamaq olar.

```Go
v, işləkdir := <-kanal

if !işləkdir {
    print("kanal bağlıdır")
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


Qeyd edək ki, kanallar həmçinin **buferləşmiş**(array) də ola bilər. Bu zaman kanala yazmalar bufer tam dolanda bloklanır.

```Go
kanal := make(chan string, 100)
```

Gəlin eyni misalda göndərmələrdə bloklanma olmadığını və daha az vaxt aldığını yoxlayaq.


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
	print("göndərmələrdə keçən vaxt ", t, "ms")
	time.Sleep(5 * time.Second)
}
```

    göndərmələrdə keçən vaxt 0ms

    Akif
    Zakir
    Samir
    Sona
    Vahid
    Əli
    Səid


Yuxarıdakı misalda channel ikitərəfli olduğundan hər iki tərəf həm göndərə, həm də oxuya bilər. Bu çox zaman arzuolunmazdır. Bəzən biz istəyirik ki, bir tərəf yalnız oxuya bilsin və bir tərəf yalnız yaza bilsin. Həmçinin bu halda ona göndərmələr və kanalın bağlanması da bir tərəfin məsuliyyətində olur.

```Go
var adiKanal chan string       // kanala hər kəs yaza, ondan oxuya və ya onu bağlaya bilər
var oxuKanalı <-chan string    // kanaldan oxuya bilər, lakin yaza və bağlaya bilməz
var yazıKanalı chan<- string   // kanala yaza və ya onu bağlaya bilər, lakin ondan oxuya bilməz

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

Qeyd edək ki, baxmayaraq ki kanal ikitərəfli elan olunub, goroutine funksiyasına verilən arqument daxildə yalnız birtərəfli oxu kanalı olur. Channel reference olduğundan aydın məsələdir ki, hər ikisinin alt kanalı eynidir.

```Go

var kanal chan string
var növbəOxuKanalı <-chan string

kanal = make(chan string)

növbəOxuKanalı = kanal //hər ikisi altda eyni kanalı saxlasa da növbəOxuKanalı yalnız oxuya bilər

```

Kanal üzərində əməliyyatlar və nəticələri:
 

"Əməliyyat" | Kanalın vəziyyəti      | Nəticə
----------|--------------------|-------------
Oxumaq      | nil                | Bloka Deadlock-a düşür
_         |Açıq və boş deyil  | qiyməti oxuyur
_         | Açıq və boş| Bloka düşür
_         | Bağlı              | default qiymət, false qaytarır
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

Select vasitəsilə biz bir neçə kanallarda baş verən kommunikasiya əməliyyatlarını izləyə və gözləyə bilərik. Həmçinin default halını da əlavə etsək, bu zaman default hissə digər əməliyyatlarda bloka düşmə olduğu halda işə düşəcək. 

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



<span style="color:red"> Diqqət edin ki select bloku heçnəsiz və ya nil kanallarla icra edilsə tam bloka və deadlock-a düşür </span>



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


```go

```
