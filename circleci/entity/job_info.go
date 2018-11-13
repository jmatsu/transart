package entity

type JobInfo struct {
	BuildNum        uint              `json:"build_num"`
	BuildParameters map[string]string `json:"build_parameters"`
	BuildUrl        string            `json:"build_url"`
	Branch          string            `json:"branch"`
	HasArtifact     bool              `json:"has_artifacts"`
	Lifecycle       string            `json:"lifecycle"`
	Subject         string            `json:"subject"`
	Username        string            `json:"username"`
	RepoName        string            `json:"reponame"`
	VcsType         string            `json:"vcs_type"`
	VcsRevision     string            `json:"vcs_revision"`
}

func (j JobInfo) GetBuildJobName() string {
	return j.BuildParameters["CIRCLE_JOB"]
}

func (j JobInfo) HasFinished() bool {
	return j.Lifecycle == "finished"
}
