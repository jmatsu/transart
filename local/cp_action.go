package local

import (
	"fmt"
	"github.com/jmatsu/artifact-transfer/config"
	"github.com/jmatsu/artifact-transfer/lib"
	"io/ioutil"
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
	fs, err := ioutil.ReadDir(localConfig.GetPath())

	if err != nil {
		return err
	}

	for _, f := range fs {
		lib.ForEachFiles(localConfig.GetPath(), f, func(dirname string, info os.FileInfo) error {
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

	return nil
}
