package template

type Engine interface {
	ProcessTemplate(template string, dataSource interface{}) (string, error)
}

//func PrintMembers(obj interface{}, memberName string) error {
//	if obj == nil {
//		return errors.New("Invalid object")
//	}
//
//	value := reflect.ValueOf(obj).FieldByName(memberName)
//	log.Println(value)
//	log.Println(value.String())
//	return nil
//}
