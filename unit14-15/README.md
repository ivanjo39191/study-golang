# Uint 14-15

# IOTA

IOTA 為 Go 內置的常量生成器，可以生成不斷增加的數字

### 使用 IOTA

一般 const 宣告

```go
package main

import (
	"fmt"
)

func main() {
	const (
		monday    = 1
		tuesday   = 2
		wednesday = 3
		thursday  = 4
		friday    = 5
		saturday  = 6
		sunday    = 7
	)

	fmt.Println(monday, tuesday, wednesday, thursday, friday, saturday, sunday)
}
```

使用 iota 宣告，自動增加數字

```go
package main

import (
	"fmt"
)

func main() {

	const (
		monday = iota + 1
		tuesday
		wednesday
		thursday
		friday
		saturday
		sunday
	)

	fmt.Println(monday, tuesday, wednesday, thursday, friday, saturday, sunday)

}
```

### 使用空白標識符

一般 const 宣告

```go
package main

import (
	"fmt"
)

func main() {
	const (
		EST = -5
		MST = -7
		PST = -8
	)
	fmt.Println(EST, MST, PST)

}
```

使用 iota 宣告，自動增加數字，並用空白標識符跳過不需要的值

```go
package main

import (
	"fmt"
)

func main() {

	const (
		EST = -(5 + iota)
		_
		MST
		PST
	)
	fmt.Println(EST, MST, PST)

}
```