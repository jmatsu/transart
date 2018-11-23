package lib

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func CopyFile(srcPath string, destPath string) error {
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

func ForEachFiles(dirname string, f os.FileInfo, action func(dirname string, info os.FileInfo) error) error {
	if f.IsDir() {
		fs, err := ioutil.ReadDir(fmt.Sprintf("%s/%s", dirname, f.Name()))

		if err != nil {
			return err
		}

		for _, f := range fs {
			if err := ForEachFiles(fmt.Sprintf("%s/%s", dirname, f.Name()), f, action); err != nil {
				return err
			}
		}
	} else if err := action(dirname, f); err != nil {
		return err
	}

	return nil
}
