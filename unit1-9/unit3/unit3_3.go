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
	_, file := path.Split("css/main.css")
	fmt.Println("file:", file)
}
