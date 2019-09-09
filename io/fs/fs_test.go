package fs_test

import (
	"testing"

	"github.com/birchwood-langham/go-toolkit/io/fs"
)

func Test_ListFiles(t *testing.T) {
	files, err := fs.ListFiles("./")

	if err != nil {
		t.Errorf("could not get a list of files - %v", err)
	}

	expected := 2
	found := 0
	for _, f := range files {
		if f == "fs_test.go" || f == "fs.go" {
			found++
		}
	}

	if found != expected {
		t.Errorf("Expected to find %d files, but found %d", expected, found)
	}
}

func Test_ListSubFolders(t *testing.T) {
	files, err := fs.ListSubFolders("../")

	if err != nil {
		t.Errorf("could not get a list of files - %v", err)
	}

	expected := 1
	found := 0
	for _, f := range files {
		if f == "fs" {
			found++
		}
	}

	if found != expected {
		t.Errorf("Expected to find %d files, but found %d", expected, found)
	}
}

func Test_FileExists(t *testing.T) {
	exists, err := fs.FileExists("fs.go")

	if err != nil {
		t.Errorf("Could not access file - %v", err)
	}

	if !exists {
		t.Errorf("fs.go does not exist")
	}
}
