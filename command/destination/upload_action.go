package destination

func NewUploadAction() Actions {
	return Actions{
		GitHubRelease: uploadToGithubRelease,
		Local:         uploadToLocal,
	}
}
