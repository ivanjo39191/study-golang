# Unit 36-39

## 簡短 IF 陳述句

### 原始 IF 陳述句

```go
func main() {
	n, err := strconv.Atoi("42")
	if err == nil {
		fmt.Println("There is no error, n is", n)
	}
}
```

### 簡短 IF 陳述句

```go
func main() {
	if n, err := strconv.Atoi("42"); err == nil {
		fmt.Println("There is no error, n is", n)
	}
}

// if n, err := strconv.Atoi("42"); 為簡單陳述句
// ; 分隔 簡單陳述句 與 條件表達式
// err == nil 為條件表達式，使用簡單陳述句已宣告的變數

```

### 簡短 IF 陳述句的作用域

在簡短陳述句宣告的變數，只能在陳述句或其分支中使用。

```go
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if a := os.Args; len(a) != 2 {
		// only: a variable
		fmt.Println("Gice me a number.")
	} else if n, err := strconv.Atoi(a[1]); err != nil {
		// only : a, n and err variables
		fmt.Printf("Cannot convert %q\n", a[1])
	} else {
		// all the variables in the if statenment
		fmt.Printf("%s * 2 is %d\n", a[1], n*2)
	}
}
```

### Shadowing

外部區域與 IF 陳述句內部區域有同樣的變數，造成遮蔽混淆

```go
func main() {
	var n int // 宣告
	if a := os.Args; len(a) != 2 {
		// only: a variable
		fmt.Println("Gice me a number.")
	} else if n, err := strconv.Atoi(a[1]); err != nil { // 重新宣告
		// only : a, n and err variables
		fmt.Printf("Cannot convert %q\n", a[1])
	} else {
		// all the variables in the if statenment
		fmt.Printf("%s * 2 is %d\n", a[1], n*2)
	}
	fmt.Printf("n is %d. QQ - you`ve been shadowed ;)\n", n) // 被遮蔽，只看到外部區域的變量
}
```

在外部宣告後，IF 陳述句不再重新宣告而是對已宣告的變數賦值，統一變數減少混淆與遮蔽

```go
func main() {
	var ( // 宣告
		n   int
		err error
	)
	if a := os.Args; len(a) != 2 {
		// only: a variable
		fmt.Println("Gice me a number.")
	} else if n, err = strconv.Atoi(a[1]); err != nil { // 重新賦值
		// only : a, n and err variables
		fmt.Printf("Cannot convert %q\n", a[1])
	} else {
		// all the variables in the if statenment
		n *= 2
		fmt.Printf("%s * 2 is %d\n", a[1], n)
	}
	fmt.Printf("n is %d. QQ - you`ve been shadowed ;)\n", n) // 正常
}
```