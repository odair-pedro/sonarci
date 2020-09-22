package sonarrestapi

import (
	"encoding/base64"
	"fmt"
	"net/url"
	"sonarci/net"
	"sonarci/net/http"
	"sonarci/sonar"
	"time"
)

type restApi struct {
	net.Connection
	Server string
}

func NewApi(server string, token string, timeout time.Duration) sonar.Api {
	return &restApi{http.NewConnection(server, getAuthentication(token), timeout), server}
}

func getAuthentication(token string) string {
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:", token)))
}

func escapeValue(value string) string {
	return url.QueryEscape(value)
}
