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

func TestAddUserEmptyString(t *testing.T) {
	var cleanup bool = false
	var err error

	test_dir := Setup()
	_, err = AddUser("", []byte("tester"))
	if nil == err {
        t.Errorf("AddUser: no error where expected")
    } else {
        cleanup = true
    }
	Teardown(cleanup, test_dir)
}

func TestAddOneUser(t *testing.T) {
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

func TestAddTwoUsers(t *testing.T) {
	var tptoken uint64 = yottadb.NOTTP
	var cleanup bool = false
	var errstr yottadb.BufferT
	var err error
	var userId1, userId2, result string

	test_dir := Setup()
	userId1, err = AddUser("jon", []byte("tester"))
	if nil != err {
		t.Errorf("AddUser failed with error: %v", err)
	} else if userId1 != "1" {
		t.Errorf("AddUser: userId = %s; expected \"1\"", userId1)
    }
	userId2, err = AddUser("bobby", []byte("buttons"))
	if nil != err {
		t.Errorf("AddUser failed with error: %v", err)
	} else if userId2 != "1" {
		t.Errorf("AddUser: userId = %s; expected \"2\"", userId2)
	}

	varname := "^%ydboctoocto"
    // Check first user
	result, err = yottadb.ValE(tptoken, &errstr, varname, []string{userId1, "rolname"})
	if nil != err {
		t.Errorf("YDBGo failed with error: %v", err)
	} else if result != "jon" {
		t.Errorf("AddUser: rolname = %s; expected \"jon\"", result)
	}
	result, err = yottadb.ValE(tptoken, &errstr, varname, []string{userId1, "rolpassword"})
	if nil != err {
		t.Errorf("YDBGo failed with error: %v", err)
	} else if result != "md5ed0c6ed88ae51106455ea90e52157be4" {
		t.Errorf("AddUser: rolpassword = %s; expected \"md5ed0c6ed88ae51106455ea90e52157be4\"", result)
	}
    // Check second user
	result, err = yottadb.ValE(tptoken, &errstr, varname, []string{userId2, "rolname"})
	if nil != err {
		t.Errorf("YDBGo failed with error: %v", err)
	} else if result != "bobby" {
		t.Errorf("AddUser: rolname = %s; expected \"bobby\"", result)
	}
	result, err = yottadb.ValE(tptoken, &errstr, varname, []string{userId2, "rolpassword"})
	if nil != err {
		t.Errorf("YDBGo failed with error: %v", err)
	} else if result != "md5ba4b0dee824f0c9d92017d4308bcc43d" {
		t.Errorf("AddUser: rolpassword = %s; expected \"md5ba4b0dee824f0c9d92017d4308bcc43d\"", result)
	} else {
		cleanup = true
	}

	Teardown(cleanup, test_dir)
}

func TestAddThreeUsers(t *testing.T) {
	var tptoken uint64 = yottadb.NOTTP
	var cleanup bool = false
	var errstr yottadb.BufferT
	var err error
	var userId1, userId2, userId3, result string

	test_dir := Setup()
	userId1, err = AddUser("jon", []byte("tester"))
	if nil != err {
		t.Errorf("AddUser failed with error: %v", err)
	} else if userId1 != "1" {
		t.Errorf("AddUser: userId = %s; expected \"1\"", userId1)
    }
	userId2, err = AddUser("bobby", []byte("buttons"))
	if nil != err {
		t.Errorf("AddUser failed with error: %v", err)
	} else if userId2 != "1" {
		t.Errorf("AddUser: userId = %s; expected \"2\"", userId2)
	}
	userId3, err = AddUser("suzy", []byte("quartz"))
	if nil != err {
		t.Errorf("AddUser failed with error: %v", err)
	} else if userId2 != "1" {
		t.Errorf("AddUser: userId = %s; expected \"3\"", userId3)
	}

	varname := "^%ydboctoocto"
    // Check first user
	result, err = yottadb.ValE(tptoken, &errstr, varname, []string{userId1, "rolname"})
	if nil != err {
		t.Errorf("YDBGo failed with error: %v", err)
	} else if result != "jon" {
		t.Errorf("AddUser: rolname = %s; expected \"jon\"", result)
	}
	result, err = yottadb.ValE(tptoken, &errstr, varname, []string{userId1, "rolpassword"})
	if nil != err {
		t.Errorf("YDBGo failed with error: %v", err)
	} else if result != "md5ed0c6ed88ae51106455ea90e52157be4" {
		t.Errorf("AddUser: rolpassword = %s; expected \"md5ed0c6ed88ae51106455ea90e52157be4\"", result)
	}
    // Check second user
	result, err = yottadb.ValE(tptoken, &errstr, varname, []string{userId2, "rolname"})
	if nil != err {
		t.Errorf("YDBGo failed with error: %v", err)
	} else if result != "bobby" {
		t.Errorf("AddUser: rolname = %s; expected \"bobby\"", result)
	}
	result, err = yottadb.ValE(tptoken, &errstr, varname, []string{userId2, "rolpassword"})
	if nil != err {
		t.Errorf("YDBGo failed with error: %v", err)
	} else if result != "md5ba4b0dee824f0c9d92017d4308bcc43d" {
		t.Errorf("AddUser: rolpassword = %s; expected \"md5ba4b0dee824f0c9d92017d4308bcc43d\"", result)
	} else {
		cleanup = true
	}
    // Check third user
	result, err = yottadb.ValE(tptoken, &errstr, varname, []string{userId3, "rolname"})
	if nil != err {
		t.Errorf("YDBGo failed with error: %v", err)
	} else if result != "suzy" {
		t.Errorf("AddUser: rolname = %s; expected \"suzy\"", result)
	}
	result, err = yottadb.ValE(tptoken, &errstr, varname, []string{userId3, "rolpassword"})
	if nil != err {
		t.Errorf("YDBGo failed with error: %v", err)
	} else if result != "md5ceb640c999ac0adfd45edbb06c7447cb" {
		t.Errorf("AddUser: rolpassword = %s; expected \"md5ceb640c999ac0adfd45edbb06c7447cb\"", result)
	} else {
		cleanup = true
	}

	Teardown(cleanup, test_dir)
}
