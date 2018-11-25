package source

import (
	"github.com/jmatsu/transart/config"
	"github.com/jmatsu/transart/local"
)

func downloadFromLocal(rootConfig config.RootConfig, localConfig config.LocalConfig) error {
	if err := localConfig.Validate(); err != nil {
		return err
	}

	return local.CopyFilesTo(localConfig, rootConfig.SaveDir)
}
