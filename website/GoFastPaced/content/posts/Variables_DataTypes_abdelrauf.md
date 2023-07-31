---
title: " Variables Data Types "
author: "abdelrauf"
date: 2023-07-31
tags: ["Variables", "Data", "Types", "Go" ]
---

### Variables, Types and

### Əsas Primitiv Tiplər. Yaddaşda ölçüsünə və ayrılmasına və yerinə təsir edir.

| Tip  | İzah | Qiymət nümunə |
| --- | --- | --- |
| int | tam ədədlər | 2, 0, -999|
|float32 | kəsr ədədlər | 3.5545, -9.5|
|float64 | kəsr ədədlər daha dəqiq hesablamalar üçün 8 baytliq  | 16000001.5, -9.5|
|complex64| kompleks ədədlər | 5+4i, 5, i|
|complex128| kompleks ədədlər | 5+4i, 5, i|
|string| dırnaq işarəsi daxilində simvol ardıcıllığı| "salam", "-sds-"|
|bool | dogru ve ya yanlis, bul mentiqi| true, false |
| byte (uint8 alias))| müsbət 8bit lik , 1baytliq ədəd| 2, 5, [0,255] arası ədəd|
|rune (int32 alias)| istenilen simvollar , unikod smilecs ve s| 'a','z', '👍' |
|int8| tam ədəd 1 baytliq 8 bitlik| -128 127 [-128,127]|
|uint8| 8bit lik müsbət ədəd | 0 255 [0,255] aralığında|
|int16, int32, int64| 16,32,64 bitlik||
|uint16,uint32,uint64| 16,32,64 bitlik müsbət| 0-dan 2^bit_sayi-1 aralığında ədədlər|

## Dəyişənlər  (Variable)

#### **Dəyişən adları üçün qayda**
- Dəyişən adı a-z A-Z 0-9 ve _  və hərf olan unikod simvolları ola bilər. Yəni digərləri boşluq və s simvollar ola bilməz.  
  Unikod hərflərə icazə olduğundan **dəyişənləri azərbaycanca da** yazmaq olar
- Dəyişən adı hərflə və ya _ altdan-xəttlə başlaya bilər. ədədlə yox
- Dəyişən adlarında böyük kiçik hərfləri fərqlilik yaradır. Yəni ( uzunluq Uzunluq uZunluq ) dəyişənləri fərqlidir
- Dəyişən adlarında Golang dilinin özəl sözlərini istifadə etmək olmaz. Məsələn: for

##### Ad yazılımında standardlar
- **Camel case**. Nümunə: mənimDəyişəninAdı, fiqurunUzunluğu
- **Pascal Case**. Nümunə: MənimDəyişəninAdı, FiqurunUzunluğu
- **Snake case**. Nümunə: mənim_dəyişənin_adı, fiqurun_uzunluğu
 
#### Dəyişən elanı

<pre>
    Qeyd: [ ] işarəsi daxilində olanlar buraxıla bilinər.
    Qeyd: Tip buraxılan elanlarda Go ən uyğun tipi özü müəyyən edir ( type inference)

    var ad [tip]   
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
  metn := "Sadə yazı"
  fmt.Println(metn, uzunluq, adamların_sayı)

}
```

    Sadə yazı 3.5 40



```go
%%
  var uzunluq,adamların_sayı,mətn   = 3.5, 40, "Sadə yazı"
  fmt.Println(mətn, uzunluq, adamların_sayı)
```

    Sadə yazı 3.5 40



```go
%%
  uzunluq,adamların_sayı,mətn   := 3.5, 40, "Sadə yazı"
  fmt.Println(uzunluq,adamların_sayı,mətn)
```

    3.5 40 Sadə yazı



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
|dəyişən elanı və ilkin qiymət ayrı ayrı qeyd etmək olur| elan və ilkin qiymət eyni sətirdə olur|

###  '=' Mənimsətmə əməliyyatı ilə dəyişənin qiymətini sonradan da dəyişə bilirik  
(*Proqramçı Qeydi: = riyazi bərabərliklə eyni anlam daşıması üçün dəyişən immutable (dəyişilməsi icazə verilməyən) olmalıdır, əks halda bu əməliyyata bərabərdir demək yanlışdır*)


```go
%%
   uzunluq,adamların_sayı,mətn   := 3.5, 40, "Sadə yazı"
  uzunluq = 99.5
  fmt.Println(mətn, uzunluq, adamların_sayı)
```

    Sadə yazı 99.5 40


Go dilində əgər dəyişən elan edilibsə mütləq istifadə edilməlidir


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


<span style="white-space: pre;"> # gonb_e6ad3a12</span>

<br/>

<span class="gonb-cell-line-info">Cell[66]: Line 2</span>
<span class="gonb-error-location">./main.go:22:2: </span> a declared and not used
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




    failed to run "/usr/local/go/bin/go build -o /tmp/gonb_e6ad3a12/gonb_e6ad3a12": exit status 1


Əgər dəyişəni istifadə etməyəcəyiksə _ adı verə bilərik. Qeyd: bu zaman iqnor etdiyimizdən əlavə dəyişən olmalıdır


```go
%%
a,_ := 77, 5
fmt.Print(a)
```

    77

#### Konstant Sabit , Proqramda sabit deyishilmeyen qiymetler uchun istifade edirik

<pre>
    const ad = qiymət  
    const ad [tip] = qiymət

    const(
        ad = qiymət1
        ad2 = qiymət2
    )
</pre>

**Konstant** sonradan dəyişmək olmaz


```go
%%
const pi = 3.14
const (
	paytaxt = "Bakı"
	ölkə = "Azərbaycan"
)
fmt.Println("Pi ededi ", pi, paytaxt, ölkə)
```

    Pi ededi  3.14 Bakı Azərbaycan


### Bir tipdən digər (çevrilə bilən) tipə aşkar keçirmək üsulu

**tip( deyishen_adi)**


```go
%%
uzunluq:=3.5
var tam_uzunluq int = int(uzunluq)
fmt.Println(tam_uzunluq)
```

    3


<h4 style="color:red"><b>var  dəyişən elanında ilkin qiymətlənmə ardıcıllığı.</b></h4>
ilkin qiymətləndirmədə  əgər bir dəyişən digər dəyişəndən asılı olduqda ardıcıllıq dəyişən elanına görə yox <b>asılıllığa görə olur</b>


```go
%%
var a int = c
var c int = 9
fmt.Println(a,c)

```

    5 9



```go
%%
var (
	a int = c
	b int = d
	d int = 6
	c int = 9
)

fmt.Println(a,b,c,d)
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

%%
fmt.Println(a,b,c,d)

```

    9 4 5 5

