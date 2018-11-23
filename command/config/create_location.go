package config

import (
	"fmt"
	"github.com/jmatsu/artifact-transfer/core"
	"gopkg.in/urfave/cli.v2"
)

const (
	sourceOptionKey      = "source"
	destinationOptionKey = "destination"
)

func CreateAddLocationFlags() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:  sourceOptionKey,
			Usage: "operate source configuration if specified",
			Value: false,
		},
		&cli.BoolFlag{
			Name:  destinationOptionKey,
			Usage: "operate source configuration if specified",
			Value: false,
		},
	}
}

func commonVerifyForAddingConfig(c *cli.Context) (*core.RootConfig, error) {
	if !c.IsSet(sourceOptionKey) && !c.IsSet(destinationOptionKey) {
		return nil, fmt.Errorf("either of --%s or --%s is required", sourceOptionKey, destinationOptionKey)
	}

	return core.LoadRootConfig()
}
