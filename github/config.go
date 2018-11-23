package github

import (
	"fmt"
	"github.com/jmatsu/artifact-transfer/core"
	"github.com/sirupsen/logrus"
	"gopkg.in/guregu/null.v3"
	"os"
)

type (
	Strategy string

	Config struct {
		values core.LocationConfig
		Err    error
	}
)

const (
	usernameKey     = "username"
	repoNameKey     = "reponame"
	apiTokenNameKey = "api-token-name"
	strategyKey     = "strategy"
)

const (
	Draft         Strategy = "draft"
	DraftOrCreate          = "draft-or-create"
	Create                 = "create"
)

func isStrategy(s string) bool {
	strategies := []Strategy{
		Draft,
		DraftOrCreate,
		Create,
	}

	for _, strategy := range strategies {
		if strategy == Strategy(s) {
			return true
		}
	}

	return false
}

func NewConfig(lc core.LocationConfig) (*Config, error) {
	if t, err := lc.GetLocationType(); err != nil || t != core.GitHubRelease {
		if err == nil {
			err = fmt.Errorf("location type is not for github releases so the caller is wrong")
		}

		return nil, err
	}

	config := &Config{
		values: lc,
	}

	return config, nil
}

func (c Config) setError(_ interface{}, err error) {
	if c.Err != nil {
		return
	}

	c.Err = err
}

func (c Config) Validate() error {
	c.setError(c.getUsername())
	c.setError(c.getRepoName())
	c.setError(c.getStrategy())

	return c.Err
}

func (c Config) getUsername() (string, error) {
	if v, prs := c.values[usernameKey]; prs {
		return v.(string), nil
	}

	return "", fmt.Errorf("%s is missinge\n", usernameKey)
}

func (c Config) GetUsername() string {
	if t, err := c.getUsername(); err != nil {
		panic(err)
	} else {
		return t
	}
}

func (c Config) SetUsername(v string) {
	c.values[usernameKey] = v
}

func (c Config) getRepoName() (string, error) {
	if v, prs := c.values[repoNameKey]; prs {
		return v.(string), nil
	}

	return "", fmt.Errorf("%s is missing\n", repoNameKey)
}

func (c Config) GetRepoName() string {
	if t, err := c.getRepoName(); err != nil {
		panic(err)
	} else {
		return t
	}
}

func (c Config) SetRepoName(v string) {
	c.values[repoNameKey] = v
}

func (c Config) GetApiToken() null.String {
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

func (c Config) SetApiTokenName(v null.String) {
	if v.Valid {
		c.values[apiTokenNameKey] = v
	} else {
		logrus.Warnf("SetApiTokenName was called but ignored because the argument is invalid string with %s\n", v.String)
	}
}

func (c Config) getStrategy() (*Strategy, error) {
	if v, prs := c.values[strategyKey]; prs {
		if !isStrategy(v.(string)) {
			return nil, fmt.Errorf("%s is not a valid strategy", v)
		}

		s := Strategy(v.(string))

		return &s, nil
	}

	return nil, fmt.Errorf("%s is missing\n", strategyKey)
}

func (c Config) GetStrategy() Strategy {
	if s, err := c.getStrategy(); err != nil {
		panic(err)
	} else {
		return *s
	}
}

func (c Config) SetStrategy(s Strategy) {
	c.values[strategyKey] = string(s)
}
