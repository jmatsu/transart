package command

import (
	"github.com/jmatsu/artifact-transfer/circleci"
	"github.com/jmatsu/artifact-transfer/core"
	"github.com/jmatsu/artifact-transfer/github"
)

type Actions struct {
	CircleCI      func(config circleci.Config) error
	GitHubRelease func(config github.Config) error
}

func (a Actions) Source(config core.RootConfig) error {
	for _, lc := range config.Source.Locations {
		if err := a.run(lc); err != nil {
			return err
		}
	}

	return nil
}

func (a Actions) Destination(config core.RootConfig) error {
	return a.run(config.Destination.Location)
}

func (a Actions) run(lc core.LocationConfig) error {
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

		if err := a.CircleCI(c); err != nil {
			return err
		}
	case core.GitHubRelease:
		c, err := github.NewConfig(lc)

		if err != nil {
			return err
		}

		if err := a.GitHubRelease(c); err != nil {
			return err
		}
	}

	return nil
}
