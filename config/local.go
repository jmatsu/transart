package config

import (
	"fmt"
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
	c.setError(c.getFileNamePattern())

	return c.Err
}

func (c LocalConfig) getPath() (string, error) {
	if c.values.Has(pathKey) {
		return c.values[pathKey].(string), nil
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
	c.values.Set(pathKey, v)
}

func (c LocalConfig) getFileNamePattern() (string, error) {
	if c.values.Has(fileNamePattern) {
		pattern := c.values[fileNamePattern].(string)

		if _, err := regexp.Compile(pattern); err != nil {
			return "", err
		} else {
			return pattern, nil
		}

		return "", nil
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

func (c LocalConfig) SetFileNamePattern(v *string) {
	c.values.Set(fileNamePattern, v)
}
