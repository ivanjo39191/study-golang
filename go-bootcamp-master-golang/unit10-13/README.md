# Go Bootcamp: Master Golang with 1000+ Exercises and Projects 筆記 2

# 字串基礎

## 一般字串與原始字串

### 說明

一般字串使用雙引號 “ 表達，只可表示單行

```go
"hi 你好"
```

原始字串使用 反引號 ` 表達，並可表示多行

```go
`hi
你好!
`
```

使用上述兩者兩者方式的儲存的變數類型皆為 string

### 範例

```go
package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Println(s)
	s = "<html>\n\t<body>\"Hello\"</body>\n</html>" // 在一般字串使用雙引號需加反斜線
	fmt.Println(s)
	s = `
<html>
	<body>"Hello"</body>
</html>
`
	fmt.Println(s)
}
```

原始字串的多行表示讓程式碼更簡潔好看

```go
package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Println("C:\\my\\dir\\file")
	fmt.Println(`C:\my\dir\file`)
}
```

原始字串不需使用雙反斜線來表達反斜線讓程式碼更容易閱讀

## 字串長度

Go 原生的 len 函數可以返回字串所使用的 byte，無法返回真實長度

Go stdlib unicode/utf8 的 RuneCountInString 可以計算字串真實長度

```go
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var name string
	name = "Ivan"
	fmt.Println("name", name, "length is", len(name))

	name = "譽瑋"
	fmt.Println("name", name, "len function is", len(name))
	fmt.Println("name", name, "length is", utf8.RuneCountInString(name))
}

// name Ivan byte is 4
// name 譽瑋 byte is 6
// name 譽瑋 length is 2
```

使用 len 函數無法正確計算非英文字母的長度，只能返回 byte

使用 utf8 的 RuneCountInString 可以返回字串真實長度

## BANGER ( 字串練習題)

get input + bang it

### strings package

strings 提供許多處理字串的實用函數

Repeat 函數：返回 count(數值) 個 s(字串) 組成的新字串

```go
func Repeat(s string, count int) string
```

ToUpper 函數：並返回輸入 s(字串) 的大寫

```go
func ToUpper(s string) string
```

### 題目

1. 使用Go 編寫一支程式，輸入 hey 返回 HEY!!!

```go
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	msg := os.Args[1]
	l := len(msg)
	s := msg + strings.Repeat("!", l)
	s = strings.ToUpper(s)
	fmt.Println(s)

}

// go run main.go hey
// HEY!!!
```

1.  使用Go 編寫一支程式，輸入 hey 返回 !!!HEY!!!，且只可使用一次 Repeat 函數

```go
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	msg := os.Args[1]
	l := len(msg)
	r := strings.Repeat("!", l)
	s := r + msg + r
	fmt.Println(s)

}

// go run main.go hey
// !!!hey!!!
```