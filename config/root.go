package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

type (
	RootConfig struct {
		ConfFileName string
		Version      uint              `yaml:"version"`
		SaveDir      string            `yaml:"save_dir"`
		Source       SourceConfig      `yaml:"source"`
		Destination  DestinationConfig `yaml:"destination"`
	}

	LocationConfig map[string]interface{}

	LocationType string

	SourceConfig struct {
		Locations []LocationConfig `yaml:"locations"`
	}

	DestinationConfig struct {
		Location LocationConfig `yaml:"location"`
	}
)

const (
	CircleCI      LocationType = "circleci"
	GitHubRelease LocationType = "github-release"
	Local         LocationType = "local"
)

func (c LocationConfig) GetLocationType() (LocationType, error) {
	if v, prs := c[locationTypeKey]; prs && v != "" {
		if v, ok := v.(string); ok {
			return NewLocationType(v)
		}
	}

	return LocationType(""), fmt.Errorf("%s is missing or an invalid value\n", locationTypeKey)
}

func (c LocationConfig) SetLocationType(t LocationType) {
	c[locationTypeKey] = string(t)
}

func NewLocationType(v string) (LocationType, error) {
	t := LocationType(v)

	switch t {
	case GitHubRelease, CircleCI, Local:
		return t, nil
	default:
		return t, fmt.Errorf("%s is invalid location type\n", v)
	}
}

func LoadRootConfig(confFileName string) (*RootConfig, error) {
	config := RootConfig{}

	if bytes, err := ioutil.ReadFile(confFileName); err != nil {
		return nil, err
	} else if err := yaml.Unmarshal(bytes, &config); err != nil {
		return nil, err
	}

	config.ConfFileName = confFileName

	return &config, nil
}

func ExistsRootConfig(confFileName string) bool {
	_, err := os.Stat(confFileName)

	return err == nil || !os.IsNotExist(err)
}

func (c *RootConfig) Save() error {
	if bytes, err := yaml.Marshal(c); err != nil {
		return err
	} else {
		return ioutil.WriteFile(c.ConfFileName, bytes, 0644)
	}
}
