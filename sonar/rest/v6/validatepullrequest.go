package v6

const routePullRequestValidation = "/api/measures/component?componentKey=%s&pullRequest=%s&metricKeys=alert_status"

func (restApi *RestApi) ValidatePullRequest(project string, pullRequest string) error {
	return restApi.ValidatePullRequestInternal(routePullRequestValidation, project, pullRequest)
}
