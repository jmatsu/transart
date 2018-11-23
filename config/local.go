package config

import (
	"fmt"
	"gopkg.in/guregu/null.v3"
	"regexp"
)

type (
	LocalConfig struct {
		values LocationConfig
		Err    error
	}
)

func NewLocalConfig(lc LocationConfig) (*LocalConfig, error) {
	if t, err := lc.GetLocationType(); err != nil || t != Local {
		if err == nil {
			err = fmt.Errorf("location type is not for local so the caller is wrong")
		}

		return nil, err
	}

	config := &LocalConfig{
		values: lc,
	}

	return config, nil
}

func (c LocalConfig) setError(_ interface{}, err error) {
	if c.Err != nil {
		return
	}

	c.Err = err
}

func (c LocalConfig) Validate() error {
	c.setError(c.getPath())

	return c.Err
}

func (c LocalConfig) getPath() (string, error) {
	if v, prs := c.values[pathKey]; prs {
		return v.(string), nil
	}

	return "", fmt.Errorf("%s is missinge\n", pathKey)
}

func (c LocalConfig) GetPath() string {
	if t, err := c.getPath(); err != nil {
		panic(err)
	} else {
		return t
	}
}

func (c LocalConfig) SetPath(v string) {
	c.values[pathKey] = v
}

func (c LocalConfig) getFileNamePattern() (string, error) {
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

func (c LocalConfig) GetFileNamePattern() string {
	if t, err := c.getFileNamePattern(); err != nil {
		panic(err)
	} else {
		return t
	}
}

func (c LocalConfig) SetFileNamePattern(v null.String) {
	c.values[fileNamePattern] = v
}
