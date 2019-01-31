package destination

import (
	"fmt"
	"github.com/jmatsu/transart/config"
	"os"
)

type Actions struct {
	// CircleCI cannot be supported
	GitHubRelease func(rootConfig config.RootConfig, gitHubConfig config.GitHubConfig) error
	Local         func(rootConfig config.RootConfig, localConfig config.LocalConfig) error
}

func (a Actions) Run(rootConfig config.RootConfig) error {
	if f, err := os.Stat(rootConfig.SaveDir); os.IsNotExist(err) {
		if err := os.MkdirAll(rootConfig.SaveDir, os.ModePerm); err != nil {
			return err
		}
	} else if !f.IsDir() {
		return fmt.Errorf("intermediate directory - %s -  already exists but it's a file", rootConfig.SaveDir)
	}

	lc := rootConfig.Destination.Location

	t, err := lc.GetLocationType()

	if err != nil {
		return err
	}

	switch t {
	case config.CircleCI:
		return fmt.Errorf("%v is not supported", t)
	case config.GitHubRelease:
		c, err := config.NewGitHubConfig(lc)

		if err != nil {
			return err
		}

		if err := a.GitHubRelease(rootConfig, *c); err != nil {
			return err
		}
	case config.Local:
		c, err := config.NewLocalConfig(lc)

		if err != nil {
			return err
		}

		if err := a.Local(rootConfig, *c); err != nil {
			return err
		}
	}

	return nil
}
