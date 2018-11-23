package config

import (
	"fmt"
	"github.com/jmatsu/transart/config"
	"gopkg.in/urfave/cli.v2"
)

const (
	githubReleaseUsernameKey     = "username"
	githubReleaseRepoNameKey     = "reponame"
	githubReleaseApiTokenNameKey = "api-token-name"
	githubReleaseStrategyKey     = "strategy"
)

func CreateGithubReleaseConfigFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:    githubReleaseUsernameKey,
			Usage:   "a username of a project",
			Aliases: []string{"u"},
		},
		&cli.StringFlag{
			Name:    githubReleaseRepoNameKey,
			Usage:   "a repository name of a project",
			Aliases: []string{"r"},
		},
		&cli.StringFlag{
			Name:    githubReleaseApiTokenNameKey,
			Usage:   "a name of a environment variable which has an api token of GitHub",
			Aliases: []string{"token-name"},
		},
		&cli.StringFlag{
			Name:    githubReleaseStrategyKey,
			Usage:   fmt.Sprintf("a strategy to select a release to be updated. either of %s, %s, %s", config.Draft, config.Create, config.DraftOrCreate),
			Aliases: []string{"s"},
		},
	}
}

func CreateGithubReleaseConfig(c *cli.Context) error {
	rootConfig, err := commonVerifyForAddingConfig(c)

	if err != nil {
		return err
	}

	lc := config.LocationConfig{}

	lc.SetLocationType(config.GitHubRelease)
	gitHubConfig, err := config.NewGitHubConfig(lc)

	if err != nil {
		return err
	}

	gitHubConfig.SetUsername(c.String(githubReleaseUsernameKey))
	gitHubConfig.SetRepoName(c.String(githubReleaseRepoNameKey))
	gitHubConfig.SetStrategy(config.GitHubReleaseCreationStrategy(c.String(githubReleaseStrategyKey)))

	if c.IsSet(githubReleaseApiTokenNameKey) {
		name := c.String(githubReleaseApiTokenNameKey)

		gitHubConfig.SetApiTokenName(&name)
	} else {
		gitHubConfig.SetApiTokenName(nil)
	}

	if err := gitHubConfig.Validate(); err != nil {
		return err
	}

	switch true {
	case c.IsSet(sourceOptionKey):
		rootConfig.Source.Locations = append(rootConfig.Source.Locations, lc)
	case c.IsSet(destinationOptionKey):
		rootConfig.Destination.Location = lc
	}

	return rootConfig.Save()
}
