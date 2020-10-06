package dummyengine

import (
	"errors"
	"reflect"
	"regexp"
	"sonarci/decoration/template"
)

type dummyEngine struct {
	template string
}

func NewEngine() template.Engine {
	return &dummyEngine{}
}

func (eng *dummyEngine) ProcessTemplate(template string, dataSource interface{}) (string, error) {
	if template == "" {
		return "", errors.New("invalid template")
	}
	if dataSource == nil {
		return "", errors.New("invalid data source")
	}

	eng.template = template
	return eng.processDataSource(dataSource)
}

func (eng *dummyEngine) processDataSource(dataSource interface{}) (string, error) {
	v := reflect.ValueOf(dataSource)
	if v.Kind() != reflect.Struct {
		return "", errors.New("invalid data source, it is not a struct")
	}

	t := reflect.TypeOf(dataSource)
	result := eng.template
	for i := 0; i < v.NumField(); i++ {
		result = processDataSourceField(t.Field(i).Tag.Get("dummy"), v.Field(i).String(), result)
	}
	return result, nil
}

func processDataSourceField(name string, value string, template string) string {
	reg := regexp.MustCompile(`\$\{` + name + `\}`)
	return reg.ReplaceAllString(template, value)
}
