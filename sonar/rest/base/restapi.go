package base

import (
	"net/url"
	"sonarci/net"
)

type RestApi struct {
	net.Connection
}

func NewRestApi(connection net.Connection) *RestApi {
	return &RestApi{connection}
}

func escapeValue(value string) string {
	return url.QueryEscape(value)
}
