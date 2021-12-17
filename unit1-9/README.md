# Go Bootcamp: Master Golang with 1000+ Exercises and Projects 筆記 1

# 環境

## 下載安裝 go (Linux)

[https://go.dev/doc/install](https://go.dev/doc/install)

```bash
# 刪除舊版本 安裝新版本
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.17.5.linux-amd64.tar.gz

# 指定路徑
export PATH=$PATH:/usr/local/go/bin

# 查看版本
go version
```

## 套件管理工具 go mod

透過go.mod檔案辨識目前要執行的專案位置(取代GOPATH設定)。並且go.mod可以紀錄所有使用的相依模組

### 初始化模組

```python
go mod init study-golang
```

### 新增模組

新增模組 unit1.go 檔案

```bash
mkdir unit1
cd unit1
touch unit1.go

```

撰寫模組

```go
package unit1

// return Unit Number
func UnitOne() string {
	return "Unit One"
}
```

### 測試引用

建立 main.go 檔案

```bash
cd ..
touch main.go
```

撰寫程式

```go
package main

import (
	"fmt"
	"study-golang/unit1"
)

func main() {
	fmt.Println(unit1.UnitOne())
}
```

執行

```go
go run main.go
```

返回值

```go
Unit One
```

### 引用外部模組

直接 import 套件來源即可使用

```go
package main

import (
	"fmt"
	"study-golang/unit1"

	"rsc.io/quote"
)

func main() {
	fmt.Println(unit1.UnitOne())
	fmt.Println(quote.Go())
}
```

返回值

```go
Unit One
Don't communicate by sharing memory, share memory by communicating.
```

## 單元測試 _test

建立單元測試檔案，在 VScode 可以使用 _test.go 做為檔名結尾

```go
cd unit1
touch unit1_test
```

撰寫單元測試， import testing 套件

```go
package unit1_test

import (
	"study-golang/unit1"
	"testing"
)

func TestUnitOne(t *testing.T) {
	if unit1.UnitOne() != "Unit One" {
		t.Fatal("Wrong number : (")
	}
}

func TestUnitOne2(t *testing.T) {
	if unit1.UnitOne() != "Unit ONe" {
		t.Fatal("Wrong number : (")
	}
}
```

# 命名規則

## 說明

```go
var speed int
var speed int = 0
// var | name | static type = value with type
```

var 為 short for : "variable"

speed 為 name of variable

int 為 static type of the variable (in GO :every variable and  value has a type)

## 可用命名

```go
// 全小寫英文
var speed int

// 頭尾大寫英文
var SpeeD int

// 底線開頭
var _speed int

// 使用中文
var 速度 int
```

## 不可用命名

```go
// 數字開頭
var 3speed int

// 驚嘆號開頭
var !speed int

// 包含驚嘆號
var spe!ed int

// 使用 var 命名
var var int
```

### 注意事項

1. 變數只可在宣告後使用

## 路徑分隔符 Path Separator

1. Go Standard Library(標準語言庫、stdlib) "path"
2.  多重賦值與返回值
3.  空白標識符定義

### 說明、

path 包提供使用 url path 的字串

Split 為輸入一個字串，返回多值 dir 與 file

```go
// path.Split()
func Split(path string) (dir, file string)
// input: path
// output: dir, file (皆為 string 時可以一起定義)
```

引用 path 多重賦值，並打印

```go
package main

import (
	"fmt"
	"path"
)

func main() {
	var dir, file string
	dir, file = path.Split("css/main.css") // 返回多值
	fmt.Println("dir:", dir)
	fmt.Println("file:", file)
}
```

使用空白標識符(blank identifier)

在 Go 宣告的變數一定要被使用，若不使用就要用空白標識符將其丟棄

```go
package main

import (
	"fmt"
	"path"
)

func main() {
	var file string
	_, file = path.Split("css/main.css")
	fmt.Println("file:", file)
}
```

簡短宣告(short declaration)

```go
package main

import (
	"fmt"
	"path"
)

func main() {
	_, file := path.Split("css/main.css")
	fmt.Println("file:", file)
}
```

## 使用不同宣告的時機

### 使用正常宣告的時機

1.  初始化未知值
    
    ```go
    package main
    
    func main() {
      // score := 0
      var score int
    }
    ```
    
2. 包作用域的變數
    
    ```go
    package main
    // version := 0 
    var version string
    ```
    
3. 群組變量，增加可讀性
    
    ```go
    package main
    
    import (
    	"fmt"
    )
    
    func main() {
    	var (
    		// related:
    		video string
    
    		// colsely related:
    		duration int
    		current  int
    	)
    }
    ```
    

### 使用簡短宣告的時機

 Gophers 更喜歡簡潔的程式碼，更傾向使用 short declaration

1. 當知道初始化值的時候
    
    ```go
    package main
    
    func main() {
      // var width, height = 100, 50
      width, height := 100, 50
    }
    ```
    
2. 保持簡潔程式碼
3. 重新宣告

## 類型轉換

1.  轉換說明
    
    ```go
    type(value)
    
    // type  要改變的類型
    // value 要被改變類型的值
    ```
    
2.  無法使用相異類型的值進行計算
    
    ```go
    package main
    
    func main() {
    	speed := 100 // speed is int
    	force := 2.5 // force is float64
    	
    	speed = speed * force
    }
    
    // invalid operation: speed * force (mismatched types int and float64)
    ```
    
3.  float 轉 int 會丟失小數位
    
    將 int 轉為 float 計算後再轉回 int
    
    ```go
    package main
    
    import "fmt"
    
    func main() {
    	speed := 100 // speed is int
    	force := 2.5 // force is float64
    	// int(force) 2.5 => 2
    	
    	speed = int(float64(speed) * force)
    	// float64(100) = 100.0
    	// 100.0 * 2.5 = 250.0
    	// int(250.0) = 250
    
    	fmt.Println(speed)
    }
    ```
    

## Greeter! 使用 os.Args 的基礎知識

從 command line 取得輸入

### package os

透過控制 os 相關操作行為

### os.Args 取得輸入參數

```go
var Args []string
// []string is a "slice of strings"
```

slice 可以儲存多個 string 的值

每個儲存其中的值為 未命名變數

### 索引表達式

```bash
go run main.go hi yo

# args[0] => 程式路徑 ...exe/main
# args[1] => 第一個參數 "hi"
# args[2] => 第二個參數 "yo"
```

### 範例

go run

```go
package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Printf("%#v\n", os.Args) // 查看物件結構的格式化輸出
	fmt.Println("Path: ", os.Args[0])
	fmt.Println("1st argument: ", os.Args[1])
	fmt.Println("2nd argument ", os.Args[2])
	fmt.Println("3rd argument ", os.Args[3])
	fmt.Println("Number of items indide os.Args: ", len(os.Args)) // 返回 slice 的長度

}

// []string{"/tmp/go-build954824368/b001/exe/unit6", "hi", "hellow", "hey"}
// Path:  /tmp/go-build954824368/b001/exe/unit6
// 1st argument:  hi
// 2nd argument  hellow
// 3rd argument  hey
// Number of items indide os.Args:  4
```

go build

```bash
go build -o greeter
./greeter hi hellow hey

# []string{"./greeter", "hi", "hellow", "hey"}
# Path:  ./greeter
# 1st argument:  hi
# 2nd argument  hellow
# 3rd argument  hey
# Number of items indide os.Args:  4
```

## 慣用命名方式

範例一、非慣用命名方式

```go
func Read(buffer *Buffer, inBuffer []byte) (size int, err error) {
    if buffer.empty() {
        buffer.Reset()
    }
    size = copy(
 inBuffer,
 buffer.buffer[buffer.offset:])
    buffer.offset += size
    return size, nil
}
```

範例二、慣用命名方式

```go
func Read(b *Buffer, p []byte) (n int, err error) {
    if b.empty() {
        b.Reset()
    }
    n = copy(p, b.buf[b.off:])
    b.off += n
    return n, nil
}

//  b  為 buffer
//  n  為 bytes 讀取的字節數
// off 為 offset，是 buffer type 的變量
```

閱讀越多 Go code，能越容易寫出慣用寫法

### 常用縮寫

```go
var buf []byte // buffer
var off int // offset
var op int // operation
var opRead int // read operation
var l int // length
var n int // number or number of
var m int // another number
var c int // capacity
var c int // character
var a int // array
var r rune // rune
var sep string // separator
var src int // source
var dst int // destination
var b byte // byte
var b []byte // buffer
var buf []byte //buffer
var w io.Writer // writer
var r io.Reader // reader
var pos int // position
```

### 命名的重要性

程式碼可讀性 = 日後的可維護性

### 使用首字母縮寫

```go
// flag value
var fv string

```

### 在小作用域時使用縮寫命名

```go
// number of bytes read
var bytesRead int // X
var n int // O
```

### 在大作用域時使用完整命名

```go
package file

var fileClosed bool
```

### 使用駝峰命名法

1.  第一個單字小寫開頭，其餘單字大寫開頭 (不可導出)
2.  全部單字大寫開頭 ( 可導出)

```go
mixedCaps or MixedCaps
```

### 使用所有首字母縮寫

```go
// API 為 Application Programming Interface 的縮寫

var localApi string  // X
var localAPI string  // O

```

### Do not stutter (不口吃)

```go
player.PlayerScore // X
player.Score // O
```

### 不使用下劃線

```go
const MAX_TIME int // X
const MaxTime int // O
var N int // O
```