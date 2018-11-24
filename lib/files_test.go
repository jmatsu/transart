package lib

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"testing"
)

func getFixturePath(filename string) string {
	if out, err := exec.Command("git", "rev-parse", "--show-toplevel").Output(); err != nil {
		panic(err)
	} else {
		// drop \n
		dir := fmt.Sprintf("%s/fixture", out[0:len(out)-1])

		if filename != "" {
			return fmt.Sprintf("%s/%s", dir, filename)
		} else {
			return dir
		}
	}
}

var testCopyFileTests = []struct {
	src  string
	dest string
}{
	{
		getFixturePath("copiee.txt"),
		getFixturePath("copiee.txt.copied"),
	},
	{
		getFixturePath("copiee2.txt"),
		getFixturePath("copiee2.txt.copied"),
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
	getFixturePath(""),
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
