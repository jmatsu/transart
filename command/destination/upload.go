package destination

import (
	"fmt"
	"github.com/jmatsu/artifact-transfer/command"
	"github.com/jmatsu/artifact-transfer/github"
	"github.com/jmatsu/artifact-transfer/github/action"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"os"
)

func NewUploadAction() command.Actions {
	return command.Actions{
		GitHubRelease: func(config github.Config) error {
			release, err := action.GetDraftRelease(config)

			if err != nil {
				return err
			}

			fs, err := ioutil.ReadDir(".transart")

			if err != nil {
				return err
			}

			for _, f := range fs {
				recursive(".transart", f, func(dirname string, info os.FileInfo) error {
					asset, err := action.UploadToRelease(config, release, fmt.Sprintf("%s/%s", dirname, info.Name()))

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
