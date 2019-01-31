package source

import (
	"fmt"
	"github.com/jmatsu/transart/client"
	"github.com/jmatsu/transart/client/entity"
	"github.com/jmatsu/transart/config"
	"github.com/jmatsu/transart/lib"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"io/ioutil"
	"path/filepath"
	"regexp"
)

func downloadFromCircleCI(rootConfig config.RootConfig, ccConfig config.CircleCIConfig) error {
	if err := ccConfig.Validate(); err != nil {
		return errors.Wrap(err, "the configuration error happened on circleci configuration")
	}

	ccClient := client.NewCircleCIClient(string(ccConfig.GetVcsType()), ccConfig.GetUsername(), ccConfig.GetRepoName(), ccConfig.GetApiToken(), ccConfig.GetBranch())

	jobInfo := ccClient.GetJobInfo(func(info entity.CircleCIJobInfo) bool {
		return info.HasFinished()
	})

	if !lib.IsNil(ccClient.Err) {
		return errors.Wrap(ccClient.Err, "retrieving job infos has failed")
	}

	var regex *regexp.Regexp

	if pattern := ccConfig.GetFileNamePattern(); pattern != "" {
		regex = regexp.MustCompile(pattern)
	}

	artifacts := ccClient.GetArtifacts(jobInfo.BuildNum, func(artifact entity.CircleCIArtifact) bool {
		return regex == nil || regex.MatchString(artifact.Path)
	})

	if !lib.IsNil(ccClient.Err) {
		return errors.Wrap(ccClient.Err, "retrieving artifact information has failed")
	}

	eg := errgroup.Group{}

	for _, artifact := range artifacts {
		a := artifact

		eg.Go(func() error {
			if bytes := ccClient.DownloadArtifact(a); ccClient.Err != nil {
				return errors.Wrap(ccClient.Err, fmt.Sprintf("downloading failed for %s", a.Path))
			} else {
				filename := filepath.Base(a.Path)

				err := ioutil.WriteFile(fmt.Sprintf("%s/%s", rootConfig.SaveDir, filename), bytes, 0644)

				if err != nil {
					return errors.Wrap(err, fmt.Sprintf("writing files failed for %s", a.Path))
				}
			}

			return nil
		})
	}

	return eg.Wait()
}
