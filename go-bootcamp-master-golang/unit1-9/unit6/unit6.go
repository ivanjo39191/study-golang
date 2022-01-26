package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Printf("%#v\n", os.Args) // 查看物件結構的格式化輸出
	fmt.Println("Path: ", os.Args[0])
	fmt.Println("1st argument: ", os.Args[1])
	fmt.Println("2nd argument ", os.Args[2])
	fmt.Println("3rd argument ", os.Args[3])
	fmt.Println("Number of items indide os.Args: ", len(os.Args)) // 返回 slice 的長度

}

// []string{"/tmp/go-build954824368/b001/exe/unit6", "hi", "hellow", "hey"}
// Path:  /tmp/go-build954824368/b001/exe/unit6
// 1st argument:  hi
// 2nd argument  hellow
// 3rd argument  hey
// Number of items indide os.Args:  4
