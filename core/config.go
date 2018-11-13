package core

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type (
	RootConfig struct {
		Version     uint              `yaml:"version"`
		Source      SourceConfig      `yaml:"source"`
		Destination DestinationConfig `yaml:"destination"`
	}

	LocationConfig map[string]interface{}

	Validatable interface {
		Validate() error
	}

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
	GitHubRelease              = "github-release"
)

const LocationTypeKey = "type"

func (c LocationConfig) GetLocationType() (LocationType, error) {
	if v, prs := c[LocationTypeKey]; prs {
		if v, ok := v.(string); ok {
			return NewLocationType(v)
		}
	}

	return LocationType(""), fmt.Errorf("%s is missing or an invalid value\n", LocationTypeKey)
}

func NewLocationType(v string) (LocationType, error) {
	t := LocationType(v)

	switch t {
	case CircleCI:
		return t, nil
	case GitHubRelease:
		return t, nil
	default:
		return t, fmt.Errorf("%s is invalid location type\n", v)
	}
}

func LoadConfig() (*RootConfig, error) {
	config := RootConfig{}

	if bytes, err := ioutil.ReadFile(".transart.yml"); err != nil {
		return nil, err
	} else if err := yaml.Unmarshal(bytes, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

func (c *RootConfig) Save() error {
	if bytes, err := yaml.Marshal(c); err != nil {
		return err
	} else {
		ioutil.WriteFile(".transart.yml", bytes, 0644)
		return nil
	}
}
