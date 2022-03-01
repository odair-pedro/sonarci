package factory

import (
	"sonarci/connection"
	"sonarci/sonar"
	"sonarci/sonar/rest/v8"
)

func CreateSonarApi(connection connection.Connection) sonar.Api {
	return v8.NewRestApi(connection)
}
