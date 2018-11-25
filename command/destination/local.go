package destination

import (
	"fmt"
	"github.com/jmatsu/transart/config"
	"github.com/jmatsu/transart/lib"
	"github.com/jmatsu/transart/local"
	"os"
)

func uploadToLocal(rootConfig config.RootConfig, localConfig config.LocalConfig) error {
	if err := localConfig.Validate(); err != nil {
		return err
	}

	return lib.ForEachFiles(rootConfig.SaveDir, func(dirname string, info os.FileInfo) error {
		if err := local.CopyFileFrom(localConfig, fmt.Sprintf("%s/%s", dirname, info.Name())); err != nil {
			return err
		}

		return nil
	})
}
