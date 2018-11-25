package destination

import (
	"fmt"
	"github.com/jmatsu/transart/client"
	"github.com/jmatsu/transart/client/entity"
	"github.com/jmatsu/transart/command"
	"github.com/jmatsu/transart/config"
	"github.com/jmatsu/transart/lib"
	"github.com/jmatsu/transart/local"
	"io/ioutil"
	"os"
)

func NewUploadAction() command.Actions {
	return command.Actions{
		GitHubRelease: uploadToGithubRelease,
		Local:         uploadToLocal,
	}
}

func uploadToGithubRelease(rootConfig config.RootConfig, ghConfig config.GitHubConfig) error {
	if err := ghConfig.Validate(); err != nil {
		return err
	}

	ghClient := client.NewGitHubClient(ghConfig.GetUsername(), ghConfig.GetRepoName(), ghConfig.GetApiToken())

	var releaseToBeUpdated entity.Release

	switch ghConfig.GetStrategy() {
	case config.Create:
		releaseToBeUpdated = ghClient.CreateDraftRelease()
	case config.Draft:
		releaseToBeUpdated = ghClient.GetDraftRelease()
	case config.DraftOrCreate:
		releaseToBeUpdated = ghClient.GetDraftRelease()

		if !lib.IsNil(ghClient.Err) {
			ghClient.Err = nil
			releaseToBeUpdated = ghClient.CreateDraftRelease()
		}
	}

	if !lib.IsNil(ghClient.Err) {
		return ghClient.Err
	}

	fs, err := ioutil.ReadDir(rootConfig.SaveDir)

	if err != nil {
		return err
	}

	for _, f := range fs {
		lib.ForEachFiles(rootConfig.SaveDir, f, func(dirname string, info os.FileInfo) error {
			_ = ghClient.UploadFileToRelease(releaseToBeUpdated, fmt.Sprintf("%s/%s", dirname, info.Name()))

			return ghClient.Err
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
