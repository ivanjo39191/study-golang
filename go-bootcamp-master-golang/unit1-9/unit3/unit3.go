package main

import (
	"fmt"
	"path"
)

// path.Split()
// input: path
// output: dir, file 皆為 string 時可以一起定義
// func Split(path string) (dir, file string)

func main() {
	var dir, file string
	dir, file = path.Split("css/main.css")
	fmt.Println("dir:", dir)
	fmt.Println("file:", file)
}
