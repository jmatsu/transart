package config

import (
	"errors"
	"github.com/jmatsu/artifact-transfer/core"
	"gopkg.in/urfave/cli.v2"
)

const (
	saveDirOptionKey = "save-dir"
	forceOptionKey   = "force"
)

func CreateRootConfig(c *cli.Context) error {
	if core.ExistsRootConfig() {
		if !c.Bool(forceOptionKey) {
			return errors.New("a config file already exists. cannot overwrite without --force option")
		}
	}

	var saveDir string

	if c.IsSet(saveDirOptionKey) {
		saveDir = c.String(saveDirOptionKey)
	} else {
		saveDir = ".transart"
	}

	if saveDir == "" {
		return errors.New("empty directory name is not allowed")
	}

	config := core.RootConfig{
		Version: 1,
		SaveDir: saveDir,
		Source: core.SourceConfig{
			Locations: []core.LocationConfig{},
		},
		Destination: core.DestinationConfig{
			Location: core.LocationConfig{},
		},
	}

	return config.Save()
}

func CreateRootConfigFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:  saveDirOptionKey,
			Usage: "a directory path of artifacts to be saved",
		},
		&cli.BoolFlag{
			Name:  forceOptionKey,
			Usage: "force to do it",
		},
	}
}
