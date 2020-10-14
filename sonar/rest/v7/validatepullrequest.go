package v7

func (restApi *RestApi) GetRouteForValidatePullRequest() string {
	return "/api/measures/component?component=%s&pullRequest=%s&metricKeys=alert_status"
}
