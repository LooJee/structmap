package structmap

import (
	"reflect"
)

func getFieldName(stField reflect.StructField) string {
	tagName, _ := stField.Tag.Lookup(TagName)

	if len(tagName) == 0 {
		tagName = stField.Name
	}

	return tagName
}

func StructToMap(st interface{}) (map[string]interface{}, error) {
	stType := reflect.TypeOf(st)
	if stType.Kind() == reflect.Ptr {
		stType = stType.Elem()
	}

	if stType.Kind() != reflect.Struct {
		return nil, ErrNeedStruct
	}

	stVal := reflect.Indirect(reflect.ValueOf(st))
	m := make(map[string]interface{})

	for i := 0; i < stType.NumField(); i++ {
		tagName := getFieldName(stType.Field(i))
		if tagName == TagIgnore {
			continue
		}

		val := stVal.Field(i)
		if !val.CanInterface() {
			continue
		}

		//TODO maybe should check pointer

		m[tagName] = val.Interface()
	}

	return m, nil
}

func MapToStruct(m map[string]interface{}, st interface{}) error {
	if reflect.TypeOf(st).Kind() != reflect.Ptr {
		return ErrNotPtr
	}

	stType := reflect.TypeOf(st).Elem()
	if stType.Kind() != reflect.Struct {
		return ErrNeedStruct
	}

	stVal := reflect.ValueOf(st).Elem()
	println(stVal.NumField())
	for i := 0; i < stType.NumField(); i++ {
		field := stType.Field(i)
		fieldVal := stVal.FieldByName(field.Name)
		if !fieldVal.CanSet() {
			continue
		}

		tagName := getFieldName(field)
		if tagName == TagIgnore {
			continue
		}

		mv, ok := m[tagName]
		if !ok {
			continue
		}

		mvType := reflect.TypeOf(mv)

		if mvType.Kind() == fieldVal.Kind() {
			fieldVal.Set(reflect.ValueOf(mv))
		} else if mvType.ConvertibleTo(field.Type) {
			fieldVal.Set(reflect.ValueOf(mv).Convert(field.Type))
		} else {
			return &ErrTypeNotMatch{field.Name, fieldVal.Kind().String(), mvType.Kind().String()}
		}

		println(reflect.ValueOf(mv).Kind().String())
	}

	return nil
}
