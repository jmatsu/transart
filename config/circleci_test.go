package config

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v3"
	"testing"
)

var testNewVcsTypeTests = []struct {
	in  string
	out VcsType
	err bool
}{
	{
		"x",
		"",
		true,
	},
	{
		"github",
		GitHub,
		false,
	},
	{
		"bitbucket",
		Bitbucket,
		false,
	},
}

func TestNewVcsType(t *testing.T) {
	for i, c := range testNewVcsTypeTests {
		t.Run(fmt.Sprintf("TestNewVcsType %d", i), func(t *testing.T) {
			vcsType, err := NewVcsType(c.in)

			if !c.err {
				assert.EqualValues(t, c.out, vcsType)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}

var testNewCircleCIConfigTests = []struct {
	in  LocationConfig
	err bool
}{
	{
		LocationConfig{
			locationTypeKey: string(GitHubRelease),
		},
		true,
	},
	{
		LocationConfig{
			locationTypeKey: string(CircleCI),
		},
		false,
	},
}

func TestNewCircleCIConfig(t *testing.T) {
	for i, c := range testNewCircleCIConfigTests {
		t.Run(fmt.Sprintf("TestNewCircleCIConfig %d", i), func(t *testing.T) {
			_, err := NewCircleCIConfig(c.in)

			if !c.err {
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}

func createCircleCIConfig(vcsType string, username string, reponame string, tokenName null.String, pattern null.String, branch null.String) CircleCIConfig {
	c, _ := NewCircleCIConfig(LocationConfig{
		locationTypeKey: string(CircleCI),
	})

	c.SetVcsType(VcsType(vcsType))
	c.SetUsername(username)
	c.SetRepoName(reponame)
	c.SetApiTokenName(tokenName.Ptr())
	c.SetFileNamePattern(pattern.Ptr())
	c.SetBranch(branch.Ptr())

	return *c
}

var testCircleCIConfig_ValidateTests = []struct {
	in  CircleCIConfig
	err bool
}{
	{
		createCircleCIConfig(string(GitHub), "username", "reponame", null.StringFrom("token"), null.StringFrom(".*"), null.StringFrom("release")),
		false,
	},
	{
		createCircleCIConfig("x", "username", "reponame", null.StringFrom("token"), null.StringFrom(".*"), null.StringFrom("release")),
		true,
	},
	{
		createCircleCIConfig(string(GitHub), "", "reponame", null.StringFrom("token"), null.StringFrom(".*"), null.StringFrom("release")),
		true,
	},
	{
		createCircleCIConfig(string(GitHub), "username", "", null.StringFrom("token"), null.StringFrom(".*"), null.StringFrom("release")),
		true,
	},
	{
		createCircleCIConfig(string(GitHub), "username", "reponame", null.StringFrom("token"), null.StringFrom("*"), null.StringFrom("release")),
		true,
	},
}

func TestCircleCIConfig_Validate(t *testing.T) {
	for i, c := range testCircleCIConfig_ValidateTests {
		t.Run(fmt.Sprintf("TestCircleCIConfig_Validate %d", i), func(t *testing.T) {
			err := c.in.Validate()

			if !c.err {
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}

var testCircleCIConfig_setErrorTests = []struct {
	in error
}{
	{
		errors.New("error"),
	},
	{
		nil,
	},
}

func TestCircleCIConfig_setError(t *testing.T) {
	for i, c := range testCircleCIConfig_setErrorTests {
		t.Run(fmt.Sprintf("TestCircleCIConfig_setError %d", i), func(t *testing.T) {
			ccConfig, _ := NewCircleCIConfig(LocationConfig{
				locationTypeKey: string(CircleCI),
			})

			ccConfig.setError(nil, c.in)

			assert.EqualValues(t, c.in, ccConfig.Err)

			if c.in != nil {
				ccConfig.setError(nil, errors.New("error2"))

				assert.EqualValues(t, c.in, ccConfig.Err)
			}
		})
	}
}
