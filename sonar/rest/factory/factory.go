package factory

import (
	"sonarci/sonar"
	latest "sonarci/sonar/rest/v8"
	"time"
)

func GetLatestSonarApi(server string, token string, timeout time.Duration) sonar.Api {
	return latest.NewRestApi(server, token, timeout)
}
