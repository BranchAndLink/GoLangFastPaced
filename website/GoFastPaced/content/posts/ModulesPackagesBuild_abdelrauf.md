---
title: " Modules Packages Build "
author: "abdelrauf"
date: 2023-07-31
tags: ["Modules", "Packages", "Build", "Go" ]
---

#### Module   
Module -  go.mod faylı ana papkada yerləşən bir neçə package-dən ibarətdir.   
Go-da ```go mod``` əmri ilə module idarə edilir.   


```go mod init module_adi```-lə module yaradaq


```go
!rm -r mymodule
!mkdir -p mymodule
!cd mymodule && rm -rf go.mod && go mod init github.com/yourname/mymodule
```

    rm: cannot remove 'mymodule': No such file or directory
    exit status 1
    go: creating new go.mod: module github.com/yourname/mymodule


```go mod init``` 
go.mod faylını yaradacaq


```go
!cat mymodule/go.mod
```

    module github.com/yourname/mymodule
    
    go 1.20


Əgər əlavə modullar varsa. 
```go mod tidy``` işə salsaq avtomatik go.mod-a əlavə edəcək

#### Package.   

 Package - bağlama - bir və bir neçə fayldan  ibarət olub eyni təyinat üçün yazılmış kodları qruplaşdırmaq üçündür.   
 Həmçinin axtarış və kompilyasıya vaxtına da xeyli qənaət edir.  
 Faylın hansı package aid olmağını bu cür qeyd edə bilərik:  
 ```Go
 package packagin_adi
 ```  

Package də olanları istifadə etmək üçün import sözündən istifadə edilir: 

```Go
   import "package_adi" // Bu zaman bütün publik (export) olunmuş elementlərə package_adi.  daxil olunur
   import alias "package_adi"// Bu zaman bütün publik (export) olunmuş elementlərə alias. ilə daxil olmaq olar
   import "packagin_urli"
   import . "package_adi" // Bu zaman bütün publik (export) olunmuş elementlərə package_adi. -siz daxil olmaq olur
   import _ "package_adi" // Yalnız init funksiyası çağrılacaq

   import (
      ...  //blok elan bir neçə package
   )
```

Package identifikatorun fərqliliyinə də səbəb olur. Yəni fərqli package də olan export olunmamış eyni adlı identifikatorlaq fərqli sayılır

Package də export olunmuş identifikatorlar:
- Böyük hərflə başlayırsa
- package bloku içindədirsə, metod və ya field adıdırsa   
sayılır.   

Export olunmamış identifikatorlar package xaricində görsənmir.

#### Package initialization (inisializasiyalar - yəni işə salmaq üçün edilən ilkin işlər)
İlk öncə eyni proses import olunmuş package-lərdə gedəcək. Daha sonra həmin proses     
- package bloku dəyişənləri inisializasiya olunacaq
- init funksiyaları çağrılacaq.    

init funksiyası məcburi deyil və **bir neçə dəfə** (hətta bir fayl daxilində) təyin oluna bilər .   
Bütün bu proses ardıcıllıqla (sequentially) icra olunacaq .    

Fiziki olaraq package adı papka adı ilə eyni olmalıdır. mymodule modulumuzda package1 adlı papka(qovluq) quraq   
və orda aşağıdakı məzmunlu fayllar əlavə edək
```Go
//mymodule/package1/numune.go
package package1
import "fmt"

func init(){
  fmt.Println("init call")
}
//bir neçə init call mümkündür
func init(){
  fmt.Println("init call ###")
}

func Salam(str string)  {
 fmt.Println( str )
}

```
```Go
//mymodule/package1/numune2.go
package package1

import "github.com/common-nighthawk/go-figure"
func Salam2(str string)  {
  myFigure := figure.NewFigure(str, "", true)
  myFigure.Print()
}

```


```go
!mkdir -p mymodule/package1
!echo -e "package package1\n \
import \"fmt\"\n  \
func init(){ \n \
	fmt.Println(\"init call\") \n \
  } \n \
  //bir neçə init call mümkündür \n \
  func init(){ \n \
	fmt.Println(\"init call ###\") \n \
  } \n \
func Salam(str string)  {\n\
\tfmt.Println( str )\n}" >mymodule/package1/numune.go


```


```go
!echo -e "package package1\n \
import \"github.com/common-nighthawk/go-figure\"\n  \
func Salam2(str string)  {\n\
\tmyFigure := figure.NewFigure(str, \"\", true) \n \
\tmyFigure.Print()\n}" > mymodule/package1/numune2.go
```

go fmt ilə kodumuzun gözəl görsənməsi üçün format edək. Bunun üçün modulumuz olan qovluğun içinə girib ```go fmt .``` çağırırıq


```go
!cd mymodule/package1 && go fmt .
```

    numune.go
    numune2.go


```go mod tidy``` ilə əlavə asılılıqları avtomatik əlavə edək. Əgər repo dursa onu go get -lə də etmək olar


```go
!cd mymodule  && go mod tidy  
```

    go: finding module for package github.com/common-nighthawk/go-figure
    go: found github.com/common-nighthawk/go-figure in github.com/common-nighthawk/go-figure v0.0.0-20210622060536-734e95fb86be


go.mod faylımızın məzmunu bu cür olacaq.


```go
!cat mymodule/go.mod
```

    module github.com/yourname/mymodule
    
    go 1.20
    
    require github.com/common-nighthawk/go-figure v0.0.0-20210622060536-734e95fb86be



```go
!cd mymodule/package1 && go build .

```


Lakin bu **yayımlanan (publish edilən)  modul üçün keçərlidir**.  Necə ki biz github.com/common-nighthawk/go-figure repodan yüklədik  
Gəlin GONB də **öz indicə yaratdığımız yayımlanmamış modulumuzu istifadə edək**.  
(Qeyd: Eynilə siz yeni modul yaradıb orda da edə bilərsiniz. Sadəcə GONB üçün !* əlavəmizi ignor edin)

Gonb bildiyimiz kimi temporary papkada ishleyir.      
orada olan go.mod -a module-uzu elave edək və daxildəki lokal papkaya yönləndirək.   
Bunun üçün biz go.mod faylında  require ilə əlavə edib həmçinin replace-direktivi ilə daxili papkaya və s yönləndirə bilərik.    
direktivlərdən **retract toolchain exclude** da var   
```
require   module versiya 
require(
    module1 versiya1
    ....
)


replace module1 versiya => module2_yeni_path versiya

``````
Dediklərimizi edit-lə də etmək olar. Əgər modul publish olsa idi onu ```go get```-lə də götürmək olardı   
```go mod edit```  help manualına əks etdirək




```go
!go help mod edit
```

    usage: go mod edit [editing flags] [-fmt|-print|-json] [go.mod]
    
    Edit provides a command-line interface for editing go.mod,
    for use primarily by tools or scripts. It reads only go.mod;
    it does not look up information about the modules involved.
    By default, edit reads and writes the go.mod file of the main module,
    but a different target file can be specified after the editing flags.
    
    The editing flags specify a sequence of editing operations.
    
    The -fmt flag reformats the go.mod file without making other changes.
    This reformatting is also implied by any other modifications that use or
    rewrite the go.mod file. The only time this flag is needed is if no other
    flags are specified, as in 'go mod edit -fmt'.
    
    The -module flag changes the module's path (the go.mod file's module line).
    
    The -require=path@version and -droprequire=path flags
    add and drop a requirement on the given module path and version.
    Note that -require overrides any existing requirements on path.
    These flags are mainly for tools that understand the module graph.
    Users should prefer 'go get path@version' or 'go get path@none',
    which make other go.mod adjustments as needed to satisfy
    constraints imposed by other modules.
    
    The -exclude=path@version and -dropexclude=path@version flags
    add and drop an exclusion for the given module path and version.
    Note that -exclude=path@version is a no-op if that exclusion already exists.
    
    The -replace=old[@v]=new[@v] flag adds a replacement of the given
    module path and version pair. If the @v in old@v is omitted, a
    replacement without a version on the left side is added, which applies
    to all versions of the old module path. If the @v in new@v is omitted,
    the new path should be a local module root directory, not a module
    path. Note that -replace overrides any redundant replacements for old[@v],
    so omitting @v will drop existing replacements for specific versions.
    
    The -dropreplace=old[@v] flag drops a replacement of the given
    module path and version pair. If the @v is omitted, a replacement without
    a version on the left side is dropped.
    
    The -retract=version and -dropretract=version flags add and drop a
    retraction on the given version. The version may be a single version
    like "v1.2.3" or a closed interval like "[v1.1.0,v1.1.9]". Note that
    -retract=version is a no-op if that retraction already exists.
    
    The -require, -droprequire, -exclude, -dropexclude, -replace,
    -dropreplace, -retract, and -dropretract editing flags may be repeated,
    and the changes are applied in the order given.
    
    The -go=version flag sets the expected Go language version.
    
    The -print flag prints the final go.mod in its text format instead of
    writing it back to go.mod.
    
    The -json flag prints the final go.mod file in JSON format instead of
    writing it back to go.mod. The JSON output corresponds to these Go types:
    
    	type Module struct {
    		Path    string
    		Version string
    	}
    
    	type GoMod struct {
    		Module  ModPath
    		Go      string
    		Require []Require
    		Exclude []Module
    		Replace []Replace
    		Retract []Retract
    	}
    
    	type ModPath struct {
    		Path       string
    		Deprecated string
    	}
    
    	type Require struct {
    		Path string
    		Version string
    		Indirect bool
    	}
    
    	type Replace struct {
    		Old Module
    		New Module
    	}
    
    	type Retract struct {
    		Low       string
    		High      string
    		Rationale string
    	}
    
    Retract entries representing a single version (not an interval) will have
    the "Low" and "High" fields set to the same value.
    
    Note that this only describes the go.mod file itself, not other modules
    referred to indirectly. For the full set of modules available to a build,
    use 'go list -m -json all'.
    
    Edit also provides the -C, -n, and -x build flags.
    
    See https://golang.org/ref/mod#go-mod-edit for more about 'go mod edit'.


##### Gonb papkasinda etmək üçün !* istifadə edəcəyik   
və əsasən asılılığı əlavə edib  
```go mod edit -require=path@version```    
sonra isə aşağıdakı ilə lokal path-a yönləndirəcəyik     
```go mod edit -replace=path1@version=path2```   
Bu cür etməklə yayımlanandan sonra replace-dən çıxara bilərik.  


```go
!echo ${PWD} >/tmp/curr_dir
//versiyani v{eded}.{eded}.{eded}-metn kimi verek
//go mod edit -require ile yeni asılılıq əlavə edirik
!*go mod edit -require=github.com/yourname/mymodule@v0.0.0-unpublished
//go mod edit -replace ile yeni asılılığı lokal papkaya yönləndirik
!*CURR_DIR=`cat /tmp/curr_dir` && go mod edit -replace="github.com/yourname/mymodule@v0.0.0-unpublished"="${CURR_DIR}/mymodule"
```

GONB-in go.mod faylı belə olacaq:


```go
!*cat go.mod
```

    module gonb_42b161cd
    
    go 1.20
    
    require github.com/yourname/mymodule v0.0.0-unpublished
    
    require github.com/common-nighthawk/go-figure v0.0.0-20210622060536-734e95fb86be // indirect
    
    replace github.com/yourname/mymodule v0.0.0-unpublished => /workspaces/GoLangFastPaced/jupyterNotebooks/abdelrauf/mymodule


Alternativ olaraq GONB-in go.mod-a replace etmədən go workspacelə daxili lokal papkaya yönləndirmək olar.   
bunun uchun ```go work``` komandasından istifadə edirik.   



```go
//gobn limitiation-a gore mecbur fayla yazb ordan oxuyacagiq
!echo ${PWD} >/tmp/curr_dir
!*rm -f go.work && CURR_DIR=`cat /tmp/curr_dir`   && go work init && go work use . "${CURR_DIR}/mymodule"
//ignore
%goworkfix
%track
```

    	- replace rule for module "github.com/yourname/mymodule" to local directory "/workspaces/GoLangFastPaced/jupyterNotebooks/abdelrauf/mymodule" already exists.


<b>List of files/directories being tracked:</b>
<ul>
<li>/workspaces/GoLangFastPaced/jupyterNotebooks/abdelrauf/mymodule</li>
<li>/workspaces/GoLangFastPaced/jupyterNotebooks/abdelrauf/mymodule/package1</li>
</ul>



export etmək üçün funksiyanın ilk hərfini böyük işarə etdik. package1 papkasında və eyni qruplaşdırılan fayllar bir package-ə aid olacaq.


```go
import (
	"github.com/yourname/mymodule"
	"fmt"
)

%%
package1.Salam("BranchAndLink")
package1.Salam2("GO FAST PACED")
```

    init call
    init call ###
    BranchAndLink
       ____    ___      _____      _      ____    _____     ____       _       ____   _____   ____
      / ___|  / _ \    |  ___|    / \    / ___|  |_   _|   |  _ \     / \     / ___| | ____| |  _ \
     | |  _  | | | |   | |_      / _ \   \___ \    | |     | |_) |   / _ \   | |     |  _|   | | | |
     | |_| | | |_| |   |  _|    / ___ \   ___) |   | |     |  __/   / ___ \  | |___  | |___  | |_| |
      \____|  \___/    |_|     /_/   \_\ |____/    |_|     |_|     /_/   \_\  \____| |_____| |____/


#### Program   
tam icra edilə bilən proqram üçün package main və func main təyin edilməlidir.  
Proqram üçün bu giriş başlanğıc nöqtəsi olacaq və yalnız bir dənə main olmalıdır   

```Go
package main


func main(){
    
}

```

İcra zamanı package inisializasiya olur və sonra main funksiyası çağrılır.   
Bu funksiyanın çağırışı qayıtdıqda, proqram çıxır. Bu zaman o (əsas olmayan) goroutinlərin tamamlanmasını gözləmir.

Daha ətraflı:   
https://go.dev/ref/mod   
https://go.dev/doc/tutorial/workspaces
