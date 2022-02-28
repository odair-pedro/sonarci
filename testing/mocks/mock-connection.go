package mocks

type MockConnection struct {
	HostServerMock string
	DeleteMock     func(route string) <-chan error
	GetMock        func(route string) (<-chan []byte, <-chan error)
	PostMock       func(endpoint string, content []byte, contentType string) (<-chan []byte, <-chan error)
}

func (connection *MockConnection) GetHostServer() string {
	return connection.HostServerMock
}

func (connection *MockConnection) Delete(route string) <-chan error {
	return connection.DeleteMock(route)
}

func (connection *MockConnection) Get(route string) (<-chan []byte, <-chan error) {
	return connection.GetMock(route)
}

func (connection *MockConnection) Post(endpoint string, content []byte, contentType string) (<-chan []byte, <-chan error) {
	return connection.PostMock(endpoint, content, contentType)
}
