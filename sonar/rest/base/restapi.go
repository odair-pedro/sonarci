package base

import (
	"net/url"
	"sonarci/sonar"
)

type RestApi struct {
	sonar.Connection
}

func NewRestApi(connection sonar.Connection) *RestApi {
	return &RestApi{connection}
}

func escapeValue(value string) string {
	return url.QueryEscape(value)
}
