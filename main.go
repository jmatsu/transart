package main

import (
	"fmt"
	"github.com/jmatsu/transart/command"
	configCommand "github.com/jmatsu/transart/command/config"
	"github.com/jmatsu/transart/command/destination"
	"github.com/jmatsu/transart/command/source"
	"github.com/jmatsu/transart/config"
	"github.com/jmatsu/transart/version"
	"github.com/sirupsen/logrus"
	"gopkg.in/urfave/cli.v2"
	"os"
	"strconv"
)

func main() {
	if b, err := strconv.ParseBool(os.Getenv("TRANSART_DEBUG")); err == nil && b {
		logrus.SetLevel(logrus.DebugLevel)
	}

	cli.InitCompletionFlag.Hidden = true

	cli.AppHelpTemplate = fmt.Sprintf(`%s
COMPLETION:
	transart --init-completion <bash|zsh>
WEBSITE:
	https://github.com/jmatsu/transart
SUPPORT:
	https://github.com/jmatsu/transart/issues
`, cli.AppHelpTemplate)

	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Println(version.Template())
	}

	app := &cli.App{}
	app.Name = "transart"
	app.Usage = "Transfer CI Artifacts"
	app.Description = "transart is a command line tool to transfer CI artifacts to anywhere like GitHub Release."
	app.Version = version.Version
	app.EnableShellCompletion = true

	hub := func(action func(context *cli.Context, confFileName string) error) func(context *cli.Context) error {
		return func(context *cli.Context) error {
			confFileName := context.String(command.ConfFileOptionKey)

			return action(context, confFileName)
		}
	}

	app.Flags = command.CommonFlags()
	app.Commands = []*cli.Command{
		{
			Name:   "init",
			Usage:  "Create an initial root configuration file",
			Action: hub(configCommand.CreateRootConfig),
			Flags:  configCommand.CreateRootConfigFlags(),
		},
		{
			Name:   "validate",
			Usage:  "Validate a configuration file",
			Action: hub(configCommand.Validate),
		},
		{
			Name:  "add",
			Usage: "Add a new configuration of a location",
			Subcommands: []*cli.Command{
				{
					Name:   "circleci",
					Usage:  "Create a configuration for CircleCI",
					Action: hub(configCommand.CreateCircleCIConfig),
					Flags:  append(configCommand.CreateAddLocationFlags(), configCommand.CreateCircleCIConfigFlags()...),
				},
				{
					Name:   "github-release",
					Usage:  "Create a configuration for GitHub Release",
					Action: hub(configCommand.CreateGithubReleaseConfig),
					Flags:  append(configCommand.CreateAddLocationFlags(), configCommand.CreateGithubReleaseConfigFlags()...),
				},
				{
					Name:   "local",
					Usage:  "Create a configuration for local file system",
					Action: hub(configCommand.CreateLocalConfig),
					Flags:  append(configCommand.CreateAddLocationFlags(), configCommand.CreateLocalConfigFlags()...),
				},
			},
		},
		{
			Name:  "transfer",
			Usage: "Download artifacts and assets from sources, and upload them to the destination",
			Action: hub(func(context *cli.Context, confFileName string) error {
				rootConfig, err := config.LoadRootConfig(confFileName)

				if err != nil {
					return err
				}

				if err := source.NewDownloadAction().Source(*rootConfig); err != nil {
					return err
				}

				if err := destination.NewUploadAction().Destination(*rootConfig); err != nil {
					return err
				}

				return nil
			}),
		},
		{
			Name:  "download",
			Usage: "Download artifacts and assets from sources",
			Action: hub(func(context *cli.Context, confFileName string) error {
				rootConfig, err := config.LoadRootConfig(confFileName)

				if err != nil {
					return err
				}

				return source.NewDownloadAction().Source(*rootConfig)
			}),
		},
		{
			Name:  "upload",
			Usage: "Upload artifacts and assets the destination",
			Action: hub(func(context *cli.Context, confFileName string) error {
				rootConfig, err := config.LoadRootConfig(confFileName)

				if err != nil {
					return err
				}

				return destination.NewUploadAction().Source(*rootConfig)
			}),
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
