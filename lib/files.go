package lib

import (
	"fmt"
	"io/ioutil"
	"os"
)

func ForEachFiles(dirPath string, action func(dirname string, info os.FileInfo) error) error {
	fs, err := ioutil.ReadDir(dirPath)

	if err != nil {
		return err
	}

	for _, f := range fs {
		applyToFileOrFindDirectory(dirPath, f, action)
	}

	return nil
}

func applyToFileOrFindDirectory(dirname string, f os.FileInfo, action func(dirname string, info os.FileInfo) error) error {
	if f.IsDir() {
		fs, err := ioutil.ReadDir(fmt.Sprintf("%s/%s", dirname, f.Name()))

		if err != nil {
			return err
		}

		for _, f := range fs {
			if err := applyToFileOrFindDirectory(fmt.Sprintf("%s/%s", dirname, f.Name()), f, action); err != nil {
				return err
			}
		}
	} else if err := action(dirname, f); err != nil {
		return err
	}

	return nil
}
