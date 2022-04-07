package dummy

import (
	"errors"
	"net/url"
	"reflect"
	"regexp"
	"sonarci/decoration/template/engine/dummy/resources"
	"strconv"
)

type Engine struct {
}

func NewEngine() *Engine {
	return &Engine{}
}

func (eng *Engine) GetQualityReportTemplate(a ...interface{}) string {
	if len(a) > 0 && a[0] == true {
		return resources.QualityReportTemplateTitleWithTag + resources.QualityReportTemplate
	}

	return resources.QualityReportTemplateTitle + resources.QualityReportTemplate
}

func (eng *Engine) ProcessTemplate(template string, dataSource interface{}) (string, error) {
	if template == "" {
		return "", errors.New("invalid template")
	}
	if dataSource == nil {
		return "", errors.New("invalid data source")
	}

	return eng.processDataSource(template, dataSource)
}

func (eng *Engine) processDataSource(template string, dataSource interface{}) (string, error) {
	v := reflect.ValueOf(dataSource)
	if v.Kind() != reflect.Struct {
		return "", errors.New("invalid data source, it is not a struct")
	}

	t := reflect.TypeOf(dataSource)
	for i := 0; i < v.NumField(); i++ {
		name := t.Field(i).Tag.Get("dummy")
		escape, err := strconv.ParseBool(t.Field(i).Tag.Get("escape"))
		if err != nil {
			escape = false
		}
		template = processDataSourceField(name, v.Field(i).String(), template, escape)
	}
	return template, nil
}

func processDataSourceField(name string, value string, template string, escape bool) string {
	reg := regexp.MustCompile(`\$\{` + name + `\}`)

	if escape {
		value = escapeValue(value)
	}

	return reg.ReplaceAllString(template, value)
}

func escapeValue(value string) string {
	return url.PathEscape(value)
}
