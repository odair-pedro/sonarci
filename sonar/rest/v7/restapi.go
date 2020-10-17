package v7

import (
	"sonarci/sonar/rest/v6"
	"time"
)

type RestApi struct {
	v6.RestApi
}

func NewRestApi(server string, token string, timeout time.Duration) *RestApi {
	return &RestApi{*v6.NewRestApi(server, token, timeout)}
}
