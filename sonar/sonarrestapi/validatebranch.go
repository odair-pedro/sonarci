package sonarrestapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

const routeValidateBranch = "/api/measures/component?componentKey=%s&branch=%s&metricKeys=alert_status"
const routeBranchDetails = "/dashboard?id=%s&branch=%s"

func (restApi *restApi) ValidateBranch(project string, branch string) error {
	chBuff, chErr := restApi.DoGet(fmt.Sprintf(routeValidateBranch, escapeValue(project), escapeValue(branch)))
	err := <-chErr
	if err != nil {
		return err
	}

	buff := <-chBuff
	wrapper := &branchStatusWrapper{}
	err = json.Unmarshal(buff, wrapper)
	if err != nil {
		return err
	}

	wrapper.checkInfo(project, branch)
	return restApi.validateBranchStatus(wrapper.Component)
}

func (restApi *restApi) validateBranchStatus(status branchStatus) error {
	const statusError = "ERROR"
	if len(status.Measures) < 1 {
		return errors.New(fmt.Sprintf("Failure on validate quality gate results\nFor more detail, visit: %s",
			strings.TrimRight(restApi.Server, "/")+fmt.Sprintf(routeBranchDetails, escapeValue(status.Project), escapeValue(status.Branch))))
	}

	isValid := strings.ToUpper(status.Measures[0].Value) != statusError
	if !isValid {
		return errors.New(fmt.Sprintf("Branch %s has not been passed on quality gate\nFor more detail, visit: %s", escapeValue(status.Branch),
			strings.TrimRight(restApi.Server, "/")+fmt.Sprintf(routeBranchDetails, escapeValue(status.Project), escapeValue(status.Branch))))
	}

	return nil
}

type branchStatusWrapper struct {
	Component branchStatus `json:"component"`
}

type branchStatus struct {
	Measures []branchStatusMeasure `json:"measures"`
	Branch   string                `json:"branch"`
	Project  string                `json:"key"`
}

type branchStatusMeasure struct {
	Value string `json:"value"`
}

func (wrp *branchStatusWrapper) checkInfo(project string, branch string) {
	if wrp.Component.Branch == "" {
		wrp.Component.Branch = branch
	}
	if wrp.Component.Project == "" {
		wrp.Component.Project = project
	}
}
