package source

import (
	"fmt"
	"github.com/jmatsu/transart/client"
	"github.com/jmatsu/transart/client/entity"
	"github.com/jmatsu/transart/config"
	"github.com/jmatsu/transart/lib"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
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
		return regex != nil && regex.MatchString(artifact.Path)
	})

	if !lib.IsNil(ccClient.Err) {
		return errors.Wrap(ccClient.Err, "retrieving artifact information has failed")
	}

	for _, artifact := range artifacts {
		if bytes := ccClient.DownloadArtifact(artifact); ccClient.Err != nil {
			return errors.Wrap(ccClient.Err, "")
		} else {
			filename := filepath.Base(artifact.Path)

			if f, err := os.Stat(rootConfig.SaveDir); os.IsNotExist(err) {
				if err := os.MkdirAll(rootConfig.SaveDir, os.ModePerm); err != nil {
					return err
				}
			} else if !f.IsDir() {
				return fmt.Errorf("%s already exists but it's a file", rootConfig.SaveDir)
			}

			err := ioutil.WriteFile(fmt.Sprintf("%s/%s", rootConfig.SaveDir, filename), bytes, 0644)

			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("downloading failed for %s", artifact.Path))
			}
		}
	}

	return nil
}
