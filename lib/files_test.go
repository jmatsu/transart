package lib

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func getTestDataPath(filename string) string {
	if filename != "" {
		return fmt.Sprintf("testdata/%s", filename)
	} else {
		return "testdata"
	}
}

var testCopyFileTests = []struct {
	src  string
	dest string
}{
	{
		getTestDataPath("copiee.txt"),
		getTestDataPath("copiee.txt.copied"),
	},
	{
		getTestDataPath("copiee2.txt"),
		getTestDataPath("copiee2.txt.copied"),
	},
}

func TestCopyFile(t *testing.T) {
	for i, c := range testCopyFileTests {
		t.Run(fmt.Sprintf("TestCopyFile %d", i), func(t *testing.T) {
			defer func() {
				if _, err := os.Stat(c.dest); err == nil {
					os.Remove(c.dest)
				}
			}()

			if _, err := os.Stat(c.dest); err == nil {
				os.Remove(c.dest)
			}

			if err := CopyFile(c.src, c.dest); err != nil {
				t.Error(err)
			}
		})
	}
}

var testForEachFiles = struct {
	dir   string
	files []string
}{
	getTestDataPath(""),
	[]string{
		"copiee.txt",
		"copiee2.txt",
		"nested_copiee.txt",
	},
}

func contains(strs []string, str string) bool {
	for _, s := range strs {
		if s == str {
			return true
		}
	}

	return false
}

func TestForEachFiles(t *testing.T) {
	if fs, err := ioutil.ReadDir(testForEachFiles.dir); err != nil {
		t.Error(err)
	} else {
		for _, f := range fs {
			err := ForEachFiles(testForEachFiles.dir, f, func(dirname string, info os.FileInfo) error {
				if !contains(testForEachFiles.files, info.Name()) {
					return fmt.Errorf("%v doesn't contain %s", testForEachFiles.files, info.Name())
				}

				return nil
			})

			if err != nil {
				t.Error(err)
			}
		}
	}
}
