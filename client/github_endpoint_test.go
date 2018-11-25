package client

import (
	"fmt"
	"github.com/jmatsu/transart/lib"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testGitHubReleaseListEndpointTests = []struct {
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

func TestGitHubReleaseListEndpoint(t *testing.T) {
	for i, c := range testGitHubReleaseListEndpointTests {
		t.Run(fmt.Sprintf("TestGitHubReleaseListEndpoint %d", i), func(t *testing.T) {
			endpoint := gitHubReleaseListEndpoint(c.username, c.repoName)
			assert.EqualValues(t, c.out, endpoint.Url)
			assert.EqualValues(t, "application/vnd.github.v3+json", endpoint.Accept)
			assert.EqualValues(t, lib.HeaderAuth, endpoint.AuthType)
		})
	}
}

var testGitHubUploadReleaseEndpointTests = []struct {
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

func TestGitHubUploadReleaseEndpoint(t *testing.T) {
	for i, c := range testGitHubUploadReleaseEndpointTests {
		t.Run(fmt.Sprintf("TestGitHubUploadReleaseEndpoint %d", i), func(t *testing.T) {
			endpoint := gitHubUploadReleaseEndpoint(c.in)
			assert.EqualValues(t, c.out, endpoint.Url)
			assert.EqualValues(t, "application/vnd.github.v3+json", endpoint.Accept)
			assert.EqualValues(t, lib.HeaderAuth, endpoint.AuthType)
		})
	}
}

var testGitHubCreateReleaseEndpointTests = testGitHubReleaseListEndpointTests

func TestGitHubCreateReleaseEndpoint(t *testing.T) {
	for i, c := range testGitHubCreateReleaseEndpointTests {
		t.Run(fmt.Sprintf("TestGitHubCreateReleaseEndpoint %d", i), func(t *testing.T) {
			endpoint := gitHubCreateReleaseEndpoint(c.username, c.repoName)
			assert.EqualValues(t, c.out, endpoint.Url)
			assert.EqualValues(t, "application/vnd.github.v3+json", endpoint.Accept)
			assert.EqualValues(t, lib.HeaderAuth, endpoint.AuthType)
		})
	}
}
