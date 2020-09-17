package rest

const routeGetServerVersion = "/api/server/version"

func (api *Api) GetServerVersion() (string, error) {
	//conn := http.NewConnection(abstract.server, abstract.token, abstract.timeout)
	chBuff, chErr := api.DoGet(routeGetServerVersion)
	err := <-chErr
	if err != nil {
		return "", err
	}

	buff := <-chBuff
	return string(buff), nil
}
