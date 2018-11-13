package action

import (
	"encoding/json"
	"fmt"
	"github.com/jmatsu/artifact-transfer/core"
	"github.com/jmatsu/artifact-transfer/github"
	"github.com/jmatsu/artifact-transfer/github/entity"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/url"
	"path/filepath"
)

func getReleases(config github.Config) ([]entity.Release, error) {
	if err := config.Validate(); err != nil {
		return nil, err
	}

	var releases []entity.Release

	apiEndpoint := github.ReleaseListEndpoint(config.GetUsername(), config.GetRepoName())

	if bytes, err := core.GetRequest(apiEndpoint, github.NewToken(config.GetApiToken()), nil); err != nil {
		return nil, err
	} else if err := json.Unmarshal(bytes, &releases); err != nil {
		logrus.Debugln(string(bytes))
		err = errors.Wrap(err, "an error happened while parsing the response as json")
		return nil, err
	}

	return releases, nil
}

func GetDraftRelease(config github.Config) (entity.Release, error) {
	var release entity.Release

	if err := config.Validate(); err != nil {
		return release, err
	}

	releases, err := getReleases(config)

	if err != nil {
		return release, err
	}

	for _, r := range releases {
		if !r.IsDraft {
			continue
		}

		return r, err
	}

	return release, fmt.Errorf("draft release is not found\n")
}

func UploadToRelease(config github.Config, release entity.Release, path string) (asset entity.Asset, err error) {
	err = config.Validate()

	if err != nil {
		return
	}

	if !config.GetApiToken().Valid {
		err = fmt.Errorf("api key is required\n")
		return
	}

	bytes, err := ioutil.ReadFile(path)

	if err != nil {
		return
	}

	params := url.Values{}
	params.Set("name", filepath.Base(path))

	endpoint := github.UploadReleaseEndpoint(release.UploadUrlInHypermedia)

	bytes, err = core.PostRequest(endpoint, github.NewToken(config.GetApiToken()), params, bytes)

	if err != nil {
		return
	}

	err = json.Unmarshal(bytes, &asset)

	if err != nil {
		return
	}

	return
}
