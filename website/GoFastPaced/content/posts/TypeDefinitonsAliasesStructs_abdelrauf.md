---
title: " Type Definitons Aliases Structs"
author: "abdelrauf"

tags: ["Type", "Definitons", "Aliases", "Structs", "Go" ]
---

### Tip elanı  
Tip elanı üçün type açar sözündən istifadə edirik. Bu zaman iki cür elan mümkündür. 

#### Tip sinonimi (Type alias)
Birinci sinonim tip vermək. Buna type alias deyilir. Bu zaman əslən yeni tip elan edilmir sadəcə sinonim kimi istifadə edilir
```go 

type  Sinonim_Ad = Başqa_tip
type (
    Sinonim_Ad1 = Başqa_tip1
    Sinonim_Ad2 = Başqa_tip2
)

```


```go
type TamƏdəd = int

type (
	TamƏdədMassivi = [5]int
	Mətn = string
)

%%
var x TamƏdəd = 99
var (
	z TamƏdədMassivi
	s Mətn = "simvollar"
)
fmt.Printf("%T %v\n", x, x )
fmt.Printf("%T %v\n", z, z )
fmt.Printf("%T %v\n", s, s )
```

    int 99
    [5]int [0 0 0 0 0]
    string simvollar


#### Tip təyini (Type definition)  
Bu zaman sinonimdən fərqli olaraq eynicinli olsa da yeni tip yaranır.
```go

type Yeni_tip  Alt_tip
type (
    Yeni_tip1  Alt_tip1
    Yeni_tip2  Alt_tip2
)

```


```go
type TamƏdəd  int

type (
	TamƏdədMassivi  [5]int
	Mətn  string
)

%%
var x TamƏdəd = 99
var (
	z TamƏdədMassivi
	s Mətn = "simvollar"
)
fmt.Printf("%T %v\n", x, x )
fmt.Printf("%T %v\n", z, z )
fmt.Printf("%T %v\n", s, s )
```

    main.TamƏdəd 99
    main.TamƏdədMassivi [0 0 0 0 0]
    main.Mətn simvollar


#### Struct  
Primitiv tiplərdən savayı biz bir neçə tiplərin birləşməsi olan ümumi tip də yarada bilərik.   
Bunun üçün struct açar sözündən istifadə edə bilərik.   
Həmin struktun eleməntləri üzv elementlər olur. (field names və ya member variables).   
Bunə həmçinin deyə bilərik ki struct bu tiplərdən kompozisiya (qurulmuş) olan tipdir. 
```go
//boş strukt
struct {}

struct {
    üzv1 Tip1 `teg`
    üzv2 Tip2
    _ Tip3 // boşluq yaratmaq üçün (padding)
    üzv3 Tip3
}

```
Struct tipində üzv (_ boşdan savayı) adları fərqli olmalıdır. 
Bundan savayı GO dilində bir strukt digər struktu öz daxilinə aşkar ad vermədən də ala bilir. Buna embedded field deyilir.  
Və o elə Tipin_adı və ya *Tipin_adı kimi əlavə edilir.  
Gəlin struct-u type -la yeni tip kimi elan edək və onun embedded və aşkar şəkildə kompozisiyasına baxaq.  
```go

type Esas struct {
    p int
}

//gizli kompozisiya T
type EmbeddedKompozisiya struct {
    Esas
    // p adlı dəyişən versək Esas-ın p-i kölgəyə salacaq
    t int
}

//gizli kompozisiya *T kimi
type EmbeddedKompozisiya2 struct {
    *Esas
    // p adlı dəyişən versək Esas-ın p-i kölgəyə salacaq
    t int
}

//aşkar
type EmbeddedKompozisiya struct {
    esas Esas
    p int
}


```
Strukt tipi üçün elan dəyişən elanı kimidir. Sadəcə ona dəyər verməsək üzvlərin ilkin dəyərini alacaq (0, nil). 
Qiymət vermək üçün isə bu cür istifadə edə bilərik 
```go
var dəyişən_adı = strukt_tipin_adı{qiymət1,..}

var dəyişən_adı = strukt_tipin_adı{ Üzv_Adı : qiymət1,..}

//Əgər embeddedirsə

var dəyişən_adı = strukt_tipin_adı{ Embeded_Tip_adı: Embeded_Tip_adı{qiymət1, ...},  Üzv_Adı : qiymət1,..}
var dəyişən_adı = strukt_tipin_adı{ Embeded_Tip_adı: Embeded_Tip_adı{embeded_üzv: qiymət1, ..},  Üzv_Adı : qiymət1,..}
// Embeded_Tip struct olduğundan dəyər vermə struct-a verdiyimiz kimidir.
```  

Struktun üzv elementlərini sonradan mənimsətmək üçün . dən istifadə edirik 
```go

var dəyişən_adı strukt_tipin_adı

dəyişən_adı.üzv1 = qiymət1
dəyişən_adı.üzv2 = qiymət2

//Əgər embedded olmuş struktdursa onun adını yazmaqla vermək olar( kölgəyə düşmürsə)
var dəyişən_adı1 embeddedli_strukt_tipin_adı

dəyişən_adı1.embedded_üzv = qiymət
```

Go dilində inheritance(varislik) yoxdur. Lakin onunla eynigüclü və alternativ olan kompozisiya və embedded kompozisiya var. 



```go
type Adam struct {
	ad string
	yaş int
	boy float32
}

%%
var Akif, Vəli Adam = Adam{"Akif", 22, 1.78}, Adam{ad: "Vəli", yaş: 22, boy: 1.78}

fmt.Println(Akif, Vəli)
fmt.Println("2 il sonra")
Akif.yaş = Akif.yaş + 2
Vəli.yaş = Vəli.yaş + 2
fmt.Println(Akif, Vəli)
```

    {Akif 22 1.78} {Vəli 22 1.78}
    2 il sonra
    {Akif 24 1.78} {Vəli 24 1.78}



```go
type Tələbə struct {
	Adam
	orta_qiymət float32
}

type TələbəTip struct {
	adam Adam
	orta_qiymət float32
}

%%
var Akif, Vəli Tələbə = Tələbə{Adam{"Akif", 22, 1.78}, 4.5}, Tələbə{Adam: Adam{ad: "Vəli", yaş: 22, boy: 1.78}, orta_qiymət:4.89}
fmt.Println(Akif, Vəli, Akif.ad)

var a, b TələbəTip = TələbəTip{Adam{"Akif", 22, 1.78}, 4.5}, TələbəTip{adam: Adam{ad: "Vəli", yaş: 22, boy: 1.78}, orta_qiymət:4.89}
//fikir versəniz embedded olmayanda .üzv_adı. ilə daxil oluruq
fmt.Println(a, b, a.adam.ad)
```

    {{Akif 22 1.78} 4.5} {{Vəli 22 1.78} 4.89} Akif
    {{Akif 22 1.78} 4.5} {{Vəli 22 1.78} 4.89} Akif


Embedded tipdə kölgəyə salma halı. bu zaman embedded-də kölgəyə düşmüş dəyişəni aşkar formada .EmbeddedTip.üzv ilə götürə bilərik. 


```go
type TələbəKolge struct {
	Adam
	boy float32 //boy Adam-da olduğundan kölgəyə salacaq
}

%%
var a = Adam{
	ad: "Vəli", 
	yaş: 22, 
	boy: 1.78}
var b = TələbəKolge{
	Adam: a, 
	boy: 1.98}

fmt.Println(a.boy, b.boy, b.Adam.boy)
```

    1.78 1.98 1.78


Qeyd edək ki embedded zamanı bir neçə tipi etmək də olur. Sadəcə fikir vermək lazımdır ki konflikt yaranmasın və unikallıq qorunsun.  
Bu zaman üzv adı kimi embedded_in tipi çıxış edir.  
```go
//  T1, *T2, P.T3  *P.T4
struct {
	T1        // üzv adı T1
	*T2       // üzv adı T2
	P.T3      // üzv adı T3
	*P.T4     // üzv adı T4
	x, y int  // üzv adı x and y
}

//Aşağıda isə unikallıq qorunmadığından konflikt olacaq
struct {
	T     // üzv adı T
	*T    // üzv adı T .artıq var və konfliktdədir
	*P.T  // üzv adı T .artıq var və konfliktdədir
}

``` 



```go

type TələbəKonflikt struct {
	Adam
	*Adam
}

type TələbəKonflikt2 struct {
	Adam
	*Tələbə.Adam
}

%%

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


<span style="white-space: pre;"> # gonb_22651709</span>

<br/>

<span class="gonb-cell-line-info">Cell[96]: Line 4</span>
<span class="gonb-error-location">./main.go:26:2: </span> Adam redeclared
<div class="gonb-error-context">
}
type TələbəKonflikt struct {
	Adam
<div class="gonb-error-line">	*Adam
</div>}
type TələbəKonflikt2 struct {

</div>

<br/>

<span class="gonb-cell-line-info">Cell[96]: Line 3</span>
<span class="gonb-error-location">	./main.go:25:2: </span> other declaration of Adam
<div class="gonb-error-context">
	boy float32 //boy Adam-da olduğundan kölgəyə salacaq
}
type TələbəKonflikt struct {
<div class="gonb-error-line">	Adam
</div>	*Adam
}

</div>

<br/>

<span class="gonb-cell-line-info">Cell[96]: Line 9</span>
<span class="gonb-error-location">./main.go:30:2: </span> Adam redeclared
<div class="gonb-error-context">
}
type TələbəKonflikt2 struct {
	Adam
<div class="gonb-error-line">	*Tələbə.Adam
</div>}
type TələbəTip struct {

</div>

<br/>

<span class="gonb-cell-line-info">Cell[96]: Line 8</span>
<span class="gonb-error-location">	./main.go:29:2: </span> other declaration of Adam
<div class="gonb-error-context">
	*Adam
}
type TələbəKonflikt2 struct {
<div class="gonb-error-line">	Adam
</div>	*Tələbə.Adam
}

</div>

<br/>

<span class="gonb-cell-line-info">Cell[96]: Line 9</span>
<span class="gonb-error-location">./main.go:30:13: </span> Tələbə.Adam is not a type
<div class="gonb-error-context">
}
type TələbəKonflikt2 struct {
	Adam
<div class="gonb-error-line">	*Tələbə.Adam
</div>}
type TələbəTip struct {

</div>

<br/>


<span style="white-space: pre;"> </span>

<br/>

</div>




    failed to run "/usr/local/go/bin/go build -o /tmp/gonb_22651709/gonb_22651709": exit status 1


type açar sözü ilə təyin olunan tiplərə metod və ona xas funksiya yazmaq mümkündür. Bunun üçün   

```go

func (a Tip) funksiya_adi( parametrlər ) nəticə {
    ...
}

func (a *Tip) funksiya_adi1( parametrlər ) nəticə {
    ...
}


```


Burada: 
(a *Tip) və ya (a Tip) - qəbul edən tip deyilir. (receiver type  )

Fikir vermək lazımdır ki, Tip ilə qəbul edən kopyası üzərində işləyir.  
Həmçinin *Tip ilə qəbul edən isə  (pointer) üzrə işləyir. Bu zaman biz struktun daxilini dəyişə bilirik.  
Bir də qeyd edək ki *Tip ilə olanda T və *T kimi elan edilmiş bütün dəyişənlər üçün işlədiyindən, təzədən T qəbul edən eyni adlı funksiya vermək olmaz.  

Hər ikisini də   

dəyişən.funksiyanın_adı(...)  kimi çağırmaq olur.   
Qeyd əslində pointer olanı (*dəyişən).funksiyanın_adı(...) kimi çağrılsa da.  
Pointer mövzusunda öyrənəcəyik ki GO dilində . avtomatik bizim üçün həll edir 





```go

func (a Adam) göstər( ) {
	fmt.Println("ad", a.ad,"yaş",a.yaş, "boy", a.boy)
}

type İl int

func (a İl) str( ) string{
	return fmt.Sprintf("%d il keçdi", a )
}

func (a *Adam) yaşını_artır(il İl) {
	a.yaş += int(il)
}

%%
var a = Adam{ ad: "Zakir", yaş: 44, boy: 1.77}
var p = &a
a.göstər()
var artir İl = 2 
fmt.Println(artir.str() )
p.yaşını_artır(artir)
p.göstər()
```

    ad Zakir yaş 44 boy 1.77
    2 il keçdi
    ad Zakir yaş 46 boy 1.77


struct-da həmçinin yuxarıda `teq` əlavəsi də qeyd etmişdik.  
Bu bizə json paketi vəs da.  Həmçinin reflection da öz məqsədlərimiz üçün istifadədə yardımçı olacaq.   
Sadə halda əgər json -a çevirdikdə fərqli adla qeyd edilməsinə imkan verəcək.
```go
type Adam struct {
	ad string `json:"name"`
	yaş int `json:"age"`
	boy float32 `json:"height"`
}



==> jsona çevirmə əməliyyatı 


Nəticə:

{
  "name" : "Zakir ",
  "age"  : 46,
  "height" : 1.77
}
```
