package local

import (
	"fmt"
	"github.com/jmatsu/artifact-transfer/config"
	"io"
	"os"
	"path/filepath"
)

func CopyFile(localConfig config.LocalConfig, srcPath string) error {
	if s, err := os.Stat(srcPath); err != nil {
		return err
	} else if !s.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", srcPath)
	}

	src, err := os.Open(srcPath)

	if err != nil {
		return err
	}

	defer src.Close()

	destPath := fmt.Sprintf("%s/%s", localConfig.GetPath(), filepath.Base(srcPath))

	dest, err := os.Create(destPath)

	if err != nil {
		return err
	}

	defer dest.Close()

	if _, err := io.Copy(dest, src); err != nil {
		return err
	}

	return nil
}
