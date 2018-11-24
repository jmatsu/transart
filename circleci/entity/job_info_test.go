package entity

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

//BuildNum        uint              `json:"build_num"`
//BuildParameters map[string]string `json:"build_parameters"`
//BuildUrl        string            `json:"build_url"`
//Branch          string            `json:"branch"`
//HasArtifact     bool              `json:"has_artifacts"`
//Lifecycle       string            `json:"lifecycle"`
//Subject         string            `json:"subject"`
//Username        string            `json:"username"`
//RepoName        string            `json:"reponame"`
//VcsType         string            `json:"vcs_type"`
//VcsRevision     string            `json:"vcs_revision"`

func TestJobInfo(t *testing.T) {
	bytes := []byte("{" +
		"\"build_num\": 1," +
		"\"build_parameters\": { \"CIRCLE_JOB\": \"build\" }," +
		"\"build_url\": \"this is a build url\"," +
		"\"branch\": \"this is a branch\"," +
		"\"has_artifacts\": false," +
		"\"lifecycle\": \"this is a lifecycle\"," +
		"\"subject\": \"this is a subject\"," +
		"\"username\": \"this is a username\"," +
		"\"reponame\": \"this is a reponame\"," +
		"\"vcs_type\": \"this is a vcs type\"," +
		"\"vcs_revision\": \"this is a vcs revision\"" +
		"}")

	jobInfo := JobInfo{}

	if err := json.Unmarshal(bytes, &jobInfo); err != nil {
		t.Error(err)
	} else {
		assert.EqualValues(t, 1, jobInfo.BuildNum)
		assert.EqualValues(t, map[string]string{"CIRCLE_JOB": "build"}, jobInfo.BuildParameters)
		assert.EqualValues(t, "this is a build url", jobInfo.BuildUrl)
		assert.EqualValues(t, "this is a branch", jobInfo.Branch)
		assert.EqualValues(t, false, jobInfo.HasArtifact)
		assert.EqualValues(t, "this is a lifecycle", jobInfo.Lifecycle)
		assert.EqualValues(t, "this is a subject", jobInfo.Subject)
		assert.EqualValues(t, "this is a username", jobInfo.Username)
		assert.EqualValues(t, "this is a reponame", jobInfo.RepoName)
		assert.EqualValues(t, "this is a vcs type", jobInfo.VcsType)
		assert.EqualValues(t, "this is a vcs revision", jobInfo.VcsRevision)
	}
}

func TestJobInfo_GetBuildJobName(t *testing.T) {
	bytes := []byte("{" +
		"\"build_parameters\": { \"CIRCLE_JOB\": \"build\" }" +
		"}")

	jobInfo := JobInfo{}
	json.Unmarshal(bytes, &jobInfo)

	assert.EqualValues(t, "build", jobInfo.GetBuildJobName())
}

var testJobInfo_HasFinishedTests = []struct {
	in  string
	out bool
}{
	{
		"finished",
		true,
	},
	{
		"no",
		false,
	},
}

func TestJobInfo_HasFinished(t *testing.T) {
	for i, c := range testJobInfo_HasFinishedTests {
		t.Run(fmt.Sprintf("TestJobInfo_HasFinished %d", i), func(t *testing.T) {
			bytes := []byte(fmt.Sprintf("{\"lifecycle\": \"%s\"}", c.in))

			jobInfo := JobInfo{}
			json.Unmarshal(bytes, &jobInfo)

			assert.EqualValues(t, c.out, jobInfo.HasFinished())
		})
	}
}
