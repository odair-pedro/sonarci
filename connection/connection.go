package connection

type Connection interface {
	GetHostServer() string
	Delete(endpoint string) <-chan error
	Get(endpoint string) (<-chan []byte, <-chan error)
	Post(endpoint string, content []byte, contentType string) (<-chan []byte, <-chan error)
}
