package v6

import (
	"net/url"
	"sonarci/net"
	connFactory "sonarci/net/factory"
	"time"
)

type RestApi struct {
	net.Connection
	Server string
}

func NewRestApi(server string, token string, timeout time.Duration) *RestApi {
	return &RestApi{Connection: connFactory.CreateHttpConnection(server, token, timeout)}
}

func escapeValue(value string) string {
	return url.QueryEscape(value)
}
