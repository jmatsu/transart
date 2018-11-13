package source

import (
	"github.com/jmatsu/artifact-transfer/circleci"
	"github.com/jmatsu/artifact-transfer/circleci/action"
	"github.com/jmatsu/artifact-transfer/command"
)

func NewDownloadAction() command.Actions {
	return command.Actions{
		CircleCI: func(config circleci.Config) error {
			artifacts, err := action.GetArtifactsFindFirst(config)

			if err != nil {
				return err
			}

			for _, artifact := range artifacts {
				if err := action.DownloadArtifact(config, artifact); err != nil {
					return err
				}
			}

			return nil
		},
	}
}
