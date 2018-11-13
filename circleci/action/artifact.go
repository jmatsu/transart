package action

import (
	"encoding/json"
	"fmt"
	"github.com/jmatsu/artifact-transfer/circleci"
	"github.com/jmatsu/artifact-transfer/circleci/entity"
	"github.com/jmatsu/artifact-transfer/core"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

func GetArtifacts(config circleci.Config, buildNum uint) ([]entity.Artifact, error) {
	if err := config.Validate(); err != nil {
		return nil, err
	}

	var artifacts []entity.Artifact

	endpoint := circleci.ArtifactListEndpoint(config.GetVcsType(), config.GetUsername(), config.GetRepoName(), buildNum)

	if bytes, err := core.GetRequest(endpoint, circleci.NewToken(config.GetApiToken()), nil); err != nil {
		return nil, err
	} else if err := json.Unmarshal(bytes, &artifacts); err != nil {
		return nil, err
	}

	var regex *regexp.Regexp

	if pattern := config.GetFileNamePattern(); pattern != "" {
		regex = regexp.MustCompile(pattern)
	}

	var filtered []entity.Artifact

	if regex != nil {
		for _, artifact := range artifacts {
			if !regex.MatchString(artifact.Path) {
				continue
			}

			filtered = append(filtered, artifact)
		}
	} else {
		filtered = artifacts
	}

	return filtered, nil
}

func GetArtifactsFindFirst(config circleci.Config) ([]entity.Artifact, error) {
	if err := config.Validate(); err != nil {
		return nil, err
	}

	// /latest/artifacts doesn't filter `has_artifacts` so it doesn't work as expected if workflow is enabled.
	jobs, err := getJobInfos(config)

	if err != nil {
		return nil, err
	}

	for _, job := range jobs {
		if !job.HasFinished() || !job.HasArtifact {
			continue
		}

		return GetArtifacts(config, job.BuildNum)
	}

	return nil, fmt.Errorf("Not found jobs which have at least one artifacts and had finished\n")
}

func DownloadArtifact(config circleci.Config, artifact entity.Artifact) error {
	endpoint := circleci.DownloadArtifactEndpoint(artifact)

	if bytes, err := core.GetRequest(endpoint, circleci.NewToken(config.GetApiToken()), nil); err != nil {
		return err
	} else {
		filename := filepath.Base(artifact.Path)

		if f, err := os.Stat(".transart"); os.IsNotExist(err) {
			if err := os.Mkdir(".transart", os.ModePerm); err != nil {
				return err
			}
		} else if !f.IsDir() {
			return fmt.Errorf(".transart file already exists")
		}

		return ioutil.WriteFile(fmt.Sprintf(".transart/%s", filename), bytes, 0644)
	}
}
