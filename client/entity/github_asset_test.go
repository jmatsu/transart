package entity

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGitHubAsset(t *testing.T) {
	bytes := []byte("{\"id\": 10, \"state\": \"this is a state\", \"name\": \"this is a name\", \"size\": 1}")

	asset := GitHubAsset{}

	if err := json.Unmarshal(bytes, &asset); err != nil {
		t.Error(err)
	} else {
		assert.EqualValues(t, 10, asset.Id)
		assert.EqualValues(t, "this is a state", asset.UploadState)
		assert.EqualValues(t, "this is a name", asset.Name)
		assert.EqualValues(t, 1, asset.Size)
	}
}

var testGitHubAsset_IsUploadedTests = []struct {
	in  string
	out bool
}{
	{
		"uploaded",
		true,
	},
	{
		"no",
		false,
	},
}

func TestGitHubAsset_IsUploaded(t *testing.T) {
	for i, c := range testGitHubAsset_IsUploadedTests {
		t.Run(fmt.Sprintf("TestGitHubAsset_IsUploaded %d", i), func(t *testing.T) {
			bytes := []byte(fmt.Sprintf("{\"state\": \"%s\"}", c.in))

			asset := GitHubAsset{}
			json.Unmarshal(bytes, &asset)

			assert.EqualValues(t, c.out, asset.IsUploaded())
		})
	}
}
