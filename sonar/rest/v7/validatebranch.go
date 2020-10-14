package v7

func (restApi *RestApi) GetRouteForValidateBranch() string {
	return "/api/measures/component?component=%s&branch=%s&metricKeys=alert_status"
}
