package factory

import (
	"sonarci/sonar"
	"sonarci/sonar/rest/v8"
	"time"
)

func CreateLatestSonarRestApi(server string, token string, timeout time.Duration) sonar.Api {
	return v8.NewRestApi(server, token, timeout)
}
