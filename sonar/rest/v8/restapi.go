package v8

import (
	"sonarci/connection"
	"sonarci/sonar/rest/v7"
)

type RestApi struct {
	v7.RestApi
}

func NewRestApi(connection connection.Connection) *RestApi {
	return &RestApi{*v7.NewRestApi(connection)}
}
