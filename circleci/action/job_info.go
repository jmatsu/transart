package action

import (
	"encoding/json"
	"github.com/jmatsu/artifact-transfer/circleci"
	"github.com/jmatsu/artifact-transfer/circleci/entity"
	"github.com/jmatsu/artifact-transfer/core"
	"github.com/pkg/errors"
)

func getJobInfos(config circleci.Config) ([]entity.JobInfo, error) {
	if err := config.Validate(); err != nil {
		return nil, err
	}

	var jobs []entity.JobInfo

	apiEndpoint := circleci.JobInfoListEndpoint(config.GetVcsType(), config.GetUsername(), config.GetRepoName(), config.GetBranch())

	if bytes, err := core.GetRequest(apiEndpoint, circleci.NewToken(config.GetApiToken()), circleci.CompletedParam); err != nil {
		return nil, err
	} else if err := json.Unmarshal(bytes, &jobs); err != nil {
		err = errors.Wrap(err, "an error happened while parsing the response as json")
		return nil, err
	}

	return jobs, nil
}
