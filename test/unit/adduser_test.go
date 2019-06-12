package adduser_test

import (
	"testing"
	"lang.yottadb.com/go/yottadb"
	. "gitlab.com/euterpe/ydbocto-admin/internal/test_helpers"
	. "gitlab.com/euterpe/ydbocto-admin/pkg/adduser"
)

func TestHashMd5Password(t *testing.T) {
	md5Password := HashMd5Password("jon", []byte("tester"))
	if md5Password != "md5ed0c6ed88ae51106455ea90e52157be4" {
		t.Errorf("HashMd5Password(\"jon\", \"tester\") = %s; expected \"md5ed0c6ed88ae51106455ea90e52157be4\"", md5Password)
	}
}

func TestAddUser(t *testing.T) {
	var tptoken uint64 = yottadb.NOTTP
	var cleanup bool = false
	var errstr yottadb.BufferT
	var err error
	var userId, result string

	test_dir := Setup()
	userId, err = AddUser("jon", []byte("tester"))
	if nil != err {
		t.Errorf("AddUser failed with error: %v", err)
	}

	varname := "^%ydboctoocto"
	result, err = yottadb.ValE(tptoken, &errstr, varname, []string{userId, "rolname"})
	if nil != err {
		t.Errorf("YDBGo failed with error: %v", err)
	} else if result != "jon" {
		t.Errorf("AddUser: rolname = %s; expected \"jon\"", result)
	}
	result, err = yottadb.ValE(tptoken, &errstr, varname, []string{userId, "rolpassword"})
	if nil != err {
		t.Errorf("YDBGo failed with error: %v", err)
	} else if result != "md5ed0c6ed88ae51106455ea90e52157be4" {
		t.Errorf("AddUser: rolpassword = %s; expected \"md5ed0c6ed88ae51106455ea90e52157be4\"", result)
	} else {
		cleanup = true
	}
	Teardown(cleanup, test_dir)
}

