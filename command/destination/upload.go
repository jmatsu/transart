package destination

func NewUploadAction() DestinationActions {
	return DestinationActions{
		GitHubRelease: uploadToGithubRelease,
		Local:         uploadToLocal,
	}
}
