package net

type Connection interface {
	GetHostServer() string
	DoGet(route string) (<-chan []byte, <-chan error)
}
