package config

import (
	"errors"
	"github.com/jmatsu/transart/config"
	"gopkg.in/urfave/cli.v2"
)

const (
	saveDirOptionKey = "save-dir"
	forceOptionKey   = "force"
)

func CreateRootConfigFlags() []cli.Flag {
	return []cli.Flag{
		&cli.PathFlag{
			Name:    saveDirOptionKey,
			Usage:   "a directory path of artifacts to be saved",
			Aliases: []string{"p"},
			Value:   ".transart",
		},
		&cli.BoolFlag{
			Name:    forceOptionKey,
			Usage:   "force to do it",
			Aliases: []string{"f"},
			Value:   false,
		},
	}
}

func CreateRootConfig(c *cli.Context, confFileName string) error {
	if config.ExistsRootConfig(confFileName) {
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

	project := config.Project{
		ConfFileName: confFileName,
		RootConfig: config.RootConfig{
			Version: 1,
			SaveDir: saveDir,
			Source: config.SourceConfig{
				Locations: []config.LocationConfig{},
			},
			Destination: config.DestinationConfig{
				Location: config.LocationConfig{},
			},
		},
	}

	return project.SaveConfig()
}
