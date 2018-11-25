package destination

import (
	"fmt"
	"github.com/jmatsu/transart/client"
	"github.com/jmatsu/transart/client/entity"
	"github.com/jmatsu/transart/config"
	"github.com/jmatsu/transart/lib"
	"github.com/pkg/errors"
	"os"
)

func uploadToGithubRelease(rootConfig config.RootConfig, ghConfig config.GitHubConfig) error {
	if err := ghConfig.Validate(); err != nil {
		return errors.Wrap(err, "the configuration error happened on github release configuration")
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
		return errors.Wrap(ghClient.Err, "cannot prepare the release to be updated")
	}

	return lib.ForEachFiles(rootConfig.SaveDir, func(dirname string, info os.FileInfo) error {
		_ = ghClient.UploadFileToRelease(releaseToBeUpdated, fmt.Sprintf("%s/%s", dirname, info.Name()))

		return ghClient.Err
	})
}