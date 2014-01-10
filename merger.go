package gofigure

import (
	"fmt"
	"reflect"
)

func mergeMapAndStruct(theMap map[string]interface{}, spec interface{}) error {
	fmt.Println("mergin'")
	s := reflect.ValueOf(spec).Elem()
	if s.Kind() != reflect.Struct {
		return fmt.Errorf("Invalid spec! Needs to be a struct.")
	}
	typeOfSpec := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		if f.CanSet() {
			fieldName := typeOfSpec.Field(i).Name
			value, ok := theMap[fieldName]
			if !ok {
				fmt.Println("Field %s not found in %s", fieldName, spec)
				continue
			}
			switch f.Kind() {
			case reflect.String:
				f.SetString(value.(string))
			case reflect.Bool:
				f.SetBool(value.(bool))
			case reflect.Int:
				switch value.(type) {
				case int:
					f.SetInt(int64(value.(int)))
				case int16:
					f.SetInt(int64(value.(int16)))
				case int32:
					f.SetInt(int64(value.(int32)))
				case int64:
					f.SetInt(value.(int64))
				default:
					panic("Non-int!")
				}
			}
		}
	}
	return nil
}
