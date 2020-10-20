package v8

import (
	"sonarci/net"
	"sonarci/sonar/rest/v7"
)

type RestApi struct {
	v7.RestApi
}

func NewRestApi(connection net.Connection) *RestApi {
	return &RestApi{*v7.NewRestApi(connection)}
}
