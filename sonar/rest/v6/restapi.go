package v6

import (
	"sonarci/sonar"
	"sonarci/sonar/rest/base"
)

type RestApi struct {
	base.RestApi
}

func NewRestApi(connection sonar.Connection) *RestApi {
	return &RestApi{*base.NewRestApi(connection)}
}
