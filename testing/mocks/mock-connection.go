package mocks

type MockConnection struct {
	HostServerMock string
	RequestMock    func(route string) (<-chan []byte, <-chan error)
	SendMock       func(endpoint string, content []byte, contentType string) (<-chan []byte, <-chan error)
}

func (connection *MockConnection) GetHostServer() string {
	return connection.HostServerMock
}

func (connection *MockConnection) Request(route string) (<-chan []byte, <-chan error) {
	return connection.RequestMock(route)
}

func (connection *MockConnection) Send(endpoint string, content []byte, contentType string) (<-chan []byte, <-chan error) {
	return connection.SendMock(endpoint, content, contentType)
}
