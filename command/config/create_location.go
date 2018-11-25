package config

import (
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

func commonVerifyForAddingConfig(c *cli.Context) error {
	if !c.IsSet(sourceOptionKey) && !c.IsSet(destinationOptionKey) {
		return locationTypeIsRequired
	}

	return nil
}
