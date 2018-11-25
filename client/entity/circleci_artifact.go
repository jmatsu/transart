package entity

type CircleCIArtifact struct {
	Path        string `json:"path"`
	PrettyPath  string `json:"pretty_path"`
	DownloadUrl string `json:"url"`
}
