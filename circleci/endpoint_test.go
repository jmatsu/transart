package circleci

import (
	"fmt"
	"github.com/jmatsu/transart/circleci/entity"
	"github.com/jmatsu/transart/lib"
	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v3"
	"testing"
)

var testBaseApiUrlTests = []struct {
	vcsType  string
	username string
	repoName string
	out      string
}{
	{
		"github",
		"username",
		"reponame",
		"https://circleci.com/api/v1.1/project/github/username/reponame",
	},
	{
		"bitbucket",
		"user name",
		"repo name",
		"https://circleci.com/api/v1.1/project/bitbucket/user%20name/repo%20name",
	},
}

func TestBaseApiUrl(t *testing.T) {
	for i, c := range testBaseApiUrlTests {
		t.Run(fmt.Sprintf("TestBaseApiUrl %d", i), func(t *testing.T) {
			url := baseApiUrl(c.vcsType, c.username, c.repoName)
			assert.EqualValues(t, c.out, url)
		})
	}
}

var testJobInfoListEndpointTests = []struct {
	vcsType  string
	username string
	repoName string
	branch   null.String
	out      string
}{
	{
		"github",
		"username",
		"reponame",
		null.StringFromPtr(nil),
		"https://circleci.com/api/v1.1/project/github/username/reponame",
	},
	{
		"github",
		"username",
		"reponame",
		null.StringFrom("release"),
		"https://circleci.com/api/v1.1/project/github/username/reponame/tree/release",
	},
}

func TestJobInfoListEndpoint(t *testing.T) {
	for i, c := range testJobInfoListEndpointTests {
		t.Run(fmt.Sprintf("TestJobInfoListEndpoint %d", i), func(t *testing.T) {
			endpoint := JobInfoListEndpoint(c.vcsType, c.username, c.repoName, c.branch)
			assert.EqualValues(t, c.out, endpoint.Url)
			assert.EqualValues(t, "application/json", endpoint.Accept)
			assert.EqualValues(t, lib.HeaderAuth, endpoint.AuthType)
		})
	}
}

var testArtifactListEndpointTests = []struct {
	vcsType  string
	username string
	repoName string
	buildNum uint
	out      string
}{
	{
		"github",
		"username",
		"reponame",
		10,
		"https://circleci.com/api/v1.1/project/github/username/reponame/10/artifacts",
	},
}

func TestArtifactListEndpoint(t *testing.T) {
	for i, c := range testArtifactListEndpointTests {
		t.Run(fmt.Sprintf("TestArtifactListEndpoint %d", i), func(t *testing.T) {
			endpoint := ArtifactListEndpoint(c.vcsType, c.username, c.repoName, c.buildNum)
			assert.EqualValues(t, c.out, endpoint.Url)
			assert.EqualValues(t, "application/json", endpoint.Accept)
			assert.EqualValues(t, lib.HeaderAuth, endpoint.AuthType)
		})
	}
}

var testDownloadArtifactEndpointTests = []struct {
	in  entity.Artifact
	out string
}{
	{
		entity.Artifact{
			DownloadUrl: "https://circleci.com/api/v1.1/project/sample",
		},
		"https://circleci.com/api/v1.1/project/sample",
	},
}

func TestDownloadArtifactEndpoint(t *testing.T) {
	for i, c := range testDownloadArtifactEndpointTests {
		t.Run(fmt.Sprintf("TestDownloadArtifactEndpoint %d", i), func(t *testing.T) {
			endpoint := DownloadArtifactEndpoint(c.in)
			assert.EqualValues(t, c.out, endpoint.Url)
			assert.EqualValues(t, lib.ParameterAuth, endpoint.AuthType)
		})
	}
}
