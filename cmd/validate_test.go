package cmd

import (
	"sonarci/sonar"
	"testing"
)

func Test_NewValidateCmd_CheckReturn(t *testing.T) {
	cmd := NewValidateCmd()
	if cmd == nil {
		t.Errorf("NewValidateCmd() = nil")
	}
}

func Test_genQualityReport(t *testing.T) {
	const want = "+------------------------------+------------+-----------------+--------------+--------+\n" +
		"| METRIC                       | COMPARATOR | ERROR THRESHOLD | ACTUAL VALUE | STATUS |\n" +
		"+------------------------------+------------+-----------------+--------------+--------+\n" +
		"| New Reliability Rating       | GT         | 30.00000        | 56.90000     | STATUS |\n" +
		"| New Security Rating          | GT         | 15.00000        | 20.50000     | STATUS |\n" +
		"| New Vulnerabilities          | GT         | 0.00000         | 3.00000      | STATUS |\n" +
		"+------------------------------+------------+-----------------+--------------+--------+\n" +
		"|                              |            |                 | QUALITY GATE | STATUS |\n" +
		"+------------------------------+------------+-----------------+--------------+--------+"

	qualityGate := sonar.QualityGate{
		Status:     "STATUS",
		LinkDetail: "http://link-detail",
		Conditions: map[string]sonar.QualityGateCondition{
			"new_reliability_rating": {Status: "STATUS", Description: "New Reliability Rating", Value: 56.9, ErrorThreshold: 30, Comparator: "GT"},
			"new_security_rating":    {Status: "STATUS", Description: "New Security Rating", Value: 20.5, ErrorThreshold: 15, Comparator: "GT"},
			"new_vulnerabilities":    {Status: "STATUS", Description: "New Vulnerabilities", Value: 3, ErrorThreshold: 0, Comparator: "GT"},
		},
	}

	report := genQualityReport(qualityGate)
	if report != want {
		t.Errorf("genQualityReport() = \n%s \n\nwant: \n%s", report, want)
	} else {
		t.Logf("\n%s", report)
	}
}
