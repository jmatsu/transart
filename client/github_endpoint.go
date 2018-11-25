package client

import (
	"fmt"
	"github.com/jmatsu/transart/lib"
	"net/url"
	"strings"
)

func gitHubReleaseListEndpoint(username string, repoName string) lib.Endpoint {
	return lib.Endpoint{
		Url:      fmt.Sprintf("https://api.github.com/repos/%s/%s/releases", url.PathEscape(username), url.PathEscape(repoName)),
		AuthType: lib.HeaderAuth,
		Accept:   "application/vnd.github.v3+json",
	}
}

func gitHubUploadReleaseEndpoint(uri string) lib.Endpoint {
	return lib.Endpoint{
		Url:      strings.Split(uri, "{")[0], // hypermedia url will be coming
		AuthType: lib.HeaderAuth,
		Accept:   "application/vnd.github.v3+json",
	}
}

func gitHubCreateReleaseEndpoint(username string, repoName string) lib.Endpoint {
	return lib.Endpoint{
		Url:      fmt.Sprintf("https://api.github.com/repos/%s/%s/releases", url.PathEscape(username), url.PathEscape(repoName)),
		AuthType: lib.HeaderAuth,
		Accept:   "application/vnd.github.v3+json",
	}
}
