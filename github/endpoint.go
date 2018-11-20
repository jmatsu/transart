package github

import (
	"fmt"
	"github.com/jmatsu/artifact-transfer/core"
	"strings"
)

func ReleaseListEndpoint(username string, repoName string) core.Endpoint {
	return core.Endpoint{
		Url:      fmt.Sprintf("https://api.github.com/repos/%s/%s/releases", username, repoName),
		AuthType: core.HeaderAuth,
		Accept:   "application/vnd.github.v3+json",
	}
}

func UploadReleaseEndpoint(uri string) core.Endpoint {
	return core.Endpoint{
		Url:      strings.Split(uri, "{")[0], // hypermedia url will be coming
		AuthType: core.HeaderAuth,
		Accept:   "application/vnd.github.v3+json",
	}
}

func CreateReleaseEndpoint(username string, repoName string) core.Endpoint {
	return core.Endpoint{
		Url:      fmt.Sprintf("https://api.github.com/repos/%s/%s/releases", username, repoName),
		AuthType: core.HeaderAuth,
		Accept:   "application/vnd.github.v3+json",
	}
}
