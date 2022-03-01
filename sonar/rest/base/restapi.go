package base

import (
	"net/url"
	"sonarci/connection"
)

type RestApi struct {
	connection.Connection
}

func NewRestApi(connection connection.Connection) *RestApi {
	return &RestApi{connection}
}

func escapeValue(value string) string {
	return url.QueryEscape(value)
}
