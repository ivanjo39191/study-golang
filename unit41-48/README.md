# Unit 41-48

## Switch 陳述句

switch 條件式會比較所有 case 條件

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	city := os.Args[1]

	switch city {
	case "Paris": // case 語句1
		fmt.Println("France") // case 條件表達式
		vip := true           //在 case 中宣告，只能在當前的 case 中使用
		fmt.Println("VIP trip?", vip)
	case "Tokyo": // case 語句2
		fmt.Println("Japan") // case 條件表達式
	}

	if city == "Paris" {  // 使用 if statements 處理相同判斷
		fmt.Println("France")
	} else if city == "Tokyo" {
		fmt.Println("Japan")
	}
}

// go run main.go Tokyo
// Japan
```

### Default 語句

若找不到符合判斷的語句，就會執行 default 語句。

deaflut 語句只能有一條。

```go
func main() {
	city := os.Args[1]

	switch city {
	case "Paris":
		fmt.Println("France")
	case "Tokyo":
		fmt.Println("Japan")
	default:
		fmt.Println("Where?")  // 預設語句
	}
}

// go run main.go aaa
// Where?
```

### 多重條件

使用 , 來新增 or 條件

```go
func main() {
	city := os.Args[1]

	switch city {
	case "Paris", "Lyon": // 多重條件
		fmt.Println("France")
	case "Tokyo":
		fmt.Println("Japan")
	default:
		fmt.Println("Where?")
	}

	if city == "Paris" || city == "Lyon" {
		fmt.Println("France")
	} else if city == "Tokyo" {
		fmt.Println("Japan")
	} else {
		fmt.Println("Where?")
	}
}

// go run main.go Lyon
// France
// France
```

### 布林表達式

直接使用 switch 語句，預設為 True

當符合判斷條件，就執行語句

switch case 由上往下執行，先符合先執行。

```go
package main

import "fmt"

func main() {
	i := -10
	switch { // 等同 switch True {
	case i > 0:
		fmt.Println("positive")
	case i < 0:
		fmt.Println("negative")
	default:
		fmt.Println("zero")
	}
}

// go run main.go 
// negative
```

### Fallthourgh

case 結尾加上 fallthrough 讓 switch 不中斷且按順序繼續進行 case。

```go
func main() {
	i := 142
	switch {
	case i > 100:
		fmt.Print("big ")
		fallthrough // 不中斷，繼續往下個 case 執行
	case i > 0:
		fmt.Print("positive ")
		fallthrough // 不中斷，繼續往下個 case 執行
	default:
		fmt.Println("number")
	}
}

// go run main.go 
// big positive number
```

### 簡短 Switch

```go
package main

import "fmt"

func main() {

	switch i := 10; { // switch i := 10; true { ，預設 true 可以不寫
	case i > 0:
		fmt.Println("positive")
	case i < 0:
		fmt.Println("negative")
	default:
		fmt.Println("zero")
	}
}

// go run main.go 
// positive
```

### 挑戰題：PARTS of a DAY

### time package

```go
time.Now() // 回傳當前時間為 Time instance

.Hour() // 從 Time instance 取得 0-23 範圍的數值

```

### 自己的版本

問題:

1.  沒有利用由上而下判斷的特性
2. 多餘的變數宣告

```go
package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Now()
	switch h := t.Hour(); {
	case 6 <= h && h < 11:
		fmt.Println("good morning")
	case 11 <= h && h < 14:
		fmt.Println("good afternoon")
	case 14 <= h && h < 22:
		fmt.Println("good evening")
	case (22 <= h && h < 24) || (0 <= h && h < 6):
		fmt.Println("good night")
	}
}

// sudo timedatectl set-time "06:00:00"
// date 
// Wed 05 Jan 2022 06:00:02 AM UTC
// go run main.go 
// good morning

// sudo timedatectl set-time "11:00:00"
// date
// Wed 05 Jan 2022 11:00:03 AM UTC
// go run main.go 
// good afternoon

// sudo timedatectl set-time "14:00:00"
// date
// Wed 05 Jan 2022 02:00:01 PM UTC
// go run main.go 
// good evening

// sudo timedatectl set-time "22:00:00"
// date
// Wed 05 Jan 2022 10:00:01 PM UTC
// go run main.go 
// good night

// sudo timedatectl set-time "02:00:00"
// date
// Wed 05 Jan 2022 02:00:01 AM UTC
// good night
```

### 正確答案

1.  使用由上而下判斷的特性，從大至小判斷來減少過多條件

```go
package main

import (
	"fmt"
	"time"
)

func main() {

	switch h := time.Now().Hour(); {
	case h >= 18:
		fmt.Println("good evening")
	case h >= 12:
		fmt.Println("good afternoon")
	case h >= 6:
		fmt.Println("good morning")
	default:
		fmt.Println("good night")
	}
}
```