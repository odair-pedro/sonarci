package base

import (
	"encoding/json"
	"sonarci/sonar"
)

const routeSearchProjects = "/api/projects/search?projects="

func (restApi *RestApi) SearchProjects(projects string) (<-chan sonar.Project, error) {
	chBuff, chErr := restApi.Get(routeSearchProjects + projects)
	err := <-chErr
	if err != nil {
		return nil, err
	}

	buff := <-chBuff
	resp := &searchProjectsWrapper{}
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

type searchProjectsWrapper struct {
	Components []searchProject `json:"components"`
}

type searchProject struct {
	Id           string `json:"id"`
	Organization string `json:"organization"`
	Key          string `json:"key"`
	Name         string `json:"name"`
	Visibility   string `json:"visibility"`
}
