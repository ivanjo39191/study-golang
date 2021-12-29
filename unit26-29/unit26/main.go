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
