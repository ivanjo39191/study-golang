# Unit180 Pointers

## Pointers 指標

指標儲存在計算機內存中的內存地址，可以進行內存地址搜尋和更新值。

```go
var counter byte = 100
P := &counter
V := *P

           6                              100                             100
_______|_______|_______|_______|_______|_______|_______|_______|_______|_______|_______
   1       2       3       4       5       6       7       8       9       10       11

              P := &counter                                         V := *P    

```

### 地址運算符 &

P 的變量儲存的為指標變量，跟普通變量的差異是它儲存順便了內存地址。

指標都會自動對應一個 type。

指針變量 P 的 type 為 *byte，不能再其中存其他類型的內存地址。

在 type 之前添加 * 時，它會變為指標類型。

### 間接運算符 *

指針變量 P 的 儲存的為 counter 變量的地址，就可以用 *P 取得 counter 變量的值。

將 *P 儲存到普通變量 V，此時 V 的 type 為 byte。

在指標值之前添加 * ，它會返回引用的值。

### 流程說明

counter 的值為 100 內存地址為 6 ，

P 變量儲存了 counter 地址並放在內存地址 2

*P 找到變量 P 儲存的內存地址為 6 ，再使用該地址取得值為 100

最後普通變量 V 進行初始化為 100 ，它為 counter 的副本。

### 透過間接運算符更新值

```go
*P = 25
```

可以直接修改 counter 中的值為 25 

### 將指標傳入函數，也可以更改地址中的數值

```go
passPtrVal(&counter)

func PassPtrVal(pn *int){
  *pn++
}
```

### 複合指標

將切片傳遞給參數並不常見，請不要將指標與切片一起使用，以下為範例

```go
func test(list *[]string)
```

map 值本身已經是指標，不須將指標傳遞給 map

map 元素是不可尋址的值，你只能複製地圖元素，獲得副本的地址

### 使用指標的時機

使用指標：要對傳入的值進行更新並返回

不使用指標：只是讀取或是打印值，不對傳入的值進行更新

### 要記住的事項

指標只是另一個可以包含值的內存地址的值

type 前面的 * 表示指標類型

* 前面的指標值獲取指標所指向的值

每次 func 運行時，它都會為其輸入參數和命名結果值（如果有）創建新變量