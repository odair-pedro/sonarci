package rest

import (
	"encoding/base64"
	"fmt"
	"net/url"
	"sonarci/http"
	"time"
)

type Api struct {
	*http.Connection
}

func NewApi(server string, token string, timeout time.Duration) *Api {
	return &Api{http.NewConnection(server, getAuthentication(token), timeout)}
}

func getAuthentication(token string) string {
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:", token)))
}

func escapeValue(value string) string {
	return url.QueryEscape(value)
}
