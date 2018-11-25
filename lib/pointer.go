package lib

import (
	"fmt"
	"reflect"
)

func IsNil(v interface{}) bool {
	ref := reflect.ValueOf(v)

	return v == nil || ref.Kind() == reflect.Ptr && ref.IsNil()
}

func IsZeroOrNil(v interface{}) bool {
	if IsNil(v) {
		return true
	}

	zero := reflect.Zero(reflect.TypeOf(v))
	value := reflect.ValueOf(v)

	// only required values for now
	switch zero.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return zero.Int() == value.Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return zero.Uint() == value.Uint()
	case reflect.String:
		return zero.String() == value.String()
	case reflect.Bool:
		return zero.Bool() == value.Bool()
	case reflect.Interface:
		return zero.Interface() == value.Interface()
	}

	reflect.ValueOf(v).Interface()

	return IsNil(v) || reflect.Zero(reflect.TypeOf(v)) == reflect.ValueOf(v)
}

func Value(v interface{}) interface{} {
	if IsNil(v) {
		panic(fmt.Errorf("%v is nil", v))
	}

	ref := reflect.ValueOf(v)

	if ref.Kind() == reflect.Ptr {
		return reflect.Indirect(ref).Interface()
	} else {
		return v
	}
}
