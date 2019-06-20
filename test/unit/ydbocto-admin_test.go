package admin_test

import (
	"os"
	"log"
	"testing"
)

func TestMain(m *testing.M) {
	test_dir := "/tmp/ydbocto-admin_test"
	err := os.Mkdir(test_dir, 0700)
	if err != nil {
		log.Fatal(err)
	}
	ret_code := m.Run()
	os.RemoveAll(test_dir)
	os.Exit(ret_code)
}
