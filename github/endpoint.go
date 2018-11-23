package github

import (
	"fmt"
	"github.com/jmatsu/artifact-transfer/lib"
	"strings"
)

func ReleaseListEndpoint(username string, repoName string) lib.Endpoint {
	return lib.Endpoint{
		Url:      fmt.Sprintf("https://api.github.com/repos/%s/%s/releases", username, repoName),
		AuthType: lib.HeaderAuth,
		Accept:   "application/vnd.github.v3+json",
	}
}

func UploadReleaseEndpoint(uri string) lib.Endpoint {
	return lib.Endpoint{
		Url:      strings.Split(uri, "{")[0], // hypermedia url will be coming
		AuthType: lib.HeaderAuth,
		Accept:   "application/vnd.github.v3+json",
	}
}

func CreateReleaseEndpoint(username string, repoName string) lib.Endpoint {
	return lib.Endpoint{
		Url:      fmt.Sprintf("https://api.github.com/repos/%s/%s/releases", username, repoName),
		AuthType: lib.HeaderAuth,
		Accept:   "application/vnd.github.v3+json",
	}
}
