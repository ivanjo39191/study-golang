# Unit181

## Function 函數

```go
func printBook(b book){
    ...       b.title, b.princ
}

printBook(mobydick)
```

運行時會將 mobydick 的副本傳送到 printBook 函數

### Method 方法

```go
func (b book) printBook(){
    ...       b.title, b.princ
}

mobydick.printBook()
```

(b book) 為一個接收參數，b 為名稱，book 為類型，

運行時會將 mobydick 作為隱藏參數傳遞給 printBook 方法。

## Funtion 與 Method 特性差異

Method 屬於一個類型，Function 屬於一個包。

在 Method 中每個類型有自己的命名空間，不同類型可以用相同名稱的方法，

在 Function 則無法使用相同的名稱。

可以將一個 Method 視為是 “將接收參數作為第一個參數傳遞的函數”。

## 將 Function 轉換為 Method

### Function

```go
// main.go
package main

func main() {
	mobydick := book{
		title: "mody dick",
		price: 10,
	}
	minecraft := game{
		title: "minecraft",
		price: 20,
	}
	printBook(mobydick)
	printGame(minecraft)
}

// book,go
package main

import "fmt"

type book struct {
	title string
	price float64
}

func printBook(b book) {
	fmt.Printf("%-15s: $%.2f\n", b.title, b.price)
}

//game.go
package main

import "fmt"

type game struct {
	title string
	price float64
}

func printGame(g game) {
	fmt.Printf("%-15s: $%.2f\n", g.title, g.price)
}
```

### Method

```go
// main.go
package main

func main() {
	mobydick := book{
		title: "mody dick",
		price: 10,
	}
	minecraft := game{
		title: "minecraft",
		price: 20,
	}
	mobydick.print()
	minecraft.print()
}

// book,go
package main

import "fmt"

type book struct {
	title string
	price float64
}

func (b book) print() {
	fmt.Printf("%-15s: $%.2f\n", b.title, b.price)
}

//game.go
package main

import "fmt"

type game struct {
	title string
	price float64
}

func (g game) print() {
	fmt.Printf("%-15s: $%.2f\n", g.title, g.price)
}
```

## 另一種調用 Method 的方式

可以將一個 Method 視為是 “將接收參數作為第一個參數傳遞的函數”。

因此對於 method 來說，mobydick.print() 與 book.print(mobydick) 有相同的效果。

```go
package main

func main() {
	var (
		mobydick  = book{title: "mody dick", price: 10}
		minecraft = game{title: "minecraft", price: 20}
		tetris    = game{title: "tetris", price: 5}
	)
	book.print(mobydick)
	game.print(minecraft)
	game.print(tetris)
}
```

## 使用 Method 接收指標

```go
package main

func main() {
	var (
		mobydick  = book{title: "mody dick", price: 10}
		minecraft = game{title: "minecraft", price: 20}
		tetris    = game{title: "tetris", price: 5}
	)
	book.print(mobydick)
	game.print(minecraft)
	game.print(tetris)
}

//game.go
package main

import "fmt"

type game struct {
	title string
	price float64
}

func (g game) print() {
	fmt.Printf("%-15s: $%.2f\n", g.title, g.price)
}

func (g *game) discount(ratio float64) {
	g.price *= (1 - ratio)
}
```

## 將方法附加到任何類型

```go
// main.go
package main

func main() {
	var (
		minecraft = game{title: "minecraft", price: 20}
		tetris    = game{title: "tetris", price: 5}
	)

	var items []*game
	items = append(items, &minecraft, &tetris)

	my := list(items)
	// my = nil
	my.print()
}

// game.go
package main

import "fmt"

type game struct {
	title string
	price money
}

func (g *game) print() {
	fmt.Printf("%-15s: %s\n", g.title, g.price.string())
}

func (g *game) discount(ratio float64) {
	g.price *= money(1 - ratio)
}

// list.go
package main

import "fmt"

type list []*game

func (l list) print() {
	if len(l) == 0 {
		fmt.Println("Sorry. We're waiting for delivery.")
		return
	}

	for _, it := range l {
		it.print()
	}

// money.go
package main

import "fmt"

type money float64

func (m money) string() string {
	// $xx.yy
	return fmt.Sprintf("$%.2f", m)
}
```