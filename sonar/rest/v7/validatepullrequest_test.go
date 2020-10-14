package v7

import (
	"testing"
)

func TestRestApi_GetRouteForValidatePullRequest(t *testing.T) {
	const want = "/api/measures/component?component=%s&pullRequest=%s&metricKeys=alert_status"

	restApi := &RestApi{}
	got := restApi.GetRouteForValidatePullRequest()

	if got != want {
		t.Errorf("GetRouteForValidatePullRequest() got %s, want %s", got, want)
	}
}
