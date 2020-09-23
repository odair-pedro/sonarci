package sonarrestapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

const routeValidateBranch = "/api/measures/component?componentKey=%s&branch=%s&metricKeys=alert_status"
const routeBranchDetails = "/dashboard?branch=%s&id=%s"

func (restApi *restApi) ValidateBranch(project string, branch string) error {
	chBuff, chErr := restApi.DoGet(fmt.Sprintf(routeValidateBranch, escapeValue(project), escapeValue(branch)))
	err := <-chErr
	if err != nil {
		return err
	}

	buff := <-chBuff
	branchSt := &branchStatus{}
	err = json.Unmarshal(buff, branchSt)
	if err != nil {
		return err
	}

	return restApi.validateBranchStatus(branchSt)
}

func (restApi *restApi) validateBranchStatus(status *branchStatus) error {
	const statusError = "ERROR"
	if len(status.Measures) < 1 {
		return errors.New(fmt.Sprintf("Failure on validate quality gate results\nFor more detail, visit: %s",
			strings.TrimRight(restApi.Server, "/")+fmt.Sprintf(routeBranchDetails, escapeValue(status.Branch), escapeValue(status.Project))))
	}

	isValid := strings.ToUpper(status.Measures[0].Value) != statusError
	if !isValid {
		return errors.New(fmt.Sprintf("Branch %s has not been passed on quality gate\nFor more detail, visit: %s", escapeValue(status.Branch),
			strings.TrimRight(restApi.Server, "/")+fmt.Sprintf(routeBranchDetails, escapeValue(status.Branch), escapeValue(status.Project))))
	}

	return nil
}

type branchStatus struct {
	Measures []branchStatusMeasure `json:"measures"`
	Branch   string                `json:"branch"`
	Project  string                `json:"key"`
}

type branchStatusMeasure struct {
	Value string `json:"value"`
}
