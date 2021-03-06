package config

import (
	"github.com/jmatsu/transart/config"
	"gopkg.in/urfave/cli.v2"
)

const (
	localPath            = "path"
	localFileNamePattern = "file-name-pattern"
)

func CreateLocalConfigFlags() []cli.Flag {
	return []cli.Flag{
		&cli.PathFlag{
			Name:    localPath,
			Usage:   "a file path",
			Aliases: []string{"p"},
		},
		&cli.StringFlag{
			Name:    localFileNamePattern,
			Usage:   "a regexp pattern for file names to filter artifacts",
			Aliases: []string{"pattern"},
		},
	}
}

func CreateLocalConfig(c *cli.Context, project config.Project) error {
	if err := commonVerifyForAddingConfig(c); err != nil {
		return err
	}

	lc := config.LocationConfig{}

	lc.SetLocationType(config.Local)
	localConfig, err := config.NewLocalConfig(lc)

	if err != nil {
		return err
	}

	localConfig.SetPath(c.String(localPath))

	if c.IsSet(localFileNamePattern) {
		pattern := c.String(localFileNamePattern)

		localConfig.SetFileNamePattern(&pattern)
	} else {
		localConfig.SetFileNamePattern(nil)
	}

	if err := localConfig.Validate(); err != nil {
		return err
	}

	switch true {
	case c.IsSet(sourceOptionKey):
		project.AddSource(lc)
	case c.IsSet(destinationOptionKey):
		project.SetDestination(lc)
	}

	return project.SaveConfig()
}
