package adduser_test

import (
	"testing"
	// "lang.yottadb.com/go/yottadb"
	. "gitlab.com/euterpe/ydbocto-admin/internal/test_helpers"
	. "gitlab.com/euterpe/ydbocto-admin/pkg/adduser"
)

func TestAddUser(t *testing.T) {
	test_dir := Setup()
	result := 0

	Teardown(result, test_dir)
}

func TestHashMd5Password(t *testing.T) {
	md5Password := HashMd5Password("jon", []byte("tester"))
	if md5Password != "md5ed0c6ed88ae51106455ea90e52157be4" {
		t.Errorf("HashMd5Password(\"jon\", \"tester\") = %s; expected \"md5ed0c6ed88ae51106455ea90e52157be4\"", md5Password)
	}
}
