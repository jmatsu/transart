package source

import (
	"fmt"
	"github.com/jmatsu/transart/client"
	"github.com/jmatsu/transart/config"
	"github.com/jmatsu/transart/lib"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"io/ioutil"
	"regexp"
)

func downloadFromGitHubRelease(rootConfig config.RootConfig, ghConfig config.GitHubConfig) error {
	if err := ghConfig.Validate(); err != nil {
		return errors.Wrap(err, "the configuration error happened on github configuration")
	}

	ghClient := client.NewGitHubClient(ghConfig.GetUsername(), ghConfig.GetRepoName(), ghConfig.GetApiToken())

	var regex *regexp.Regexp

	if pattern := ghConfig.GetFileNamePattern(); pattern != "" {
		regex = regexp.MustCompile(pattern)
	}

	release := ghClient.GetLatestRelease()

	if !lib.IsNil(ghClient.Err) {
		return errors.Wrap(ghClient.Err, "retrieving the latest release has failed")
	}

	assets := ghClient.GetAssets(release)

	if !lib.IsNil(ghClient.Err) {
		return errors.Wrap(ghClient.Err, "retrieving assets has failed")
	}

	eg := errgroup.Group{}

	for _, asset := range assets {
		if regex != nil && !regex.MatchString(asset.Name) {
			continue
		}

		a := asset

		eg.Go(func() error {
			if bytes := ghClient.DownloadAsset(a); ghClient.Err != nil {
				return errors.Wrap(ghClient.Err, fmt.Sprintf("downloading failed for %s", a.Name))
			} else {
				err := ioutil.WriteFile(fmt.Sprintf("%s/%s", rootConfig.SaveDir, a.Name), bytes, 0644)

				if err != nil {
					return errors.Wrap(err, fmt.Sprintf("writing files failed for %s", a.Name))
				}
			}

			return nil
		})
	}

	return eg.Wait()
}
