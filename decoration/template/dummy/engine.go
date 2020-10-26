package dummy

import (
	"errors"
	"reflect"
	"regexp"
)

type Engine struct {
}

func NewEngine() *Engine {
	return &Engine{}
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
		template = processDataSourceField(t.Field(i).Tag.Get("dummy"), v.Field(i).String(), template)
	}
	return template, nil
}

func processDataSourceField(name string, value string, template string) string {
	reg := regexp.MustCompile(`\$\{` + name + `\}`)
	return reg.ReplaceAllString(template, value)
}
