---
title: Map
author:abdelrauf
date: 2023-07-31
tags: ["Map", "Go" ]
---

### Map. Assosativ Array. Dictionary. HashMap

Go dilində bu data struktur dilin özündə gəlir. Map dedikdə key (açar) və ona bağlı value (qiymət) nəzərdə tutulur.   
Bir növ açar sözü bilməklə ona bağlı qiyməti götürmək olur.   
Bu map **nizamsızdır**. Əslən Array özü də map-ə bənzəyir sadəcə orada indeks açar rolunda çıxış edir və data struktur çox sadə olur.     
map-də key tip **== !=** əməliyyatlarına malik tip ola bilər.  
Bu səbədən dildə olan tiplərdən key olaraq slice, map, funksiya tipi çıxış edə bilmir. 



<pre>


 var  ad1  map [key tipi] value_tipi
 var  ad1  map [key tipi] value_tipi = map [key tipi] value_tipi { [achar1 : qiymet1, achar2 : qiymet2, ..]}
 var  ad1 = map [key tipi] value_tipi { [achar1 : qiymet1, achar2 : qiymet2, ..]}
 ad1 := map [key tipi] value_tipi { [achar1 : qiymet1, achar2 : qiymet2, ..]}

 key tipi [ ] mötərizəsi arasında qyed edilir.   
 Burda  map [key tipi] xaric [ ] mötərizə işarəsi istək və buraxıla bilinməyə işarədir

</pre>

map **reference** tipdir. Ona görə adi elan edildikdə nil ola bilər.  
map make ilə də etmək olar. make vasitəsilə bildiyiniz kimi tutumu əvvəlcədən vermək olur. və make initialize etdiyindən nil olmur.  
```Go
make(map[key_tipi] value_tipi, [tutum] )  
```  
tutum - buraxıla bilinər




```go
import "fmt"
%%
var paytaxt map[string] string 

var paytaxt2  = map[string] string {"Azerbaycan" : "Baki", "Irlandiya": "Dublin"}
var paytaxt3 = make(map[string] string)
fmt.Println(paytaxt, "nill?", paytaxt == nil)
fmt.Println(paytaxt2, "nill?", paytaxt2 == nil)
fmt.Println(paytaxt3, "nill?", paytaxt3 == nil)
```

    map[] nill? true
    map[Azerbaycan:Baki Irlandiya:Dublin] nill? false
    map[] nill? false


map -ə əlavə yeni açar və qiymət əlavə etmək üçün, həmçinin açar-ı yeni qiymətə bağlamaq üçün:
```Go
  deyishen[achar] = qiymet
```


```go
%%
paytaxt := map[string] string {"Azerbaycan" : "Baki", "Irlandiya": "Dublin"}
fmt.Println(paytaxt)
//yeni elave edek
paytaxt["Ingiltere"] = "London"
//kohneni deyishek
paytaxt["Azerbaycan"] = "Gozel Baki"
fmt.Println(paytaxt)
```

    map[Azerbaycan:Baki Irlandiya:Dublin]
    map[Azerbaycan:Gozel Baki Ingiltere:London Irlandiya:Dublin]


Bezen map-ə nəsə əlavə etmədən onda həmin açar sözün olub olmadığını yoxlamaq lazım olur. 
Bunun üçün 
```Go
qiymet, movcuddurmu = deyishen[achar]

_, movcuddurmu = deyishen[achar]

//movcuddurmu bool tipdə olur
```


```go
%%
paytaxt := map[string] string {"Azerbaycan" : "Baki", "Irlandiya": "Dublin"}
fmt.Println(paytaxt)
qiymet, varmi := paytaxt["Ingiltere"]
fmt.Printf("%T '%s' \n", qiymet, qiymet)
if !varmi {
	fmt.Println("elave edek")
	paytaxt["Ingiltere"] = "London"
}
fmt.Println(paytaxt)
//skop qydalarina göre if daxilində etsək. hərdəfə yeni dəyişən adına ehtiyac olmayacaq
if _, varmi2 := paytaxt["Brazilya"]; varmi2 {
	fmt.Println("elave edek")
	paytaxt["Brazilya"] = "Braziliya"
}
//skop qaydalarina göre if daxilində etsək. hərdəfə yeni dəyişən adına ehtiyac olmayacaq
if _, varmi2 := paytaxt["Argentina"]; varmi2 {
	fmt.Println("elave edek")
	paytaxt["Argentina"] = "Beunos Aeros"
}
fmt.Println(paytaxt)
```

    map[Azerbaycan:Baki Irlandiya:Dublin]
    string '' 
    elave edek
    map[Azerbaycan:Baki Ingiltere:London Irlandiya:Dublin]
    map[Azerbaycan:Baki Ingiltere:London Irlandiya:Dublin]


map-dən hər hansı açar sözü pozmaq üçün delete builtin funksiyasını istifadə edə bilərik
```Go
delete(deyishen, achar)
```


```go
%%
paytaxt := map[string] string {"Azerbaycan" : "Baki", "Irlandiya": "Dublin"}
fmt.Println(paytaxt)
delete(paytaxt, "Irlandiya")
fmt.Println(paytaxt)
```

    map[Azerbaycan:Baki Irlandiya:Dublin]
    map[Azerbaycan:Baki]


map for range ilə:
```Go
for achar, qiymet := map_deyisheni{
	...
}
//tekce achar
for achar  := map_deyisheni{
	...
}

//qiymet
for achar, _  := map_deyisheni{
	...
}
```

map-i **set** kimi unikal list kimi istifadə misalı  



```go
%%
set:= map[string] bool {}
for _, x := range []string{"alma", "armud", "alma", "armud", "banan", "kivi", "kivi", "apelsin"}{
	set[x] = true
}
//unikallar
for achar := range set {
	fmt.Println(achar)
}
 
```

    apelsin
    alma
    armud
    banan
    kivi


map of slice misalı


```go
%%
map_list:= map[string][]string {}
map_list["Zakir"]=[]string {"Bekir"}

//append nil uzre ishleyir
map_list["Eli"]=append(map_list["Eli"], "Orxan", "Zakir")
map_list["Xabib"]=append(map_list["Xabib"], "Ismayil", "Zakir")
map_list["Zakir"]=append(map_list["Zakir"], "Ismayil")

for achar, qiymet := range map_list {
	fmt.Print(achar,": ")
	for _, dost := range qiymet {
		fmt.Print(dost, " ")
	}
	fmt.Println()
}
```

    Zakir: Bekir Ismayil 
    Eli: Orxan Zakir 
    Xabib: Ismayil Zakir 


<span style="color:red"><b>Qeyd etdiyimiz kimi map nizamsızdır. Bu səbəbdən print fərqli çıxa bilər. Nizamlı map üçün əlavələrə(Go Ordered Map) baxmaq lazımdır. </b></style>
