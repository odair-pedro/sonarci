package net

type Connection interface {
	GetHostServer() string
	Request(endpoint string) (<-chan []byte, <-chan error)
	Send(data []byte, endpoint string) (<-chan []byte, <-chan error)
}
