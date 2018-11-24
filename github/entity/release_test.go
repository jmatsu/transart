package entity

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArtifact(t *testing.T) {
	bytes := []byte("{\"id\": 1, \"url\": \"this is a url\",\"upload_url\": \"this is a upload_url\",\"tag_name\": \"this is a tag_name\", \"draft\": true }")

	release := Release{}

	if err := json.Unmarshal(bytes, &release); err != nil {
		t.Error(err)
	} else {
		assert.EqualValues(t, 1, release.Id)
		assert.EqualValues(t, "this is a url", release.Url)
		assert.EqualValues(t, "this is a upload_url", release.UploadUrlInHypermedia)
		assert.EqualValues(t, "this is a tag_name", release.TagName)
		assert.EqualValues(t, true, release.IsDraft)
	}
}
