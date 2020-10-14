package v6

const routeGetServerVersion = "/api/server/version"

func (restApi *RestApi) GetServerVersion() (string, error) {
	chBuff, chErr := restApi.DoGet(restApi.GetRouteServerVersion())
	err := <-chErr
	if err != nil {
		return "", err
	}

	buff := <-chBuff
	return string(buff), nil
}

func (restApi *RestApi) GetRouteServerVersion() string {
	return routeGetServerVersion
}
