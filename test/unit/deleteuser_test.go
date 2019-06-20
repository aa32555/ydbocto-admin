package admin_test

import (
	"testing"
	"lang.yottadb.com/go/yottadb"
	. "gitlab.com/euterpe/ydbocto-admin/internal/test_helpers"
	. "gitlab.com/euterpe/ydbocto-admin/pkg/adduser"
	. "gitlab.com/euterpe/ydbocto-admin/pkg/deleteuser"
)

func TestDeleteUserEmptyString(t *testing.T) {
	var err error

	Setup()
	err = DeleteUser("")
	if nil == err {
		t.Errorf("DeleteUser: no error where expected")
	}
	Teardown()
}

func TestDeleteOneUser(t *testing.T) {
	var tptoken uint64 = yottadb.NOTTP
	var errstr yottadb.BufferT
	var err error
	var result uint32

	Setup()
	_, err = AddUser("jon", []byte("tester"))
	if nil != err {
		t.Errorf("AddUser failed with error: %v", err)
	}

	err = DeleteUser("jon")
	if nil != err {
		t.Errorf("DeleteUser failed with error: %v", err)
	}

	varname := "^%ydboctoocto"
	result, err = yottadb.DataE(tptoken, &errstr, varname, []string{"users", "jon"})
	if nil != err {
		t.Errorf("YDBGo failed with error: %v", err)
	} else {
		if 0 != result {
			t.Errorf("DeleteUser failed to delete user: %v", "jon")
		}
	}
	Teardown()
}

/*
func TestDeleteTwoUsers(t *testing.T) {
	var tptoken uint64 = yottadb.NOTTP
	var cleanup bool = false
	var errstr yottadb.BufferT
	var err error
	var userId1, userId2, row string

	test_dir := Setup()
	userId1, err = AddUser("jon", []byte("tester"))
	if nil != err {
		t.Errorf("AddUser failed with error: %v", err)
	}
	userId2, err = AddUser("bobby", []byte("buttons"))
	if nil != err {
		t.Errorf("AddUser failed with error: %v", err)
	}

	varname := "^%ydboctoocto"
	// Check first user
	row, err = yottadb.ValE(tptoken, &errstr, varname, []string{"users", "jon"})
	if nil != err {
		t.Errorf("YDBGo failed with error: %v", err)
		cleanup = false
	} else {
		columns := strings.Split(row, "|")
		if columns[Oid] != userId1 {
			t.Errorf("AddUser: oid = %s; expected \"%s\"", columns[Oid], userId1)
			cleanup = false
		}
		if columns[Rolname] != "jon" {
			t.Errorf("AddUser: rolname = %s; expected \"jon\"", columns[Rolname])
			cleanup = false
		}
		if columns[Rolpassword] != "md5ed0c6ed88ae51106455ea90e52157be4" {
			t.Errorf("AddUser: rolpassword = %s; expected \"md5ed0c6ed88ae51106455ea90e52157be4\"", columns[Rolpassword])
			cleanup = false
		}
	}
	// Check second user
	row, err = yottadb.ValE(tptoken, &errstr, varname, []string{"users", "bobby"})
	if nil != err {
		t.Errorf("YDBGo failed with error: %v", err)
		cleanup = false
	} else {
		columns := strings.Split(row, "|")
		if columns[Oid] != userId2 {
			t.Errorf("AddUser: oid = %s; expected \"%s\"", columns[Oid], userId2)
			cleanup = false
		}
		if columns[Rolname] != "bobby" {
			t.Errorf("AddUser: rolname = %s; expected \"bobby\"", columns[Rolname])
			cleanup = false
		}
		if columns[Rolpassword] != "md5ba4b0dee824f0c9d92017d4308bcc43d" {
			t.Errorf("AddUser: rolpassword = %s; expected \"md5ba4b0dee824f0c9d92017d4308bcc43d\"", columns[Rolpassword])
			cleanup = false
		}
	}
	Teardown(cleanup, test_dir)
}

func TestAddThreeUsers(t *testing.T) {
	var tptoken uint64 = yottadb.NOTTP
	var cleanup bool = false
	var errstr yottadb.BufferT
	var err error
	var userId1, userId2, userId3, row string

	test_dir := Setup()
	userId1, err = AddUser("jon", []byte("tester"))
	if nil != err {
		t.Errorf("AddUser failed with error: %v", err)
	}
	userId2, err = AddUser("bobby", []byte("buttons"))
	if nil != err {
		t.Errorf("AddUser failed with error: %v", err)
	}
	userId3, err = AddUser("suzy", []byte("quartz"))
	if nil != err {
		t.Errorf("AddUser failed with error: %v", err)
	}

	varname := "^%ydboctoocto"
	// Check first user
	row, err = yottadb.ValE(tptoken, &errstr, varname, []string{"users", "jon"})
	if nil != err {
		t.Errorf("YDBGo failed with error: %v", err)
		cleanup = false
	} else {
		columns := strings.Split(row, "|")
		if columns[Oid] != userId1 {
			t.Errorf("AddUser: oid = %s; expected \"%s\"", columns[Oid], userId1)
			cleanup = false
		}
		if columns[Rolname] != "jon" {
			t.Errorf("AddUser: rolname = %s; expected \"jon\"", columns[Rolname])
			cleanup = false
		}
		if columns[Rolpassword] != "md5ed0c6ed88ae51106455ea90e52157be4" {
			t.Errorf("AddUser: rolpassword = %s; expected \"md5ed0c6ed88ae51106455ea90e52157be4\"", columns[Rolpassword])
			cleanup = false
		}
	}
	// Check second user
	row, err = yottadb.ValE(tptoken, &errstr, varname, []string{"users", "bobby"})
	if nil != err {
		t.Errorf("YDBGo failed with error: %v", err)
		cleanup = false
	} else {
		columns := strings.Split(row, "|")
		if columns[Oid] != userId2 {
			t.Errorf("AddUser: oid = %s; expected \"%s\"", columns[Oid], userId2)
			cleanup = false
		}
		if columns[Rolname] != "bobby" {
			t.Errorf("AddUser: rolname = %s; expected \"bobby\"", columns[Rolname])
			cleanup = false
		}
		if columns[Rolpassword] != "md5ba4b0dee824f0c9d92017d4308bcc43d" {
			t.Errorf("AddUser: rolpassword = %s; expected \"md5ba4b0dee824f0c9d92017d4308bcc43d\"", columns[Rolpassword])
			cleanup = false
		}
	}
	// Check third user
	row, err = yottadb.ValE(tptoken, &errstr, varname, []string{"users", "suzy"})
	if nil != err {
		t.Errorf("YDBGo failed with error: %v", err)
		cleanup = false
	} else {
		columns := strings.Split(row, "|")
		if columns[Oid] != userId3 {
			t.Errorf("AddUser: oid = %s; expected \"%s\"", columns[Oid], userId3)
			cleanup = false
		}
		if columns[Rolname] != "suzy" {
			t.Errorf("AddUser: rolname = %s; expected \"suzy\"", columns[Rolname])
			cleanup = false
		}
		if columns[Rolpassword] != "md5ceb640c999ac0adfd45edbb06c7447cb" {
			t.Errorf("AddUser: rolpassword = %s; expected \"md5ceb640c999ac0adfd45edbb06c7447cb\"", columns[Rolpassword])
			cleanup = false
		}
	}
	Teardown(cleanup, test_dir)
}
*/
