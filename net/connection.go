package net

type Connection interface {
	DoGet(route string) (<-chan []byte, <-chan error)
}
