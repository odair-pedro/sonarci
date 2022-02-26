package azuredevops

import (
	"encoding/json"
	"fmt"
	"sonarci/decoration/azuredevops/models"
	"strings"
)

const routeListThreadsPullRequest = "%s/_apis/git/repositories/%s/pullRequests/%s/threads?api-version=6.0"

func (decorator *PullRequestDecorator) ClearPreviousComments(pullRequest string) error {
	// TODO: to implements
	return nil
}

func (decorator *PullRequestDecorator) loadMyPullRequestThreads(pullRequest string) ([]string, error) {
	chBuff, chErr := decorator.Request(fmt.Sprintf(routeListThreadsPullRequest, formatPath(decorator.Project), formatPath(decorator.Repository), pullRequest))
	err := <-chErr
	if err != nil {
		return nil, err
	}

	buff := <-chBuff
	threadsWrapper := &models.ThreadModelWrapper{}
	err = json.Unmarshal(buff, threadsWrapper)
	if err != nil {
		return nil, err
	}

	var threadsRet []string
	for _, t := range threadsWrapper.Value {
		if !t.IsDeleted && strings.ToLower(t.Properties.GeneratedBySonarCI.Value) == "true" {
			threadsRet = append(threadsRet, t.Id)
		}
	}

	return threadsRet, nil
}
