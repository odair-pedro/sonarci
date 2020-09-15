package rest

import (
	"encoding/base64"
	"fmt"
	"time"
)

type config struct {
	server  string
	token   string
	timeout time.Duration
}

func (cfg config) getAuthentication() string {
	return base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:", cfg.token)))
}

//
//var (
//	defaultTimeout      time.Duration = DefaultTimeout
//	tokenAuthentication string
//	server              string
//)
//
//func SetDefaultTimeout(time time.Duration) {
//	defaultTimeout = time
//}
//
//func SetAuthentication(token string) {
//	tokenAuthentication = base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:", token)))
//}
//
//func SetServer(serverUrl string) {
//	server = serverUrl
//}
