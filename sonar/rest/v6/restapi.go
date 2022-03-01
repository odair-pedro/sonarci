package v6

import (
	"sonarci/connection"
	"sonarci/sonar/rest/base"
)

type RestApi struct {
	base.RestApi
}

func NewRestApi(connection connection.Connection) *RestApi {
	return &RestApi{*base.NewRestApi(connection)}
}
