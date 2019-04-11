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
	CreateDraftRelease(username string, reponame string, token lib.Token) (entity.GitHubRelease, error)
	GetReleases(username string, reponame string, token lib.Token) ([]entity.GitHubRelease, error)
	GetAssets(username string, reponame string, token lib.Token, release entity.GitHubRelease) ([]entity.GitHubAsset, error)
	AttachFileToRelease(username string, reponame string, token lib.Token, release entity.GitHubRelease, path string) (entity.GitHubAsset, error)
	DeleteAsset(username string, reponame string, token lib.Token, asset entity.GitHubAsset) error
	DownloadAsset(username string, reponame string, token lib.Token, asset entity.GitHubAsset) ([]byte, error)
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

func (gc *GitHubClient) GetDraftRelease() entity.GitHubRelease {
	var release entity.GitHubRelease

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

func (gc *GitHubClient) CreateDraftRelease() entity.GitHubRelease {
	var release entity.GitHubRelease

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

func (gc *GitHubClient) GetLatestRelease() entity.GitHubRelease {
	var release entity.GitHubRelease

	if !lib.IsNil(gc.Err) {
		return release
	}

	releases, err := gc.c.GetReleases(gc.username, gc.reponame, gc.token)

	if err != nil {
		gc.Err = err
		return release
	}

	if len(releases) == 0 {
		gc.Err = errors.New("no release found")
		return release
	}

	release = releases[0]

	return release
}

func (gc *GitHubClient) GetAssets(release entity.GitHubRelease) []entity.GitHubAsset {
	if !lib.IsNil(gc.Err) {
		return nil
	}

	assets, err := gc.c.GetAssets(gc.username, gc.reponame, gc.token, release)

	if err != nil {
		gc.Err = err
		return nil
	}

	return assets
}

func (gc *GitHubClient) UploadFileToRelease(release entity.GitHubRelease, path string) entity.GitHubAsset {
	var asset entity.GitHubAsset

	if !lib.IsNil(gc.Err) {
		return asset
	}

	asset, err := gc.c.AttachFileToRelease(gc.username, gc.reponame, gc.token, release, path)

	if err != nil {
		gc.Err = err
		return asset
	}

	logrus.Infof("%s has been uploaded\n", asset.Name)
	logrus.Debugf("%v\n", asset)

	return asset
}

func (gc *GitHubClient) DeleteAssetFromRelease(asset entity.GitHubAsset) {
	if !lib.IsNil(gc.Err) {
		return
	}

	err := gc.c.DeleteAsset(gc.username, gc.reponame, gc.token, asset)

	if err != nil {
		gc.Err = err
		return
	}

	logrus.Infof("%s has been deleted\n", asset.Name)
	logrus.Debugf("%v\n", asset)
}

func (gc *GitHubClient) DownloadAsset(asset entity.GitHubAsset) []byte {
	if !lib.IsNil(gc.Err) {
		return nil
	}

	bytes, err := gc.c.DownloadAsset(gc.username, gc.reponame, gc.token, asset)

	if err != nil {
		gc.Err = err
		return nil
	}

	return bytes
}

type gitHubImpl struct {
}

func (gh gitHubImpl) CreateDraftRelease(username string, reponame string, token lib.Token) (entity.GitHubRelease, error) {
	var release entity.GitHubRelease

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

func (gh gitHubImpl) GetReleases(username string, reponame string, token lib.Token) ([]entity.GitHubRelease, error) {
	var releases []entity.GitHubRelease

	apiEndpoint := gitHubReleaseListEndpoint(username, reponame)

	if bytes, err := lib.GetRequest(apiEndpoint, token, nil); err != nil {
		return nil, err
	} else if err := json.Unmarshal(bytes, &releases); err != nil {
		err = errors.Wrap(err, "an error happened while parsing the response as json")
		return nil, err
	}

	return releases, nil
}

func (gh gitHubImpl) GetAssets(username string, reponame string, token lib.Token, release entity.GitHubRelease) ([]entity.GitHubAsset, error) {
	var assets []entity.GitHubAsset

	apiEndpoint := gitHubAssetListEndpoint(username, reponame, release.Id)

	if bytes, err := lib.GetRequest(apiEndpoint, token, nil); err != nil {
		return nil, err
	} else if err := json.Unmarshal(bytes, &assets); err != nil {
		err = errors.Wrap(err, "an error happened while parsing the response as json")
		return nil, err
	}

	return assets, nil

}

func (gh gitHubImpl) AttachFileToRelease(username string, reponame string, token lib.Token, release entity.GitHubRelease, path string) (entity.GitHubAsset, error) {
	var asset entity.GitHubAsset

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

func (gh gitHubImpl) DeleteAsset(username string, reponame string, token lib.Token, asset entity.GitHubAsset) error {
	endpoint := gitHubAssetEndpoint(username, reponame, asset.Id)

	_, err := lib.DeleteRequest(endpoint, token, nil)

	return err
}

func (gh gitHubImpl) DownloadAsset(username string, reponame string, token lib.Token, asset entity.GitHubAsset) ([]byte, error) {
	endpoint := gitHubAssetDownloadEndpoint(asset.DownloadBrowserUrl)

	if bytes, err := lib.GetRequest(endpoint, token, nil); err != nil {
		return nil, err
	} else {
		return bytes, nil
	}
}
