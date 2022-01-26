# Unit 51-58

## FOR 語句

### FOR 語句

```go
  初始化語句  狀態
      |        |  
for i := 1; i <= 5; i++ {
  sum += 1            |
}                   發送語句
```

初始化語句：宣告初始化值，只會執行一次

狀態：為布林表達式，為 true 時進入迴圈

發送語句：每次迴圈結束時執行

### 縮短 FOR 語句

若已初始化值，可省略初始化語句

可將發送語句的放進迴圈中，省略發送語句

```go
i := 1
for i <= 5 {
  sum += 1
  i++
}            
```

### BREAK 語句

無窮迴圈

沒有設定狀態，預設為 true

```go
i := 1
for {
  i++
}    
```

使用 break 中止迴圈

```go
package main

import (
	"fmt"
)

func main() {
	var (
		sum int
		i   = 1
	)
	for {
		if i > 5 {
			break  // 當 i 大於 5 時中止迴圈
		}
		sum += i
		fmt.Println(i, "-->", sum)
		i++
	}
	fmt.Println(sum)
}
```

### CONTINUE 語句

continue結束當次迴圈，值接進入下一輪迴圈

```go
package main

import (
	"fmt"
)

func main() {
	var (
		sum int
		i   = 1
	)

	for {
		if i > 10 {
			break
		}

		if i%2 != 0 {
			i++       // 調整 i ，避免無線迴圈
			continue  // 結束當次迴圈
		}

		sum += i
		fmt.Println(i, "-->", sum)
		i++
	}
	fmt.Println(sum)
}
```

## 嵌套迴圈

在迴圈裡面再次使用迴圈

```go
package main

import (
	"fmt"
)

const max = 5

func main() {
	fmt.Printf("%5s", "X")
	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)

		for j := 0; j <= max; j++ {
			fmt.Printf("%5d", i*j)
		}

		fmt.Println()
	}
}

// go run main.go 
//     X    0    1    2    3    4    5
//     0    0    0    0    0    0    0
//     1    0    1    2    3    4    5
//     2    0    2    4    6    8   10
//     3    0    3    6    9   12   15
//     4    0    4    8   12   16   20
//     5    0    5   10   15   20   25
```

**挑戰題，動態乘法表**

將傳入值判斷後轉為數值進行迴圈

```go
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[:]
	if len(args) != 2 {
		fmt.Println("Usage: [nummber]")
		return
	}
	max, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}

	fmt.Printf("%5s", "X")
	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)
	}
	fmt.Println()

	for i := 0; i <= max; i++ {
		fmt.Printf("%5d", i)

		for j := 0; j <= max; j++ {
			fmt.Printf("%5d", i*j)
		}

		fmt.Println()
	}
}
```

### FOR Slice(切片)

一般迴圈打字字串

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("%q\n", os.Args[i])
	}
}

// go run main.go hi hey hello
// "hi"
// "hey"
// "hello"
```

使用 strings.Fields，將 string 轉為 slice string

```go
Fields func (s string) []string
```

範例、

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	words := strings.Fields(
		"lazy car jumps again and again and again",
	)
	for j := 0; j < len(words); j++ {
		fmt.Printf("#%-2d: %q\n", j+1, words[j])
	}
}

// go run main.go
// #1 : "lazy"
// #2 : "car"
// #3 : "jumps"
// #4 : "again"
// #5 : "and"
// #6 : "again"
// #7 : "and"
// #8 : "again"
```

### FOR Range

```go
index value  range keyword
    |  |       |
for i, v := range words {
    fmt.Println(i, v) | 
}                  range expression
```

只需 value 時， index 要設為空白標識符

只需 index 時，可移除 value

range keyword 遞歸返回下一個索引與值

range 表達式可以為 array, slice, string, map, channel 類型

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	words := strings.Fields(
		"lazy car jumps again and again and again",
	)
	for i, v := range words {
		fmt.Printf("#%-2d: %q\n", i+1, v)
	}
}

// go run main1.go
// #1 : "lazy"
// #2 : "car"
// #3 : "jumps"
// #4 : "again"
// #5 : "and"
// #6 : "again"
// #7 : "and"
// #8 : "again"
```