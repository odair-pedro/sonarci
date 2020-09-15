package sonar

import (
	"sonarci/sonar/abstract"
	"sonarci/sonar/rest"
	"time"
)

func NewApi(server string, token string, timeout time.Duration) abstract.Api {
	return rest.NewApi(server, token, timeout)
}
