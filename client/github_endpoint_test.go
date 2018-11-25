package client

import (
	"fmt"
	"github.com/jmatsu/transart/lib"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testReleaseListEndpointTests = []struct {
	username string
	repoName string
	out      string
}{
	{
		"username",
		"reponame",
		"https://api.github.com/repos/username/reponame/releases",
	},
	{
		"user name",
		"repo name",
		"https://api.github.com/repos/user%20name/repo%20name/releases",
	},
}

func TestReleaseListEndpoint(t *testing.T) {
	for i, c := range testReleaseListEndpointTests {
		t.Run(fmt.Sprintf("TestReleaseListEndpoint %d", i), func(t *testing.T) {
			endpoint := gitHubReleaseListEndpoint(c.username, c.repoName)
			assert.EqualValues(t, c.out, endpoint.Url)
			assert.EqualValues(t, "application/vnd.github.v3+json", endpoint.Accept)
			assert.EqualValues(t, lib.HeaderAuth, endpoint.AuthType)
		})
	}
}

var testUploadReleaseEndpointTests = []struct {
	in  string
	out string
}{
	{
		"https://api.github.com",
		"https://api.github.com",
	},
	{
		"https://api.github.com{x,y,z",
		"https://api.github.com",
	},
}

func TestUploadReleaseEndpoint(t *testing.T) {
	for i, c := range testUploadReleaseEndpointTests {
		t.Run(fmt.Sprintf("TestUploadReleaseEndpoint %d", i), func(t *testing.T) {
			endpoint := gitHubUploadReleaseEndpoint(c.in)
			assert.EqualValues(t, c.out, endpoint.Url)
			assert.EqualValues(t, "application/vnd.github.v3+json", endpoint.Accept)
			assert.EqualValues(t, lib.HeaderAuth, endpoint.AuthType)
		})
	}
}

var testCreateReleaseEndpointTests = testReleaseListEndpointTests

func TestCreateReleaseEndpoint(t *testing.T) {
	for i, c := range testCreateReleaseEndpointTests {
		t.Run(fmt.Sprintf("TestCreateReleaseEndpoint %d", i), func(t *testing.T) {
			endpoint := gitHubCreateReleaseEndpoint(c.username, c.repoName)
			assert.EqualValues(t, c.out, endpoint.Url)
			assert.EqualValues(t, "application/vnd.github.v3+json", endpoint.Accept)
			assert.EqualValues(t, lib.HeaderAuth, endpoint.AuthType)
		})
	}
}
