package config

import (
	"fmt"
	"gopkg.in/guregu/null.v3"
	"os"
	"regexp"
)

type VcsType string

const (
	GitHub    VcsType = "github"
	Bitbucket         = "bitbucket"
)

func NewVcsType(v string) (VcsType, error) {
	t := VcsType(v)

	switch t {
	case GitHub:
		return t, nil
	case Bitbucket:
		return t, nil
	default:
		return t, fmt.Errorf("%s is invalid vcs type\n", v)
	}
}

type CircleCIConfig struct {
	values LocationConfig
	Err    error
}

func NewCircleCIConfig(lc LocationConfig) (*CircleCIConfig, error) {
	if t, err := lc.GetLocationType(); err != nil || t != CircleCI {
		if err == nil {
			err = fmt.Errorf("location type is not for circleci so the caller is wrong")
		}

		return nil, err
	}

	config := &CircleCIConfig{
		values: lc,
	}

	return config, nil
}

func (c CircleCIConfig) Validate() error {
	c.setErr(c.getVcsType())
	c.setErr(c.getUsername())
	c.setErr(c.getRepoName())
	c.setErr(c.getFileNamePattern())

	return c.Err
}

func (c CircleCIConfig) setErr(_ interface{}, err error) {
	if c.Err != nil {
		return
	}

	c.Err = err
}

func (c CircleCIConfig) getVcsType() (VcsType, error) {
	if v, prs := c.values[vcsTypeKey]; prs && v != "" {
		if v, ok := v.(string); ok {
			return NewVcsType(v)
		}
	}

	return VcsType(""), fmt.Errorf("%s is missing or an invalid value\n", vcsTypeKey)
}

func (c CircleCIConfig) GetVcsType() VcsType {
	if t, err := c.getVcsType(); err != nil {
		panic(err)
	} else {
		return t
	}
}

func (c CircleCIConfig) SetVcsType(v VcsType) {
	c.values[vcsTypeKey] = string(v)
}

func (c CircleCIConfig) getUsername() (string, error) {
	if v, prs := c.values[usernameKey]; prs && v != "" {
		return v.(string), nil
	}

	return "", fmt.Errorf("%s is missinge\n", usernameKey)
}

func (c CircleCIConfig) GetUsername() string {
	if t, err := c.getUsername(); err != nil {
		panic(err)
	} else {
		return t
	}
}

func (c CircleCIConfig) SetUsername(v string) {
	c.values[usernameKey] = v
}

func (c CircleCIConfig) getRepoName() (string, error) {
	if v, prs := c.values[repoNameKey]; prs && v != "" {
		return v.(string), nil
	}

	return "", fmt.Errorf("%s is missing\n", repoNameKey)
}

func (c CircleCIConfig) GetRepoName() string {
	if t, err := c.getRepoName(); err != nil {
		panic(err)
	} else {
		return t
	}
}

func (c CircleCIConfig) SetRepoName(v string) {
	c.values[repoNameKey] = v
}

func (c CircleCIConfig) GetBranch() null.String {
	if v, prs := c.values[branchKey]; prs && v != nil {
		return null.StringFrom(v.(string))
	} else {
		return null.StringFromPtr(nil)
	}
}

func (c CircleCIConfig) SetBranch(v *string) {
	c.values[branchKey] = v
}

func (c CircleCIConfig) GetApiToken() null.String {
	if v, prs := c.values[apiTokenNameKey]; prs && v != nil {
		if v, ok := os.LookupEnv(v.(string)); ok {
			return null.StringFrom(v)
		} else {
			return null.StringFromPtr(nil)
		}
	} else {
		return null.StringFromPtr(nil)
	}
}

func (c CircleCIConfig) SetApiTokenName(v *string) {
	c.values[apiTokenNameKey] = v
}

func (c CircleCIConfig) getFileNamePattern() (string, error) {
	if v, prs := c.values[fileNamePattern]; prs && v != nil {
		if _, err := regexp.Compile(v.(string)); err != nil {
			return "", err
		} else {
			return v.(string), nil
		}
	} else {
		return "", nil
	}
}

func (c CircleCIConfig) GetFileNamePattern() string {
	if t, err := c.getFileNamePattern(); err != nil {
		panic(err)
	} else {
		return t
	}
}

func (c CircleCIConfig) SetFileNamePattern(v *string) {
	c.values[fileNamePattern] = v
}
