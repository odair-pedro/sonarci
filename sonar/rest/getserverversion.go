package rest

import "sonarci/http"

const routeGetServerVersion = "/api/server/version"

func (api restApi) GetServerVersion() (string, error) {
	conn := http.NewConnection(api.server, api.token, api.timeout)
	chBuff, chErr := conn.DoGet(routeGetServerVersion)
	err := <-chErr
	if err != nil {
		return "", err
	}

	buff := <-chBuff
	return string(buff), nil
}
