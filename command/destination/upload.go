package destination

import (
	"fmt"
	"github.com/jmatsu/artifact-transfer/command"
	"github.com/jmatsu/artifact-transfer/config"
	"github.com/jmatsu/artifact-transfer/github"
	"github.com/jmatsu/artifact-transfer/github/entity"
	"github.com/jmatsu/artifact-transfer/lib"
	"github.com/jmatsu/artifact-transfer/local"
	"github.com/sirupsen/logrus"
	"gopkg.in/urfave/cli.v2"
	"io/ioutil"
	"os"
)

func NewUploadAction(c cli.Context) command.Actions {
	return command.Actions{
		GitHubRelease: uploadToGithubRelease,
		Local:         uploadToLocal,
	}
}

func uploadToGithubRelease(rootConfig config.RootConfig, gitHubConfig config.GitHubConfig) error {
	if err := gitHubConfig.Validate(); err != nil {
		return err
	}

	var releaseToBeUpdated *entity.Release

	switch gitHubConfig.GetStrategy() {
	case config.Create:
		if release, err := github.CreateDraftRelease(gitHubConfig); err == nil {
			releaseToBeUpdated = &release
		} else {
			return err
		}
	case config.Draft:
		if release, err := github.GetDraftRelease(gitHubConfig); err == nil {
			releaseToBeUpdated = &release
		} else {
			return err
		}
	case config.DraftOrCreate:
		if release, err := github.GetDraftRelease(gitHubConfig); err == nil {
			releaseToBeUpdated = &release
		} else if release, err := github.CreateDraftRelease(gitHubConfig); err == nil {
			releaseToBeUpdated = &release
		} else {
			return err
		}
	}

	if releaseToBeUpdated == nil {
		panic(fmt.Errorf("implementation error"))
	}

	fs, err := ioutil.ReadDir(rootConfig.SaveDir)

	if err != nil {
		return err
	}

	for _, f := range fs {
		lib.ForEachFiles(rootConfig.SaveDir, f, func(dirname string, info os.FileInfo) error {
			asset, err := github.UploadToRelease(gitHubConfig, *releaseToBeUpdated, fmt.Sprintf("%s/%s", dirname, info.Name()))

			if err != nil {
				return err
			}

			logrus.Debugln(asset.Name)

			return nil
		})
	}

	return nil
}

func uploadToLocal(rootConfig config.RootConfig, localConfig config.LocalConfig) error {
	if err := localConfig.Validate(); err != nil {
		return err
	}

	fs, err := ioutil.ReadDir(rootConfig.SaveDir)

	if err != nil {
		return err
	}

	for _, f := range fs {
		lib.ForEachFiles(rootConfig.SaveDir, f, func(dirname string, info os.FileInfo) error {
			if err := local.CopyFileFrom(localConfig, fmt.Sprintf("%s/%s", dirname, info.Name())); err != nil {
				return err
			}

			return nil
		})
	}

	return nil
}
