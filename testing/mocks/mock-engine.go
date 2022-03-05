package mocks

type MockEngine struct {
	GetQualityReportTemplateMock func(a ...interface{}) string
	ProcessTemplateMock          func(template string, dataSource interface{}) (string, error)
}

func (engine *MockEngine) GetQualityReportTemplate(a ...interface{}) string {
	return engine.GetQualityReportTemplateMock(a)
}

func (engine *MockEngine) ProcessTemplate(template string, dataSource interface{}) (string, error) {
	return engine.ProcessTemplateMock(template, dataSource)
}
