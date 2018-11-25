package source

import (
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

func downloadFromLocal(rootConfig config.RootConfig, localConfig config.LocalConfig) error {
	if err := localConfig.Validate(); err != nil {
		return err
	}

	return local.CopyFilesTo(localConfig, rootConfig.SaveDir)
}
