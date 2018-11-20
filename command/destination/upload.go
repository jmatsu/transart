package destination

import (
	"fmt"
	"github.com/jmatsu/artifact-transfer/command"
	"github.com/jmatsu/artifact-transfer/github"
	"github.com/jmatsu/artifact-transfer/github/action"
	"github.com/jmatsu/artifact-transfer/github/entity"
	"github.com/sirupsen/logrus"
	"gopkg.in/urfave/cli.v2"
	"io/ioutil"
	"os"
)

func NewUploadAction(c cli.Context) command.Actions {
	return command.Actions{
		GitHubRelease: func(config github.Config) error {
			if err := config.Validate(); err != nil {
				return err
			}

			tagName := c.String(github.TagNameFlag().Name)
			ref := c.String(github.RefFlag().Name)

			var releaseToBeUpdated *entity.Release

			switch config.GetStrategy() {
			case github.Create:
				if release, err := action.CreateDraftRelease(config, tagName, ref); err == nil {
					releaseToBeUpdated = &release
				} else {
					return err
				}
			case github.Draft:
				if release, err := action.GetDraftRelease(config); err == nil {
					releaseToBeUpdated = &release
				} else {
					return err
				}
			case github.DraftOrCreate:
				if release, err := action.GetDraftRelease(config); err == nil {
					releaseToBeUpdated = &release
				} else if release, err := action.CreateDraftRelease(config, tagName, ref); err == nil {
					releaseToBeUpdated = &release
				} else {
					return err
				}
			}

			if releaseToBeUpdated == nil {
				panic(fmt.Errorf("implementation error"))
			}

			fs, err := ioutil.ReadDir(".transart")

			if err != nil {
				return err
			}

			for _, f := range fs {
				recursive(".transart", f, func(dirname string, info os.FileInfo) error {
					asset, err := action.UploadToRelease(config, *releaseToBeUpdated, fmt.Sprintf("%s/%s", dirname, info.Name()))

					if err != nil {
						return err
					}

					logrus.Debugln(asset.Name)

					return nil
				})
			}

			return nil
		},
	}
}

func recursive(dirname string, f os.FileInfo, action func(dirname string, info os.FileInfo) error) error {
	if f.IsDir() {
		fs, err := ioutil.ReadDir(fmt.Sprintf("%s/%s", dirname, f.Name()))

		if err != nil {
			return err
		}

		for _, _f := range fs {
			if err := recursive(fmt.Sprintf("%s/%s", dirname, _f.Name()), _f, action); err != nil {
				return err
			}
		}
	} else if err := action(dirname, f); err != nil {
		return err
	}

	return nil
}
