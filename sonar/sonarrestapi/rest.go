package sonarrestapi

import (
	"net/url"
	"sonarci/net"
	"sonarci/net/http"
	"sonarci/sonar"
	"time"
)

type restApi struct {
	net.Connection
	Server string
}

func NewApi(server string, token string, timeout time.Duration) sonar.Api {
	return &restApi{http.NewConnection(server, token, timeout), server}
}

func escapeValue(value string) string {
	return url.QueryEscape(value)
}
