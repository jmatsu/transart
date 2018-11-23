package source

import (
	"github.com/jmatsu/artifact-transfer/circleci"
	"github.com/jmatsu/artifact-transfer/circleci/action"
	"github.com/jmatsu/artifact-transfer/command"
	"github.com/jmatsu/artifact-transfer/core"
)

func NewDownloadAction() command.Actions {
	return command.Actions{
		CircleCI: func(rootConfig core.RootConfig, config circleci.Config) error {
			artifacts, err := action.GetArtifactsFindFirst(config)

			if err != nil {
				return err
			}

			for _, artifact := range artifacts {
				if err := action.DownloadArtifact(rootConfig, config, artifact); err != nil {
					return err
				}
			}

			return nil
		},
	}
}
