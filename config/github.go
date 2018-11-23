package config

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"gopkg.in/guregu/null.v3"
	"os"
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
	DraftOrCreate                               = "draft-or-create"
	Create                                      = "create"
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

func (c GitHubConfig) setError(_ interface{}, err error) {
	if c.Err != nil {
		return
	}

	c.Err = err
}

func (c GitHubConfig) Validate() error {
	c.setError(c.getUsername())
	c.setError(c.getRepoName())
	c.setError(c.getStrategy())

	return c.Err
}

func (c GitHubConfig) getUsername() (string, error) {
	if v, prs := c.values[usernameKey]; prs {
		return v.(string), nil
	}

	return "", fmt.Errorf("%s is missinge\n", usernameKey)
}

func (c GitHubConfig) GetUsername() string {
	if t, err := c.getUsername(); err != nil {
		panic(err)
	} else {
		return t
	}
}

func (c GitHubConfig) SetUsername(v string) {
	c.values[usernameKey] = v
}

func (c GitHubConfig) getRepoName() (string, error) {
	if v, prs := c.values[repoNameKey]; prs {
		return v.(string), nil
	}

	return "", fmt.Errorf("%s is missing\n", repoNameKey)
}

func (c GitHubConfig) GetRepoName() string {
	if t, err := c.getRepoName(); err != nil {
		panic(err)
	} else {
		return t
	}
}

func (c GitHubConfig) SetRepoName(v string) {
	c.values[repoNameKey] = v
}

func (c GitHubConfig) GetApiToken() null.String {
	if v, prs := c.values[apiTokenNameKey]; prs {
		if v, ok := os.LookupEnv(v.(string)); ok {
			return null.StringFrom(v)
		} else {
			return null.StringFromPtr(nil)
		}
	} else {
		return null.StringFromPtr(nil)
	}
}

func (c GitHubConfig) SetApiTokenName(v null.String) {
	if v.Valid {
		c.values[apiTokenNameKey] = v
	} else {
		logrus.Warnf("SetApiTokenName was called but ignored because the argument is invalid string with %s\n", v.String)
	}
}

func (c GitHubConfig) getStrategy() (*GitHubReleaseCreationStrategy, error) {
	if v, prs := c.values[strategyKey]; prs {
		if !isGitHubReleaseCreationStrategy(v.(string)) {
			return nil, fmt.Errorf("%s is not a valid strategy", v)
		}

		s := GitHubReleaseCreationStrategy(v.(string))

		return &s, nil
	}

	return nil, fmt.Errorf("%s is missing\n", strategyKey)
}

func (c GitHubConfig) GetStrategy() GitHubReleaseCreationStrategy {
	if s, err := c.getStrategy(); err != nil {
		panic(err)
	} else {
		return *s
	}
}

func (c GitHubConfig) SetStrategy(s GitHubReleaseCreationStrategy) {
	c.values[strategyKey] = string(s)
}
