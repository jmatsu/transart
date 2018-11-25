package source

func NewDownloadAction() SourceActions {
	return SourceActions{
		CircleCI: downloadFromCircleCI,
		Local:    downloadFromLocal,
	}
}
