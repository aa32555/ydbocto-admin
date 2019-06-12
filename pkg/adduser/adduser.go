package adduser

import (
	"strings"
	"crypto/md5"
	"encoding/hex"
)

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

/*
func AddUser(user string, rawPassword []byte) int {
}
*/
