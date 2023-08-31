---
title: " Variables Data Types Constants"
author: "abdelrauf"
date: 2023-07-29
tags: ["Variables", "Data", "Types", "Constants", "Go" ]
---

### Variables, Types and

### Əsas Primitiv Tiplər. Yaddaşda ölçüsünə və ayrılmasına və yerinə təsir edir.

| Tip  | İzah | Qiymət nümunə |
| --- | --- | --- |
| int | tam ədədlər | 2, 0, -999|
|float32 | kəsr ədədlər | 3.5545, -9.5|
|float64 | kəsr ədədlər, daha dəqiq hesablamalar üçün 8-baytlıq  | 16000001.5, -9.5|
|complex64| kompleks ədədlər | 5+4i, 5, i|
|complex128| kompleks ədədlər | 5+4i, 5, i|
|string| dırnaq işarəsi " və ya \` daxilində simvol ardıcıllığı| "salam", "-sds-", \`bir neçə sətir\`|
|bool | doğru və ya yanlış, bul məntiqi| true, false |
| byte (uint8 alias))| müsbət 8-bitlik , 1-baytlıq ədəd| 2, 5, [0,255] arası ədəd|
|rune (int32 alias)| istənilən simvollar, unikod və s| 'a','z', '👍' |
|int8| tam ədəd 1-baytlıq 8-bitlik| -128 127 [-128,127]|
|uint8| 8-bitlik müsbət ədəd | 0 255 [0,255] aralığında|
|int16, int32, int64| 16,32,64 bitlik tam ədəd||
|uint16,uint32,uint64| 16,32,64 bitlik müsbət| 0-dan 2^bit_sayı-1 aralığında ədədlər|

## Dəyişənlər  (Variable)

#### **Dəyişən adları üçün qayda**
- Dəyişən adı a-z A-Z 0-9 ve _  və hərf olan unikod simvolları ola bilər. Yəni, digərləri, boşluq və s simvollar ola bilməz.  
  Unikod hərflərə icazə olduğundan **dəyişənləri azərbaycanca da** yazmaq olar
- Dəyişən adı hərflə və ya _ altdan-xəttlə başlaya bilər. ədədlə yox
- Dəyişən adlarında böyük kiçik hərflər fərqlilik yaradır. Yəni, ( uzunluq Uzunluq uZunluq ) dəyişənləri fərqlidir
- Dəyişən adlarında Golang dilinin özəl sözlərini istifadə etmək olmaz. Məsələn: for

##### Ad yazılımında standardlar
- **Camel case**. Nümunə: mənimDəyişəninAdı, fiqurunUzunluğu
- **Pascal Case**. Nümunə: MənimDəyişəninAdı, FiqurunUzunluğu
- **Snake case**. Nümunə: mənim_dəyişənin_adı, fiqurun_uzunluğu
 
#### Dəyişən elanı

<pre>
    Qeyd: [ ] işarəsi daxilində olanlar buraxıla bilinər.
    Qeyd: Tip buraxılan elanlarda Go ən uyğun tipi özü müəyyən edir ( type inference)

    var ad tip   
    var ad [tip] = qiymet    
    var ad = qiymet  
    ad := qiymet

    var ad1, ad2 ,.. = ad1_qiymət, ad2_qiymət, ..  
    ad1, ad2 ,.. := ad1_qiymət, ad2_qiymət, ..                                      

    var (
        ad1 [tip1]
        ad2 [tip2] = ad2_qiymət
        ad3 [tip3] = ad3_qiymət  
        ...
    )

    //eyni tipli bir neçəsi
    var ad1, ad2 ,.. [eyni_tip] = ad1_qiymət, ad2_qiymət, ..
    Nümunə: var a,b,c int = 4,5,7

</pre>



```go
import "fmt"

func main(){
  var uzunluq float32 = 3.5
  var adamların_sayı = 40
  mətn := "Sadə yazı"
  fmt.Println(mətn, uzunluq, adamların_sayı)

}
```

    Sadə yazı 3.5 40



```go
%%
  var uzunluq, adamların_sayı, mətn = 3.5, 40, "Sadə yazı"
  fmt.Println(mətn, uzunluq, adamların_sayı)
```

    Sadə yazı 3.5 40



```go
%%
  uzunluq, adamların_sayı, mətn := 3.5, 40, "Sadə yazı"
  fmt.Println(mətn, uzunluq, adamların_sayı)
```

    Sadə yazı 3.5 40



```go
%%
var (
	uzunluq float32 = 3.5
	adamların_sayı = 40
	mətn string = "Sadə yazı" 
)
fmt.Println(mətn, uzunluq, adamların_sayı)
```

    Sadə yazı 3.5 40


#### var və := arasında fərqlər
| var | := |
| --- |---|
| funksiya daxilində və xaricində istifadə etmək olur | Yalnız funksiya daxilində istifadə etmək olur|
|dəyişən elanı və ilkin qiymət verməni ayrı-ayrı qeyd etmək olur| elan və ilkin qiymət eyni sətirdə olur|

###  '=' Mənimsətmə əməliyyatı ilə dəyişənin qiymətini sonradan da dəyişə bilirik  
(*Proqramçı Qeydi: = riyazi bərabərliklə eyni anlam daşıması üçün dəyişən const və ya readonly immutable (dəyişilməsi icazə verilməyən) olmalıdır, əks halda bu əməliyyata bərabərdir demək yanlışdır*)


```go
%%
  uzunluq, adamların_sayı, mətn := 3.5, 40, "Sadə yazı"
  uzunluq = 99.5
  fmt.Println(mətn, uzunluq, adamların_sayı)
```

    Sadə yazı 99.5 40


Go dilində əgər dəyişən elan edilibsə mütləq istifadə edilməlidir.


```go
%%
a:=8
```



<style>
.gonb-error-location {
	background: var(--jp-error-color2);  
	border-radius: 3px;
	border-style: dotted;
	border-width: 1px;
	border-color: var(--jp-border-color2);
}
.gonb-error-location:hover {
	border-width: 2px;
	border-style: solid;
	border-color: var(--jp-border-color2);
}
.gonb-error-context {
	display: none;
}
.gonb-error-location:hover + .gonb-error-context {
	background: var(--jp-dialog-background);  
	border-radius: 3px;
	border-style: solid;
	border-width: 1px;
	border-color: var(--jp-border-color2);
	display: block;
	white-space: pre;
	font-family: monospace;
}
.gonb-error-line {
	border-radius: 3px;
	border-style: dotted;
	border-width: 1px;	
	border-color: var(--jp-border-color2);
	background-color: var(--jp-rendermime-error-background);
	font-weight: bold;
}
.gonb-cell-line-info {
	background: var(--jp-layout-color2);
	color: #999;
	margin: 0.1em;
	border: 1px solid var(--jp-border-color1);
	padding-left: 0.2em;
	padding-right: 0.2em;
}
</style>
<div class="lm-Widget p-Widget lm-Panel p-Panel jp-OutputArea-child">
<div class="lm-Widget p-Widget jp-RenderedText jp-mod-trusted jp-OutputArea-output" data-mime-type="application/vnd.jupyter.stderr" style="font-family: monospace;">


<span style="white-space: pre;"> # gonb_a54331cb</span>

<br/>

<span class="gonb-cell-line-info">Cell[76]: Line 2</span>
<span class="gonb-error-location">./main.go:58:2: </span> a declared and not used
<div class="gonb-error-context">

func main() {
	flag.Parse()
<div class="gonb-error-line">	a:=8
</div>
}

</div>

<br/>


<span style="white-space: pre;"> </span>

<br/>

</div>




    failed to run "/usr/local/go/bin/go build -o /tmp/gonb_a54331cb/gonb_a54331cb": exit status 1


Əgər dəyişəni istifadə etməyəcəyiksə _ adı verə bilərik. Qeyd: bu zaman istifadə olunmayan dəyişəndən əlavə dəyişən də olmalıdır.


```go
%%
a, _ := 77, 5
fmt.Print(a)
```

    77

#### Konstant və ya Sabit
Proqramda icra zamanı sabit dəyişilməyən qiymətlərin istifadəsi üçün işlədirik.  


<pre>
    const ad = qiymət  
    const ad [tip] = qiymət

    const(
        ad = qiymət1
        ad1   // blok içində dəyərsiz vermək olur. ya əvvəlkinin qiymətini alacaq, ya da iota varsa sıra dəyəri alacaq.
        ad2 = qiymət2
    )
</pre>

**Konstant** sonradan dəyişmək olmaz


```go
%%
const pi = 3.14159265359
const (
	paytaxt = "Bakı"
	ölkə = "Azərbaycan"
)
fmt.Println("Pi ədədi", pi, paytaxt, ölkə)
```

    Pi ədədi 3.14159265359 Bakı Azərbaycan


**iota** konstant generatoru. Bəzən biz ardıcıl əlaqəli konstant elan etmək istəyirik.   
Bunun üçün biz indeks ədəd (>=0 müsbət tam) verən iota identifikatorundan istifadə edə bilərik.   
iota konstant bloku daxilində 0 dəyəri alır və ondan ardıcıllıq vermək üçün istifadə edirik.   
İstifadə nümunələrinə baxaq:




```go
const (
	Bazar  = iota  // 0
	BazarErtəsi    // 1
	ÇərşənbəAxşamı // 2
	Çərşənbə       // 3
	CüməAxşamı     // 4
	Cümə           // 5
	Şənbə          // 6
)

%%
fmt.Println(Bazar, BazarErtəsi, ÇərşənbəAxşamı, Şənbə)
```

    0 1 2 6


blok daxilində bir neçə iota verdiyimizdə olan davranışa baxsaq görərik ki, ardıcıllıq **sıraya** əsasən təyin edilir.


```go
const (
	heçnə = 55
	təkrar    // blok içində heçnə vermədiyimizdən əvvəlkinin dəyərini götürüb
	iki  = iota  // 2
	_    // 3 // _ işarəsi lazımsız deməkdir. və bir neçə dəfə istifadə edə bilirik
	dörd // 4
	_       // 5
	_     // 6
	yeddi  = iota         // 7
	_  //8
	doqquz          // 9
)

%%
fmt.Println(heçnə, təkrar,  iki , dörd, yeddi, doqquz)
```

    55 55 2 4 7 9


iota ilə birgə riyazi əməllər edib fərqli konstant də vermək olar. sadəlik üçün toplama istifadə edək.


```go
const (
	g = iota + 99  //0 + 99
	t = iota + 1   // 1 + 1
)
%%
fmt.Println(g, t)
```

    99 2


blok içində olmayan konstantlar üçün iota elə ilkin sıra dəyərini alır. yəni 0


```go
const w = iota  // w == 0
const x = 33
const y = iota  // y == 0
const z = x
%%
fmt.Println(w, x, y, z)
```

    0 33 0 33


### Bir tipdən digər (çevrilə bilən) tipə aşkar keçirmək üsulu

<pre>
tip( dəyişənin_adı)
</pre>


```go
%%
uzunluq := 3.5
var tam_uzunluq int = int(uzunluq)
fmt.Println(tam_uzunluq)
```

    3


<h4 style="color:red"><b>Package blokunda (package-level) var dəyişən elanında ilkin qiymətlənmə ardıcıllığı.</b></h4>
<p>ilkin qiymətləndirmədə  əgər bir dəyişən digər dəyişəndən asılı olduqda </p> 
<p>ardıcıllıq dəyişən elanına görə yox <b>asılıllığa görə olur</b>.</p>   
<p>Burada söhbət <b>fayl içində, funksiya xaricində</b> verilən var elanından gedir.</p>


```go

var (
	a int = c
	b int = d
	d int = 6
	c int = 9
)

func main() {
  fmt.Println(a, b, c, d)
}


```

    9 6 9 6



```go
var (
	a = c + b  // == 9
	b = f()    // == 4
	c = f()    // == 5
	d = 3      // == 5  
)

func f() int {
	d++
	return d
}

func main() {
	fmt.Println(a, b, c, d)
}

```

    9 4 5 5

