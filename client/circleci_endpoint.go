package client

import (
	"fmt"
	"github.com/jmatsu/transart/client/entity"
	"github.com/jmatsu/transart/lib"
	"gopkg.in/guregu/null.v3"
	"net/url"
)

var CompletedParam = url.Values(
	map[string][]string{
		"filter": {"completed"},
	})

func baseCircleCIApiUrl(vcs string, username string, repoName string) string {
	return fmt.Sprintf("https://circleci.com/api/v1.1/project/%s/%s/%s", vcs, url.PathEscape(username), url.PathEscape(repoName))
}

func circleCIJobInfoListEndpoint(vcs string, username string, repoName string, branch null.String) lib.Endpoint {
	var uri string

	if branch.Valid {
		uri = fmt.Sprintf("%s/tree/%s", baseCircleCIApiUrl(vcs, username, repoName), branch.String)
	} else {
		uri = baseCircleCIApiUrl(vcs, username, repoName)
	}

	return lib.Endpoint{
		Url:      uri,
		AuthType: lib.HeaderAuth,
		Accept:   "application/json",
	}
}

func circleCIArtifactListEndpoint(vcs string, username string, repoName string, buildNum uint) lib.Endpoint {
	return lib.Endpoint{
		Url:      fmt.Sprintf("%s/%d/artifacts", baseCircleCIApiUrl(vcs, username, repoName), buildNum),
		AuthType: lib.HeaderAuth,
		Accept:   "application/json",
	}
}

func circleCIDownloadArtifactEndpoint(artifact entity.CircleCIArtifact) lib.Endpoint {
	return lib.Endpoint{
		Url:      artifact.DownloadUrl,
		AuthType: lib.ParameterAuth,
	}
}
