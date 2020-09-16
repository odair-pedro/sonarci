package rest

import (
	"sonarci/sonar/abstract"
	"time"
)

type restApi struct {
	config
}

func NewApi(server string, token string, timeout time.Duration) abstract.Api {
	return restApi{config{server: server, token: token, timeout: timeout}}
}

//func (api restApi) GetServerVersion() (string, error) {
//	return "1.0.0", nil
//}
