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

var testForEachFilesTests = []struct {
	in  string
	out []string
}{
	{
		getTestDataPath(""),
		[]string{
			"copiee.txt",
			"copiee2.txt",
			"nested_copiee.txt",
			"nested2_copiee.txt",
		},
	},
	{
		getTestDataPath("nested"),
		[]string{
			"nested_copiee.txt",
			"nested2_copiee.txt",
		},
	},
	{
		getTestDataPath("nested/nested2"),
		[]string{
			"nested2_copiee.txt",
		},
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
	for i, c := range testForEachFilesTests {
		t.Run(fmt.Sprintf("TestForEachFiles %d", i), func(t *testing.T) {
			err := ForEachFiles(c.in, func(dirname string, info os.FileInfo) error {
				if !contains(c.out, info.Name()) {
					return fmt.Errorf("%v doesn't contain %s", c.out, info.Name())
				}

				return nil
			})

			if err != nil {
				t.Error(err)
			}
		})
	}
}
