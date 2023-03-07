package structmap

import (
	"reflect"
)

func Decode(st interface{}) (map[string]interface{}, error) {
	stType := reflect.TypeOf(st)
	if stType.Kind() == reflect.Ptr {
		stType = stType.Elem()
	}

	if stType.Kind() != reflect.Struct {
		return nil, ErrNotValidElem
	}

	stVal := reflect.Indirect(reflect.ValueOf(st))
	m := make(map[string]interface{})

	for i := 0; i < stType.NumField(); i++ {
		stField := stType.Field(i)
		tagName, _ := stField.Tag.Lookup(TagName)

		if len(tagName) == 0 {
			tagName = stField.Name
		} else if tagName == TagIgnore {
			continue
		}

		val := stVal.Field(i)
		if !val.CanInterface() {
			continue
		}
		m[tagName] = val.Interface()
	}

	return m, nil
}
