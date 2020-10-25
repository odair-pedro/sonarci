package v8

import (
	"sonarci/sonar"
	"sonarci/sonar/rest/v7"
)

type RestApi struct {
	v7.RestApi
}

func NewRestApi(connection sonar.Connection) *RestApi {
	return &RestApi{*v7.NewRestApi(connection)}
}
