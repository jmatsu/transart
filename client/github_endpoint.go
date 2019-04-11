package client

import (
	"fmt"
	"github.com/jmatsu/transart/lib"
	"github.com/sirupsen/logrus"
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
	logrus.Debugln(uri)

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

func gitHubAssetListEndpoint(username string, repoName string, releaseId uint) lib.Endpoint {
	return lib.Endpoint{
		Url:      fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/%d/assets", url.PathEscape(username), url.PathEscape(repoName), releaseId),
		AuthType: lib.HeaderAuth,
		Accept:   "application/vnd.github.v3+json",
	}
}

func gitHubAssetEndpoint(username string, repoName string, assetId uint) lib.Endpoint {
	return lib.Endpoint{
		Url:      fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/assets/%d", url.PathEscape(username), url.PathEscape(repoName), assetId),
		AuthType: lib.HeaderAuth,
		Accept:   "application/vnd.github.v3+json",
	}
}

func gitHubAssetDownloadEndpoint(browserDownloadUrl string) lib.Endpoint {
	return lib.Endpoint{
		Url:      browserDownloadUrl,
		AuthType: lib.ParameterAuth,
		Accept:   "application/octet-stream",
	}
}
