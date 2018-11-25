package entity

type Release struct {
	Id                    uint   `json:"id"`
	Url                   string `json:"url"`
	UploadUrlInHypermedia string `json:"upload_url"`
	TagName               string `json:"tag_name"`
	IsDraft               bool   `json:"draft"`
}
