package factory

import (
	"sonarci/net"
	"sonarci/sonar"
	"sonarci/sonar/rest/v8"
)

func CreateSonarApi(connection net.Connection) sonar.Api {
	return v8.NewRestApi(connection)
}
