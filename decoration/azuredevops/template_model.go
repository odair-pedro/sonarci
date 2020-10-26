package azuredevops

import (
	"sonarci/sonar"
	"strconv"
	"strings"
)

const (
	keyNewReliabilityRating      = "new_reliability_rating"
	keyNewSecurityRating         = "new_security_rating"
	keyNewMaintainabilityRating  = "new_maintainability_rating"
	keyNewCoverage               = "new_coverage"
	keyNewDuplicatedLinesDensity = "new_duplicated_lines_density"
)

type templateModel struct {
	host                       string `dummy:"host"`
	project                    string `dummy:"project"`
	pullRequest                string `dummy:"pullrequest"`
	status                     string `dummy:"status"`
	statusColor                string `dummy:"status-color"`
	coverage                   string `dummy:"cov"`
	coverageStatus             string `dummy:"cov-status"`
	coverageStatusColor        string `dummy:"cov-status-color"`
	duplication                string `dummy:"dup"`
	duplicationStatus          string `dummy:"dup-status"`
	duplicationStatusColor     string `dummy:"dup-status-color"`
	reliabilityStatus          string `dummy:"rel-status"`
	reliabilityStatusColor     string `dummy:"rel-status-color"`
	securityStatus             string `dummy:"sec-status"`
	securityStatusColor        string `dummy:"sec-status-color"`
	maintainabilityStatus      string `dummy:"mtb-status"`
	maintainabilityStatusColor string `dummy:"mtb-status-color"`
}

func parseTemplateModel(qualityGate sonar.QualityGate) templateModel {
	cov := qualityGate.Conditions[keyNewCoverage]
	dup := qualityGate.Conditions[keyNewDuplicatedLinesDensity]
	rel := qualityGate.Conditions[keyNewReliabilityRating]
	sec := qualityGate.Conditions[keyNewSecurityRating]
	mtb := qualityGate.Conditions[keyNewMaintainabilityRating]

	return templateModel{
		host:                       qualityGate.Host,
		project:                    qualityGate.Project,
		pullRequest:                qualityGate.Source,
		status:                     convertStatus(qualityGate.Status),
		statusColor:                convertStatusColor(qualityGate.Status),
		coverage:                   strconv.FormatFloat(float64(cov.Value), 'f', 2, 32),
		coverageStatus:             convertStatus(cov.Status),
		coverageStatusColor:        convertStatusColor(cov.Status),
		duplication:                strconv.FormatFloat(float64(dup.Value), 'f', 2, 32),
		duplicationStatus:          convertStatus(dup.Status),
		duplicationStatusColor:     convertStatusColor(dup.Status),
		reliabilityStatus:          convertStatus(rel.Status),
		reliabilityStatusColor:     convertStatusColor(rel.Status),
		securityStatus:             convertStatus(sec.Status),
		securityStatusColor:        convertStatusColor(sec.Status),
		maintainabilityStatus:      convertStatus(mtb.Status),
		maintainabilityStatusColor: convertStatusColor(mtb.Status),
	}
}

func convertStatus(status string) string {
	status = strings.ToUpper(status)
	switch status {
	case "OK":
		return "SUCCESS"
	case "ERROR":
		return "FAILED"
	default:
		return status
	}
}

func convertStatusColor(status string) string {
	status = strings.ToUpper(status)
	switch status {
	case "OK":
		return "green"
	case "ERROR":
		return "red"
	default:
		return "yellow"
	}
}
