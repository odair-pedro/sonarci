package base

const routeServerVersion = "/api/server/version"

func (restApi *RestApi) GetServerVersion() (string, error) {
	chBuff, chErr := restApi.DoGet(routeServerVersion)
	err := <-chErr
	if err != nil {
		return "", err
	}

	buff := <-chBuff
	return string(buff), nil
}
