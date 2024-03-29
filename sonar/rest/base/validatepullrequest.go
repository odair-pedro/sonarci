package base

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

const routePullRequestDetails = "/dashboard?id=%s&pullRequest=%s"

func (restApi *RestApi) ValidatePullRequestInternal(routeApi string, project string, pullRequest string) error {
	chBuff, chErr := restApi.Get(fmt.Sprintf(routeApi, escapeValue(project), pullRequest))
	err := <-chErr
	if err != nil {
		return err
	}

	buff := <-chBuff
	wrapper := &pullRequestStatusWrapper{}
	err = json.Unmarshal(buff, wrapper)
	if err != nil {
		return err
	}

	wrapper.checkInfo(project, pullRequest)
	return restApi.validatePullRequestStatus(wrapper.Component)
}

func (restApi *RestApi) validatePullRequestStatus(status pullRequestStatus) error {
	const statusError = "ERROR"
	if len(status.Measures) < 1 {
		return errors.New(fmt.Sprintf("Failure on validate quality gate results\nFor more detail, visit: %s",
			strings.TrimRight(restApi.GetHostServer(), "/")+fmt.Sprintf(routePullRequestDetails, escapeValue(status.Project), status.PullRequest)))
	}

	isValid := strings.ToUpper(status.Measures[0].Value) != statusError
	if !isValid {
		return errors.New(fmt.Sprintf("PullRequest %s has not been passed on quality gate\nFor more detail, visit: %s", status.PullRequest,
			strings.TrimRight(restApi.GetHostServer(), "/")+fmt.Sprintf(routePullRequestDetails, escapeValue(status.Project), status.PullRequest)))
	}

	return nil
}

type pullRequestStatusWrapper struct {
	Component pullRequestStatus `json:"component"`
}

type pullRequestStatus struct {
	Measures    []pullRequestStatusMeasure `json:"measures"`
	PullRequest string                     `json:"pullRequest"`
	Project     string                     `json:"key"`
}

type pullRequestStatusMeasure struct {
	Value string `json:"value"`
}

func (wrp *pullRequestStatusWrapper) checkInfo(project string, pullRequest string) {
	if wrp.Component.PullRequest == "" {
		wrp.Component.PullRequest = pullRequest
	}
	if wrp.Component.Project == "" {
		wrp.Component.Project = project
	}
}
