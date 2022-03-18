# Unit 168-180 Fuction

## Functions 函數基礎

```go
     (name type, name type)
               |
func name (inputParams) (resultParams) {
	// your code                |
}                        (type, type)

```

每個函數都有 func 關鍵字

每個包層級的函數都有名稱，init, main 為保留字

可以宣告 0 個或多個參數值

return 語句使函數立即停止並結束

可以宣告 0 個或多個返回值，return 語句的返回值須與此宣告相匹配

當返回值只有一個時，不須括號

### Naked Return

當調用 return 語句時，會自動返回預先定義的返回值

```go
func parse(p parser, line string) (parsed result, err error){
  var parsed result
  var err error

  return
}
```

### 將變量限制在函數中

避免包層級的變量互相影響，盡可能的將變量本地化並限制在函數中。

### 參數值與返回值都是用複製的方式傳遞

當你改變一個函數中的切片值時，你需要返回值。

以下的範例 incrWrong 只會將複製後的 n 在函數中進行增加，所以不會影響到 main 中的 local 變量。

```go
package main

import "fmt"

func main() {
	local := 10
	show(local)
	incrWrong(local)
	fmt.Printf("show -> local = %d\n", local)
}

func show(n int) {
	fmt.Printf("show -> n = %d\n", n)
}

func incrWrong(n int) {
	n++
}

####################################
show -> n = 10
show -> local = 10
```

調用的函數要回傳返回值，且儲存回 main 中的 local 變量

```go
package main

import "fmt"

func main() {
	local := 10
	show(local)
	incrWrong(local)
	fmt.Printf("show -> local = %d\n", local)
	local = incr(local)
	show(local)

}

func show(n int) {
	fmt.Printf("show -> n = %d\n", n)
}

func incrWrong(n int) {
	n++
}

func incr(n int) int {
	n++
	return n
}

####################################
show -> n = 10
show -> local = 10
show -> local = 11
```

### 字典值為指針

map 為指針，傳遞時不會進行複製，會直接更新。