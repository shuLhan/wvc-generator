// Copyright 2016 Mhd Sulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

/*
Package revision provide the module for working with dump files of Wikipedia
revision.
*/
package revision

import (
	"errors"
	"io/ioutil"
	"os"
)

var (
	// Dir define where the revision directory located.
	Dir = ""
	// CleanDir define directory where revision that has been cleaned
	// up located.
	CleanDir = ""
)

/*
SetDir will set revision directory to `path`.
*/
func SetDir(path string) {
	Dir = path
}

/*
SetCleanDir set directory where revision that has been cleaned up located.
*/
func SetCleanDir(path string) {
	CleanDir = path
}

/*
GetContent will return content of revision based on specific `id`.
*/
func GetContent(id string) (string, error) {
	if Dir == "" {
		return "", errors.New("Revision directory is not set!")
	}

	path := Dir + "/" + id + ".txt"

	b, e := ioutil.ReadFile(path)

	return string(b), e
}

/*
GetContentClean return content of revision that has been cleaning up.
*/
func GetContentClean(id string) (string, error) {
	if CleanDir == "" {
		return "", errors.New("Clean revision directory is not set!")
	}

	path := CleanDir + "/" + id + ".txt"

	b, e := ioutil.ReadFile(path)

	return string(b), e
}

/*
GetSize return the file size of revision file.
*/
func GetSize(id string) int64 {
	if Dir == "" {
		return 0
	}

	path := Dir + "/" + id + ".txt"

	finfo, e := os.Stat(path)
	if e != nil {
		return 0
	}

	return finfo.Size()
}
