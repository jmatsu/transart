package client

import (
	"encoding/json"
	"github.com/jmatsu/transart/client/entity"
	"github.com/jmatsu/transart/lib"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/guregu/null.v3"
)

type CircleCI interface {
	GetArtifacts(vcsType string, username string, reponame string, token lib.Token, buildNum uint) ([]entity.CircleCIArtifact, error)
	DownloadArtifact(vcsType string, username string, reponame string, token lib.Token, artifact entity.CircleCIArtifact) ([]byte, error)
	GetJobInfos(vcsType string, username string, reponame string, token lib.Token, branch null.String) ([]entity.CircleCIJobInfo, error)
}

type CircleCIClient struct {
	vcsType  string
	username string
	reponame string
	token    lib.Token
	branch   null.String
	c        CircleCI
	Err      error
}

func NewCircleCIClient(vcsType string, username string, reponame string, token null.String, branch null.String) CircleCIClient {
	return CircleCIClient{
		vcsType:  vcsType,
		username: username,
		reponame: reponame,
		token:    newCircleCIToken(token),
		branch:   branch,
		c:        &circleCIImpl{},
	}
}

func (cc *CircleCIClient) GetArtifacts(buildNum uint, prod func(entity.CircleCIArtifact) bool) []entity.CircleCIArtifact {
	if !lib.IsNil(cc.Err) {
		return nil
	}

	artifacts, err := cc.c.GetArtifacts(cc.vcsType, cc.username, cc.reponame, cc.token, buildNum)

	if err != nil {
		cc.Err = err
		return nil
	}

	if len(artifacts) == 0 {
		cc.Err = errors.New("no artifact was found")
		return nil
	}

	var results []entity.CircleCIArtifact

	for _, a := range artifacts {
		if prod(a) {
			results = append(results, a)
		} else {
			logrus.Debugf("%s was skipped\n", a.Path)
		}
	}

	if len(results) == 0 {
		cc.Err = errors.New("no target artifact was found")
		return nil
	}

	return results
}

func (cc *CircleCIClient) DownloadArtifact(artifact entity.CircleCIArtifact) []byte {
	if !lib.IsNil(cc.Err) {
		return nil
	}

	bytes, err := cc.c.DownloadArtifact(cc.vcsType, cc.username, cc.reponame, cc.token, artifact)

	if err != nil {
		cc.Err = err
		return nil
	}

	return bytes
}

func (cc *CircleCIClient) GetJobInfo(prod func(entity.CircleCIJobInfo) bool) entity.CircleCIJobInfo {
	var jobInfo entity.CircleCIJobInfo

	if !lib.IsNil(cc.Err) {
		return jobInfo
	}

	jobInfos, err := cc.c.GetJobInfos(cc.vcsType, cc.username, cc.reponame, cc.token, cc.branch)

	if err != nil {
		cc.Err = err
		return jobInfo
	}

	if len(jobInfos) == 0 {
		cc.Err = errors.New("no job info was found")
		return jobInfo
	}

	for _, j := range jobInfos {
		if !j.HasArtifact {
			continue
		} else {
			logrus.Debugf("Skipped due to the job info has no artifacts. Job no. is %d\n", j.BuildNum)
		}

		if prod(j) {
			return j
		}
	}

	cc.Err = errors.New("no target job info was found")

	return jobInfo
}

type circleCIImpl struct {
}

func (c circleCIImpl) GetArtifacts(vcsType string, username string, reponame string, token lib.Token, buildNum uint) ([]entity.CircleCIArtifact, error) {
	var artifacts []entity.CircleCIArtifact

	endpoint := circleCIArtifactListEndpoint(vcsType, username, reponame, buildNum)

	if bytes, err := lib.GetRequest(endpoint, token, nil); err != nil {
		return nil, err
	} else if err := json.Unmarshal(bytes, &artifacts); err != nil {
		return nil, err
	}

	return artifacts, nil
}

func (c circleCIImpl) DownloadArtifact(vcsType string, username string, reponame string, token lib.Token, artifact entity.CircleCIArtifact) ([]byte, error) {
	endpoint := circleCIDownloadArtifactEndpoint(artifact)

	if bytes, err := lib.GetRequest(endpoint, token, nil); err != nil {
		return nil, err
	} else {
		return bytes, nil
	}
}

func (c circleCIImpl) GetJobInfos(vcsType string, username string, reponame string, token lib.Token, branch null.String) ([]entity.CircleCIJobInfo, error) {
	var jobs []entity.CircleCIJobInfo

	apiEndpoint := circleCIJobInfoListEndpoint(vcsType, username, reponame, branch)

	if bytes, err := lib.GetRequest(apiEndpoint, token, CompletedParam); err != nil {
		return nil, err
	} else if err := json.Unmarshal(bytes, &jobs); err != nil {
		err = errors.Wrap(err, "an error happened while parsing the response as json")
		return nil, err
	}

	return jobs, nil
}
