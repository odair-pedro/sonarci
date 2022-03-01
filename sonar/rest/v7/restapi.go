package v7

import (
	"sonarci/connection"
	v6 "sonarci/sonar/rest/v6"
)

type RestApi struct {
	v6.RestApi
}

func NewRestApi(connection connection.Connection) *RestApi {
	return &RestApi{*v6.NewRestApi(connection)}
}
