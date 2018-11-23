package source

import (
	"github.com/jmatsu/transart/circleci"
	"github.com/jmatsu/transart/command"
	"github.com/jmatsu/transart/config"
	"github.com/jmatsu/transart/local"
)

func NewDownloadAction() command.Actions {
	return command.Actions{
		CircleCI: downloadFromCircleCI,
		Local:    downloadFromLocal,
	}
}

func downloadFromCircleCI(rootConfig config.RootConfig, circleCIConfig config.CircleCIConfig) error {
	artifacts, err := circleci.GetArtifactsFindFirst(circleCIConfig)

	if err != nil {
		return err
	}

	for _, artifact := range artifacts {
		if err := circleci.DownloadArtifact(rootConfig, circleCIConfig, artifact); err != nil {
			return err
		}
	}

	return nil
}

func downloadFromLocal(rootConfig config.RootConfig, localConfig config.LocalConfig) error {
	if err := localConfig.Validate(); err != nil {
		return err
	}

	return local.CopyFilesTo(localConfig, rootConfig.SaveDir)
}
