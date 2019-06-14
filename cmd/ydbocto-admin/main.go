package main

import (
	"fmt"
	// "strings"
	// "syscall"
	// "crypto/md5"
	// "encoding/hex"
	"readpassword"
	"gitlab.com/euterpe/ydbocto-admin/pkg/adduser"
	"github.com/docopt/docopt-go"
	// "golang.org/x/crypto/ssh/terminal"
)

func main() {
	usage := `ydbocto-admin.

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
            prompt := fmt.Sprintf("Enter password for user %v: ", user)
            // rawPassword, err := terminal.ReadPassword(int(syscall.Stdin))
            rawPassword, err := readpassword.ReadPassword(prompt)
            if err != nil {
                fmt.Println(err)
                return
            }
            _, err = adduser.AddUser(user, rawPassword)
            if err != nil {
                fmt.Println(err)
                return
            }
            fmt.Printf("Successfully added user: \"%s\"\n", user)
        }
	} else if opts["remove"] == true {
         if opts["user"] == true {
         }
	} else if opts["show"] == true {
         if opts["users"] == true {
         }
    }
}
