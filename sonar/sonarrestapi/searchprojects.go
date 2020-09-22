package sonarrestapi

import (
	"encoding/json"
	"sonarci/sonar"
)

const routeSearchProjects = "/api/projects/search?projects="

func (restApi *restApi) SearchProjects(projects string) (<-chan sonar.Project, error) {
	chBuff, chErr := restApi.DoGet(routeSearchProjects + projects)
	err := <-chErr
	if err != nil {
		return nil, err
	}

	buff := <-chBuff
	resp := &searchProjectsResp{}
	err = json.Unmarshal(buff, resp)
	if err != nil {
		return nil, err
	}

	chOut := make(chan sonar.Project, len(resp.Components))
	go func() {
		defer close(chOut)
		for _, comp := range resp.Components {
			chOut <- sonar.Project{
				Id:           comp.Id,
				Organization: comp.Organization,
				Key:          comp.Key,
				Name:         comp.Name,
				Visibility:   comp.Visibility,
			}
		}
	}()

	return chOut, nil
}

type searchProjectsResp struct {
	Components []searchProjectsRespComp `json:"components"`
}

type searchProjectsRespComp struct {
	Id           string `json:"id"`
	Organization string `json:"organization"`
	Key          string `json:"key"`
	Name         string `json:"name"`
	Visibility   string `json:"visibility"`
}
