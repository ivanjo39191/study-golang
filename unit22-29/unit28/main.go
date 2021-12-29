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
