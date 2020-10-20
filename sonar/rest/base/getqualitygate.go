package base

import (
	"encoding/json"
	"fmt"
	"sonarci/sonar"
	"strconv"
	"strings"
)

const routeQualityGate = "api/qualitygates/project_status?projectKey=%s&%s=%s"
const routeQualityGateDetails = "/dashboard?id=%s&%s=%s"

func (restApi *RestApi) GetBranchQualityGate(project string, branch string) (sonar.QualityGate, error) {
	return restApi.getQualityGate(project, "branch", branch)
}

func (restApi *RestApi) GetPullRequestQualityGate(project string, pullRequest string) (sonar.QualityGate, error) {
	return restApi.getQualityGate(project, "pullRequest", pullRequest)
}

func (restApi *RestApi) getQualityGate(project string, filter string, value string) (sonar.QualityGate, error) {
	chBuff, chErr := restApi.DoGet(fmt.Sprintf(routeQualityGate, project, filter, value))
	err := <-chErr
	if err != nil {
		return sonar.QualityGate{}, err
	}

	buff := <-chBuff
	wrapper := &qualityGateWrapper{}
	err = json.Unmarshal(buff, wrapper)
	if err != nil {
		return sonar.QualityGate{}, err
	}

	qualityGate := wrapper.convert()
	qualityGate.LinkDetail = strings.TrimRight(restApi.GetHostServer(), "/") +
		fmt.Sprintf(routeQualityGateDetails, escapeValue(project), escapeValue(filter), escapeValue(value))
	return qualityGate, nil
}

type qualityGateWrapper struct {
	QualityGate qualityGateStatus `json:"projectStatus"`
}

type qualityGateStatus struct {
	Status     string                 `json:"status"`
	Conditions []qualityGateCondition `json:"conditions"`
}

type qualityGateCondition struct {
	Status         string `json:"status"`
	MetricKey      string `json:"metricKey"`
	Comparator     string `json:"comparator"`
	ErrorThreshold string `json:"errorThreshold"`
	ActualValue    string `json:"actualValue"`
}

func (wrapper qualityGateWrapper) convert() sonar.QualityGate {
	result := sonar.QualityGate{Status: wrapper.QualityGate.Status, Conditions: map[string]sonar.QualityGateCondition{}}
	if wrapper.QualityGate.Conditions == nil {
		return result
	}

	for _, condition := range wrapper.QualityGate.Conditions {
		result.Conditions[condition.MetricKey] = condition.convert()
	}

	return result
}

var conditionNames = map[string]string{
	"new_reliability_rating":       "New Reliability Rating",
	"new_security_rating":          "New Security Rating",
	"new_maintainability_rating":   "New Maintainability Rating",
	"new_coverage":                 "New Coverage",
	"new_duplicated_lines_density": "New Duplicated Lines Density",
}

func (condition qualityGateCondition) convert() sonar.QualityGateCondition {
	value, _ := strconv.ParseFloat(condition.ActualValue, 32)
	errorThreshold, _ := strconv.ParseFloat(condition.ErrorThreshold, 32)

	return sonar.QualityGateCondition{
		Status:         condition.Status,
		Description:    conditionNames[condition.MetricKey],
		Value:          float32(value),
		ErrorThreshold: float32(errorThreshold),
		Comparator:     condition.Comparator,
	}
}
