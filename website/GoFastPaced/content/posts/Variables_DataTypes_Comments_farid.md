---
title: " Variables Data Types Comments"
author: "farid"
date: 2023-07-31
tags: ["Variables", "Data", "Types", "Comments", "Go" ]
---

GO dilinin sintaksisi
==
Go programlaşdırma dilində kodun mətni */.go formatında faylda olur.
Go faylında package elanı, modul və paketlərin daxil edilməsi(import), funksiyalar, əmrlər(statements) və ifadələr(expressions) olur. Məsələn:
        
        package main
        import "fmt"

        func main() {
            fmt.Println("Hello World!")
        }


Buradakı misalda kodun "main" paketinə aid olduğu göstərilir. Daha sonra koda standart "fmt" paketi daxil edilir. Bu paket Go dilinin əsas kitabxanası (library) hesab olunur və içərisində mühim funksiyalar və standart giriş/çıxış (std input/output) mövcuddur.
Sonra main funksiyası başladılır və içərisində standart çıxışla ekrana salamalama mətni çıxarılır. Sonrakı bölmələrdə bu mövzular daha çox izah olunacaq.



Comments (Şərhlər)
===

Şərhlər (comments) hər bir proqramlaşdırma dilində kodun işləməyinə heçbir təsiri olmayan, lakin çox önəmli bir hissədir.
Onun vasitəsilə kodun mühim hissələrində proqramçı öz qeydlərini yazır. Bu qeydlər həm özünə, həm də koda baxan digər şəxslərə bir xatırlatma, aydınlatma rolunu oynayır.
Şərhlər proqram işləyərkən heç bir rol oynamır. Ancaq kodun içində qeyri məhdud sayda yazıla bilər.
İki forması mövcuddur: 
- ***Birsətirli***: // (iki çəp xətt) işarəsindən sonra sətrin sonuna qədər yazılan mətn şərh sayılır.
- ***Çoxsətirli***: /* (çəp xətt və ulduz) işarəsindən sonra */ işarəsi gələnə qədər olan blok şərh sayılır. Bir neçə sətri əhatə edə bilər. Məsələn:
 


```go
// Bu birsətirli şərhdir.
package main
import ("fmt")

func main() {
  
  /* Bu isə
  çoxsətirli şərhdir.
  Hələ də davam edir... 
  Burada şərhi bitiririk */

  fmt.Println("Hello World!")
}
```

Data Types   *(Verilənlər Tipi)*
===

Verilənlər tipi hər bir proqramlaşdırma dilinin ən mühim mövzularındandır. 
Go dili statik tiplidir. Bu o deməkdir ki, bir dəyişən(variable) müəyyən tipdə elan olunubsa yalnız o tip verilən(data) daşıya bilər.
Dəyişən içində məlumat daşıyan bir qutudur. Yaddaşda müəyyən yer tutur. Onun ölçüsü içində daşıdığı tipə görə təyin olunur.
Standarta uyğun olaraq bundan sonra izahı verilməklə beynəlxalq(ingiliscə) terminlərdən istifadə olunacaq.


***GO*** dilində əsas 3 baza(primitiv) tip var:
- **Bool** : Özündə iki boolean dəyər daşıyır. True və False (True həqiqət, müsbət dəyər, mövcudluq, False isə yalan, mənfi dəyər və yoxluq ifadə edir)
- **Numeric** (ədədlər): Bura daxildir integer (tam ədədlər), float (kəsr ədədlər) və complex (kompleks ədədlər).
- **String** (sətir/mətn): iki və daha artıq simvol, hərf string hesab olunur və kodda dırnaq işarəsi içində yazılır.


Boolean Data Type
----
Bu tipi elan etmək üçün *bool* açar sözündən istifadə olunur.
 >Açar sözlər dildə rezerv olunmuş kəlimələrdir. Dəyişən və funksiya adı kimi istifadə oluna bilməz. Açar sözlər dilin sintaksisinə aiddir və kodun hər bir hissəsində istifadə olunur. Məsələn: *import*, *if*, *func*.
 
Bool tipi iki dəyər daşıyır: true və false. Əsasən şərt operatorunda ifadənin həqiqi olub olmadığını yoxlamaq üçün istifadə olunur.
Default (başlanğıc sərbəst) dəyəri *false*-dır. 



```go
package main
import ("fmt")

func main() {
  var dəyər1 bool = true // dəyər1 dəyişəni bool tipində elan olunub və true dəyəri mənimsədilib.
  var dəyər2 bool // dəyər2 bool tipində elan olunub amma dəyər verilməyib.
  
  fmt.Println(dəyər1) // ekranda true çıxacaq
  fmt.Println(dəyər2) // ekranda false çıxacaq (Başlanğıc dəyəri)
}
```

Integer Data Types  *(Tam ədədlər)*
---

Integer tipinə tam ədədlər daxildir. Məsələn: 54, 0, -37, 945 və s.

İnteger tipi iki cür olur: 
- signed integer, tam ədəd üçün default tipdir və *int* sözü ilə elan olunur. Bu tipə mənfi və müsbət tam ədədlər daxildir.
- unsigned integer, bu tip *uint* sözü ilə elan olunur və yalnız mənfi olmayan tam ədədlər üçün istifadə olunur.

Təfsilatı cədvəldə göstərilib:

|Tip  |Ölçü  |Dəyər |
|--|--|--|
|int  |32 yaxud 64bit sistemə görə dəyişir  |-2147483648 ~ 2147483647 32bit sistemdə, <br> -9223372036854775808 ~ 9223372036854775807 64bit sistemdə |
|int8|8bit / 1bayt|-128 ~ 127  |
| int16 |16bit / 2bayt  |-32768 ~ 32767  |
|int32  |32bit / 4bayt  |-2147483648 ~ 2147483647  |
| int64 |64bit / 8bayt  | -9223372036854775808 ~ 9223372036854775807|
|--|--|--|
|uint  |32 yaxud 64bit sistemə görə dəyişir  |0 ~ 4294967295 32bit sistemdə, <br> 0 ~ 18446744073709551615 64bit sistemdə |
|uint8|8bit / 1bayt|0 ~ 255  |
| uint16 |16bit / 2bayt  |0 ~ 65535  |
|uint32  |32bit / 4bayt  |0 ~ 4294967295  |
| uint64 |64bit / 8bayt  | 0 ~ 18446744073709551615|

Cədvəldə qeyd olunmuş tiplər dəyişənə veriləcək dəyərə görə təyin olunur.
Əgər konkret tip verilməzsə tam ədədlər üçün standart tip **int** seçiləcək. integer tipi üçün standart dəyər sıfırdır.


```go
package main
import ("fmt")

func main() {
  var a int8 = 100
  var b int16 = 5000
  var c int
  d := 3500
  fmt.Printf("Type: %T, value: %v\n", a, a) // Type: int8, value: 100
  fmt.Printf("Type: %T, value: %v\n", b, b) // Type: int16, value: 5000
  fmt.Printf("Type: %T, value: %v\n", c, c) // Type: int, value: 0 (default dəyər)
  fmt.Printf("Type: %T, value: %v", d, d) // Type: int, value: 3500 (default tip)
}
```

Float Data Types  *(Onluq kəsrlər)*
---

Float data tipləri mənfi və müsbət onluq kəsrləri qeyd etmək üçün istifadə olunur. Tam hissəni ayırmaq üçün nöqtədən istifadə olunur.

Float ədədlər üçün iki tip istifadə olunur:
| Tip | Ölçü  | Dəyər  |
|--|--|--|
|float32  |32bit  | -3.4e+38 ~ 3.4e+38. |
| float64 |64bit  |-1.7e+308 ~ +1.7e+308.  |


Veriləcək dəyərə görə *float32* və ya *float64* tipi seçilir. Default tip *float64*-dür. O daha çox dəyər aralığına malikdir. Float tipləri üçün default dəyər sıfırdır.


```go
package main
import ("fmt")

func main() {
  var x float32 = 123.78
  var y float32
  z := 79.65
  fmt.Printf("Type: %T, value: %v\n", x, x) // Type: float32, value: 123.78
  fmt.Printf("Type: %T, value: %v\n", y, y) // Type: float32, value: 0 (default dəyər 0)
  fmt.Printf("Type: %T, value: %v", z, z) // Type: float64, value: 79.65 (default tip float64)
}
```

String Data Type  *(Mətn tipi)*
---

String tipi simvol ardıcıllığı (mətn) yadda saxlamaq üçün istifadə olunur. Mütləq qoşa dırnaq işarəsi içində yazılır. Default dəyəri boş sətirdir.


```go
package main
import ("fmt")

func main() {
  var txt1 string = "Salam!"
  var txt2 string
  txt3 := "Dünya"

  fmt.Printf("Type: %T, value: %v\n", txt1, txt1) // Type: string, value: Salam!
  fmt.Printf("Type: %T, value: %v\n", txt2, txt2) // Type: string, value:  (default dəyər boş sətir)
  fmt.Printf("Type: %T, value: %v\n", txt3, txt3) // Type: string, value: Dünya (mətn üçün default tip avtomatik seçilib)
}  
```

Variables  *(Dəyişənlər)*
===

Dəyişənlər (variables) verilənlərin (data) dəyərini yadda saxlamaq üçün istifadə olunan konteynerdir. Dəyişən elan olunarkən onun tipi və ölçüsü elan olunan datanın tipinə uyğun təyin olunur. Müəyyən tip dəyər dəyişənə mənimsədilir və daha sonra dəyişən ilkin olaraq hansı tipdə elan olunubsa yalnız o tipdə dəyər qəbul edir.

Dəyişənin adı
---
Dəyişənə qısa (a, b, c və s.) yaxud geniş və izahlı ad (soyad, yaş, doğulduğu_yer və s.) verilə bilər.  
İkinci variant daha münasibdir. Məna ifadə etməyən qısa adlar kod böyüdükcə unudula, çaşqınlıq yarada bilər.

Dəyişən adları aşağıdakı qaydalara uyğun olmalıdır:
- Dəyişən adı ancaq hərf, rəqəm və alt_xətt( _ ) işarəsindən ibarət ola bilər, lakin rəqəmlə başlaya bilməz.
- Dəyişən adları reqistrə həssasdır. Yəni Şifrə, şifrə və ŞİFRƏ sistem tərəfindən fərqli dəyişən kimi qəbul olunur.
- Dəyişən adlarında boşluq (space) ola bilməz.
- Dildəki açar sözlər dəyişən adı kimi istifadə oluna bilməz. Məsələn if, import, func, string və s.
- Dəyişən adı istənilən uzunluqda ola bilər.

İki və daha çox sözdən ibarət adların oxunaqlı olması üçün 3 forma tövsiyyə olunur:
- **Camel Case** = İlk söz istisna bütün sözlər böyük hərflə başlayır: ***mürəkkəbDəyişənAdı***
- **Pascal Case** = Bütün sözlər böyük hərflə başlayır: ***MürəkkəbDəyişənAdı***
- **Snake Case** = sözlər alt_xətt (underscore) vasitəsilə ayrılır: ***mürəkkəb_dəyişən_adı***

Dəyişənlərin elan olunması(yaradılması)
---

GO proqramlaşdırma dilində dəyişənlər iki cür elan olunur:

1. **var** açar sözü vasitəsilə.
2. **:=** işarəsi vasitəsilə. 

#### **var** sözü ilə dəyişən üç formada elan oluna bilər: ####

1.  ***var** dəyişənin_adı dəyişənin_tipi = dəyər* . 
> Burada var açar sözündən sonra dəyişənin adı, tipi qeyd olunur və mənimsətmə("=") işarəsi ilə dəyər yazılır.<br>Dəyər əvvəlcədən bilinirsə dəyişən məhz bu formada təyin olunur.
      
        var avtomobil_1 string = "Mercedes" 
        var avtomobil_2 string = "BMW"
        var buraxılış_ili int = 2015 
2. ***var** dəyişənin_adı dəyişənin_tipi*. 
> Bu zaman dəyişən elan olunur, lakin ona heçbir dəyər mənimsədilmir. Bu səbəbdən dəyişən tipinə uyğun olaraq default dəyər alır.  <br> **int** və **float** üçün bu dəyər sıfır, **string** üçün boş sətir(""), **bool** üçün isə false-dır. 

        var sayğac uint /* sayğac adlı dəyişən elan olunur və dəyər verilmir. Default dəyər olaraq sıfır təyin olunur.*/
        var sayğac_dəyişib bool /* sayğac_dəyişib adlı dəyişənə bool tipi verilib. Dəyər verilmədiyinə görə default dəyər olan false təyin olunub.*/

3. ***var** dəyişənin_adı = dəyər*
> Bu formada dəyişənin adına birbaşa dəyər mənimsədilib. Belə halda dəyişənin tipi mənimsədilən dəyərə uyğun sistem tərəfindən təxmin edilir(inferred from value).
     
        var oğlan = "Əlövsət"  /* dəyişənin tipi string təyin olunub. */
        var qız = "Bənövşə"  /* eyni şəkildə */
        var oğlan_yaş = 15  /* dəyişənin tipi int təyin olunub */
        var qız_yaş = 13  /* eyni şəkildə */ 


#### **:=** işarəsi ilə dəyişənin elan edilməsi ####
*dəyişənin_adı := dəyər* 
> Bu formada **var** sözü istifadə olunmur, dəyər **:=** işarəsi ilə birbaşa dəyişənə mənimsədilir. Tip dəyərə görə kompilyator tərəfindən təyin olunur(inferred)

        ölkə := "Azərbaycan" /* tip string olaraq təyin olunub */
        tarixi := 1918 /* tip int olaraq təyin olunub. */
> **var** açar sözündən fərqli olaraq **:=** işarəsi ilə elan olunduqda dəyişən yalnız funksiyanın içində elan edilə bilər və həmçinin dəyər dərhal qeyd edilməlidir.
        
        

### **Çox sayda dəyişənin bir sətirdə elan edilməsi**  ###

        var ad1, ad2, ad3, ad4 string = "Hikmət", "Roma", "Vasif", "Validə"

> Misaldakı kimi tip qeyd edilibə bütün dəyişənlər bu tipdə yaradılır və mənimsədilən dəyər yalnız bu tipdə ola bilər.<br>
Əgər tip qeyd edilməzsə istənilən tipdə dəyər verilə bilər. Hər dəyişənin tipi ona verilən dəyərə uyğun təyin olunur. := işarəsindən də istifadə oluna bilər.

        var ad1, ad2, yaş1, yaş2 = "Nərgiz", "Xəyal", 17, 20
        şəxs, boy, çəki := "Fərid", 186, 106

Həmçinin **var** bloku açılaraq istənilən sayda dəyişən blokun içində elan edilə bilər. Bu forma daha çox dekorativ xarakterlidir və kodun oxunaqlığını artırır.
    
        package main
        import ("fmt")

        func main() {
           var (
           a int
           b int = 1
           c string = "hello"
           )

        fmt.Println(a)
        fmt.Println(b)
        fmt.Println(c)
        }

