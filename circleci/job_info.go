package circleci

import (
	"encoding/json"
	"github.com/jmatsu/transart/circleci/entity"
	"github.com/jmatsu/transart/config"
	"github.com/jmatsu/transart/lib"
	"github.com/pkg/errors"
)

func getJobInfos(circleCIConfig config.CircleCIConfig) ([]entity.JobInfo, error) {
	if err := circleCIConfig.Validate(); err != nil {
		return nil, err
	}

	var jobs []entity.JobInfo

	apiEndpoint := JobInfoListEndpoint(string(circleCIConfig.GetVcsType()), circleCIConfig.GetUsername(), circleCIConfig.GetRepoName(), circleCIConfig.GetBranch())

	if bytes, err := lib.GetRequest(apiEndpoint, NewToken(circleCIConfig.GetApiToken()), CompletedParam); err != nil {
		return nil, err
	} else if err := json.Unmarshal(bytes, &jobs); err != nil {
		err = errors.Wrap(err, "an error happened while parsing the response as json")
		return nil, err
	}

	return jobs, nil
}
