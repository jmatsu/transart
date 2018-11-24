package command

import "gopkg.in/urfave/cli.v2"

const ConfFileOptionKey = "file"

func CommonFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    ConfFileOptionKey,
			Aliases: []string{"f"},
			Usage:   "a file path to a configuration file",
			Value:   ".transart.yml",
		},
	}
}
