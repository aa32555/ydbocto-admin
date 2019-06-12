//////////////////////////////////////////////////////////////////
//								//
// Copyright (c) 2019 YottaDB LLC. and/or its subsidiaries.	//
// All rights reserved.						//
//								//
//	This source code contains the intellectual property	//
//	of its copyright holder(s), and is made available	//
//	under a license.  If you do not know the terms of	//
//	the license, please stop and do not read further.	//
//								//
//////////////////////////////////////////////////////////////////

package test_helpers

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

// Setup prepares a test directory for running tests by configuring the environment
// and creating a database.
func Setup() (test_dir string) {
	// Get a temporary directory to put the database in
	test_dir, err := ioutil.TempDir("", "ydbgo")
	if err != nil {
		log.Fatal(err)
	}

	// Configure standard logger
	log.SetFlags(log.Lshortfile)

	// Setup environment variables
	log.Printf("Test directory is %s", test_dir)
	ydb_gbldir := filepath.Join(test_dir, "mumps.gld")
	ydb_datfile := filepath.Join(test_dir, "mumps.dat")
	// Save current global directory for post-test restoration, if set
	temp_gbldir := os.Getenv("ydb_gbldir")
	if temp_gbldir != "" {
		os.Setenv("temp_gbldir", temp_gbldir)
	}
	os.Setenv("ydb_gbldir", ydb_gbldir)
	ydb_dist := os.Getenv("ydb_dist")
	if ydb_dist == "" {
		log.Fatal("ydb_dist not set")
	}
	mumps_exe := filepath.Join(ydb_dist, "mumps")
	mupip_exe := filepath.Join(ydb_dist, "mupip")

	// Create global directory
	cmd := exec.Command(mumps_exe, "-run", "^GDE",
		"change -seg DEFAULT -file="+ydb_datfile)
	output, err := cmd.CombinedOutput()
	log.Printf("%s\n", output)
	if err != nil {
		log.Fatal(err)
	}

	// Create database itself
	cmd = exec.Command(mupip_exe, "create")
	output, err = cmd.CombinedOutput()
	log.Printf("%s\n", output)
	if err != nil {
		log.Fatal(err)
	}
	return test_dir
}

// Teardown resets modified environment variables and cleans up the test directory on failure, if desired.
func Teardown(result int, test_dir string) {
	// Restore previous global directory, if any
	ydb_gbldir := os.Getenv("temp_gbldir")
	if ydb_gbldir != "" {
		os.Setenv("ydb_gbldir", ydb_gbldir)
	}
	// Cleanup test directory on success, retain on failure
	if 0 == result {
		log.Printf("Cleaning up test directory")
		os.RemoveAll(test_dir)
	}
	os.Exit(result)
}
