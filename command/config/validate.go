package config

import (
	"errors"
	"github.com/jmatsu/transart/config"
	"gopkg.in/urfave/cli.v2"
)

func Validate(_ *cli.Context, project config.Project) error {
	rootConfig := project.RootConfig

	if rootConfig.Version < 1 {
		return errors.New("version must be greater than 0")
	}

	for _, lc := range rootConfig.Source.Locations {
		if t, err := lc.GetLocationType(); err != nil {
			return err
		} else {
			switch t {
			case config.CircleCI:
				if c, err := config.NewCircleCIConfig(lc); err != nil {
					return err
				} else if err := c.Validate(); err != nil {
					return err
				}
			case config.GitHubRelease:
				if c, err := config.NewGitHubConfig(lc); err != nil {
					return err
				} else if err := c.Validate(); err != nil {
					return err
				}
			case config.Local:
				if c, err := config.NewLocalConfig(lc); err != nil {
					return err
				} else if err := c.Validate(); err != nil {
					return err
				}
			}
		}
	}

	return nil
}
