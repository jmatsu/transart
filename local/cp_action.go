package local

import (
	"fmt"
	"github.com/jmatsu/transart/config"
	"github.com/jmatsu/transart/lib"
	"os"
	"path/filepath"
	"regexp"
)

func CopyFileFrom(localConfig config.LocalConfig, srcPath string) error {
	if pattern := localConfig.GetFileNamePattern(); pattern != "" {
		if !regexp.MustCompile(pattern).MatchString(srcPath) {
			return nil
		}
	}

	return lib.CopyFile(srcPath, fmt.Sprintf("%s/%s", localConfig.GetPath(), filepath.Base(srcPath)))
}

func CopyFilesTo(localConfig config.LocalConfig, destDirPath string) error {
	return lib.ForEachFiles(localConfig.GetPath(), func(dirname string, info os.FileInfo) error {
		srcPath := fmt.Sprintf("%s/%s", dirname, info.Name())

		if pattern := localConfig.GetFileNamePattern(); pattern != "" {
			if !regexp.MustCompile(pattern).MatchString(srcPath) {
				return nil
			}
		}

		destPath := fmt.Sprintf("%s/%s", destDirPath, info.Name())

		if err := lib.CopyFile(srcPath, destPath); err != nil {
			return err
		}

		return nil
	})
}
