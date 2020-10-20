package v7

const routeBranchValidation = "/api/measures/component?component=%s&branch=%s&metricKeys=alert_status"

func (restApi *RestApi) ValidateBranch(project string, branch string) error {
	return restApi.ValidateBranchInternal(routeBranchValidation, project, branch)
}
