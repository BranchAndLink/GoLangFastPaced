---
title: " Build Install Test"
author: "abdelrauf"
date: 2023-07-29
tags: ["Build", "Install", "Test", "Go" ]
---

#### Build Install Test


```go
!rm -r example
!mkdir -p example
!cd example && go mod init example.com/hello
```

    rm: cannot remove 'example': No such file or directory
    exit status 1
    go: creating new go.mod: module example.com/hello



```go

!echo -e "package main\n \
import \"fmt\"\n  \
func main()  {\n\
\tfmt.Println( \"Salam\" )\n}" >example/hello.go
```


```go
!cd example && go fmt .
```

    hello.go


Proqramımızı qovluğun içindən birbaşa run da edə bilərik ```go run .```


```go
!cd example && go run .
```

    Salam


Proqrami build etmək üçün ```go build .``` istifadə edirik


```go
!cd example && go build .
```

-o option-ı verməklə çıxış faylının adını özümüzdə verə bilərik  
```go build -o exe_path_ad```


```go
!cd example && go build -o myhello
!cd example && ./myhello
```

    Salam


Build zamanı -n option-ı və -x verərək go faylını build prosesinə baxmaq olur.   
-n optionı ilə bütün gedən prosesi görmək olsa da fərqli olaraq build  getmir.


```go
!cd example && go build -n hello.go
```

    
    #
    # command-line-arguments
    #
    
    mkdir -p $WORK/b001/
    cat >$WORK/b001/importcfg << 'EOF' # internal
    # import config
    packagefile fmt=/home/codespace/.cache/go-build/5a/5a85274bbefd6702b95fe9edb05543e0787d85e1a7dd4a638459c0d94279347e-d
    packagefile runtime=/home/codespace/.cache/go-build/47/47475b234b211c63b7b5248c5ffe186380c67b3a013d5e0a9bfc38ff287322ce-d
    EOF
    cd /workspaces/GoLangFastPaced/jupyterNotebooks/abdelrauf/example
    /usr/local/go/pkg/tool/linux_amd64/compile -o $WORK/b001/_pkg_.a -trimpath "$WORK/b001=>" -p main -complete -buildid 0lZJx4XgzTiq9kwogQYm/0lZJx4XgzTiq9kwogQYm -goversion go1.20.5 -c=2 -nolocalimports -importcfg $WORK/b001/importcfg -pack ./hello.go
    /usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b001/_pkg_.a # internal
    cat >$WORK/b001/importcfg.link << 'EOF' # internal
    packagefile command-line-arguments=$WORK/b001/_pkg_.a
    packagefile fmt=/home/codespace/.cache/go-build/5a/5a85274bbefd6702b95fe9edb05543e0787d85e1a7dd4a638459c0d94279347e-d
    packagefile runtime=/home/codespace/.cache/go-build/47/47475b234b211c63b7b5248c5ffe186380c67b3a013d5e0a9bfc38ff287322ce-d
    packagefile errors=/home/codespace/.cache/go-build/16/1614d1e860779bdd7e63b9ace048ed5c7a97b04fbf64c21e3d5d9ed474dc6127-d
    packagefile internal/fmtsort=/home/codespace/.cache/go-build/38/387acd18f92d205195c39a0cd19fccf97740aca13e8ff884e8380ee620a42625-d
    packagefile io=/home/codespace/.cache/go-build/db/db976f4d0ab0631112d3df9d84973615392ec449f518ab2fdd80f012add5e0ea-d
    packagefile math=/home/codespace/.cache/go-build/9b/9b5cfacac4b76a1cd5bffe8ddf1205bd3311b3351271ef0478df6b24a330aac5-d
    packagefile os=/home/codespace/.cache/go-build/54/543e805161b3797fca3138c85e870dae0b983b5820a9d1d27dfb32787dfb17a8-d
    packagefile reflect=/home/codespace/.cache/go-build/ba/ba60b6d4a1e062811c877f7525496d90904df31d574ccddcafd18ab0765fd3e6-d
    packagefile sort=/home/codespace/.cache/go-build/9c/9c22541614fe877a88f3a267fb9a282f736945c545901a4ab92f6c87295594f0-d
    packagefile strconv=/home/codespace/.cache/go-build/61/61af2248c1b10fa37e5d7da74ce63bc2f3fc4b85d71c8a8d2472c6841263dbbc-d
    packagefile sync=/home/codespace/.cache/go-build/af/af0c9e835f263130088c123bb71d9f528ead461fe081b137543a23a2e6880e0c-d
    packagefile unicode/utf8=/home/codespace/.cache/go-build/55/55c78141439967c497ef1d8bcc1f724dd5b1fab6131ce80d6bc7ab4d338ac783-d
    packagefile internal/abi=/home/codespace/.cache/go-build/b9/b9fe5f7abd9676969136c0fda934dee047633039cae77c70d282251609443bce-d
    packagefile internal/bytealg=/home/codespace/.cache/go-build/b8/b87cc310d783a3a27ab80f0165651d417a52848d12f7adf3003d305eb87f3158-d
    packagefile internal/coverage/rtcov=/home/codespace/.cache/go-build/bc/bc4d12e895fb104d560e6b4192e952d2829d8194cf8a274c5f104c35d80b9774-d
    packagefile internal/cpu=/home/codespace/.cache/go-build/ae/ae76a5e84cb76a1bf351f5ab1b823a97e845dfbdea48a8c3bd64010d853b20ee-d
    packagefile internal/goarch=/home/codespace/.cache/go-build/e3/e3dc83e2820ca7ed5dbc544f0181cfd8829e3357e99876f40e5454c7abd3d7f7-d
    packagefile internal/goexperiment=/home/codespace/.cache/go-build/ac/acdb750114f8d21a951104fb87a789b829c55e1269d8e2d5ede8e8109dceb514-d
    packagefile internal/goos=/home/codespace/.cache/go-build/13/13c70de12e7d1cf949413b66cf278a25a60205c9fe2a48884df4cc75bf099dce-d
    packagefile runtime/internal/atomic=/home/codespace/.cache/go-build/02/021d982fd31ceb193751ff630bd0a0708326e65bad78c10d35fe284dc2ddb667-d
    packagefile runtime/internal/math=/home/codespace/.cache/go-build/e9/e9c1b787dae856726480d898196fec4bae10539b338bbb8373d00fe395fbb426-d
    packagefile runtime/internal/sys=/home/codespace/.cache/go-build/f4/f4b5f90f461eb546868335e492567ea93fafc7616321907072ff597df87dd8bc-d
    packagefile runtime/internal/syscall=/home/codespace/.cache/go-build/ba/baf2d32d6593f77b395fed92cf640776faf6fb971842ce63916341639d6d80e1-d
    packagefile internal/reflectlite=/home/codespace/.cache/go-build/01/012b08e178a6c8e813d3c48a5a52d73e938eeec3182f79894f2ba716aae9412c-d
    packagefile math/bits=/home/codespace/.cache/go-build/cb/cb87d7fb3a3d8b81084908cf252115ec4d0345b313c07eb3ee9929df39ac53f0-d
    packagefile internal/itoa=/home/codespace/.cache/go-build/2e/2e1ffdf15257231907ed22bea794e71342406c8066cf719ed7c3c1a295926451-d
    packagefile internal/poll=/home/codespace/.cache/go-build/dd/ddfa9bcc31f046a25a975333486f4ecddec59fb396184cbb6d2761492fc79e18-d
    packagefile internal/safefilepath=/home/codespace/.cache/go-build/f1/f186c3bc25497814c5681b3bbbd5b8e4bbcf455c65dc4a833e9446db8ac538a3-d
    packagefile internal/syscall/execenv=/home/codespace/.cache/go-build/8a/8a956c2b317f1c0d69bc02cbb6c6afb4fa7db43990bb723820e49c4949747892-d
    packagefile internal/syscall/unix=/home/codespace/.cache/go-build/0d/0dd2c0be76cd7c72277a5b8391db4040f85182d4204ac8d43ae635730dc7fa67-d
    packagefile internal/testlog=/home/codespace/.cache/go-build/34/34811ed3b55f7e3febc65f3d800e76be4adc05cfedea8fa873ace7b3092a6c07-d
    packagefile io/fs=/home/codespace/.cache/go-build/31/3138290bc0a2bc085e7839ef4ab7ba901be8b10750897a7e671a23a48a8c1e41-d
    packagefile sync/atomic=/home/codespace/.cache/go-build/5b/5bdd0d873301df7a24b744f11f78b678e44a684b6b4cd83829136aa69cac73f0-d
    packagefile syscall=/home/codespace/.cache/go-build/3a/3a6c5bb22516cd974a1b3a14c01500d92d73d33785b098d6a79de53586937925-d
    packagefile time=/home/codespace/.cache/go-build/7f/7f201b0f381f3a1af27c0a073b9f62698ab1d771b35a678e97f662bbbf1d30e0-d
    packagefile internal/unsafeheader=/home/codespace/.cache/go-build/a5/a5d0da800ef285eedd4592c5c0343f7ca289a8e9611c69d90f42fae894fe4580-d
    packagefile unicode=/home/codespace/.cache/go-build/2b/2be17f8a5ea974fc4bc4e294eb8d59a605226bb7b8ffa829b8417425bb792ff4-d
    packagefile internal/race=/home/codespace/.cache/go-build/09/090603aeb000493ecd31846063b3fc479da8f548f6c94ecd6524ee0d3eee540f-d
    packagefile internal/oserror=/home/codespace/.cache/go-build/25/25e65d202f732dc1c662bddad3dbbc0f733872a3033bf105f7734e562010d3f4-d
    packagefile path=/home/codespace/.cache/go-build/2e/2ef225fea25de35eeec1b70589fa74f75a89f9aedf763fc4e507e81ee24e9e63-d
    modinfo "0w\xaf\f\x92t\b\x02A\xe1\xc1\a\xe6\xd6\x18\xe6path\tcommand-line-arguments\nbuild\t-buildmode=exe\nbuild\t-compiler=gc\nbuild\tCGO_ENABLED=1\nbuild\tCGO_CFLAGS=\nbuild\tCGO_CPPFLAGS=\nbuild\tCGO_CXXFLAGS=\nbuild\tCGO_LDFLAGS=\nbuild\tGOARCH=amd64\nbuild\tGOOS=linux\nbuild\tGOAMD64=v1\n\xf92C1\x86\x18 r\x00\x82B\x10A\x16\xd8\xf2"
    EOF
    mkdir -p $WORK/b001/exe/
    cd .
    /usr/local/go/pkg/tool/linux_amd64/link -o $WORK/b001/exe/a.out -importcfg $WORK/b001/importcfg.link -buildmode=exe -buildid=CUDQaEwoSWoMd-aemZem/0lZJx4XgzTiq9kwogQYm/0lZJx4XgzTiq9kwogQYm/CUDQaEwoSWoMd-aemZem -extld=gcc $WORK/b001/_pkg_.a
    /usr/local/go/pkg/tool/linux_amd64/buildid -w $WORK/b001/exe/a.out # internal
    mv $WORK/b001/exe/a.out hello


build commandası genişdir.   
bununla əlavə kompilyasiya və link və assembly flag-lər (-gcflags=  -ldflags= və -asmflags=) vermək olur.   
başqa arxitektura üçün build etmək mümkündür.     
Sadə halda generasiya olunmuş assembly-ə baxaq (Qeyd cross kompilyasiya və s gələcəkdə keçəcəyik)


```go
!cd example &&  go build -gcflags="-S" hello.go
```

    # command-line-arguments
    main.main STEXT size=103 args=0x0 locals=0x40 funcid=0x0 align=0x0
    	0x0000 00000 (/workspaces/GoLangFastPaced/jupyterNotebooks/abdelrauf/example/hello.go:5)	TEXT	main.main(SB), ABIInternal, $64-0
    	0x0000 00000 (/workspaces/GoLangFastPaced/jupyterNotebooks/abdelrauf/example/hello.go:5)	CMPQ	SP, 16(R14)
    	0x0004 00004 (/workspaces/GoLangFastPaced/jupyterNotebooks/abdelrauf/example/hello.go:5)	PCDATA	$0, $-2
    	0x0004 00004 (/workspaces/GoLangFastPaced/jupyterNotebooks/abdelrauf/example/hello.go:5)	JLS	92
    	0x0006 00006 (/workspaces/GoLangFastPaced/jupyterNotebooks/abdelrauf/example/hello.go:5)	PCDATA	$0, $-1
    	0x0006 00006 (/workspaces/GoLangFastPaced/jupyterNotebooks/abdelrauf/example/hello.go:5)	SUBQ	$64, SP
    	0x000a 00010 (/workspaces/GoLangFastPaced/jupyterNotebooks/abdelrauf/example/hello.go:5)	MOVQ	BP, 56(SP)
    	0x000f 00015 (/workspaces/GoLangFastPaced/jupyterNotebooks/abdelrauf/example/hello.go:5)	LEAQ	56(SP), BP
    	0x0014 00020 (/workspaces/GoLangFastPaced/jupyterNotebooks/abdelrauf/example/hello.go:5)	FUNCDATA	$0, gclocals·g2BeySu+wFnoycgXfElmcg==(SB)
    	0x0014 00020 (/workspaces/GoLangFastPaced/jupyterNotebooks/abdelrauf/example/hello.go:5)	FUNCDATA	$1, gclocals·EaPwxsZ75yY1hHMVZLmk6g==(SB)
    	0x0014 00020 (/workspaces/GoLangFastPaced/jupyterNotebooks/abdelrauf/example/hello.go:5)	FUNCDATA	$2, main.main.stkobj(SB)
    	0x0014 00020 (/workspaces/GoLangFastPaced/jupyterNotebooks/abdelrauf/example/hello.go:6)	MOVUPS	X15, main..autotmp_8+40(SP)
    	0x001a 00026 (/workspaces/GoLangFastPaced/jupyterNotebooks/abdelrauf/example/hello.go:6)	LEAQ	type:string(SB), DX
    	0x0021 00033 (/workspaces/GoLangFastPaced/jupyterNotebooks/abdelrauf/example/hello.go:6)	MOVQ	DX, main..autotmp_8+40(SP)
    	0x0026 00038 (/workspaces/GoLangFastPaced/jupyterNotebooks/abdelrauf/example/hello.go:6)	LEAQ	main..stmp_0(SB), DX
    	0x002d 00045 (/workspaces/GoLangFastPaced/jupyterNotebooks/abdelrauf/example/hello.go:6)	MOVQ	DX, main..autotmp_8+48(SP)
    	0x0032 00050 (<unknown line number>)	NOP
    	0x0032 00050 (/usr/local/go/src/fmt/print.go:314)	MOVQ	os.Stdout(SB), BX
    	0x0039 00057 (/usr/local/go/src/fmt/print.go:314)	LEAQ	go:itab.*os.File,io.Writer(SB), AX
    	0x0040 00064 (/usr/local/go/src/fmt/print.go:314)	LEAQ	main..autotmp_8+40(SP), CX
    	0x0045 00069 (/usr/local/go/src/fmt/print.go:314)	MOVL	$1, DI
    	0x004a 00074 (/usr/local/go/src/fmt/print.go:314)	MOVQ	DI, SI
    	0x004d 00077 (/usr/local/go/src/fmt/print.go:314)	PCDATA	$1, $0
    	0x004d 00077 (/usr/local/go/src/fmt/print.go:314)	CALL	fmt.Fprintln(SB)
    	0x0052 00082 (/workspaces/GoLangFastPaced/jupyterNotebooks/abdelrauf/example/hello.go:7)	MOVQ	56(SP), BP
    	0x0057 00087 (/workspaces/GoLangFastPaced/jupyterNotebooks/abdelrauf/example/hello.go:7)	ADDQ	$64, SP
    	0x005b 00091 (/workspaces/GoLangFastPaced/jupyterNotebooks/abdelrauf/example/hello.go:7)	RET
    	0x005c 00092 (/workspaces/GoLangFastPaced/jupyterNotebooks/abdelrauf/example/hello.go:7)	NOP
    	0x005c 00092 (/workspaces/GoLangFastPaced/jupyterNotebooks/abdelrauf/example/hello.go:5)	PCDATA	$1, $-1
    	0x005c 00092 (/workspaces/GoLangFastPaced/jupyterNotebooks/abdelrauf/example/hello.go:5)	PCDATA	$0, $-2
    	0x005c 00092 (/workspaces/GoLangFastPaced/jupyterNotebooks/abdelrauf/example/hello.go:5)	NOP
    	0x0060 00096 (/workspaces/GoLangFastPaced/jupyterNotebooks/abdelrauf/example/hello.go:5)	CALL	runtime.morestack_noctxt(SB)
    	0x0065 00101 (/workspaces/GoLangFastPaced/jupyterNotebooks/abdelrauf/example/hello.go:5)	PCDATA	$0, $-1
    	0x0065 00101 (/workspaces/GoLangFastPaced/jupyterNotebooks/abdelrauf/example/hello.go:5)	JMP	0
    	0x0000 49 3b 66 10 76 56 48 83 ec 40 48 89 6c 24 38 48  I;f.vVH..@H.l$8H
    	0x0010 8d 6c 24 38 44 0f 11 7c 24 28 48 8d 15 00 00 00  .l$8D..|$(H.....
    	0x0020 00 48 89 54 24 28 48 8d 15 00 00 00 00 48 89 54  .H.T$(H......H.T
    	0x0030 24 30 48 8b 1d 00 00 00 00 48 8d 05 00 00 00 00  $0H......H......
    	0x0040 48 8d 4c 24 28 bf 01 00 00 00 48 89 fe e8 00 00  H.L$(.....H.....
    	0x0050 00 00 48 8b 6c 24 38 48 83 c4 40 c3 0f 1f 40 00  ..H.l$8H..@...@.
    	0x0060 e8 00 00 00 00 eb 99                             .......
    	rel 2+0 t=23 type:string+0
    	rel 2+0 t=23 type:*os.File+0
    	rel 29+4 t=14 type:string+0
    	rel 41+4 t=14 main..stmp_0+0
    	rel 53+4 t=14 os.Stdout+0
    	rel 60+4 t=14 go:itab.*os.File,io.Writer+0
    	rel 78+4 t=7 fmt.Fprintln+0
    	rel 97+4 t=7 runtime.morestack_noctxt+0
    go:cuinfo.producer.main SDWARFCUINFO dupok size=0
    	0x0000 72 65 67 61 62 69                                regabi
    go:cuinfo.packagename.main SDWARFCUINFO dupok size=0
    	0x0000 6d 61 69 6e                                      main
    go:info.fmt.Println$abstract SDWARFABSFCN dupok size=42
    	0x0000 05 66 6d 74 2e 50 72 69 6e 74 6c 6e 00 01 01 13  .fmt.Println....
    	0x0010 61 00 00 00 00 00 00 13 6e 00 01 00 00 00 00 13  a.......n.......
    	0x0020 65 72 72 00 01 00 00 00 00 00                    err.......
    	rel 0+0 t=22 type:[]interface {}+0
    	rel 0+0 t=22 type:error+0
    	rel 0+0 t=22 type:int+0
    	rel 19+4 t=31 go:info.[]interface {}+0
    	rel 27+4 t=31 go:info.int+0
    	rel 37+4 t=31 go:info.error+0
    go:itab.*os.File,io.Writer SRODATA dupok size=32
    	0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
    	0x0010 5a 22 ee 60 00 00 00 00 00 00 00 00 00 00 00 00  Z".`............
    	rel 0+8 t=1 type:io.Writer+0
    	rel 8+8 t=1 type:*os.File+0
    	rel 24+8 t=-32767 os.(*File).Write+0
    main..inittask SNOPTRDATA size=32
    	0x0000 00 00 00 00 00 00 00 00 01 00 00 00 00 00 00 00  ................
    	0x0010 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
    	rel 24+8 t=1 fmt..inittask+0
    go:string."Salam" SRODATA dupok size=5
    	0x0000 53 61 6c 61 6d                                   Salam
    main..stmp_0 SRODATA static size=16
    	0x0000 00 00 00 00 00 00 00 00 05 00 00 00 00 00 00 00  ................
    	rel 0+8 t=1 go:string."Salam"+0
    runtime.nilinterequal·f SRODATA dupok size=8
    	0x0000 00 00 00 00 00 00 00 00                          ........
    	rel 0+8 t=1 runtime.nilinterequal+0
    runtime.memequal64·f SRODATA dupok size=8
    	0x0000 00 00 00 00 00 00 00 00                          ........
    	rel 0+8 t=1 runtime.memequal64+0
    runtime.gcbits.0100000000000000 SRODATA dupok size=8
    	0x0000 01 00 00 00 00 00 00 00                          ........
    type:.namedata.*[1]interface {}- SRODATA dupok size=18
    	0x0000 00 10 2a 5b 31 5d 69 6e 74 65 72 66 61 63 65 20  ..*[1]interface 
    	0x0010 7b 7d                                            {}
    type:*[1]interface {} SRODATA dupok size=56
    	0x0000 08 00 00 00 00 00 00 00 08 00 00 00 00 00 00 00  ................
    	0x0010 a8 0e 57 36 08 08 08 36 00 00 00 00 00 00 00 00  ..W6...6........
    	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
    	0x0030 00 00 00 00 00 00 00 00                          ........
    	rel 24+8 t=1 runtime.memequal64·f+0
    	rel 32+8 t=1 runtime.gcbits.0100000000000000+0
    	rel 40+4 t=5 type:.namedata.*[1]interface {}-+0
    	rel 48+8 t=1 type:[1]interface {}+0
    runtime.gcbits.0200000000000000 SRODATA dupok size=8
    	0x0000 02 00 00 00 00 00 00 00                          ........
    type:[1]interface {} SRODATA dupok size=72
    	0x0000 10 00 00 00 00 00 00 00 10 00 00 00 00 00 00 00  ................
    	0x0010 6e 20 6a 3d 02 08 08 11 00 00 00 00 00 00 00 00  n j=............
    	0x0020 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
    	0x0030 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
    	0x0040 01 00 00 00 00 00 00 00                          ........
    	rel 24+8 t=1 runtime.nilinterequal·f+0
    	rel 32+8 t=1 runtime.gcbits.0200000000000000+0
    	rel 40+4 t=5 type:.namedata.*[1]interface {}-+0
    	rel 44+4 t=-32763 type:*[1]interface {}+0
    	rel 48+8 t=1 type:interface {}+0
    	rel 56+8 t=1 type:[]interface {}+0
    type:.importpath.fmt. SRODATA dupok size=5
    	0x0000 00 03 66 6d 74                                   ..fmt
    gclocals·g2BeySu+wFnoycgXfElmcg== SRODATA dupok size=8
    	0x0000 01 00 00 00 00 00 00 00                          ........
    gclocals·EaPwxsZ75yY1hHMVZLmk6g== SRODATA dupok size=9
    	0x0000 01 00 00 00 02 00 00 00 00                       .........
    main.main.stkobj SRODATA static size=24
    	0x0000 01 00 00 00 00 00 00 00 f0 ff ff ff 10 00 00 00  ................
    	0x0010 10 00 00 00 00 00 00 00                          ........
    	rel 20+4 t=5 runtime.gcbits.0200000000000000+0


GO build komandası manualı


```go
!go help build
```

    usage: go build [-o output] [build flags] [packages]
    
    Build compiles the packages named by the import paths,
    along with their dependencies, but it does not install the results.
    
    If the arguments to build are a list of .go files from a single directory,
    build treats them as a list of source files specifying a single package.
    
    When compiling packages, build ignores files that end in '_test.go'.
    
    When compiling a single main package, build writes
    the resulting executable to an output file named after
    the first source file ('go build ed.go rx.go' writes 'ed' or 'ed.exe')
    or the source code directory ('go build unix/sam' writes 'sam' or 'sam.exe').
    The '.exe' suffix is added when writing a Windows executable.
    
    When compiling multiple packages or a single non-main package,
    build compiles the packages but discards the resulting object,
    serving only as a check that the packages can be built.
    
    The -o flag forces build to write the resulting executable or object
    to the named output file or directory, instead of the default behavior described
    in the last two paragraphs. If the named output is an existing directory or
    ends with a slash or backslash, then any resulting executables
    will be written to that directory.
    
    The build flags are shared by the build, clean, get, install, list, run,
    and test commands:
    
    	-C dir
    		Change to dir before running the command.
    		Any files named on the command line are interpreted after
    		changing directories.
    	-a
    		force rebuilding of packages that are already up-to-date.
    	-n
    		print the commands but do not run them.
    	-p n
    		the number of programs, such as build commands or
    		test binaries, that can be run in parallel.
    		The default is GOMAXPROCS, normally the number of CPUs available.
    	-race
    		enable data race detection.
    		Supported only on linux/amd64, freebsd/amd64, darwin/amd64, darwin/arm64, windows/amd64,
    		linux/ppc64le and linux/arm64 (only for 48-bit VMA).
    	-msan
    		enable interoperation with memory sanitizer.
    		Supported only on linux/amd64, linux/arm64, freebsd/amd64
    		and only with Clang/LLVM as the host C compiler.
    		PIE build mode will be used on all platforms except linux/amd64.
    	-asan
    		enable interoperation with address sanitizer.
    		Supported only on linux/arm64, linux/amd64.
    		Supported only on linux/amd64 or linux/arm64 and only with GCC 7 and higher
    		or Clang/LLVM 9 and higher.
    	-cover
    		enable code coverage instrumentation (requires
    		that GOEXPERIMENT=coverageredesign be set).
    	-coverpkg pattern1,pattern2,pattern3
    		For a build that targets package 'main' (e.g. building a Go
    		executable), apply coverage analysis to each package matching
    		the patterns. The default is to apply coverage analysis to
    		packages in the main Go module. See 'go help packages' for a
    		description of package patterns.  Sets -cover.
    	-v
    		print the names of packages as they are compiled.
    	-work
    		print the name of the temporary work directory and
    		do not delete it when exiting.
    	-x
    		print the commands.
    
    	-asmflags '[pattern=]arg list'
    		arguments to pass on each go tool asm invocation.
    	-buildmode mode
    		build mode to use. See 'go help buildmode' for more.
    	-buildvcs
    		Whether to stamp binaries with version control information
    		("true", "false", or "auto"). By default ("auto"), version control
    		information is stamped into a binary if the main package, the main module
    		containing it, and the current directory are all in the same repository.
    		Use -buildvcs=false to always omit version control information, or
    		-buildvcs=true to error out if version control information is available but
    		cannot be included due to a missing tool or ambiguous directory structure.
    	-compiler name
    		name of compiler to use, as in runtime.Compiler (gccgo or gc).
    	-gccgoflags '[pattern=]arg list'
    		arguments to pass on each gccgo compiler/linker invocation.
    	-gcflags '[pattern=]arg list'
    		arguments to pass on each go tool compile invocation.
    	-installsuffix suffix
    		a suffix to use in the name of the package installation directory,
    		in order to keep output separate from default builds.
    		If using the -race flag, the install suffix is automatically set to race
    		or, if set explicitly, has _race appended to it. Likewise for the -msan
    		and -asan flags. Using a -buildmode option that requires non-default compile
    		flags has a similar effect.
    	-ldflags '[pattern=]arg list'
    		arguments to pass on each go tool link invocation.
    	-linkshared
    		build code that will be linked against shared libraries previously
    		created with -buildmode=shared.
    	-mod mode
    		module download mode to use: readonly, vendor, or mod.
    		By default, if a vendor directory is present and the go version in go.mod
    		is 1.14 or higher, the go command acts as if -mod=vendor were set.
    		Otherwise, the go command acts as if -mod=readonly were set.
    		See https://golang.org/ref/mod#build-commands for details.
    	-modcacherw
    		leave newly-created directories in the module cache read-write
    		instead of making them read-only.
    	-modfile file
    		in module aware mode, read (and possibly write) an alternate go.mod
    		file instead of the one in the module root directory. A file named
    		"go.mod" must still be present in order to determine the module root
    		directory, but it is not accessed. When -modfile is specified, an
    		alternate go.sum file is also used: its path is derived from the
    		-modfile flag by trimming the ".mod" extension and appending ".sum".
    	-overlay file
    		read a JSON config file that provides an overlay for build operations.
    		The file is a JSON struct with a single field, named 'Replace', that
    		maps each disk file path (a string) to its backing file path, so that
    		a build will run as if the disk file path exists with the contents
    		given by the backing file paths, or as if the disk file path does not
    		exist if its backing file path is empty. Support for the -overlay flag
    		has some limitations: importantly, cgo files included from outside the
    		include path must be in the same directory as the Go package they are
    		included from, and overlays will not appear when binaries and tests are
    		run through go run and go test respectively.
    	-pgo file
    		specify the file path of a profile for profile-guided optimization (PGO).
    		Special name "auto" lets the go command select a file named
    		"default.pgo" in the main package's directory if that file exists.
    		Special name "off" turns off PGO.
    	-pkgdir dir
    		install and load all packages from dir instead of the usual locations.
    		For example, when building with a non-standard configuration,
    		use -pkgdir to keep generated packages in a separate location.
    	-tags tag,list
    		a comma-separated list of additional build tags to consider satisfied
    		during the build. For more information about build tags, see
    		'go help buildconstraint'. (Earlier versions of Go used a
    		space-separated list, and that form is deprecated but still recognized.)
    	-trimpath
    		remove all file system paths from the resulting executable.
    		Instead of absolute file system paths, the recorded file names
    		will begin either a module path@version (when using modules),
    		or a plain import path (when using the standard library, or GOPATH).
    	-toolexec 'cmd args'
    		a program to use to invoke toolchain programs like vet and asm.
    		For example, instead of running asm, the go command will run
    		'cmd args /path/to/asm <arguments for asm>'.
    		The TOOLEXEC_IMPORTPATH environment variable will be set,
    		matching 'go list -f {{.ImportPath}}' for the package being built.
    
    The -asmflags, -gccgoflags, -gcflags, and -ldflags flags accept a
    space-separated list of arguments to pass to an underlying tool
    during the build. To embed spaces in an element in the list, surround
    it with either single or double quotes. The argument list may be
    preceded by a package pattern and an equal sign, which restricts
    the use of that argument list to the building of packages matching
    that pattern (see 'go help packages' for a description of package
    patterns). Without a pattern, the argument list applies only to the
    packages named on the command line. The flags may be repeated
    with different patterns in order to specify different arguments for
    different sets of packages. If a package matches patterns given in
    multiple flags, the latest match on the command line wins.
    For example, 'go build -gcflags=-S fmt' prints the disassembly
    only for package fmt, while 'go build -gcflags=all=-S fmt'
    prints the disassembly for fmt and all its dependencies.
    
    For more about specifying packages, see 'go help packages'.
    For more about where packages and binaries are installed,
    run 'go help gopath'.
    For more about calling between Go and C/C++, run 'go help c'.
    
    Note: Build adheres to certain conventions such as those described
    by 'go help gopath'. Not all projects can follow these conventions,
    however. Installations that have their own conventions or that use
    a separate software build system may choose to use lower-level
    invocations such as 'go tool compile' and 'go tool link' to avoid
    some of the overheads and design decisions of the build tool.
    
    See also: go install, go get, go clean.


```go install``` $GOPATH path-da bin alt qovluğuna icra proqramlarını install edir


```go
!echo ${GOPATH}
!cd example && go install
!ls ${GOPATH}/bin/* | grep hello
```

    /go
    /go/bin/hello


Go da test fayllar-da düzəltmək olur . Test faylların adı konvensiya olaraq **_test.go** ilə bitməlidir.   
```go build``` əmrində bu cür fayllar ignor olacaq
Bu faylları ```go test```   ilə icra etmək olar   
və ya ```go test -c``` ilə test icra faylını yaratmaq olar.  



```go
!echo -e "package main\n \
import \"testing\"\n  \
func simpleTest(t *testing.T)  {\n\
\tres:=2+2\nif res!=4 {\n  t.Errorf(\"%d olmalidir\",4)\n}\n}" >example/simple_test.go
!cd example && go fmt . && go mod tidy
```

    simple_test.go



```go
!cd example && go test -v
```

    testing: warning: no tests to run
    PASS
    ok  	example.com/hello	0.005s



Bu məsələni **tests and bench** bölməsində təzədən keçəcəyik
