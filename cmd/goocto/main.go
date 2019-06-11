package main

import (
	"fmt"
	"syscall"
	"github.com/docopt/docopt-go"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	usage := `goocto.

Usage:
	goocto add user <name>
`

	opts, err := docopt.ParseDoc(usage)
	if err != nil {
		fmt.Println("%v", err)
		return
	}

	if opts["add"] == true && opts["user"] == true {
		fmt.Printf("Enter password for user %v: ", opts["<name>"])
		rawPassword, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			fmt.Println("%v", err)
			return
		}
		password := string(rawPassword)
		fmt.Println(password)
	}
	// fmt.Println(opts)
}
