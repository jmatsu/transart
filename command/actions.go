package command

import (
	"github.com/jmatsu/artifact-transfer/circleci"
	"github.com/jmatsu/artifact-transfer/core"
	"github.com/jmatsu/artifact-transfer/github"
)

type Actions struct {
	CircleCI      func(rootConfig core.RootConfig, config circleci.Config) error
	GitHubRelease func(rootConfig core.RootConfig, config github.Config) error
}

func (a Actions) Source(rootConfig core.RootConfig) error {
	for _, lc := range rootConfig.Source.Locations {
		if err := a.run(rootConfig, lc); err != nil {
			return err
		}
	}

	return nil
}

func (a Actions) Destination(rootConfig core.RootConfig) error {
	return a.run(rootConfig, rootConfig.Destination.Location)
}

func (a Actions) run(rootConfig core.RootConfig, lc core.LocationConfig) error {
	t, err := lc.GetLocationType()

	if err != nil {
		return err
	}

	switch t {
	case core.CircleCI:
		c, err := circleci.NewConfig(lc)

		if err != nil {
			return err
		}

		if err := a.CircleCI(rootConfig, *c); err != nil {
			return err
		}
	case core.GitHubRelease:
		c, err := github.NewConfig(lc)

		if err != nil {
			return err
		}

		if err := a.GitHubRelease(rootConfig, *c); err != nil {
			return err
		}
	}

	return nil
}
