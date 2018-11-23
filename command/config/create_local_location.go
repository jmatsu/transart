package config

import (
	"github.com/jmatsu/artifact-transfer/config"
	"gopkg.in/guregu/null.v3"
	"gopkg.in/urfave/cli.v2"
)

const (
	localPath            = "username"
	localFileNamePattern = "file-name-pattern"
)

func CreateLocalConfigFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:  localPath,
			Usage: "a file path",
		},
		&cli.StringFlag{
			Name:  localFileNamePattern,
			Usage: "a regexp pattern for file names to filter artifacts",
		},
	}
}

func CreateLocalConfig(c *cli.Context) error {
	rootConfig, err := commonVerifyForAddingConfig(c)

	if err != nil {
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
		localConfig.SetFileNamePattern(null.StringFrom(c.String(localFileNamePattern)))
	} else {
		localConfig.SetFileNamePattern(null.StringFromPtr(nil))
	}

	if err := localConfig.Validate(); err != nil {
		return err
	}

	switch true {
	case c.IsSet(sourceOptionKey):
		rootConfig.Source.Locations = append(rootConfig.Source.Locations, lc)
	case c.IsSet(destinationOptionKey):
		rootConfig.Destination.Location = lc
	}

	return rootConfig.Save()
}
