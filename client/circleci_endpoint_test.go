package client

import (
	"fmt"
	"github.com/jmatsu/transart/client/entity"
	"github.com/jmatsu/transart/lib"
	"github.com/stretchr/testify/assert"
	"gopkg.in/guregu/null.v3"
	"testing"
)

var testCircleCIBaseApiUrlTests = []struct {
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

func TestCircleCIBaseApiUrl(t *testing.T) {
	for i, c := range testCircleCIBaseApiUrlTests {
		t.Run(fmt.Sprintf("TestCircleCIBaseApiUrl %d", i), func(t *testing.T) {
			url := baseCircleCIApiUrl(c.vcsType, c.username, c.repoName)
			assert.EqualValues(t, c.out, url)
		})
	}
}

var testCircleCIJobInfoListEndpointTests = []struct {
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

func TestCircleCIJobInfoListEndpoint(t *testing.T) {
	for i, c := range testCircleCIJobInfoListEndpointTests {
		t.Run(fmt.Sprintf("TestCircleCIJobInfoListEndpoint %d", i), func(t *testing.T) {
			endpoint := circleCIJobInfoListEndpoint(c.vcsType, c.username, c.repoName, c.branch)
			assert.EqualValues(t, c.out, endpoint.Url)
			assert.EqualValues(t, "application/json", endpoint.Accept)
			assert.EqualValues(t, lib.HeaderAuth, endpoint.AuthType)
		})
	}
}

var testCircleCIArtifactListEndpointTests = []struct {
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

func TestCircleCIArtifactListEndpoint(t *testing.T) {
	for i, c := range testCircleCIArtifactListEndpointTests {
		t.Run(fmt.Sprintf("TestCircleCIArtifactListEndpoint %d", i), func(t *testing.T) {
			endpoint := circleCIArtifactListEndpoint(c.vcsType, c.username, c.repoName, c.buildNum)
			assert.EqualValues(t, c.out, endpoint.Url)
			assert.EqualValues(t, "application/json", endpoint.Accept)
			assert.EqualValues(t, lib.HeaderAuth, endpoint.AuthType)
		})
	}
}

var testCircleCIDownloadArtifactEndpointTests = []struct {
	in  entity.CircleCIArtifact
	out string
}{
	{
		entity.CircleCIArtifact{
			DownloadUrl: "https://circleci.com/api/v1.1/project/sample",
		},
		"https://circleci.com/api/v1.1/project/sample",
	},
}

func TestCircleCIDownloadArtifactEndpoint(t *testing.T) {
	for i, c := range testCircleCIDownloadArtifactEndpointTests {
		t.Run(fmt.Sprintf("TestCircleCIDownloadArtifactEndpoint %d", i), func(t *testing.T) {
			endpoint := circleCIDownloadArtifactEndpoint(c.in)
			assert.EqualValues(t, c.out, endpoint.Url)
			assert.EqualValues(t, lib.ParameterAuth, endpoint.AuthType)
		})
	}
}
