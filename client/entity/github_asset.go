package entity

type GitHubAsset struct {
	Id                 uint   `json:"id"`
	UploadState        string `json:"state"`
	Name               string `json:"name"`
	Size               uint   `json:"size"`
	DownloadBrowserUrl string `json:"browser_download_url"`
}

func (a GitHubAsset) IsUploaded() bool {
	return a.UploadState == "uploaded"
}
