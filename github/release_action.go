package github

import (
	"encoding/json"
	"fmt"
	"github.com/jmatsu/artifact-transfer/config"
	"github.com/jmatsu/artifact-transfer/lib"
	"github.com/jmatsu/artifact-transfer/github/entity"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/url"
	"path/filepath"
)

func getReleases(gitHubConfig config.GitHubConfig) ([]entity.Release, error) {
	if err := gitHubConfig.Validate(); err != nil {
		return nil, err
	}

	var releases []entity.Release

	apiEndpoint := ReleaseListEndpoint(gitHubConfig.GetUsername(), gitHubConfig.GetRepoName())

	if bytes, err := lib.GetRequest(apiEndpoint, NewToken(gitHubConfig.GetApiToken()), nil); err != nil {
		return nil, err
	} else if err := json.Unmarshal(bytes, &releases); err != nil {
		err = errors.Wrap(err, "an error happened while parsing the response as json")
		return nil, err
	}

	return releases, nil
}

func GetDraftRelease(gitHubConfig config.GitHubConfig) (entity.Release, error) {
	var release entity.Release

	if err := gitHubConfig.Validate(); err != nil {
		return release, err
	}

	releases, err := getReleases(gitHubConfig)

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

func CreateDraftRelease(gitHubConfig config.GitHubConfig) (release entity.Release, err error) {
	err = gitHubConfig.Validate()

	if err != nil {
		return
	}

	body := struct {
		A string `json:"tag_name"`
		B string `json:"target_commitish"`
		C bool   `json:"draft"`
	}{
		"", // this is required but the GitHub API can accept the request
		"",
		true,
	}

	bodyBytes, err := json.Marshal(body)

	if err != nil {
		return
	}

	apiEndpoint := CreateReleaseEndpoint(gitHubConfig.GetUsername(), gitHubConfig.GetRepoName())

	bytes, err := lib.PostRequest(apiEndpoint, NewToken(gitHubConfig.GetApiToken()), nil, bodyBytes)

	if err != nil {
		return
	}

	err = json.Unmarshal(bytes, &release)

	if err != nil {
		return
	}

	return
}

func UploadToRelease(gitHubConfig config.GitHubConfig, release entity.Release, path string) (asset entity.Asset, err error) {
	err = gitHubConfig.Validate()

	if err != nil {
		return
	}

	if !gitHubConfig.GetApiToken().Valid {
		err = fmt.Errorf("api key is required\n")
		return
	}

	bytes, err := ioutil.ReadFile(path)

	if err != nil {
		return
	}

	params := url.Values{}
	params.Set("name", filepath.Base(path))

	endpoint := UploadReleaseEndpoint(release.UploadUrlInHypermedia)

	bytes, err = lib.PostRequest(endpoint, NewToken(gitHubConfig.GetApiToken()), params, bytes)

	if err != nil {
		return
	}

	err = json.Unmarshal(bytes, &asset)

	if err != nil {
		return
	}

	return
}
