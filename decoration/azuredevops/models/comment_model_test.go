package models

import (
	"sonarci/sonar"
	"testing"
)

func Test_ParseCommentModel_CheckStatus(t *testing.T) {
	type args struct {
		qualityGate sonar.QualityGate
		report      string
	}
	tests := []struct {
		name       string
		args       args
		wantStatus string
	}{
		{name: "test-success", args: args{qualityGate: sonar.QualityGate{Status: "OK"}}, wantStatus: "closed"},
		{name: "test-failed", args: args{qualityGate: sonar.QualityGate{Status: "ANY THING ELSE"}}, wantStatus: "active"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseCommentModel(tt.args.qualityGate, tt.args.report, "any-tag"); got.Status != tt.wantStatus {
				t.Errorf("parseCommentModel() Status = %v, want %v", got.Status, tt.wantStatus)
			}
		})
	}
}

func Test_ParseCommentModel_CheckPropertiesGeneratedBySonarCI_IsTrue(t *testing.T) {
	got := ParseCommentModel(sonar.QualityGate{Status: "anything"}, "report-test", "any-tag").Properties.GeneratedBySonarCI
	if got != true {
		t.Errorf("parseCommentModel() Properties.GeneratedBySonarCI want true")
	}
}
