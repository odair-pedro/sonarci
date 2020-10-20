package v7

import (
	"sonarci/net"
	v6 "sonarci/sonar/rest/v6"
)

type RestApi struct {
	v6.RestApi
}

func NewRestApi(connection net.Connection) *RestApi {
	return &RestApi{*v6.NewRestApi(connection)}
}
