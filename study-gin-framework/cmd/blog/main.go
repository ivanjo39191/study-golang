package main

import (
	"blog/internal/routers"
	"os"
)

const version = "v0.1.0"

func main() {
	os.Setenv("TZ", "Asia/Taipei")
	routers.ServerCore()
}
