package command

import (
	"github.com/jmatsu/transart/config"
)

type Actions struct {
	CircleCI      func(rootConfig config.RootConfig, circleCIConfig config.CircleCIConfig) error
	GitHubRelease func(rootConfig config.RootConfig, gitHubConfig config.GitHubConfig) error
	Local         func(rootConfig config.RootConfig, localConfig config.LocalConfig) error
}

func (a Actions) Source(rootConfig config.RootConfig) error {
	for _, lc := range rootConfig.Source.Locations {
		if err := a.run(rootConfig, lc); err != nil {
			return err
		}
	}

	return nil
}

func (a Actions) Destination(rootConfig config.RootConfig) error {
	return a.run(rootConfig, rootConfig.Destination.Location)
}

func (a Actions) run(rootConfig config.RootConfig, lc config.LocationConfig) error {
	t, err := lc.GetLocationType()

	if err != nil {
		return err
	}

	switch t {
	case config.CircleCI:
		c, err := config.NewCircleCIConfig(lc)

		if err != nil {
			return err
		}

		if err := a.CircleCI(rootConfig, *c); err != nil {
			return err
		}
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
