# Unit26-29

## Tiny Challenge: Validate a single user

### 挑戰：使用者驗證

透過 command line 輸入帳號密碼，

判斷是否為錯誤的使用者，

判斷是否為錯誤的密碼，

成功通過。

### 自己寫的初始版本：

發現問題：

1. 沒有例外處理
2. 沒有簡短命名
3. 沒有用格式化打印
4. 沒有使用 const 儲存常量

```go
package main

import (
	"fmt"
	"os"
)

func main() {

	correct_username := "ivankao"
	correct_password := "password"
	input_username := os.Args[1]
	input_password := os.Args[2]

	if input_username != correct_username {
		fmt.Println("wrong user")
	} else if input_password != correct_password {
		fmt.Println("wrong password")
	} else {
		fmt.Println("access granted!")
	}
}
```

### 解答：

1. 增加例外處理
2. 使用簡短命名
3. 使用格式化打印
4. 使用 const 儲存常量

```go
package main

import (
	"fmt"
	"os"
)

const (
	usage    = "Usage: [username] [password]"
	errUser  = "Access denied for %q\n"
	errPwd   = "Invalid password for %q\n"
	accessOK = "Access granted to %q\n"

	user = "ivankao"
	pass = "password"
)

func main() {

	args := os.Args
	if len(args) != 3 {
		fmt.Println(usage)
		return
	}

	u, p := args[1], args[2]
	if u != user {
		fmt.Printf(errUser, u)
	} else if p != pass {
		fmt.Printf(errPwd, p)
	} else {
		fmt.Printf(accessOK, u)
	}
}
```

### 多位使用者：

```go
package main

import (
	"fmt"
	"os"
)

const (
	usage    = "Usage: [username] [password]"
	errPwd   = "Invalid password for %q\n"
	accessOK = "Access granted to %q\n"

	user, user2 = "ivankao", "ivankao2"
	pass, pass2 = "password", "password2"
)

func main() {

	args := os.Args
	if len(args) != 3 {
		fmt.Println(usage)
		return
	}

	u, p := args[1], args[2]
	if u == user && p == pass {
		fmt.Printf(accessOK, u)
	} else if u == user2 && p == pass2 {
		fmt.Printf(accessOK, u)
	} else {
		fmt.Printf(errPwd, u)
	}
}
```

bonus: 使用 map 判斷

```go
package main

import (
	"fmt"
	"os"
)

const (
	usage    = "Usage: [username] [password]"
	errUser  = "Access denied for %q\n"
	errPwd   = "Invalid password for %q\n"
	accessOK = "Access granted to %q\n"
)

func main() {
	m := map[string]string{
		"ivankao":  "password",
		"ivankao2": "password2",
	}
	args := os.Args
	if len(args) != 3 {
		fmt.Println(usage)
		return
	}

	u, p := args[1], args[2]
	if _, ok := m[u]; !ok {
		fmt.Printf(errUser, u)
	} else if p == m[u] {
		fmt.Printf(accessOK, u)
	} else {
		fmt.Printf(errPwd, u)
	}
}
```