package config

import (
	"github.com/jmatsu/artifact-transfer/circleci"
	"github.com/jmatsu/artifact-transfer/core"
	"gopkg.in/guregu/null.v3"
	"gopkg.in/urfave/cli.v2"
)

const (
	circleciVcsTypeKey      = "vcs-type"
	circleciUsernameKey     = "username"
	circleciRepoNameKey     = "reponame"
	circleciBranchKey       = "branch"
	circleciApiTokenNameKey = "api-token-name"
	circleciFileNamePattern = "file-name-pattern"
)

func CreateCircleCIConfigFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:  circleciVcsTypeKey,
			Usage: "a directory path of artifacts to be saved",
		},
		&cli.StringFlag{
			Name:  circleciUsernameKey,
			Usage: "a username of a project",
		},
		&cli.StringFlag{
			Name:  circleciRepoNameKey,
			Usage: "a repository name of a project",
		},
		&cli.StringFlag{
			Name:  circleciApiTokenNameKey,
			Usage: "a name of a environment variable which has an api token of CircleCI",
		},
		&cli.StringFlag{
			Name:  circleciBranchKey,
			Usage: "a branch to be filtered",
		},
		&cli.StringFlag{
			Name:  circleciFileNamePattern,
			Usage: "a regexp pattern for file names to filter artifacts",
		},
	}
}

func CreateCircleCIConfig(c *cli.Context) error {
	rootConfig, err := commonVerifyForAddingConfig(c)

	if err != nil {
		return err
	}

	lc := core.LocationConfig{}

	lc.SetLocationType(core.CircleCI)
	config, err := circleci.NewConfig(lc)

	if err != nil {
		return err
	}

	config.SetVcsType(circleci.VcsType(c.String(circleciVcsTypeKey)))
	config.SetUsername(c.String(circleciUsernameKey))
	config.SetRepoName(c.String(circleciRepoNameKey))

	if c.IsSet(circleciApiTokenNameKey) {
		config.SetApiTokenName(null.StringFrom(c.String(circleciApiTokenNameKey)))
	} else {
		config.SetApiTokenName(null.StringFromPtr(nil))
	}

	if c.IsSet(circleciBranchKey) {
		config.SetBranch(null.StringFrom(c.String(circleciBranchKey)))
	} else {
		config.SetBranch(null.StringFromPtr(nil))
	}

	if c.IsSet(circleciFileNamePattern) {
		config.SetFileNamePattern(null.StringFrom(c.String(circleciFileNamePattern)))
	} else {
		config.SetFileNamePattern(null.StringFromPtr(nil))
	}

	config.Validate()

	if config.Err != nil {
		return config.Err
	}

	switch true {
	case c.IsSet(sourceOptionKey):
		rootConfig.Source.Locations = append(rootConfig.Source.Locations, lc)
	case c.IsSet(destinationOptionKey):
		rootConfig.Destination.Location = lc
	}

	return rootConfig.Save()
}
