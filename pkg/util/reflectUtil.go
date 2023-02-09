package util

import (
	"errors"
	"reflect"
)

func GetFieldsOfStruct(v any) ([]string, error) {

	refObj := reflect.ValueOf(v)
	if refObj.Kind() == reflect.Ptr {
		refObj = reflect.Indirect(refObj)
	}

	if refObj.Kind() != reflect.Struct {
		return nil, errors.New("argument isn't struct")
	}
	structType := refObj.Type()
	res := make([]string, structType.NumField())
	for i := 0; i < refObj.NumField(); i++ {
		res[i] = structType.Field(i).Name
	}
	return res, nil
}
