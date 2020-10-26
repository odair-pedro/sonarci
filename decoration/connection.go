package decoration

type Connection interface {
	GetHostServer() string
	Send(endpoint string, content []byte, contentType string) (<-chan []byte, <-chan error)
}
