package entity

type Asset struct {
	Id          uint   `json:"id"`
	UploadState string `json:"state"`
	Name        string `json:"name"`
	Size        uint   `json:"size"`
}

func (a Asset) IsUploaded() bool {
	return a.UploadState == "uploaded"
}
