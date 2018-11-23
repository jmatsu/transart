package source

import (
	"github.com/jmatsu/artifact-transfer/circleci"
	"github.com/jmatsu/artifact-transfer/command"
	"github.com/jmatsu/artifact-transfer/config"
)

func NewDownloadAction() command.Actions {
	return command.Actions{
		CircleCI: func(rootConfig config.RootConfig, circleCIConfig config.CircleCIConfig) error {
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
		},
	}
}
