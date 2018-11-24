package lib

import (
	"fmt"
	"reflect"
)

func isNil(v interface{}) bool {
	ref := reflect.ValueOf(v)

	return v == nil || ref.Kind() == reflect.Ptr && ref.IsNil()
}

func IsZeroOrNil(v interface{}) bool {
	return isNil(v) || reflect.Zero(reflect.TypeOf(v)) == reflect.ValueOf(v)
}

func Value(v interface{}) interface{} {
	if isNil(v) {
		panic(fmt.Errorf("%v is nil", v))
	}

	ref := reflect.ValueOf(v)

	if ref.Kind() == reflect.Ptr {
		return reflect.Indirect(ref).Interface()
	} else {
		return v
	}
}
