package config

import (
	"github.com/jmatsu/artifact-transfer/config"
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

	lc := config.LocationConfig{}

	lc.SetLocationType(config.CircleCI)
	circleCIConfig, err := config.NewCircleCIConfig(lc)

	if err != nil {
		return err
	}

	circleCIConfig.SetVcsType(config.VcsType(c.String(circleciVcsTypeKey)))
	circleCIConfig.SetUsername(c.String(circleciUsernameKey))
	circleCIConfig.SetRepoName(c.String(circleciRepoNameKey))

	if c.IsSet(circleciApiTokenNameKey) {
		circleCIConfig.SetApiTokenName(null.StringFrom(c.String(circleciApiTokenNameKey)))
	} else {
		circleCIConfig.SetApiTokenName(null.StringFromPtr(nil))
	}

	if c.IsSet(circleciBranchKey) {
		circleCIConfig.SetBranch(null.StringFrom(c.String(circleciBranchKey)))
	} else {
		circleCIConfig.SetBranch(null.StringFromPtr(nil))
	}

	if c.IsSet(circleciFileNamePattern) {
		circleCIConfig.SetFileNamePattern(null.StringFrom(c.String(circleciFileNamePattern)))
	} else {
		circleCIConfig.SetFileNamePattern(null.StringFromPtr(nil))
	}

	circleCIConfig.Validate()

	if circleCIConfig.Err != nil {
		return circleCIConfig.Err
	}

	switch true {
	case c.IsSet(sourceOptionKey):
		rootConfig.Source.Locations = append(rootConfig.Source.Locations, lc)
	case c.IsSet(destinationOptionKey):
		rootConfig.Destination.Location = lc
	}

	return rootConfig.Save()
}
