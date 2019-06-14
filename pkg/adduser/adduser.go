package adduser

import (
    "errors"
	"strings"
	"crypto/md5"
	"encoding/hex"
	"lang.yottadb.com/go/yottadb"
)

/*
err = yottadb.SetValE(tptoken, &errstr, "users", varname, nil)
if nil != err {
	return err
}
*/

// HashMd5Password hashes a password using a username as a salt to produce
// an md5 password string in accordance with the PostgreSQL spec.
func HashMd5Password(user string, rawPassword []byte) string {
	var password strings.Builder
	password.Write(rawPassword)
	password.WriteString(user)
	hashedPassword := md5.Sum([]byte(password.String()))

	var md5Password strings.Builder
	md5Password.WriteString("md5")
	md5Password.WriteString(hex.EncodeToString(hashedPassword[:]))
	return md5Password.String()
}

// AddUser creates a new database user, hashes the user's password, and stores it in the database.
// Assumes existence of the relevant global variable.
func AddUser(user string, rawPassword []byte) (newId string, err error) {
    if user == "" {
        err = errors.New("AddUser: user name cannot be empty string")
        return "", err
    }

	var tptoken uint64 = yottadb.NOTTP
	var errstr yottadb.BufferT
	varname := "^%ydboctoocto"

	newId, err = yottadb.IncrE(tptoken, &errstr, "", varname, []string{"users", "userIdCount"})
	if nil != err {
		return "", err
	}

	err = yottadb.SetValE(tptoken, &errstr, newId, varname, []string{"users", user, "id"})
	if nil != err {
		return "", err
	}

	err = yottadb.SetValE(tptoken, &errstr, user, varname, []string{"users", user, "rolname"})
	if nil != err {
		return "", err
	}

	md5Password := HashMd5Password(user, rawPassword)
	err = yottadb.SetValE(tptoken, &errstr, md5Password, varname, []string{"users", user, "rolpassword"})
	if nil != err {
		return "", err
	}
	return newId, nil
}
