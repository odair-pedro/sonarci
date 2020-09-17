package rest

import (
	"encoding/json"
	"fmt"
	"strings"
)

const routeValidateBranch = "/api/measures/component?componentKey=%s&branch=%s&metricKeys=alert_status"

func (api *Api) ValidateBranch(project string, branch string) (bool, error) {
	//conn := http.NewConnection(abstract.server, abstract.token, abstract.timeout)
	chBuff, chErr := api.DoGet(fmt.Sprintf(routeValidateBranch, project, branch))
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

	return branchSt.isValid(), nil
}

type branchStatus struct {
	Measures []branchStatusMeasure `json:"measures"`
}

type branchStatusMeasure struct {
	Value string `json:"value"`
}

func (status branchStatus) isValid() bool {
	const statusError = "ERROR"
	if len(status.Measures) < 1 {
		return false
	}

	return strings.ToUpper(status.Measures[0].Value) == statusError

}
