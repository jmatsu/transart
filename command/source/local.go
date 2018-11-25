package source

import (
	"github.com/jmatsu/transart/client"
	"github.com/jmatsu/transart/config"
	"github.com/jmatsu/transart/lib"
	"github.com/pkg/errors"
	"regexp"
)

func downloadFromLocal(rootConfig config.RootConfig, lConfig config.LocalConfig) error {
	if err := lConfig.Validate(); err != nil {
		return errors.Wrap(err, "the configuration error happened on local configuration")
	}

	var regex *regexp.Regexp

	if pattern := lConfig.GetFileNamePattern(); pattern != "" {
		regex = regexp.MustCompile(pattern)
	}

	lc := client.NewLocalClient(lConfig.GetPath())

	lc.CopyDirTo(rootConfig.SaveDir, func(s string) bool {
		return regex != nil && regex.MatchString(s)
	})

	if !lib.IsNil(lc.Err) {
		return lc.Err
	}

	return nil
}
