package config

import (
	"fmt"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testLocalConfig_setErrorTests = []struct {
	in error
}{
	{
		errors.New("error"),
	},
	{
		nil,
	},
}

func TestLocalConfig_setError(t *testing.T) {
	for i, c := range testGitHubConfig_setErrorTests {
		t.Run(fmt.Sprintf("TestLocalConfig_setError %d", i), func(t *testing.T) {
			lConfig, _ := NewLocalConfig(LocationConfig{
				locationTypeKey: string(Local),
			})

			lConfig.setError(nil, c.in)

			assert.EqualValues(t, c.in, lConfig.Err)

			if c.in != nil {
				lConfig.setError(nil, errors.New("error2"))

				assert.EqualValues(t, c.in, lConfig.Err)
			}
		})
	}
}
