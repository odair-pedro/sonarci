package connection

type Connection interface {
	GetHostServer() string
	Request(endpoint string) (<-chan []byte, <-chan error)
	Send(endpoint string, content []byte, contentType string) (<-chan []byte, <-chan error)
}
