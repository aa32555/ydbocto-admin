package adduser

import (
    "fmt"
    "strconv"
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

	varname := "^%ydboctoocto"
    users := make(map[int]string)
    for subserr == nil {
        user, subserr = yottadb.SubNextE(tptoken, &errstr, varname, []string{"user"})
        userId, err := yottadb.ValE(tptoken, &errstr, varname, []string{user, "id"})
        if nil != err {
            return 0, err
        }
        i, err := strconv.ParseInt(userId, 10, 64)
        m[i] = user
    }

    totalUsers := len(users)
    if  totalUsers <= 0 {
        fmt.Println("No YDBOcto users found.")
    } else {
        fmt.Println("Current YDBOcto users, by ID:")
		var keys []int
		for k := range users {
			keys = append(keys, k)
		}
		sort.Ints(keys)
        for _, i := range keys {
            fmt.Printf("%8d%s\n", i, users[i])
        }
    }
    return totalUsers, nil
}
