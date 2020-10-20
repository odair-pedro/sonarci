package factory

import (
	"sonarci/net"
	"sonarci/net/http"
	"time"
)

func CreateConnection(server string, token string, timeout time.Duration) net.Connection {
	return http.NewConnection(server, token, timeout)
}
