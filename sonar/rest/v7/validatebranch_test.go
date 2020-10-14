package v7

import (
	"testing"
)

func TestRestApi_GetRouteForValidateBranch(t *testing.T) {
	const want = "/api/measures/component?component=%s&branch=%s&metricKeys=alert_status"

	restApi := &RestApi{}
	got := restApi.GetRouteForValidateBranch()

	if got != want {
		t.Errorf("GetRouteForValidateBranch() got %s, want %s", got, want)
	}
}
