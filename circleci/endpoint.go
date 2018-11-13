package circleci

import (
	"fmt"
	"github.com/jmatsu/artifact-transfer/circleci/entity"
	"github.com/jmatsu/artifact-transfer/core"
	"gopkg.in/guregu/null.v3"
	"net/url"
)

type VcsType string

const (
	GitHub    VcsType = "github"
	Bitbucket         = "bitbucket"
)

func newVcsType(v string) (VcsType, error) {
	t := VcsType(v)

	switch t {
	case GitHub:
		return t, nil
	case Bitbucket:
		return t, nil
	default:
		return t, fmt.Errorf("%s is invalid vcs type\n", v)
	}
}

var CompletedParam = url.Values(
	map[string][]string{
		"filter": {"completed"},
	})

func baseApiUrl(vcs VcsType, username string, repoName string) string {
	return fmt.Sprintf("https://circleci.com/api/v1.1/project/%s/%s/%s", vcs, username, repoName)
}

func JobInfoListEndpoint(vcs VcsType, username string, repoName string, branch null.String) core.Endpoint {
	var uri string

	if branch.Valid {
		uri = fmt.Sprintf("%s/tree/%s", baseApiUrl(vcs, username, repoName), branch.String)
	} else {
		uri = baseApiUrl(vcs, username, repoName)
	}

	return core.Endpoint{
		Url:      uri,
		AuthType: core.HeaderAuth,
	}
}

func ArtifactListEndpoint(vcs VcsType, username string, repoName string, buildNum uint) core.Endpoint {
	return core.Endpoint{
		Url:      fmt.Sprintf("%s/%d/artifacts", baseApiUrl(vcs, username, repoName), buildNum),
		AuthType: core.HeaderAuth,
	}
}

func DownloadArtifactEndpoint(artifact entity.Artifact) core.Endpoint {
	return core.Endpoint{
		Url:      artifact.DownloadUrl,
		AuthType: core.ParameterAuth,
	}
}
