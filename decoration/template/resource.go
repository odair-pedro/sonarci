package template

type Resource interface {
	GetQualityReportTemplate(a ...interface{}) string
}
