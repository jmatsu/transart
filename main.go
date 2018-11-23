package main

import (
	"fmt"
	configCommand "github.com/jmatsu/artifact-transfer/command/config"
	"github.com/jmatsu/artifact-transfer/command/destination"
	"github.com/jmatsu/artifact-transfer/command/source"
	"github.com/jmatsu/artifact-transfer/config"
	"github.com/jmatsu/artifact-transfer/version"
	"github.com/sirupsen/logrus"
	"gopkg.in/urfave/cli.v2"
	"os"
	"strconv"
)

func main() {
	if b, err := strconv.ParseBool(os.Getenv("ARTIFACT_TRANSFER_DEBUG")); err == nil && b {
		logrus.SetLevel(logrus.DebugLevel)
	}

	cli.InitCompletionFlag.Hidden = true

	cli.AppHelpTemplate = fmt.Sprintf(`%s
COMPLETION:
	transart --init-completion <bash|zsh>
WEBSITE:
	https://github.com/jmatsu/artifact-transfer
SUPPORT:
	https://github.com/jmatsu/artifact-transfer/issues
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

	app.Commands = []*cli.Command{
		{
			Name:   "init",
			Usage:  "Create an initial root configuration file",
			Action: configCommand.CreateRootConfig,
			Flags:  configCommand.CreateRootConfigFlags(),
		},
		{
			Name:  "add",
			Usage: "Add a new configuration of a location",
			Flags: configCommand.CreateAddLocationFlags(),
			Subcommands: []*cli.Command{
				{
					Name:   "circleci",
					Usage:  "Create a configuration for CircleCI",
					Action: configCommand.CreateCircleCIConfig,
					Flags:  configCommand.CreateCircleCIConfigFlags(),
				},
				{
					Name:   "github-release",
					Usage:  "Create a configuration for GitHub Release",
					Action: configCommand.CreateGithubReleaseConfig,
					Flags:  configCommand.CreateGithubReleaseConfigFlags(),
				},
			},
		},
		{
			Name:  "transfer",
			Usage: "Download artifacts and assets from sources, and upload them to the destination",
			Action: func(context *cli.Context) error {
				rootConfig, err := config.LoadRootConfig()

				if err != nil {
					return err
				}

				if err := source.NewDownloadAction().Source(*rootConfig); err != nil {
					return err
				}

				if err := destination.NewUploadAction(*context).Destination(*rootConfig); err != nil {
					return err
				}

				return nil
			},
		},
		{
			Name:  "download",
			Usage: "Download artifacts and assets from sources",
			Action: func(context *cli.Context) error {
				rootConfig, err := config.LoadRootConfig()

				if err != nil {
					return err
				}

				return source.NewDownloadAction().Source(*rootConfig)
			},
		},
		{
			Name:  "upload",
			Usage: "Upload artifacts and assets the destination",
			Action: func(context *cli.Context) error {
				rootConfig, err := config.LoadRootConfig()

				if err != nil {
					return err
				}

				return destination.NewUploadAction(*context).Destination(*rootConfig)
			},
		},
	}

	err := app.Run(os.Args)

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
