package main

import (
	"fmt"
	// "strings"
	"syscall"
	// "crypto/md5"
	// "encoding/hex"
	"gitlab.com/euterpe/ydbocto-admin/pkg/adduser"
	"github.com/docopt/docopt-go"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	usage := `goocto.

Usage:
	ydbocto-admin add user <name>
`

	opts, err := docopt.ParseDoc(usage)
	if err != nil {
		fmt.Println(err)
		return
	}

	if opts["add"] == true && opts["user"] == true {
		user := opts["<name>"].(string)
		fmt.Printf("Enter password for user %v: ", user)
		rawPassword, err := terminal.ReadPassword(int(syscall.Stdin))
		if err != nil {
			fmt.Println(err)
			return
		}
		md5Password := adduser.HashMd5Password(user, rawPassword)
		fmt.Println(md5Password)
	}
	// fmt.Println(opts)
}
