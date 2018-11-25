package config

import (
	"github.com/jmatsu/transart/config"
	"gopkg.in/urfave/cli.v2"
)

const (
	circleciVcsTypeKey      = "vcs-type"
	circleciUsernameKey     = "username"
	circleciRepoNameKey     = "reponame"
	circleciBranchKey       = "branch"
	circleciApiTokenNameKey = "api-token-name"
	circleciFileNamePattern = localFileNamePattern
)

func CreateCircleCIConfigFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    circleciVcsTypeKey,
			Usage:   "a directory path of artifacts to be saved. either of github or bitbucket",
			Aliases: []string{"v"},
		},
		&cli.StringFlag{
			Name:    circleciUsernameKey,
			Usage:   "a username of a project",
			Aliases: []string{"u"},
		},
		&cli.StringFlag{
			Name:    circleciRepoNameKey,
			Usage:   "a repository name of a project",
			Aliases: []string{"r"},
		},
		&cli.StringFlag{
			Name:    circleciApiTokenNameKey,
			Usage:   "a name of a environment variable which has an api token of CircleCI",
			Aliases: []string{"token-name"},
		},
		&cli.StringFlag{
			Name:    circleciBranchKey,
			Usage:   "a branch to be filtered",
			Aliases: []string{"b"},
		},
		&cli.StringFlag{
			Name:    circleciFileNamePattern,
			Usage:   "a regexp pattern for file names to filter artifacts",
			Aliases: []string{"pattern"},
		},
	}
}

func CreateCircleCIConfig(c *cli.Context, project config.Project) error {
	if err := commonVerifyForAddingConfig(c); err != nil {
		return err
	}

	if c.IsSet(destinationOptionKey) {
		return destinationNotSupported(config.CircleCI)
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
		name := c.String(circleciApiTokenNameKey)

		circleCIConfig.SetApiTokenName(&name)
	} else {
		circleCIConfig.SetApiTokenName(nil)
	}

	if c.IsSet(circleciBranchKey) {
		branch := c.String(circleciBranchKey)

		circleCIConfig.SetBranch(&branch)
	} else {
		circleCIConfig.SetBranch(nil)
	}

	if c.IsSet(circleciFileNamePattern) {
		pattern := c.String(circleciFileNamePattern)

		circleCIConfig.SetFileNamePattern(&pattern)
	} else {
		circleCIConfig.SetFileNamePattern(nil)
	}

	circleCIConfig.Validate()

	if circleCIConfig.Err != nil {
		return circleCIConfig.Err
	}

	switch true {
	case c.IsSet(sourceOptionKey):
		project.AddSource(lc)
	case c.IsSet(destinationOptionKey):
		project.SetDestination(lc)
	}

	return project.SaveConfig()
}
