package source

func NewDownloadAction() Actions {
	return Actions{
		CircleCI:      downloadFromCircleCI,
		GitHubRelease: downloadFromGitHubRelease,
		Local:         downloadFromLocal,
	}
}
