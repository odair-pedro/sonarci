package v7

import (
	"sonarci/net"
	"sonarci/net/http"
	"sonarci/sonar"
	"sonarci/sonar/rest/v6"
	"time"
)

type RestApi struct {
	v6.RestApi
	net.Connection
	Server string
}

func NewRestApi(server string, token string, timeout time.Duration) sonar.Api {
	return &RestApi{Connection: http.NewConnection(server, token, timeout), Server: server}
}
