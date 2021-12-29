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
