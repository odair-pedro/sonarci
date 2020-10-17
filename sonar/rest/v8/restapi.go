package v8

import (
	"sonarci/sonar/rest/v7"
	"time"
)

type RestApi struct {
	v7.RestApi
}

func NewRestApi(server string, token string, timeout time.Duration) *RestApi {
	return &RestApi{*v7.NewRestApi(server, token, timeout)}
}
