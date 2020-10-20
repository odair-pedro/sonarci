package sonar

import "strings"

type Project struct {
	Id           string
	Organization string
	Key          string
	Name         string
	Visibility   string
}

type QualityGate struct {
	Status     string
	LinkDetail string
	Conditions map[string]QualityGateCondition
}

type QualityGateCondition struct {
	Status         string
	Description    string
	Value          float32
	ErrorThreshold float32
	Comparator     string
}

const statusSuccess = "OK"

func (qualityGate QualityGate) HasPassed() bool {
	return strings.ToUpper(qualityGate.Status) == statusSuccess
}

func (condition QualityGateCondition) HasPassed() bool {
	return strings.ToUpper(condition.Status) == statusSuccess
}
