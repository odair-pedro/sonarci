package rest

import (
	"encoding/json"
	"fmt"
	"github.com/prometheus/common/log"
	"strings"
)

const routeValidateBranch = "/api/measures/component?componentKey=%s&branch=%s&metricKeys=alert_status"
const routeBranchDetails = "/dashboard?branch=%s&id=%s"

func (api *Api) ValidateBranch(project string, branch string) (bool, error) {
	chBuff, chErr := api.DoGet(fmt.Sprintf(routeValidateBranch, escapeValue(project), escapeValue(branch)))
	err := <-chErr
	if err != nil {
		return false, err
	}

	buff := <-chBuff
	branchSt := &branchStatus{}
	err = json.Unmarshal(buff, branchSt)
	if err != nil {
		return false, err
	}

	return api.validateBranchStatus(branchSt), nil
}

func (api *Api) validateBranchStatus(status *branchStatus) bool {
	const statusError = "ERROR"
	if len(status.Measures) < 1 {
		log.Errorf("Failure on validate quality gate results\nFor more detail, visit: %s", escapeValue(status.Branch),
			fmt.Sprintf("%s/%s", strings.TrimRight(api.Connection.Server, "/"),
				fmt.Sprintf(routeBranchDetails, escapeValue(status.Branch), escapeValue(status.Project))))
		return false
	}

	isValid := strings.ToUpper(status.Measures[0].Value) != statusError
	if !isValid {
		log.Errorf("Branch %s has not been passed on quality gate\nFor more detail, visit: %s", escapeValue(status.Branch),
			fmt.Sprintf("%s/%s", strings.TrimRight(api.Connection.Server, "/"),
				fmt.Sprintf(routeBranchDetails, escapeValue(status.Branch), escapeValue(status.Project))))
	}

	return isValid
}

type branchStatus struct {
	Measures []branchStatusMeasure `json:"measures"`
	Branch   string                `json:"branch"`
	Project  string                `json:"key"`
}

type branchStatusMeasure struct {
	Value string `json:"value"`
}
