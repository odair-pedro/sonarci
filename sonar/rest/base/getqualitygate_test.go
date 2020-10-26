package base

import "testing"

func Test_parseMetricKeyInDescription(t *testing.T) {
	type args struct {
		metricKey string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "new_reliability_rating", args: args{metricKey: "new_reliability_rating"}, want: "New Reliability Rating"},
		{name: "new_vulnerabilities", args: args{metricKey: "new_vulnerabilities"}, want: "New Vulnerabilities"},
		{name: "coverage", args: args{metricKey: "coverage"}, want: "Coverage"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseMetricKeyInDescription(tt.args.metricKey); got != tt.want {
				t.Errorf("parseMetricKeyInDescription() = %v, want %v", got, tt.want)
			}
		})
	}
}
