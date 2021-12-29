# Unit 22-25

# If Statement

### 範例、當分數大於3的時候 print “good”

```go
if score > 3 {  // 這裡的表達式會是一個 bool 值
    fmt.Println("good")
}
```

在 if 區塊中的聲明無法被外部區塊訪問

### 布林判斷條件

用 && 兩個 bool 值都要成立

可以直接使用 bool 判斷該值是否為 true ，不須比較

```go
if score > 3 && valid == true {  // 用 && 兩個 bool 值都要成立
    fmt.Println("good")
}

if score > 3 && valid {  // 使用 bool 判斷 valid 是否為 true
    fmt.Println("good")
}

if score > 3 && !valid {  // 使用 bool 判斷 valid 是否為 false
    fmt.Println("good")
}
```

使用 else if 分支判斷

```go
package main

import "fmt"

func main() {
	score := 2
	if score > 3 {  // if 第一次判斷，只能使用一次
		fmt.Println("good")
	} else if score == 3 {  // else if 可有多次分支判斷
		fmt.Println("on the edge")
	} else if score == 2 {
		fmt.Println("meh...")
	} else {        // else 在所有分支皆為 false 時判斷
		fmt.Println("low")
	}
}
```