package client

import (
	"fmt"
	"github.com/jmatsu/transart/lib"
	"golang.org/x/sync/errgroup"
	"io"
	"os"
)

type Local interface {
	CopyFile(srcPath string, destPath string) error
}

type LocalClient struct {
	dirPath string
	c       Local
	Err     error
}

func NewLocalClient(dirPath string) LocalClient {
	return LocalClient{
		dirPath: dirPath,
		c:       localImpl{},
	}
}

func (lc *LocalClient) CopyFileFrom(srcPath string) {
	if !lib.IsNil(lc.Err) {
		return
	}

	lc.Err = lc.c.CopyFile(srcPath, lc.dirPath)
}

func (lc *LocalClient) CopyDirFrom(srcPath string, prod func(string) bool) {
	if !lib.IsNil(lc.Err) {
		return
	}

	lc.Err = lc.copyDir(srcPath, lc.dirPath, prod)
}

func (lc *LocalClient) CopyDirTo(destPath string, prod func(string) bool) {
	if !lib.IsNil(lc.Err) {
		return
	}

	lc.Err = lc.copyDir(lc.dirPath, destPath, prod)
}

func (lc *LocalClient) copyDir(srcPath string, destPath string, prod func(string) bool) error {
	eg := errgroup.Group{}

	err := lib.ForEachFiles(srcPath, func(dirname string, info os.FileInfo) error {
		eg.Go(func() error {
			srcFilePath := fmt.Sprintf("%s/%s", dirname, info.Name())

			if !prod(srcFilePath) {
				return nil
			}

			destFilePath := fmt.Sprintf("%s/%s", destPath, info.Name())
			return lc.c.CopyFile(srcFilePath, destFilePath)
		})

		return nil
	})

	if err != nil {
		return err
	}

	return eg.Wait()
}

type localImpl struct {
}

func (l localImpl) CopyFile(srcPath string, destPath string) error {
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
