package source

import (
	"fmt"
	"github.com/jmatsu/transart/config"
	"golang.org/x/sync/errgroup"
)

type Actions struct {
	CircleCI func(rootConfig config.RootConfig, circleCIConfig config.CircleCIConfig) error
	// TODO GitHubRelease
	Local func(rootConfig config.RootConfig, localConfig config.LocalConfig) error
}

func (a Actions) Run(rootConfig config.RootConfig) error {
	eg := errgroup.Group{}

	for _, location := range rootConfig.Source.Locations {
		l := location

		eg.Go(func() error {
			return a.run(rootConfig, l)
		})
	}

	return eg.Wait()
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
		return fmt.Errorf("%v is not supported", t)
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
