package source

func NewDownloadAction() Actions {
	return Actions{
		CircleCI: downloadFromCircleCI,
		Local:    downloadFromLocal,
	}
}
