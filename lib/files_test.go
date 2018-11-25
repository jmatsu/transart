package lib

import (
	"fmt"
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
	err := ForEachFiles(testForEachFiles.dir, func(dirname string, info os.FileInfo) error {
		if !contains(testForEachFiles.files, info.Name()) {
			return fmt.Errorf("%v doesn't contain %s", testForEachFiles.files, info.Name())
		}

		return nil
	})

	if err != nil {
		t.Error(err)
	}
}
