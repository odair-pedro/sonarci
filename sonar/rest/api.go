package rest

import (
	"encoding/base64"
	"fmt"
	"sonarci/http"
	"time"
)

type Api struct {
	*http.Connection
}

//func (a Api) GetServerVersion() (string, error) {
//	panic("implement me")
//}
//
//func (a Api) SearchProjects(projects string) (<-chan sonar.Project, error) {
//	panic("implement me")
//}
//
//func (a Api) ValidateBranch(project string, branch string) (bool, error) {
//	panic("implement me")
//}

func NewApi(server string, token string, timeout time.Duration) *Api {
	return &Api{http.NewConnection(server, getAuthentication(token), timeout)}
}

func getAuthentication(token string) string {
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:", token)))
}
