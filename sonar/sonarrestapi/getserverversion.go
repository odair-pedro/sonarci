package sonarrestapi

const routeGetServerVersion = "/api/server/version"

func (restApi *restApi) GetServerVersion() (string, error) {
	chBuff, chErr := restApi.DoGet(routeGetServerVersion)
	err := <-chErr
	if err != nil {
		return "", err
	}

	buff := <-chBuff
	return string(buff), nil
}
