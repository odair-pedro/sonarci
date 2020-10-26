package v7

import (
	"sonarci/sonar"
	v6 "sonarci/sonar/rest/v6"
)

type RestApi struct {
	v6.RestApi
}

func NewRestApi(connection sonar.Connection) *RestApi {
	return &RestApi{*v6.NewRestApi(connection)}
}
