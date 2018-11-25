package client

import (
	"fmt"
	"github.com/jmatsu/transart/lib"
	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v3"
	"net/http"
	"testing"
)

var testGitHubToken_SetToHeaderTests = []struct {
	in  lib.Token
	out string
}{
	{
		newGitHubToken(null.StringFrom("XYZ")),
		"token XYZ",
	},
	{
		newGitHubToken(null.StringFromPtr(nil)),
		"",
	},
}

func TestGitHubToken_SetToHeader(t *testing.T) {
	for i, c := range testGitHubToken_SetToHeaderTests {
		t.Run(fmt.Sprintf("TestGitHubToken_SetToHeader %d", i), func(t *testing.T) {
			request := http.Request{
				Header: make(http.Header),
			}

			c.in.SetToHeader(&request)

			assert.EqualValues(t, c.out, request.Header.Get("Authorization"))
		})
	}
}

func TestGitHubToken_ToParam(t *testing.T) {
	// no-op
}
