package models

import (
	"sonarci/sonar"
	"testing"
)

func Test_convertStatus(t *testing.T) {
	type args struct {
		status string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test-ok", args: args{status: "ok"}, want: "SUCCESS"},
		{name: "test-OK", args: args{status: "OK"}, want: "SUCCESS"},
		{name: "test-error", args: args{status: "error"}, want: "FAILED"},
		{name: "test-ERROR", args: args{status: "ERROR"}, want: "FAILED"},
		{name: "test-other", args: args{status: "other"}, want: "OTHER"},
		{name: "test-OTHER", args: args{status: "OTHER"}, want: "OTHER"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertStatus(tt.args.status); got != tt.want {
				t.Errorf("convertStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_convertStatusColor(t *testing.T) {
	type args struct {
		status string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "test-ok", args: args{status: "ok"}, want: "brightgreen"},
		{name: "test-OK", args: args{status: "OK"}, want: "brightgreen"},
		{name: "test-error", args: args{status: "error"}, want: "red"},
		{name: "test-ERROR", args: args{status: "ERROR"}, want: "red"},
		{name: "test-other", args: args{status: "other"}, want: "yellow"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertStatusColor(tt.args.status); got != tt.want {
				t.Errorf("convertStatusColor() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_ParseTemplateModel_Ok_WithCoverage(t *testing.T) {
	qualityGate := sonar.QualityGate{Host: "http://localhost", Project: "Project", Source: "123",
		SourceType: "pullrequest", Status: "OK", LinkDetail: "http://localhost/detail", Conditions: map[string]sonar.QualityGateCondition{
			"new_reliability_rating":       {Status: "OK", Description: "new_reliability_rating", Value: 0, ErrorThreshold: 0, Comparator: "GT"},
			"new_security_rating":          {Status: "OK", Description: "new_security_rating", Value: 0, ErrorThreshold: 0, Comparator: "GT"},
			"new_maintainability_rating":   {Status: "OK", Description: "new_maintainability_rating", Value: 0, ErrorThreshold: 0, Comparator: "GT"},
			"new_coverage":                 {Status: "OK", Description: "new_coverage", Value: 10.86789, ErrorThreshold: 0, Comparator: "GT"},
			"new_duplicated_lines_density": {Status: "OK", Description: "new_duplicated_lines_density", Value: 10.86789, ErrorThreshold: 0, Comparator: "GT"},
		}}

	want := TemplateModel{host: "http://localhost", project: "Project", pullRequest: "123", status: "SUCCESS",
		statusColor: "brightgreen", coverage: "10.87%", coverageStatus: "SUCCESS", coverageStatusColor: "brightgreen",
		duplication: "10.87%", duplicationStatus: "SUCCESS", duplicationStatusColor: "brightgreen",
		reliabilityStatus: "SUCCESS", reliabilityStatusColor: "brightgreen", securityStatus: "SUCCESS",
		securityStatusColor: "brightgreen", maintainabilityStatus: "SUCCESS", maintainabilityStatusColor: "brightgreen"}

	got := ParseTemplateModel(qualityGate)
	if got != want {
		t.Errorf("parseTemplateModel() returned unexpected value")
	}
}

func Test_ParseTemplateModel_Error_WithCoverage(t *testing.T) {
	qualityGate := sonar.QualityGate{Host: "http://localhost", Project: "Project", Source: "123",
		SourceType: "pullrequest", Status: "ERROR", LinkDetail: "http://localhost/detail", Conditions: map[string]sonar.QualityGateCondition{
			"new_reliability_rating":       {Status: "OK", Description: "new_reliability_rating", Value: 0, ErrorThreshold: 0, Comparator: "GT"},
			"new_security_rating":          {Status: "ERROR", Description: "new_security_rating", Value: 0, ErrorThreshold: 0, Comparator: "GT"},
			"new_maintainability_rating":   {Status: "OK", Description: "new_maintainability_rating", Value: 0, ErrorThreshold: 0, Comparator: "GT"},
			"new_coverage":                 {Status: "OK", Description: "new_coverage", Value: 10.86789, ErrorThreshold: 0, Comparator: "GT"},
			"new_duplicated_lines_density": {Status: "OK", Description: "new_duplicated_lines_density", Value: 10.86789, ErrorThreshold: 0, Comparator: "GT"},
		}}

	want := TemplateModel{host: "http://localhost", project: "Project", pullRequest: "123", status: "FAILED",
		statusColor: "red", coverage: "10.87%", coverageStatus: "SUCCESS", coverageStatusColor: "brightgreen",
		duplication: "10.87%", duplicationStatus: "SUCCESS", duplicationStatusColor: "brightgreen",
		reliabilityStatus: "SUCCESS", reliabilityStatusColor: "brightgreen", securityStatus: "FAILED",
		securityStatusColor: "red", maintainabilityStatus: "SUCCESS", maintainabilityStatusColor: "brightgreen"}

	got := ParseTemplateModel(qualityGate)
	if got != want {
		t.Errorf("parseTemplateModel() returned unexpected value")
	}
}

func Test_ParseTemplateModel_Ok_WithoutCoverage(t *testing.T) {
	qualityGate := sonar.QualityGate{Host: "http://localhost", Project: "Project", Source: "123",
		SourceType: "pullrequest", Status: "OK", LinkDetail: "http://localhost/detail", Conditions: map[string]sonar.QualityGateCondition{
			"new_reliability_rating":       {Status: "OK", Description: "new_reliability_rating", Value: 0, ErrorThreshold: 0, Comparator: "GT"},
			"new_security_rating":          {Status: "OK", Description: "new_security_rating", Value: 0, ErrorThreshold: 0, Comparator: "GT"},
			"new_maintainability_rating":   {Status: "OK", Description: "new_maintainability_rating", Value: 0, ErrorThreshold: 0, Comparator: "GT"},
			"new_coverage":                 {Status: "", Description: "new_coverage", Value: 0, ErrorThreshold: 0, Comparator: ""},
			"new_duplicated_lines_density": {Status: "OK", Description: "new_duplicated_lines_density", Value: 10.86789, ErrorThreshold: 0, Comparator: "GT"},
		}}

	want := TemplateModel{host: "http://localhost", project: "Project", pullRequest: "123", status: "SUCCESS",
		statusColor: "brightgreen", coverage: "N/A", coverageStatus: "SUCCESS", coverageStatusColor: "lightgray",
		duplication: "10.87%", duplicationStatus: "SUCCESS", duplicationStatusColor: "brightgreen",
		reliabilityStatus: "SUCCESS", reliabilityStatusColor: "brightgreen", securityStatus: "SUCCESS",
		securityStatusColor: "brightgreen", maintainabilityStatus: "SUCCESS", maintainabilityStatusColor: "brightgreen"}

	got := ParseTemplateModel(qualityGate)
	if got != want {
		t.Errorf("parseTemplateModel() returned unexpected value")
	}
}

func Test_ParseTemplateModel_Error_WithoutCoverage(t *testing.T) {
	qualityGate := sonar.QualityGate{Host: "http://localhost", Project: "Project", Source: "123",
		SourceType: "pullrequest", Status: "ERROR", LinkDetail: "http://localhost/detail", Conditions: map[string]sonar.QualityGateCondition{
			"new_reliability_rating":       {Status: "OK", Description: "new_reliability_rating", Value: 0, ErrorThreshold: 0, Comparator: "GT"},
			"new_security_rating":          {Status: "ERROR", Description: "new_security_rating", Value: 0, ErrorThreshold: 0, Comparator: "GT"},
			"new_maintainability_rating":   {Status: "OK", Description: "new_maintainability_rating", Value: 0, ErrorThreshold: 0, Comparator: "GT"},
			"new_coverage":                 {Status: "", Description: "new_coverage", Value: 0, ErrorThreshold: 0, Comparator: ""},
			"new_duplicated_lines_density": {Status: "OK", Description: "new_duplicated_lines_density", Value: 10.86789, ErrorThreshold: 0, Comparator: "GT"},
		}}

	want := TemplateModel{host: "http://localhost", project: "Project", pullRequest: "123", status: "FAILED",
		statusColor: "red", coverage: "N/A", coverageStatus: "SUCCESS", coverageStatusColor: "lightgray",
		duplication: "10.87%", duplicationStatus: "SUCCESS", duplicationStatusColor: "brightgreen",
		reliabilityStatus: "SUCCESS", reliabilityStatusColor: "brightgreen", securityStatus: "FAILED",
		securityStatusColor: "red", maintainabilityStatus: "SUCCESS", maintainabilityStatusColor: "brightgreen"}

	got := ParseTemplateModel(qualityGate)
	if got != want {
		t.Errorf("parseTemplateModel() returned unexpected value")
	}
}

func Test_ParseTemplateModel_Ok_WithoutDuplication(t *testing.T) {
	qualityGate := sonar.QualityGate{Host: "http://localhost", Project: "Project", Source: "123",
		SourceType: "pullrequest", Status: "OK", LinkDetail: "http://localhost/detail", Conditions: map[string]sonar.QualityGateCondition{
			"new_reliability_rating":       {Status: "OK", Description: "new_reliability_rating", Value: 0, ErrorThreshold: 0, Comparator: "GT"},
			"new_security_rating":          {Status: "OK", Description: "new_security_rating", Value: 0, ErrorThreshold: 0, Comparator: "GT"},
			"new_maintainability_rating":   {Status: "OK", Description: "new_maintainability_rating", Value: 0, ErrorThreshold: 0, Comparator: "GT"},
			"new_coverage":                 {Status: "", Description: "new_coverage", Value: 0, ErrorThreshold: 0, Comparator: ""},
			"new_duplicated_lines_density": {Status: "", Description: "new_duplicated_lines_density", Value: 0, ErrorThreshold: 0, Comparator: ""},
		}}

	want := TemplateModel{host: "http://localhost", project: "Project", pullRequest: "123", status: "SUCCESS",
		statusColor: "brightgreen", coverage: "N/A", coverageStatus: "SUCCESS", coverageStatusColor: "lightgray",
		duplication: "N/A", duplicationStatus: "SUCCESS", duplicationStatusColor: "lightgray",
		reliabilityStatus: "SUCCESS", reliabilityStatusColor: "brightgreen", securityStatus: "SUCCESS",
		securityStatusColor: "brightgreen", maintainabilityStatus: "SUCCESS", maintainabilityStatusColor: "brightgreen"}

	got := ParseTemplateModel(qualityGate)
	if got != want {
		t.Errorf("parseTemplateModel() returned unexpected value")
	}
}

func Test_ParseTemplateModel_Error_WithoutDuplication(t *testing.T) {
	qualityGate := sonar.QualityGate{Host: "http://localhost", Project: "Project", Source: "123",
		SourceType: "pullrequest", Status: "ERROR", LinkDetail: "http://localhost/detail", Conditions: map[string]sonar.QualityGateCondition{
			"new_reliability_rating":       {Status: "OK", Description: "new_reliability_rating", Value: 0, ErrorThreshold: 0, Comparator: "GT"},
			"new_security_rating":          {Status: "ERROR", Description: "new_security_rating", Value: 0, ErrorThreshold: 0, Comparator: "GT"},
			"new_maintainability_rating":   {Status: "OK", Description: "new_maintainability_rating", Value: 0, ErrorThreshold: 0, Comparator: "GT"},
			"new_coverage":                 {Status: "", Description: "new_coverage", Value: 0, ErrorThreshold: 0, Comparator: ""},
			"new_duplicated_lines_density": {Status: "", Description: "new_duplicated_lines_density", Value: 0, ErrorThreshold: 0, Comparator: ""},
		}}

	want := TemplateModel{host: "http://localhost", project: "Project", pullRequest: "123", status: "FAILED",
		statusColor: "red", coverage: "N/A", coverageStatus: "SUCCESS", coverageStatusColor: "lightgray",
		duplication: "N/A", duplicationStatus: "SUCCESS", duplicationStatusColor: "lightgray",
		reliabilityStatus: "SUCCESS", reliabilityStatusColor: "brightgreen", securityStatus: "FAILED",
		securityStatusColor: "red", maintainabilityStatus: "SUCCESS", maintainabilityStatusColor: "brightgreen"}

	got := ParseTemplateModel(qualityGate)
	if got != want {
		t.Errorf("parseTemplateModel() returned unexpected value")
	}
}
