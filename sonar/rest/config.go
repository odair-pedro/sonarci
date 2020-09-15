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
