package mocks

type MockEngine struct {
	ProcessTemplateMock func(template string, dataSource interface{}) (string, error)
}

func (engine *MockEngine) ProcessTemplate(template string, dataSource interface{}) (string, error) {
	return engine.ProcessTemplateMock(template, dataSource)
}
