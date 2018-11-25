package client

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

var testLocalImpl_CopyFileTests = []struct {
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

func TestLocalImpl_CopyFile(t *testing.T) {
	localImpl := localImpl{}

	for i, c := range testLocalImpl_CopyFileTests {
		t.Run(fmt.Sprintf("TestCopyFile %d", i), func(t *testing.T) {
			defer func() {
				if _, err := os.Stat(c.dest); err == nil {
					os.Remove(c.dest)
				}
			}()

			if _, err := os.Stat(c.dest); err == nil {
				os.Remove(c.dest)
			}

			if err := localImpl.CopyFile(c.src, c.dest); err != nil {
				t.Error(err)
			}
		})
	}
}
