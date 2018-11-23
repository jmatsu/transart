package circleci

import (
	"encoding/json"
	"fmt"
	"github.com/jmatsu/transart/circleci/entity"
	"github.com/jmatsu/transart/config"
	"github.com/jmatsu/transart/lib"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
)

func GetArtifacts(circleCIConfig config.CircleCIConfig, buildNum uint) ([]entity.Artifact, error) {
	if err := circleCIConfig.Validate(); err != nil {
		return nil, err
	}

	var artifacts []entity.Artifact

	endpoint := ArtifactListEndpoint(string(circleCIConfig.GetVcsType()), circleCIConfig.GetUsername(), circleCIConfig.GetRepoName(), buildNum)

	if bytes, err := lib.GetRequest(endpoint, NewToken(circleCIConfig.GetApiToken()), nil); err != nil {
		return nil, err
	} else if err := json.Unmarshal(bytes, &artifacts); err != nil {
		return nil, err
	}

	var regex *regexp.Regexp

	if pattern := circleCIConfig.GetFileNamePattern(); pattern != "" {
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

func GetArtifactsFindFirst(circleCIConfig config.CircleCIConfig) ([]entity.Artifact, error) {
	if err := circleCIConfig.Validate(); err != nil {
		return nil, err
	}

	// /latest/artifacts doesn't filter `has_artifacts` so it doesn't work as expected if workflow is enabled.
	jobs, err := getJobInfos(circleCIConfig)

	if err != nil {
		return nil, err
	}

	for _, job := range jobs {
		if !job.HasFinished() || !job.HasArtifact {
			continue
		}

		return GetArtifacts(circleCIConfig, job.BuildNum)
	}

	return nil, fmt.Errorf("Not found jobs which have at least one artifacts and had finished\n")
}

func DownloadArtifact(rootConfig config.RootConfig, circleCIConfig config.CircleCIConfig, artifact entity.Artifact) error {
	endpoint := DownloadArtifactEndpoint(artifact)

	if bytes, err := lib.GetRequest(endpoint, NewToken(circleCIConfig.GetApiToken()), nil); err != nil {
		return err
	} else {
		filename := filepath.Base(artifact.Path)

		if f, err := os.Stat(rootConfig.SaveDir); os.IsNotExist(err) {
			if err := os.MkdirAll(rootConfig.SaveDir, os.ModePerm); err != nil {
				return err
			}
		} else if !f.IsDir() {
			return fmt.Errorf("%s file already exists", rootConfig.SaveDir)
		}

		return ioutil.WriteFile(fmt.Sprintf("%s/%s", rootConfig.SaveDir, filename), bytes, 0644)
	}
}
