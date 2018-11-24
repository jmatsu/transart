package config

import (
	"fmt"
	"github.com/jmatsu/transart/config"
	"gopkg.in/urfave/cli.v2"
)

const (
	sourceOptionKey      = "source"
	destinationOptionKey = "destination"
)

func CreateAddLocationFlags() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:    sourceOptionKey,
			Usage:   "operate source configuration if specified",
			Value:   false,
			Aliases: []string{"src"},
		},
		&cli.BoolFlag{
			Name:    destinationOptionKey,
			Usage:   "operate source configuration if specified",
			Value:   false,
			Aliases: []string{"dest"},
		},
	}
}

func commonVerifyForAddingConfig(c *cli.Context, confFileName string) (*config.RootConfig, error) {
	if !c.IsSet(sourceOptionKey) && !c.IsSet(destinationOptionKey) {
		return nil, fmt.Errorf("either of --%s or --%s is required", sourceOptionKey, destinationOptionKey)
	}

	return config.LoadRootConfig(confFileName)
}
