package v8

import (
	"sonarci/net"
	"sonarci/net/http"
	"sonarci/sonar"
	"sonarci/sonar/rest/v7"
	"time"
)

type RestApi struct {
	v7.RestApi
	net.Connection
	Server string
}

func NewRestApi(server string, token string, timeout time.Duration) sonar.Api {
	return &RestApi{Connection: http.NewConnection(server, token, timeout), Server: server}
}
