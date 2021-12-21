# Unit 16-21

# Printf

打印格式化輸出

### Printf

```
          格式化文本
              |         
fmt.Printf("%q\n", brand)
             |        |
            動詞    替換值
```

格式化文本：決定用何種方式打印

替換值：替換格式文本中的動詞

動詞：被傳遞的值替換

\n：為跳脫字元，打印新的一行

### Println vs Printf

```go
package main

import (
	"fmt"
)

func main() {
	var (
		ops  = 2350
		ok   = 543
		fail = 433
	)
	fmt.Println("total:", ops, "succuss", ok, "/", fail)
	fmt.Printf(
		"total: %d success %d / %d\n", ops, ok, fail,
	)
}
```

Printf 更容易閱讀與運作，講者較喜歡使用 Printf

### 跳脫字元

```go
func main() {
  fmt.Println("hi\nhi")
  fmt.Println("hi\\n\"hi\"")
}
```

當字串開頭為反斜線代表這是一個跳脫字元，整組字元只占一個字串的長度

要打印反斜線 可以使用雙反斜線 \\

要打印雙引號 可以使用反斜線+雙引號  \”

## 打印類型

%T 打印”值”的類型

```go
package main

import "fmt"

func main() {
	var speed int
	var heat float64
	var off bool
	var brand string

	fmt.Printf("%T\n", speed)
	fmt.Printf("%T\n", heat)
	fmt.Printf("%T\n", off)
	fmt.Printf("%T\n", brand)
}

// int
// float64
// bool
// string
```

%v 打印出各種類型的值

%[index]verb 依據索引打印參數

```go
package main

import "fmt"

func main() {
	var (
		planet   = "venus"
		distance = 261
		orbital  = 224.701
		hasLife  = false
	)

	fmt.Printf("Planet: %v\n", planet)
	fmt.Printf("Distance: %v millons kms\n", distance)
	fmt.Printf("Orbital Period: %v days\n", orbital)
	fmt.Printf("Does %v have life? %v\n", planet, hasLife)
	fmt.Printf(
		"%v is %v away. Think! %[2]v kms! %[1]v OMG!\n",
		planet, distance,
	)

  fmt.Printf("Type of %d is %[1]T\n", 3)

}

// Planet: venus
// Distance: 261 millons kms
// Orbital Period: 224.701000 days
// Does venus have life? false
// venus is 261 away. Think! 261 kms! venus OMG!
// Type of 3 is int
```

### 安全類型

Go 會比較值的類型和它對應的動詞

如果類型不匹配會有警告提示，可以防止意外錯誤

建議多使用安全類型，少用 %v 

%s  打印字串

%d 打印十進位數值

%f  打印十進位數值

%t 打印布林值

```go
package main

import "fmt"

func main() {
	var (
		planet   = "venus"
		distance = 261
		orbital  = 224.701
		hasLife  = false
	)

	fmt.Printf("Planet: %s\n", planet)
	fmt.Printf("Distance: %d millons kms\n", distance)
	fmt.Printf("Orbital Period: %f days\n", orbital)
	fmt.Printf("Does %s have life? %t\n", planet, hasLife)
}

// Planet: venus
// Distance: 261 millons kms
// Orbital Period: 224.701000 days
// Does venus have life? false
```

### 小數點打印

在 %f 調整顯示小數點位數

```go
package main

import "fmt"

func main() {
	var orbital = 224.701

	fmt.Printf("Orbital Period: %f days\n", orbital)
	fmt.Printf("Orbital Period: %.0f days\n", orbital)
	fmt.Printf("Orbital Period: %.1f days\n", orbital)
	fmt.Printf("Orbital Period: %.2f days\n", orbital)
	fmt.Printf("Orbital Period: %.3f days\n", orbital)
	fmt.Printf("Orbital Period: %.4f days\n", orbital)
	fmt.Printf("Orbital Period: %.5f days\n", orbital)
}

// Orbital Period: 224.701000 days
// Orbital Period: 225 days
// Orbital Period: 224.7 days
// Orbital Period: 224.70 days
// Orbital Period: 224.701 days
// Orbital Period: 224.7010 days
// Orbital Period: 224.70100 days
```