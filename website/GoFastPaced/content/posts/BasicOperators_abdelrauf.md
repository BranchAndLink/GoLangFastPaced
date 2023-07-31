---
title: BasicOperators
author:abdelrauf
date: 2023-07-31
tags: ["Basic", "Operators", "Go" ]
---

## Sadə Operatorlar


iki operand (əməl həyata keçirilən dəyişən ədəd və s) işlədilən operatorlara **binary operator**lar deyirik.     
Məsələn toplama iki ədədin arasında işlədilir 5 + 9   
tək operandın önündə işlənə bilən operatorlara **unary operator**lar deyirik.    
Məsələn mənfi işarəsi  -8

#### Hesablama operatorları

+,-,*,/,%, ++, --

+ bütün ədəd tipləri üçün işləyir və toplamadır.  <span style="color:red">string tipi üçün də işləyir  lakin **birləşdirmə** əməliyyatı aparır</span>


```go
%%
var (
	a ,b int = 5 ,6
	c, d float32 =  4.5, 7.87
	e, h complex64 = 5, 7+9i
	m1, m2 = "Learn-", "Go"
)

fmt.Println(a+b, c + d, m1 + m2, e  + h)
```

    11 12.37 Learn-Go (12+9i)


% modul bildirir və yalnız **tam ədədlər** üçün işləyir


```go
%%
d, m := 12 , 5

fmt.Println(d%m)
```

    2



```go
%%
a,b,c,d := 5, 4, 7, 14

fmt.Println(a*b-d/c)
```

    18


**C dilindən fərqli olaraq ++ və -- operatoru **postfix** dir və ifadə içində yazılmır**. Dəyişənin qiymətini 1 vahid artırır və ya azaldır


```go
%%
i:=0
i++
i++
i--
fmt.Println(i)
```

    1


Bit üzrə operatorlar. Bu operatorlar yalnız **tam ədədlər üzərində** işləyir

| Bit (İkili sayda) operatoru |İzahı|
|---|---|
| m & n| bit and, ədədin bitləri üzrə hər ikisi 1 olduqda 1 əksa halda 0|
| m \| n| bit or, hər ikisi 0 olduqda 0, qalan hallarda 1|
|m ^ n | xor eyni olduqda 0 fərqli olduqda 1|
|  ^n | n dəki bitlər əksinə çevrilir|
|m &^ n | bit clear (and  not m & (^n)) n də 1 olan mərtəbələr m ədədində sıfırlanır|
|m << n | m ədədini n sayda sola sürüşdürür|
|m >> n | m ədədini n sayda sağa sürüşdürür|


```go
%%
var a,b uint16 = 1235, 4536
fmt.Printf("%016b(onluq: %d)\n&\n%016b(onluq: %d)\n__________________\n%016b  (onluq: %d)\n\n\n",a,a,b,b,a&b,a&b)
fmt.Printf("%016b(onluq: %d)\n|\n%016b(onluq: %d)\n__________________\n%016b  (onluq: %d)\n\n\n",a,a,b,b,a|b,a|b)
fmt.Printf("%016b(onluq: %d)\n^\n%016b(onluq: %d)\n__________________\n%016b  (onluq: %d)\n\n\n",a,a,b,b,a^b,a^b)
fmt.Printf("%016b(onluq: %d)\n&^\n%016b(onluq: %d)\n__________________\n%016b  (onluq: %d)\n\n\n",a,a,b,b,a&^b,a&^b)

```

    0000010011010011(onluq: 1235)
    &
    0001000110111000(onluq: 4536)
    __________________
    0000000010010000  (onluq: 144)
    
    
    0000010011010011(onluq: 1235)
    |
    0001000110111000(onluq: 4536)
    __________________
    0001010111111011  (onluq: 5627)
    
    
    0000010011010011(onluq: 1235)
    ^
    0001000110111000(onluq: 4536)
    __________________
    0001010101101011  (onluq: 5483)
    
    
    0000010011010011(onluq: 1235)
    &^
    0001000110111000(onluq: 4536)
    __________________
    0000010001000011  (onluq: 1091)
    
    



```go
%%
var a,b, n uint16 = 1235, 4536, 3
fmt.Printf("^\n%016b(onluq: %d)\n__________________\n%016b  (onluq: %d)\n\n\n",b,b,^b,^b)
fmt.Printf("%016b(onluq: %d)\n>>\n%16d\n__________________\n%016b  (onluq: %d)\n\n\n",a,a,n,a>>n,a>>n)
fmt.Printf("%016b(onluq: %d)\n<<\n%16d\n__________________\n%016b  (onluq: %d)\n\n\n",a,a,n,a<<n,a<<n)
```

    ^
    0001000110111000(onluq: 4536)
    __________________
    1110111001000111  (onluq: 60999)
    
    
    0000010011010011(onluq: 1235)
    >>
                   3
    __________________
    0000000010011010  (onluq: 154)
    
    
    0000010011010011(onluq: 1235)
    <<
                   3
    __________________
    0010011010011000  (onluq: 9880)
    
    


Müqayisə operatorları   

|operator|izahı|
|---|---|
|== |   bərabərdir|
|!=  |  bərabər deyil|
|<  |   kiçikdir|
|<=  |  kiçik bərabərdir|
|>  |   böyükdür|
|>=  |  böyük bərabərdir|

Ədəd tipləri və string müqayisə ediləndir. (string tipində kodlaşdırma hərf sırası müqayisəyə təsir edir).  
Həmçinin array də müqayisə ediləndir.  
Əgər struct tipinin bütün elementləri müqayisə ediləndirsə onlar da müqayisə ediləndir  
Daha ətraflı spesifikasiyaya baxa bilərsiniz


```go
%%
d,e := 102, 100
s,v := "Ananas", "Anar"
fmt.Println (d,">",e, ": ", d > e)
fmt.Println (s,">", v, ": ", s > v)
```

    102 > 100 :  true
    Ananas > Anar :  false


Məntiqi operatorlar   




|operator|izahı|
|---|---|
| && |   məntiqi və. hər ikisi true olduqda true|
| \|\|  |  məntiqi və ya. ikisindən biri true olduqda true.|
| !  |   əksi true olduqda | 


```go
%%
a,b := 12,11
c,d := 9, 13

fmt.Println(a>b && c>d)
fmt.Println(a>b || c>d)
fmt.Println( !(c>d) )
```

    false
    true
    true


Mənimsətmə operatorları



|operator |	nümunə |izah |
|---|---|---|
|=	|x = y	| x-ə y-in qiyməti mənimsənir	|
|+=	|x += y	| x = x + y	|
|-=	|x -= y	| x = x - y	|
|*=	|x *= y	| x = x * y	|
|/=	|x /= y	| x = x / y	|
|%=	|x %= y	| x = x % y	|
|&=	|x &= y	| x = x & y	|
|\|=	|x \|= y	| x = x \| y	|
|^=	|x ^= y	| x = x ^ y	|
|>>=	|x >>= y	| x = x >> y	|
|<<=	|x <<= y	| x = x << y|

Mənimsətmə , vergul ilə


```go
%%
a,b:=0,5
a,b = a+1, b-a
fmt.Println(a,b)
```

    1 5



```go
%%
a,b:=0,5
a,b = b,a
fmt.Println(a,b)
```

    5 0


Əgər funksiya tuple (bir neçə tip) qaytarırsa. <span style="color:red">Funksiyaları növbəti dərslərdə keçəcəyik</span>


```go
func tt(x int) (int, string)  {
	return 5 + x , "hello"
}

func main(){
	a,b := tt(2)
	fmt.Println(a,b)
	_,c := tt(3)
	fmt.Println(c)
	d,_ := tt(4)
	fmt.Println(d)
}
```

    7 hello
    hello
    9


Operatorların üstünlüyü və assosasitivliyi və tiplərə bağlılığı

 Ən üstdəki daha güclü operator sayılır. Məsələn * +-dən güclü olduğu üçün ilk öncə o yerinə yetirilir   
 Üstünlük yuxarıdan aşağıya doğru


|kateqoriya |operator|assosasitivliyi|
|---|---|---|
|5. Əsasən Multiplikativ|	*  /   % <<   >> & &^	|Soldan Sağa|
|4. Əsasən Additiv|	+ -	\| ^ |Soldan Sağa|
|3. Müqayisə|	== != <  <=  >  >=	|Soldan Sağa| 
|2. Məntiqi AND|	&&	|Soldan Sağa|
|1. Məntiqi OR	|\|\|	|Soldan Sağa| 

Assosativlik eyni güclü operatorlarda hesablama və  həyata keçirmə ardıcıllığını göstərir. 


<span style="color:red">**Proqramçı qeydi.** Bu cədvəl dillərdən dillərə dəyişilir. Ona görə digər dildən Go dilinə çevirəndə ordakı üstünlük cədvəlini nəzərdən keçirmək lazımdır.  
Ümumən məsləhətdir ki mötərizələrdən istifadə edilsin.</span>


```go
%%
var a,b,c  = 18, 2, 3

netice := a / b *c  // (a/b)*c  yoxsa a/(b*c) ?
fmt.Printf("%d/%d*%d ifadesinde (%d/%d)*%d = %d yoxsa %d/(%d*%d) = %d ?\n",a,b,c, a,b,c ,(a/b)*c ,a,b,c,a/(b*c))

fmt.Println("Asosativlik soldan saga oldugu uchun\nnetice: ", netice)
```

    18/2*3 ifadesinde (18/2)*3 = 27 yoxsa 18/(2*3) = 3 ?
    Asosativlik soldan saga oldugu uchun
    netice:  27

