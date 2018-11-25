package client

import (
	"encoding/json"
	"fmt"
	"github.com/jmatsu/transart/client/entity"
	"github.com/jmatsu/transart/lib"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/guregu/null.v3"
	"io/ioutil"
	"net/url"
	"path/filepath"
)

type GitHub interface {
	CreateDraftRelease(username string, reponame string, token lib.Token) (entity.Release, error)
	GetReleases(username string, reponame string, token lib.Token) ([]entity.Release, error)
	UploadToRelease(username string, reponame string, token lib.Token, release entity.Release, path string) (entity.Asset, error)
}

type GitHubClient struct {
	username string
	reponame string
	token    lib.Token
	c        GitHub
	Err      error
}

func NewGitHubClient(username string, reponame string, token null.String) GitHubClient {
	return GitHubClient{
		username: username,
		reponame: reponame,
		token:    newGitHubToken(token),
		c:        &gitHubImpl{},
	}
}

func (gc *GitHubClient) GetDraftRelease() entity.Release {
	var release entity.Release

	if !lib.IsNil(gc.Err) {
		return release
	}

	releases, err := gc.c.GetReleases(gc.username, gc.reponame, gc.token)

	if err != nil {
		gc.Err = err
		return release
	}

	for _, r := range releases {
		if !r.IsDraft {
			continue
		}

		return r
	}

	gc.Err = fmt.Errorf("draft release is not found")

	return release
}

func (gc *GitHubClient) CreateDraftRelease() entity.Release {
	var release entity.Release

	if !lib.IsNil(gc.Err) {
		return release
	}

	release, err := gc.c.CreateDraftRelease(gc.username, gc.reponame, gc.token)

	if err != nil {
		gc.Err = err
		return release
	}

	return release
}

func (gc *GitHubClient) UploadFileToRelease(release entity.Release, path string) entity.Asset {
	var asset entity.Asset

	if !lib.IsNil(gc.Err) {
		return asset
	}

	asset, err := gc.c.UploadToRelease(gc.username, gc.reponame, gc.token, release, path)

	if err != nil {
		gc.Err = err
		return asset
	}

	logrus.Debugf("%s has been uploaded\n", asset.Name)

	return asset
}

type gitHubImpl struct {
}

func (gh gitHubImpl) CreateDraftRelease(username string, reponame string, token lib.Token) (entity.Release, error) {
	var release entity.Release

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
		return release, err
	}

	apiEndpoint := gitHubCreateReleaseEndpoint(username, reponame)

	bytes, err := lib.PostRequest(apiEndpoint, token, nil, bodyBytes)

	if err != nil {
		return release, err
	}

	if err := json.Unmarshal(bytes, &release); err != nil {
		return release, err
	}

	return release, nil
}

func (gh gitHubImpl) GetReleases(username string, reponame string, token lib.Token) ([]entity.Release, error) {
	var releases []entity.Release

	apiEndpoint := gitHubReleaseListEndpoint(username, reponame)

	if bytes, err := lib.GetRequest(apiEndpoint, token, nil); err != nil {
		return nil, err
	} else if err := json.Unmarshal(bytes, &releases); err != nil {
		err = errors.Wrap(err, "an error happened while parsing the response as json")
		return nil, err
	}

	return releases, nil
}

func (gh gitHubImpl) UploadToRelease(username string, reponame string, token lib.Token, release entity.Release, path string) (entity.Asset, error) {
	var asset entity.Asset

	fileBytes, err := ioutil.ReadFile(path)

	if err != nil {
		return asset, err
	}

	params := url.Values{}
	params.Set("name", filepath.Base(path))

	endpoint := gitHubUploadReleaseEndpoint(release.UploadUrlInHypermedia)

	bytes, err := lib.PostRequest(endpoint, token, params, fileBytes)

	if err != nil {
		return asset, err
	}

	if err := json.Unmarshal(bytes, &asset); err != nil {
		return asset, err
	}

	return asset, nil
}
