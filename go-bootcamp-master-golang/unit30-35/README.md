# Unit 30-35

# nil value

nil 為可預先聲明的標識符，代表 沒有 、尚未定義、尚未初始化等。

nil 沒有特定值，因為本身就是一個值，代表沒有值的特殊值。

在 Go ，尚未初始化的變數會有一個零值，基於指針的類型 pointers 、slices、maps、interfaces、channels 的零值皆為 nil。

賦新值可以是無 type 型別或有型別的，且廣泛運用了 nil 作為錯誤處理，

### 用 nil 判斷 function 無錯誤

當不該返回值的 functions 返回了值 ，代表程式發生了錯誤，

而如果返回了 nil 代表程式一切正常進行

```go
err := do()
if err != nil {
  // 不該有返回值，進行錯誤處理並終止
}

// 返回值為 nil ，程式正確執行
```

## Error Value

strconv.Itoa 可以將數值轉換為字串

```go
func Itoa(i int) string
```

strconv.Atoi 可以將字串轉換為數值

```go
func Atoi(s string) (int, error) // 會返回一個錯誤類型的值

// success => nil    (未初始化錯誤值)
// error   => no-nil (已初始化錯誤值)
```

Itoa 必定會成功，不須做錯誤處理，但是 Atoi 有時候會失敗，因此要做錯誤處理。

### 打印錯誤

```go
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	n, err := strconv.Atoi(os.Args[1])
	fmt.Println("Converted number     :", n)
	fmt.Println("Returnbed error value:", err)
}

// go run main.go 42
// Converted number     : 42
// Returnbed error value: <nil>

// go run main.go hahaha
// Converted number     : 0
// Returnbed error value: strconv.Atoi: parsing "hahaha": invalid syntax
```

### 加上錯誤處理

```go
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	age := os.Args[1]
	n, err := strconv.Atoi(age)
	if err != nil { // 判斷是否有錯誤值
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Printf("SUCCESS: Converted %q to %d.\n:", age, n)
}

// go run main.go 26
// SUCCESS: Converted "26" to 26.

// go run main.go a
// ERROR: strconv.Atoi: parsing "a": invalid syntax

```

## 題目：feet to meter

## 自己寫的初始的版本

發現問題：

1.  使用 strconv.Atoi 
2. 使用 %d %f 格式化輸出

```go
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: [feet]")
		return
	}
	a := os.Args[1]
	f, err := strconv.Atoi(a)
	if err != nil {
		fmt.Printf("'%s' is not a number.\n", a)
		return
	}
	m := float64(f) * 0.3048
	fmt.Printf("%d feet is %.2f meters.\n", f, m)
}
```

### 正確解答：

1.  移除題目不需要的判斷
2. 使用 strconv.ParseFloat 字串轉換浮點數
3. 使用 %g 格式化輸出 (根據情況顯示科學計數法或浮點數)

```go
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arg := os.Args[1]
	f, err := strconv.ParseFloat(arg, 64)
	if err != nil {
		fmt.Printf("Error: %q is not a number.\n", arg)
		return
	}
	m := f * 0.3048
	fmt.Printf("%g feet is %g meters.\n", f, m) // %g 根據情況顯示科學計數法或浮點數
}

// go run main.go 100
// 100 feet is 30.48 meters.

// go run main.go hahaha
// Error: "hahaha" is not a number.
```