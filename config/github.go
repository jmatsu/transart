package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/guregu/null.v3"
	"os"
	"regexp"
)

type (
	GitHubReleaseCreationStrategy string

	GitHubConfig struct {
		values LocationConfig
		Err    error
	}
)

const (
	Draft         GitHubReleaseCreationStrategy = "draft"
	DraftOrCreate GitHubReleaseCreationStrategy = "draft-or-create"
	Create        GitHubReleaseCreationStrategy = "create"
)

func isGitHubReleaseCreationStrategy(s string) bool {
	strategies := []GitHubReleaseCreationStrategy{
		Draft,
		DraftOrCreate,
		Create,
	}

	for _, strategy := range strategies {
		if strategy == GitHubReleaseCreationStrategy(s) {
			return true
		}
	}

	return false
}

func NewGitHubConfig(lc LocationConfig) (*GitHubConfig, error) {
	if t, err := lc.GetLocationType(); err != nil || t != GitHubRelease {
		if err == nil {
			err = fmt.Errorf("location type is not for github releases so the caller is wrong")
		}

		return nil, err
	}

	config := &GitHubConfig{
		values: lc,
	}

	return config, nil
}

func (c *GitHubConfig) setError(_ interface{}, err error) {
	if c.Err != nil {
		return
	}

	c.Err = err
}

func (c *GitHubConfig) Validate() error {
	c.setError(c.getUsername())
	c.setError(c.getRepoName())
	c.setError(c.getStrategy())

	return c.Err
}

func (c *GitHubConfig) getUsername() (string, error) {
	if c.values.Has(usernameKey) {
		return c.values[usernameKey].(string), nil
	}

	return "", fmt.Errorf("%s is missinge\n", usernameKey)
}

func (c *GitHubConfig) GetUsername() string {
	if t, err := c.getUsername(); err != nil {
		panic(err)
	} else {
		return t
	}
}

func (c *GitHubConfig) SetUsername(v string) {
	c.values.Set(usernameKey, v)
}

func (c *GitHubConfig) getRepoName() (string, error) {
	if c.values.Has(repoNameKey) {
		return c.values[repoNameKey].(string), nil
	}

	return "", fmt.Errorf("%s is missing\n", repoNameKey)
}

func (c *GitHubConfig) GetRepoName() string {
	if t, err := c.getRepoName(); err != nil {
		panic(err)
	} else {
		return t
	}
}

func (c *GitHubConfig) SetRepoName(v string) {
	c.values.Set(repoNameKey, v)
}

func (c *GitHubConfig) GetApiToken() null.String {
	var name string

	if c.values.Has(apiTokenNameKey) {
		name = c.values[apiTokenNameKey].(string)
	} else {
		name = "GITHUB_TOKEN"
	}

	if v, ok := os.LookupEnv(name); ok {
		return null.StringFrom(v)
	}

	return null.StringFromPtr(nil)
}

func (c *GitHubConfig) SetApiTokenName(v *string) {
	c.values.Set(apiTokenNameKey, v)
}

func (c *GitHubConfig) getStrategy() (*GitHubReleaseCreationStrategy, error) {
	var s GitHubReleaseCreationStrategy

	if c.values.Has(strategyKey) {
		v := c.values[strategyKey].(string)

		if !isGitHubReleaseCreationStrategy(c.values[strategyKey].(string)) {
			return nil, fmt.Errorf("%s is not a valid strategy", v)
		}

		s = GitHubReleaseCreationStrategy(v)
	} else {
		logrus.Warnf("%s is missing so set `%s` as a default value", strategyKey, DraftOrCreate)
		s = DraftOrCreate
	}

	return &s, nil
}

func (c *GitHubConfig) GetStrategy() GitHubReleaseCreationStrategy {
	if s, err := c.getStrategy(); err != nil {
		panic(err)
	} else {
		return *s
	}
}

func (c *GitHubConfig) SetStrategy(s GitHubReleaseCreationStrategy) {
	c.values.Set(strategyKey, string(s))
}

func (c *GitHubConfig) getFileNamePattern() (string, error) {
	if c.values.Has(fileNamePattern) {
		pattern := c.values[fileNamePattern].(string)

		if _, err := regexp.Compile(pattern); err != nil {
			return "", err
		} else {
			return pattern, nil
		}
	} else {
		return "", nil
	}
}

func (c *GitHubConfig) GetFileNamePattern() string {
	if t, err := c.getFileNamePattern(); err != nil {
		panic(err)
	} else {
		return t
	}
}

func (c *GitHubConfig) SetFileNamePattern(v *string) {
	c.values.Set(fileNamePattern, v)
}
