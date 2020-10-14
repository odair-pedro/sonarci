package v6

import (
	"net/url"
	"sonarci/net"
	"sonarci/net/http"
	"sonarci/sonar"
	"time"
)

type RestApi struct {
	net.Connection
	Server string
}

func NewRestApi(server string, token string, timeout time.Duration) sonar.Api {
	return &RestApi{Connection: http.NewConnection(server, token, timeout), Server: server}
}

func escapeValue(value string) string {
	return url.QueryEscape(value)
}
