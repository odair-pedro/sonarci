package v6

import (
	"sonarci/net"
	"sonarci/sonar/rest/base"
)

type RestApi struct {
	base.RestApi
}

func NewRestApi(connection net.Connection) *RestApi {
	return &RestApi{*base.NewRestApi(connection)}
}
