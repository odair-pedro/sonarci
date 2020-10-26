package sonar

type Connection interface {
	GetHostServer() string
	Request(endpoint string) (<-chan []byte, <-chan error)
}
