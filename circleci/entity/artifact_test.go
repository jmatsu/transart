package entity

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArtifact(t *testing.T) {
	bytes := []byte("{\"pretty_path\": \"this is a pretty path\",\"path\": \"this is a path\",\"url\": \"this is a url\"}")

	artifact := Artifact{}

	if err := json.Unmarshal(bytes, &artifact); err != nil {
		t.Error(err)
	} else {
		assert.EqualValues(t, "this is a path", artifact.Path)
		assert.EqualValues(t, "this is a pretty path", artifact.PrettyPath)
		assert.EqualValues(t, "this is a url", artifact.DownloadUrl)
	}
}
