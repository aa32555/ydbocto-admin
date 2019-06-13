package adduser

import (
    "fmt"
	"lang.yottadb.com/go/yottadb"
)

// ShowUsers retrieves all user entries from the database, prints them to stdout,
// and returns the total number retrieved for convenience.
// Assumes existence of the relevant global variable.
func ShowUsers() (int, error) {
	var tptoken uint64 = yottadb.NOTTP
	var errstr yottadb.BufferT
	var subserr error
    var userId string
    var users []string

	varname := "^%ydboctoocto"
    for subserr == nil {
        userId, subserr = yottadb.SubNextE(tptoken, &errstr, varname, []string{"user"})
        user, err := yottadb.ValE(tptoken, &errstr, varname, []string{userId, "rolname"})
        if nil != err {
            return 0, err
        }
        users = append(users, user)
    }
    totalUsers := len(users)
    if  totalUsers <= 0 {
        fmt.Println("No YDBOcto users found.")
    } else {
        fmt.Println("Current YDBOcto users, by ID:")
        for i, user := range users {
            fmt.Printf("%8d%s\n", i, user)
        }
    }
    return totalUsers, nil
}
