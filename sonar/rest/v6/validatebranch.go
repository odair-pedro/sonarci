package v6

const routeBranchValidation = "/api/measures/component?componentKey=%s&branch=%s&metricKeys=alert_status"

func (restApi *RestApi) ValidateBranch(project string, branch string) error {
	return restApi.ValidateBranchInternal(routeBranchValidation, project, branch)
}
