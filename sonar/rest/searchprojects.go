package rest

import (
	"encoding/json"
	"sonarci/http"
	"sonarci/sonar/abstract"
)

const route = "/api/projects/search?projects="

func (api restApi) SearchProjects(projects string) (<-chan abstract.Project, error) {
	conn := http.NewConnection(api.server, api.token, api.timeout)
	chBuff, chErr := conn.DoGetWithChannels(route + projects)
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

	chOut := make(chan abstract.Project, len(resp.Components))
	go func() {
		defer close(chOut)
		for _, comp := range resp.Components {
			chOut <- abstract.Project{
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
