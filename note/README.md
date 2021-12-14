# Golang Note

## 1.  nil

### nil 的零值

按照Go語言規範，任何型別在未初始化時都對應一個零值：布林型別是false，整型是0，字串是""，而指標、函式、interface、slice、channel和map的零值都是nil。

### nil 型別的地址和值大小

`nil`型別的所有值的記憶體佈局始終相同,換一句話說就是：不同型別`nil`的記憶體地址是一樣的。

### nil值比較

1.不同型別的`nil`是不能比較的。

2.同一型別的兩個nil值可能無法比較，因為golang中存在map、slice和函式型別是不可比較型別，它們有一個別稱為不可比擬的型別，所以比較它們的nil亦是非法的。

3.兩nil值可能不相等

### 參考資料：

[https://codertw.com/程式語言/5262/](https://codertw.com/%E7%A8%8B%E5%BC%8F%E8%AA%9E%E8%A8%80/5262/)

## 2. make new 的差別

**new**

new 可以用來初始化泛型，**並且返回儲存位址**。

**make**

`make` 與 `new` 不同，是用來初始化一些特別的型別，像是 channel, map, slice 等等。另外特別要注意的是 **make 並不會回傳指標**，

### 參考資料：

[https://medium.com/d-d-mag/golang-筆記-make-與-new-的差別-68b05c7ce016](https://medium.com/d-d-mag/golang-%E7%AD%86%E8%A8%98-make-%E8%88%87-new-%E7%9A%84%E5%B7%AE%E5%88%A5-68b05c7ce016)

## 3. '...' 語法糖

第一個用法主要是用於函數有多個不定參數的情況，可以接受多個不確定數量的參數。

第二個用法是slice可以被打散進行傳遞。

```go
func test1(args ...string) { //可以接受任意个string参数
    for _, v:= range args{
        fmt.Println(v)
    }
}

func main(){
var strss= []string{
        "qwr",
        "234",
        "yui",
        "cvbc",
    }
    test1(strss...) //切片被打散传入
}
//output:
//qwr
//234
//yui
//cvbc

var strss= []string{
        "qwr",
        "234",
        "yui",

    }
    var strss2= []string{
        "qqq",
        "aaa",
        "zzz",
        "zzz",
    }
strss=append(strss,strss2...) //strss2的元素被打散一个个append进strss
fmt.Println(strss)
//output:[qwr 234 yui qqq aaa zzz zzz]
```

### 參考資料：

[https://juejin.cn/post/6844903873409204232](https://juejin.cn/post/6844903873409204232)

## 4. forloop string

迭代字串

index, 字元 （不使用 index 時使用 _ )

```go
package main

import "fmt"

func main() {
        for _, s := range "Hello, 世界" {
                fmt.Printf("%c\n", s)
        }
}
```

### 參考資料：

[https://stackoverflow.com/questions/18130859/how-can-i-iterate-over-a-string-by-runes-in-go](https://stackoverflow.com/questions/18130859/how-can-i-iterate-over-a-string-by-runes-in-go)

## 5. 迴圈

Golang 沒有 while，只有 for

有三個參數時當 for ，只有一個參數時當 while。

```go
package main
import "fmt"
func main(){
		// 一個參數
    i := 1           // i 預設為 1
    for i <= 10 {     // i <= 10 為 true 才執行
        fmt.Println("hello")
        i = i + 1    // 每次執行時 i 加 1
    }                // 回到 for 那行 (第5行)

		// 三個參數中以 ; 為分隔
    for i := 1; i <= 10; i = i+1 {
        fmt.Printf("%d ", i)
    }
}

```

### 參考資料：

[https://ithelp.ithome.com.tw/articles/10234260](https://ithelp.ithome.com.tw/articles/10234260)

## 6. struct

**定義 struct 型別**

```go
type Person struct {
	FirstName string
	LastName  string
	Age       int
}
```

type 關鍵字引進了新型別。其後是型別（Person）的名稱和關鍵字 struct，表示我們正在定義一個 struct。struct 包含了一個在大括號內的欄位列表。每個欄位都有名稱和型別。

**宣告 struct**

```python
// Declares a variable of type 'Person'
var p Person // All the struct fields are initialized with their zero value
```

上面的程式碼建立了一個類型為 Person 的變數，預設情況下將其設為零。對於 struct，零表示所有欄位均設定為其對應的零值。因此，欄位 FirstName 和 LastName 設定為 ""，且 Age 設定為 0。

**範例**

```python
package main

import "fmt"

// Defining a struct type
type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	// Declaring a variable of a `struct` type
	var p Person // // All the struct fields are initialized with their zero value
	fmt.Println(p)

	// Declaring and initializing a struct using a struct literal
	p1 := Person{"Rajeev", "Singh", 26}
	fmt.Println("Person1: ", p1)

	// Naming fields while initializing a struct
	p2 := Person{
		FirstName: "John",
		LastName:  "Snow",
		Age:       45,
	}
	fmt.Println("Person2: ", p2)

	// Uninitialized fields are set to their corresponding zero-value
	p3 := Person{FirstName: "Robert"}
	fmt.Println("Person3: ", p3)
}
```

```python
# Output
{  0}
Person1:  {Rajeev Singh 26}
Person2:  {John Snow 45}
Person3:  {Robert  0}
```

### 參考資料：

[https://calvertyang.github.io/2019/11/15/golang-structs-tutorial-with-examples/](https://calvertyang.github.io/2019/11/15/golang-structs-tutorial-with-examples/)

## 7. struct 繼承與指標

使用指標作為接收者，傳遞的是一個指向原指標的副本，指標的副本指向的就是原類型的值，修改時也就能影響原類型的值

```go
func main() {
	p:=person{name:"張三"}
	p.modify() 
	fmt.Println(p.String())
}

type person struct {
	name string
}

func (p person) String() string{
	return "the person name is "+p.name
}

func (p person) modify(){ //值接收者，修改無效
	p.name = "李四"
}

func (p *person) modify(){ //值接收者，修改無效
	p.name = "李四"
}
```

### 參考資料：

[https://cyent.github.io/golang/method/overview/](https://cyent.github.io/golang/method/overview/)

[https://www.flysnow.org/2017/03/31/go-in-action-go-method.html](https://www.flysnow.org/2017/03/31/go-in-action-go-method.html)

## 8. Package list

Package list implements a doubly linked list.

To iterate over a list (where l is a *List):

```go
for e := l.Front(); e != nil; e = e.Next() {
	// do something with e.Value
}
```

list 套件相關 func

```go
type Element
	func (e *Element) Next() *Element
	func (e *Element) Prev() *Element
type List
	func New() *List
	func (l *List) Back() *Element
	func (l *List) Front() *Element
	func (l *List) Init() *List
	func (l *List) InsertAfter(v interface{}, mark *Element) *Element
	func (l *List) InsertBefore(v interface{}, mark *Element) *Element
	func (l *List) Len() int
	func (l *List) MoveAfter(e, mark *Element)
	func (l *List) MoveBefore(e, mark *Element)
	func (l *List) MoveToBack(e *Element)
	func (l *List) MoveToFront(e *Element)
	func (l *List) PushBack(v interface{}) *Element
	func (l *List) PushBackList(other *List)
	func (l *List) PushFront(v interface{}) *Element
	func (l *List) PushFrontList(other *List)
	func (l *List) Remove(e *Element) interface{}
```

### 參考資料：

[https://pkg.go.dev/container/list](https://pkg.go.dev/container/list)

## 9. Slice and Array

### 參考資料：

[https://pjchender.dev/golang/slice-and-array/](https://pjchender.dev/golang/slice-and-array/)

## 延伸閱讀

徹底弄清Golang中[]byte與string轉換

[https://www.gushiciku.cn/pl/pDSl/zh-tw](https://www.gushiciku.cn/pl/pDSl/zh-tw)