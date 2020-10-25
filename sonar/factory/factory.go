package factory

import (
	"sonarci/sonar"
	"sonarci/sonar/rest/v8"
)

func CreateSonarApi(connection sonar.Connection) sonar.Api {
	return v8.NewRestApi(connection)
}
