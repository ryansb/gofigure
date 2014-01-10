package gofigure

import (
	"fmt"
	"reflect"
)

func mergeMapAndStruct(theMap map[string]interface{}, spec interface{}) error {
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
				// the value isn't in the map. Skip dat.
				continue
			}
			switch f.Kind() {
			case reflect.String:
				f.SetString(value.(string))
			case reflect.Bool:
				f.SetBool(value.(bool))
			case reflect.Float32, reflect.Float64:
				switch value.(type) {
				case float32:
					f.SetFloat(float64(value.(float32)))
				case float64:
					f.SetFloat(value.(float64))
				}
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
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
					panic("This should never happen. A non-int has been detected as an int by 'reflect'")
				}
			}
		}
	}
	return nil
}
