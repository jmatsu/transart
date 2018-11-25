package config

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testGitHubConfig_setErrorTests = []struct {
	in error
}{
	{
		errors.New("error"),
	},
	{
		nil,
	},
}

func TestGitHubConfig_setError(t *testing.T) {
	for i, c := range testGitHubConfig_setErrorTests {
		t.Run(fmt.Sprintf("TestGitHubConfig_setError %d", i), func(t *testing.T) {
			ghConfig, _ := NewGitHubConfig(LocationConfig{
				locationTypeKey: string(GitHubRelease),
			})

			ghConfig.setError(nil, c.in)

			assert.EqualValues(t, c.in, ghConfig.Err)

			if c.in != nil {
				ghConfig.setError(nil, errors.New("error2"))

				assert.EqualValues(t, c.in, ghConfig.Err)
			}
		})
	}
}
