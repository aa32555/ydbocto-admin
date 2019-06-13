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
	ydbocto-admin remove user <name>
	ydbocto-admin show users <name>
`

	opts, err := docopt.ParseDoc(usage)
	if err != nil {
		fmt.Println(err)
		return
	}

	if opts["add"] == true {
        if opts["user"] == true {
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
	} else if opts["remove"] == true {
         if opts["user"] == true {
         }
	} else if opts["show"] == true {
         if opts["users"] == true {
         }
    }
	// fmt.Println(opts)
}
