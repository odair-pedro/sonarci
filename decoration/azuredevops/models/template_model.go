package models

import (
	"sonarci/sonar"
	"strconv"
	"strings"
)

const (
	_keyNewReliabilityRating      = "new_reliability_rating"
	_keyNewSecurityRating         = "new_security_rating"
	_keyNewMaintainabilityRating  = "new_maintainability_rating"
	_keyNewCoverage               = "new_coverage"
	_keyNewDuplicatedLinesDensity = "new_duplicated_lines_density"
	_keyNewCodeSmells             = "new_code_smells"
)

type TemplateModel struct {
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
	codeSmells                 string `dummy:"smells"`
	codeSmellsStatus           string `dummy:"smells-status"`
	codeSmellsStatusColor      string `dummy:"smells-status-color"`
	tag                        string `dummy:"tag"`
}

func ParseTemplateModel(qualityGate sonar.QualityGate, tag string) TemplateModel {
	cov := qualityGate.Conditions[_keyNewCoverage]
	dup := qualityGate.Conditions[_keyNewDuplicatedLinesDensity]
	rel := qualityGate.Conditions[_keyNewReliabilityRating]
	sec := qualityGate.Conditions[_keyNewSecurityRating]
	mtb := qualityGate.Conditions[_keyNewMaintainabilityRating]
	smells := qualityGate.Conditions[_keyNewCodeSmells]

	model := TemplateModel{
		host:                       qualityGate.Host,
		project:                    qualityGate.Project,
		pullRequest:                qualityGate.Source,
		status:                     _convertStatus(qualityGate.Status),
		statusColor:                _convertStatusColor(qualityGate.Status),
		coverage:                   strconv.FormatFloat(float64(cov.Value), 'f', 2, 32) + "%",
		coverageStatus:             _convertStatus(cov.Status),
		coverageStatusColor:        _convertStatusColor(cov.Status),
		duplication:                strconv.FormatFloat(float64(dup.Value), 'f', 2, 32) + "%",
		duplicationStatus:          _convertStatus(dup.Status),
		duplicationStatusColor:     _convertStatusColor(dup.Status),
		reliabilityStatus:          _convertStatus(rel.Status),
		reliabilityStatusColor:     _convertStatusColor(rel.Status),
		securityStatus:             _convertStatus(sec.Status),
		securityStatusColor:        _convertStatusColor(sec.Status),
		maintainabilityStatus:      _convertStatus(mtb.Status),
		maintainabilityStatusColor: _convertStatusColor(mtb.Status),
		codeSmells:                 strconv.FormatFloat(float64(smells.Value), 'f', 0, 32),
		codeSmellsStatus:           _convertStatus(smells.Status),
		codeSmellsStatusColor:      _convertStatusColor(smells.Status),
		tag:                        tag,
	}

	model = _checkCoverageData(model)
	model = _checkDuplicationData(model)

	return model
}

func _convertStatus(status string) string {
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

func _convertStatusColor(status string) string {
	status = strings.ToUpper(status)
	switch status {
	case "OK":
		return "brightgreen"
	case "ERROR":
		return "red"
	default:
		return "yellow"
	}
}

func _checkCoverageData(model TemplateModel) TemplateModel {
	if model.coverageStatus == "" {
		model.coverage = "N/A"
		model.coverageStatus = "SUCCESS"
		model.coverageStatusColor = "lightgray"
	}
	return model
}

func _checkDuplicationData(model TemplateModel) TemplateModel {
	if model.duplicationStatus == "" {
		model.duplication = "N/A"
		model.duplicationStatus = "SUCCESS"
		model.duplicationStatusColor = "lightgray"
	}
	return model
}
