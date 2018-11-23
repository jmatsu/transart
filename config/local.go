package config

import (
	"fmt"
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
