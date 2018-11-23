package circleci

import (
	"fmt"
	"github.com/jmatsu/artifact-transfer/core"
	"github.com/sirupsen/logrus"
	"gopkg.in/guregu/null.v3"
	"os"
	"regexp"
)

type Config struct {
	values core.LocationConfig
	Err    error
}

const (
	vcsTypeKey      = "vcs-type"
	usernameKey     = "username"
	repoNameKey     = "reponame"
	branchKey       = "branch"
	apiTokenNameKey = "api-token-name"
	fileNamePattern = "file-name-pattern"
)

func NewConfig(lc core.LocationConfig) (*Config, error) {
	if t, err := lc.GetLocationType(); err != nil || t != core.CircleCI {
		if err == nil {
			err = fmt.Errorf("location type is not for circleci so the caller is wrong")
		}

		return nil, err
	}

	config := &Config{
		values: lc,
	}

	return config, nil
}

func (c Config) Validate() error {
	c.setErr(c.getVcsType())
	c.setErr(c.getUsername())
	c.setErr(c.getRepoName())
	c.setErr(c.getFileNamePattern())

	return c.Err
}

func (c Config) setErr(_ interface{}, err error) {
	if c.Err != nil {
		return
	}

	c.Err = err
}

func (c Config) getVcsType() (VcsType, error) {
	if v, prs := c.values[vcsTypeKey]; prs {
		if v, ok := v.(string); ok {
			return NewVcsType(v)
		}
	}

	return VcsType(""), fmt.Errorf("%s is missing or an invalid value\n", vcsTypeKey)
}

func (c Config) GetVcsType() VcsType {
	if t, err := c.getVcsType(); err != nil {
		panic(err)
	} else {
		return t
	}
}

func (c Config) SetVcsType(v VcsType) {
	c.values[vcsTypeKey] = v
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

func (c Config) GetBranch() null.String {
	if v, prs := c.values[branchKey]; prs {
		return null.StringFrom(v.(string))
	} else {
		return null.StringFromPtr(nil)
	}
}

func (c Config) SetBranch(v null.String) {
	if v.Valid {
		c.values[branchKey] = v
	} else {
		logrus.Warnf("SetBranch was called but ignored because the argument is invalid string with %s\n", v.String)
	}
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

func (c Config) getFileNamePattern() (string, error) {
	if v, prs := c.values[fileNamePattern]; prs {
		if _, err := regexp.Compile(v.(string)); err != nil {
			return "", err
		} else {
			return v.(string), nil
		}
	} else {
		return "", nil
	}
}

func (c Config) GetFileNamePattern() string {
	if t, err := c.getFileNamePattern(); err != nil {
		panic(err)
	} else {
		return t
	}
}

func (c Config) SetFileNamePattern(v null.String) {
	c.values[fileNamePattern] = v
}
