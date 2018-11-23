package config

import (
	"github.com/jmatsu/artifact-transfer/config"
	"gopkg.in/guregu/null.v3"
	"gopkg.in/urfave/cli.v2"
)

const (
	githubReleaseUsernameKey     = "username"
	githubReleaseRepoNameKey     = "reponame"
	githubReleaseStrategyKey     = "strategy"
	githubReleaseApiTokenNameKey = "api-token-name"
)

func CreateGithubReleaseConfigFlags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:  githubReleaseUsernameKey,
			Usage: "a username of a project",
		},
		&cli.StringFlag{
			Name:  githubReleaseRepoNameKey,
			Usage: "a repository name of a project",
		},
		&cli.StringFlag{
			Name:  githubReleaseApiTokenNameKey,
			Usage: "a name of a environment variable which has an api token of githubRelease",
		},
		&cli.StringFlag{
			Name:  githubReleaseStrategyKey,
			Usage: "a strategy to select a release to be updated",
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
		gitHubConfig.SetApiTokenName(null.StringFrom(c.String(githubReleaseApiTokenNameKey)))
	} else {
		gitHubConfig.SetApiTokenName(null.StringFromPtr(nil))
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
